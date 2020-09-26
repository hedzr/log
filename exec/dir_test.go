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
	//t.Logf("InTesting: %v", cmdr.InTesting())
	//t.Logf("InDebugging: %v", cmdr.InDebugging())

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

	err := exec.ForDir("$HOME/.config", func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
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

	err := exec.ForDirMax("$HOME/.config", 0, 2, func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
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

	dn := path.Join(exec.GetCurrentDir(), ".github")
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
