package exec

import (
	"bytes"
	"context"
	"fmt"
	"github.com/hedzr/log"
	"gopkg.in/hedzr/errors.v2"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// Call executes the command line via system (OS).
//
// DO NOT QUOTE: in 'cmd', A command line shouldn't has quoted parts.
// These are bad:
//
//    cmd := "ls '/usr/bin'"
//    cmd := `tar "c:/My Documents/"`
//
// Uses CallSlice if your args includes space (like 'c:/My Documents/')
func Call(cmd string, fn func(retCode int, stdoutText string)) (err error) {
	a := strings.Split(cmd, " ")
	err = internalCallImpl(a, fn, true)
	return
}

// CallQuiet executes the command line via system (OS) without error printing.
//
// DO NOT QUOTE: in 'cmd', A command line shouldn't has quoted parts.
// These are bad:
//
//    cmd := "ls '/usr/bin'"
//    cmd := `tar "c:/My Documents/"`
//
// Uses CallSliceQuiet if your args includes space (like 'c:/My Documents/')
func CallQuiet(cmd string, fn func(retCode int, stdoutText string)) (err error) {
	a := strings.Split(cmd, " ")
	err = internalCallImpl(a, fn, false)
	return
}

// CallSlice executes the command line via system (OS).
//
func CallSlice(cmd []string, fn func(retCode int, stdoutText string)) (err error) {
	err = internalCallImpl(cmd, fn, true)
	return
}

// CallSliceQuiet executes the command line via system (OS) without error printing.
//
func CallSliceQuiet(cmd []string, fn func(retCode int, stdoutText string)) (err error) {
	err = internalCallImpl(cmd, fn, false)
	return
}

// internalCallImpl executes the command line via system (OS) without error printing.
func internalCallImpl(cmd []string, fn func(retCode int, stdoutText string), autoErrReport bool) (err error) {
	var (
		str string
		rc  int
	)

	_, str, err = RunWithOutput(cmd[0], cmd[1:]...)
	if err != nil {
		if autoErrReport {
			log.Errorf("Error on launching '%v': %v", cmd, err)
		}
		return
	}
	fn(rc, str)
	return
}

type calling struct {
	*exec.Cmd

	err          error
	wg           sync.WaitGroup
	stdout       io.ReadCloser
	stderr       io.ReadCloser
	stdoutWriter io.Writer
	stderrWriter io.Writer
	envCopied    bool

	retCode int
	output  bytes.Buffer
	slurp   bytes.Buffer

	quiet bool

	onOK    func(retCode int, stdoutText string)
	onError func(err error, retCode int, stdoutText, stderrText string)
}

func New(opts ...Opt) *calling {
	c := &calling{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *calling) RunAndCheckError() error {
	c.Run()
	return c.err
}

func (c *calling) Run() {
	err := c.runNow()
	if err == nil {
		if c.onOK != nil {
			c.onOK(c.retCode, c.output.String())
		}
	} else {
		if c.onError != nil {
			c.onError(err, c.retCode, c.output.String(), c.slurp.String())
		} else if !c.quiet {
			if c.output.Len() > 0 {
				log.Printf("OUTPUT:\n%v\n", c.output.String())
			}
			if c.slurp.Len() > 0 {
				log.Errorf("SLURP:\n%v\n", c.slurp.String())
			}
			log.Errorf("system call failed: %v, command-line: %q", err, c.Path)
		}
	}
	return
}

func (c *calling) runNow() error {
	if c.stdout != nil {
		defer c.stdout.Close()
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			if c.stdoutWriter != nil {
				_, _ = io.Copy(c.stdoutWriter, c.stderr)
			} else {
				_, _ = io.Copy(&c.output, c.stdout)
			}
		}()
	} else {
		c.Stdout = os.Stdout
	}

	//c.stderr, c.err = c.Cmd.StderrPipe()
	//if c.err != nil {
	//	// Failed to connect pipe
	//	c.err = fmt.Errorf("failed to connect stderr pipe: %v, cmd: %q", c.err, c.Path)
	//	return c.err
	//}

	if c.stderr != nil {
		defer c.stderr.Close()
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			if c.stderrWriter != nil {
				_, _ = io.Copy(c.stderrWriter, c.stderr)
			} else {
				_, _ = io.Copy(&c.slurp, c.stderr)
			}
		}()
	} else {
		c.Stderr = os.Stderr
	}

	if c.err = c.Cmd.Start(); c.err != nil {
		// Problem while copying stdin, stdout, or stderr
		c.err = fmt.Errorf("failed: %v, cmd: %q", c.err, c.Path)
		return c.err
	}

	// Zero exit status
	// Darwin: launchctl can fail with a zero exit status,
	// so check for emtpy stderr
	if c.Path == "launchctl" {
		slurpText, _ := ioutil.ReadAll(c.stderr)
		if len(slurpText) > 0 && !bytes.HasSuffix(slurpText, []byte("Operation now in progress\n")) {
			c.err = fmt.Errorf("failed with stderr: %s, cmd: %q", slurpText, c.Path)
			return c.err
		}
	}

	// slurp, _ := ioutil.ReadAll(stderr)

	c.wg.Wait()

	if c.err = c.Cmd.Wait(); c.err != nil {
		exitStatus, ok := IsExitError(c.err)
		if ok {
			// Command didn't exit with a zero exit status.
			c.retCode = exitStatus
			c.err = errors.New("%q failed with stderr:\n%v\n  ", c.Path, c.slurp.String()).Attach(c.err)
			return c.err
		}

		// An error occurred and there is no exit status.
		//return 0, output, fmt.Errorf("%q failed: %v |\n  stderr: %s", command, err.Error(), slurp)
		c.err = errors.New("%q failed with stderr:\n%v\n  ", c.Path, c.slurp.String()).Attach(c.err)
		return c.err
	}

	return nil
}

func (c *calling) OutputText() string { return c.output.String() }
func (c *calling) SlurpText() string  { return c.slurp.String() }
func (c *calling) RetCode() int       { return c.retCode }

func (c *calling) WithCommand(cmd string, args ...string) *calling {
	c.Cmd = exec.Command(cmd, args...)
	return c
}

func (c *calling) WithCommandString(cmd string) *calling {
	a := strings.Split(cmd, " ")
	c.Cmd = exec.Command(a[0], a[1:]...)
	return c
}

func (c *calling) WithCommandSlice(cmd ...string) *calling {
	c.Cmd = exec.Command(cmd[0], cmd[1:]...)
	return c
}

func (c *calling) WithEnv(key, value string) *calling {
	if key != "" {
		if c.envCopied == false {
			c.envCopied = true
			c.Cmd.Env = append(c.Cmd.Env, os.Environ()...)
		}
		c.Cmd.Env = append(c.Cmd.Env, key+"="+value)
	}
	return c
}

func (c *calling) WithWorkDir(dir string) *calling {
	c.Cmd.Dir = dir
	return c
}

func (c *calling) WithExtraFiles(files ...*os.File) *calling {
	c.Cmd.ExtraFiles = files
	return c
}

func (c *calling) WithContext(ctx context.Context) *calling {
	c.Cmd = exec.CommandContext(ctx, c.Cmd.Path, c.Cmd.Args...)
	return c
}

func (c *calling) WithStdoutCaught(writer ...io.Writer) *calling {
	for _, w := range writer {
		c.stdoutWriter = w
	}
	c.stdout, c.err = c.Cmd.StdoutPipe()
	if c.err != nil {
		// Failed to connect pipe
		c.err = fmt.Errorf("failed to connect stdout pipe: %v, cmd: %q", c.err, c.Path)
	}
	return c
}

func (c *calling) WithStderrCaught(writer ...io.Writer) *calling {
	for _, w := range writer {
		c.stderrWriter = w
	}
	c.stderr, c.err = c.Cmd.StderrPipe()
	if c.err != nil {
		// Failed to connect pipe
		c.err = fmt.Errorf("failed to connect stderr pipe: %v, cmd: %q", c.err, c.Path)
	}
	return c
}

func (c *calling) WithOnOK(onOK func(retCode int, stdoutText string)) *calling {
	c.onOK = onOK
	return c
}

func (c *calling) WithOnError(onError func(err error, retCode int, stdoutText, stderrText string)) *calling {
	c.onError = onError
	return c
}

type Opt func(*calling)

func WithCommand(cmd string, args ...string) Opt {
	return func(c *calling) {
		c.Cmd = exec.Command(cmd, args...)
	}
}

func WithCommandString(cmd string) Opt {
	return func(c *calling) {
		a := strings.Split(cmd, " ")
		c.Cmd = exec.Command(a[0], a[1:]...)
	}
}

func WithCommandSlice(cmd ...string) Opt {
	return func(c *calling) {
		c.Cmd = exec.Command(cmd[0], cmd[1:]...)
	}
}

func WithEnv(env ...string) Opt {
	return func(c *calling) {
		c.Cmd.Env = append(c.Cmd.Env, env...)
	}
}

func WithWorkDir(dir string) Opt {
	return func(c *calling) {
		c.Cmd.Dir = dir
	}
}

func WithExtraFiles(files ...*os.File) Opt {
	return func(c *calling) {
		c.Cmd.ExtraFiles = files
	}
}

func WithContext(ctx context.Context) Opt {
	return func(c *calling) {
		c.Cmd = exec.CommandContext(ctx, c.Cmd.Path, c.Cmd.Args...)
	}
}

func WithStdoutCaught() Opt {
	return func(c *calling) {
		c.stdout, c.err = c.Cmd.StdoutPipe()
		if c.err != nil {
			// Failed to connect pipe
			c.err = fmt.Errorf("failed to connect stdout pipe: %v, cmd: %q", c.err, c.Path)
			return
		}
	}
}

func WithOnOK(onOK func(retCode int, stdoutText string)) Opt {
	return func(c *calling) {
		c.onOK = onOK
	}
}

func WithOnError(onError func(err error, retCode int, stdoutText, stderrText string)) Opt {
	return func(c *calling) {
		c.onError = onError
	}
}
