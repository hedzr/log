/*
 * Copyright © 2019 Hedzr Yeh.
 */

package exec_test

import (
	"github.com/hedzr/log/exec"
	"os"
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

	err := exec.ForDir("$HOME/.local", func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
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

	err := exec.ForDirMax("$HOME/.local", 0, 2, func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
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
