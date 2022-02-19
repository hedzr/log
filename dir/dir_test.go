/*
 * Copyright © 2019 Hedzr Yeh.
 */

package dir_test

import (
	"github.com/hedzr/log/dir"
	"gopkg.in/hedzr/errors.v3"
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

	dir.NormalizeDir("")

	if yes, err := dir.IsDirectory("./conf.d1"); yes {
		t.Fatal(err)
	}
	if yes, err := dir.IsDirectory("../dir"); !yes {
		t.Fatal(err)
	}
	if yes, err := dir.IsRegularFile("./doc1.golang"); yes {
		t.Fatal(err)
	}
	if yes, err := dir.IsRegularFile("./dir.go"); !yes {
		t.Fatal(err)
	}
}

func TestForDir(t *testing.T) {
	// defer logex.CaptureLog(t).Release()

	dirName := "$HOME/.local"
	if !dir.FileExists(dirName) {
		dirName = "$HOME/.config"
	}

	err := dir.ForDir(dirName, func(depth int, dirname string, fi os.FileInfo) (stop bool, err error) {
		if fi.IsDir() {
			t.Logf("  - dir: %v - [%v]", dirname, fi.Name())
		} else {
			t.Logf("  - file: %v - %v", dirname, fi.Name())
		}
		return
	})

	if err != nil && !errors.TypeIs(err, &os.PathError{}) {
		t.Errorf("wrong for ForDir(): %v", err)
	}
}

func TestForDirMax(t *testing.T) {
	// defer logex.CaptureLog(t).Release()

	dirName := "$HOME/.local"
	if !dir.FileExists(dirName) {
		dirName = "$HOME/.config"
	}

	err := dir.ForDirMax(dirName, 0, 2, func(depth int, dirname string, fi os.FileInfo) (stop bool, err error) {
		if fi.IsDir() {
			t.Logf("  - dir: %v - [%v]", dirname, fi.Name())
		} else {
			t.Logf("  - file: %v - %v", dirname, fi.Name())
		}
		return
	})

	if err != nil {
		t.Errorf("wrong for ForDir(): %v", err)
	}
}

//func TestWalk(t *testing.T) {
//	dirName := "$HOME/.local"
//	if !dir.FileExists(dirName) {
//		dirName = "$HOME/.config"
//	}
//	err := filepath.Walk(os.ExpandEnv(dirName), func(path string, info fs.FileInfo, err error) error {
//		if info == nil {
//			t.Logf("  - file: %v - ERR: %v", path, err)
//		} else {
//			t.Logf("  - file: %v - %v - %v", path, info.Name(), err)
//		}
//		return err
//	})
//	if err != nil {
//		t.Error(err)
//	}
//}

func TestForFileMax(t *testing.T) {
	// defer logex.CaptureLog(t).Release()

	//dirName := "$HOME/.local"
	//if !dir.FileExists(dirName) {
	//	dirName = "$HOME/.config"
	//}
	dirName := "$HOME/.config"
	if !dir.FileExists(dirName) {
		dirName = "$HOME"
	}

	err := dir.ForFileMax(dirName, 0, 6, func(depth int, dirname string, fi os.FileInfo) (stop bool, err error) {
		if fi.IsDir() {
			t.Logf("  - dir: %v - [%v]", dirname, fi.Name())
		} else {
			t.Logf("  - file: %v - %v", dirname, fi.Name())
		}
		return
	}, "*/1.x/*", "*/1.x", "*/2.c", "*/node_modules", "*/.git", "*/usr*", "*/share")

	if err != nil {
		t.Errorf("wrong for ForDir(): %v", err)
	}
}

func TestGetExecutableDir(t *testing.T) {
	t.Logf("GetExecutablePath = %v", dir.GetExecutablePath())
	t.Logf("GetExecutableDir = %v", dir.GetExecutableDir())
	t.Logf("GetCurrentDir = %v", dir.GetCurrentDir())

	fn := path.Join(dir.GetCurrentDir(), "dir.go")
	if ok, err := dir.IsRegularFile(fn); err != nil || !ok {
		t.Fatal("expecting regular file detected.")
	}
	if !dir.FileExists(fn) {
		t.Fatal("expecting regular file existed.")
	}

	fileInfo, err := os.Stat(fn)
	if err != nil {
		t.Fatal(err)
	}

	dir.FileModeIs(fn, dir.IsModeIrregular)
	dir.FileModeIs(fn, dir.IsModeRegular)
	dir.FileModeIs(fn, dir.IsModeDirectory)
	dir.FileModeIs("/etc", dir.IsModeDirectory)
	dir.FileModeIs("/etc", dir.IsModeIrregular)
	dir.FileModeIs("/etc/not-existence", dir.IsModeIrregular)

	dir.IsModeExecOwner(fileInfo.Mode())
	dir.IsModeExecGroup(fileInfo.Mode())
	dir.IsModeExecOther(fileInfo.Mode())
	dir.IsModeExecAny(fileInfo.Mode())
	dir.IsModeExecAll(fileInfo.Mode())

	dir.IsModeWriteOwner(fileInfo.Mode())
	dir.IsModeWriteGroup(fileInfo.Mode())
	dir.IsModeWriteOther(fileInfo.Mode())
	dir.IsModeWriteAny(fileInfo.Mode())
	dir.IsModeWriteAll(fileInfo.Mode())

	dir.IsModeReadOwner(fileInfo.Mode())
	dir.IsModeReadGroup(fileInfo.Mode())
	dir.IsModeReadOther(fileInfo.Mode())
	dir.IsModeReadAny(fileInfo.Mode())
	dir.IsModeReadAll(fileInfo.Mode())

	dir.IsModeDirectory(fileInfo.Mode())
	dir.IsModeSymbolicLink(fileInfo.Mode())
	dir.IsModeDevice(fileInfo.Mode())
	dir.IsModeNamedPipe(fileInfo.Mode())
	dir.IsModeSocket(fileInfo.Mode())
	dir.IsModeSetuid(fileInfo.Mode())
	dir.IsModeSetgid(fileInfo.Mode())
	dir.IsModeCharDevice(fileInfo.Mode())
	dir.IsModeSticky(fileInfo.Mode())
	dir.IsModeIrregular(fileInfo.Mode())
}

func TestEnsureDir(t *testing.T) {
	//

	if err := dir.EnsureDir(""); err == nil {
		t.Fatal("expecting an error.")
	}

	if err := dir.EnsureDirEnh(""); err == nil {
		t.Fatal("expecting an error.")
	}

	//

	dn := path.Join(dir.GetCurrentDir(), ".tmp.1")
	if err := dir.EnsureDir(dn); err != nil {
		t.Fatal(err)
	}
	if err := dir.RemoveDirRecursive(dn); err != nil {
		t.Fatal(err)
	}

	//

	dn = path.Join(dir.GetCurrentDir(), ".github")
	if err := dir.EnsureDir(dn); err != nil {
		t.Fatal(err)
	}
	if err := dir.EnsureDirEnh(dn); err != nil {
		t.Fatal(err)
	}

	dn = path.Join(dir.GetCurrentDir(), ".tmp1")
	if err := dir.EnsureDirEnh(dn); err != nil {
		t.Fatal(err)
	}
	if err := dir.RemoveDirRecursive(dn); err != nil {
		t.Fatal(err)
	}

	dn = path.Join(dn, ".tmp2")
	if err := dir.EnsureDirEnh(dn); err != nil {
		t.Fatal(err)
	}
	if err := dir.RemoveDirRecursive(dn); err != nil {
		t.Fatal(err)
	}
}

func TestNormalizeDir(t *testing.T) {
	dir.NormalizeDir("")
	dir.NormalizeDir(".")
	dir.NormalizeDir("./ad/./c")
	dir.NormalizeDir("./ad/../c")
	dir.NormalizeDir("/ad/./c")
	dir.NormalizeDir("../ad/./c")
	dir.NormalizeDir("~/ad/./c")
}

func TestDirTimestamps(t *testing.T) {
	fileInfo, err := os.Stat("/tmp")
	if err != nil {
		return
	}
	t.Logf("create time: %v", dir.FileCreatedTime(fileInfo))
	t.Logf("access time: %v", dir.FileAccessedTime(fileInfo))
	t.Logf("modified time: %v", dir.FileModifiedTime(fileInfo))
}
