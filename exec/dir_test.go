/*
 * Copyright © 2019 Hedzr Yeh.
 */

package exec_test

import (
	"github.com/hedzr/log/exec"
	"os"
	"path"
	"testing"
)

// TestIsDirectory tests more
//
// usage:
//   go test ./... -v -test.run '^TestIsDirectory$'
//
func TestIsDirectory(t *testing.T) {
	t.Logf("osargs[0] = %v", os.Args[0])
	// t.Logf("InTesting: %v", cmdr.InTesting())
	// t.Logf("InDebugging: %v", cmdr.InDebugging())

	exec.NormalizeDir("")

	if yes, err := exec.IsDirectory("./conf.d1"); yes {
		t.Fatal(err)
	}
	if yes, err := exec.IsDirectory("../exec"); !yes {
		t.Fatal(err)
	}
	if yes, err := exec.IsRegularFile("./doc1.golang"); yes {
		t.Fatal(err)
	}
	if yes, err := exec.IsRegularFile("./dir.go"); !yes {
		t.Fatal(err)
	}
}

func TestForDir(t *testing.T) {
	// defer logex.CaptureLog(t).Release()

	dir := "$HOME/.local"
	if !exec.FileExists(dir) {
		dir = "$HOME/.config"
	}

	err := exec.ForDir(dir, func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
		if fi.IsDir() {
			t.Logf("  - dir: %v/[%v]", cwd, fi.Name())
		} else {
			t.Logf("  - file: %v/%v", cwd, fi.Name())
		}
		return
	})

	if err != nil {
		t.Errorf("wrong for ForDir(): %v", err)
	}
}

func TestForDirMax(t *testing.T) {
	// defer logex.CaptureLog(t).Release()

	dir := "$HOME/.local"
	if !exec.FileExists(dir) {
		dir = "$HOME/.config"
	}

	err := exec.ForDirMax(dir, 0, 2, func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
		if fi.IsDir() {
			t.Logf("  - dir: %v/[%v]", cwd, fi.Name())
		} else {
			t.Logf("  - file: %v/%v", cwd, fi.Name())
		}
		return
	})

	if err != nil {
		t.Errorf("wrong for ForDir(): %v", err)
	}
}

func TestGetExecutableDir(t *testing.T) {
	t.Logf("GetExecutablePath = %v", exec.GetExecutablePath())
	t.Logf("GetExecutableDir = %v", exec.GetExecutableDir())
	t.Logf("GetCurrentDir = %v", exec.GetCurrentDir())

	fn := path.Join(exec.GetCurrentDir(), "dir.go")
	if ok, err := exec.IsRegularFile(fn); err != nil || !ok {
		t.Fatal("expecting regular file detected.")
	}
	if !exec.FileExists(fn) {
		t.Fatal("expecting regular file existed.")
	}

	fileInfo, err := os.Stat(fn)
	if err != nil {
		t.Fatal(err)
	}

	_ = exec.FileModeIs(fn, exec.IsModeIrregular)
	_ = exec.FileModeIs(fn, exec.IsModeRegular)
	_ = exec.FileModeIs(fn, exec.IsModeDirectory)
	_ = exec.FileModeIs("/etc", exec.IsModeDirectory)
	_ = exec.FileModeIs("/etc", exec.IsModeIrregular)
	_ = exec.FileModeIs("/etc/not-existence", exec.IsModeIrregular)

	_ = exec.IsModeExecOwner(fileInfo.Mode())
	_ = exec.IsModeExecGroup(fileInfo.Mode())
	_ = exec.IsModeExecOther(fileInfo.Mode())
	_ = exec.IsModeExecAny(fileInfo.Mode())
	_ = exec.IsModeExecAll(fileInfo.Mode())

	_ = exec.IsModeWriteOwner(fileInfo.Mode())
	_ = exec.IsModeWriteGroup(fileInfo.Mode())
	_ = exec.IsModeWriteOther(fileInfo.Mode())
	_ = exec.IsModeWriteAny(fileInfo.Mode())
	_ = exec.IsModeWriteAll(fileInfo.Mode())

	_ = exec.IsModeReadOwner(fileInfo.Mode())
	_ = exec.IsModeReadGroup(fileInfo.Mode())
	_ = exec.IsModeReadOther(fileInfo.Mode())
	_ = exec.IsModeReadAny(fileInfo.Mode())
	_ = exec.IsModeReadAll(fileInfo.Mode())

	_ = exec.IsModeDirectory(fileInfo.Mode())
	_ = exec.IsModeSymbolicLink(fileInfo.Mode())
	_ = exec.IsModeDevice(fileInfo.Mode())
	_ = exec.IsModeNamedPipe(fileInfo.Mode())
	_ = exec.IsModeSocket(fileInfo.Mode())
	_ = exec.IsModeSetuid(fileInfo.Mode())
	_ = exec.IsModeSetgid(fileInfo.Mode())
	_ = exec.IsModeCharDevice(fileInfo.Mode())
	_ = exec.IsModeSticky(fileInfo.Mode())
	_ = exec.IsModeIrregular(fileInfo.Mode())
}

func TestEnsureDir(t *testing.T) {
	//

	if err := exec.EnsureDir(""); err == nil {
		t.Fatal("expecting an error.")
	}

	if err := exec.EnsureDirEnh(""); err == nil {
		t.Fatal("expecting an error.")
	}

	//

	dn := path.Join(exec.GetCurrentDir(), ".tmp.1")
	if err := exec.EnsureDir(dn); err != nil {
		t.Fatal(err)
	}
	if err := exec.RemoveDirRecursive(dn); err != nil {
		t.Fatal(err)
	}

	//

	dn = path.Join(exec.GetCurrentDir(), ".github")
	if err := exec.EnsureDir(dn); err != nil {
		t.Fatal(err)
	}
	if err := exec.EnsureDirEnh(dn); err != nil {
		t.Fatal(err)
	}

	dn = path.Join(exec.GetCurrentDir(), ".tmp1")
	if err := exec.EnsureDirEnh(dn); err != nil {
		t.Fatal(err)
	}
	if err := exec.RemoveDirRecursive(dn); err != nil {
		t.Fatal(err)
	}

	dn = path.Join(dn, ".tmp2")
	if err := exec.EnsureDirEnh(dn); err != nil {
		t.Fatal(err)
	}
	if err := exec.RemoveDirRecursive(dn); err != nil {
		t.Fatal(err)
	}
}

func TestNormalizeDir(t *testing.T) {
	exec.NormalizeDir("")
	exec.NormalizeDir(".")
	exec.NormalizeDir("./ad/./c")
	exec.NormalizeDir("./ad/../c")
	exec.NormalizeDir("/ad/./c")
	exec.NormalizeDir("../ad/./c")
	exec.NormalizeDir("~/ad/./c")
}
