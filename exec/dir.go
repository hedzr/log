// Copyright © 2020 Hedzr Yeh.

package exec

import (
	"errors"
	"fmt"
	"github.com/hedzr/log"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"syscall"
)

// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// GetExecutableDir returns the executable file directory
// Deprecated see also dir.GetExecutableDir
func GetExecutableDir() string {
	// _ = ioutil.WriteFile("/tmp/11", []byte(strings.Join(os.Args,",")), 0644)
	// fmt.Printf("os.Args[0] = %v\n", os.Args[0])

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Println(dir)
	return dir
}

// GetExecutablePath returns the executable file path
// Deprecated see also dir.GetExecutablePath
func GetExecutablePath() string {
	p, _ := filepath.Abs(os.Args[0])
	return p
}

// GetCurrentDir returns the current workingFlag directory
// it should be equal with os.Getenv("PWD")
// Deprecated see also dir.GetCurrentDir
func GetCurrentDir() string {
	dir, _ := os.Getwd()
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Println(dir)
	return dir
}

// IsDirectory tests whether `path` is a directory or not
// Deprecated see also dir.IsDirectory
func IsDirectory(filepath string) (bool, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

// IsRegularFile tests whether `path` is a normal regular file or not
// Deprecated see also dir.IsRegularFile
func IsRegularFile(filepath string) (bool, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	return fileInfo.Mode().IsRegular(), err
}

// FileModeIs tests the mode of 'filepath' with 'tester'. Examples:
//
//     var yes = exec.FileModeIs("/etc/passwd", exec.IsModeExecAny)
//     var yes = exec.FileModeIs("/etc/passwd", exec.IsModeDirectory)
//
// Deprecated see also dir.FileModeIs
func FileModeIs(filepath string, tester func(mode os.FileMode) bool) (ret bool) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return
	}
	ret = tester(fileInfo.Mode())
	return
}

// IsModeRegular give the result of whether a file is a regular file
// Deprecated see also dir.IsModeRegular
func IsModeRegular(mode os.FileMode) bool { return mode.IsRegular() }

// IsModeDirectory give the result of whether a file is a directory
// Deprecated see also dir.IsModeDirectory
func IsModeDirectory(mode os.FileMode) bool { return mode&os.ModeDir != 0 }

// IsModeSymbolicLink give the result of whether a file is a symbolic link
// Deprecated see also dir.IsModeSymbolicLink
func IsModeSymbolicLink(mode os.FileMode) bool { return mode&os.ModeSymlink != 0 }

// IsModeDevice give the result of whether a file is a device
// Deprecated see also dir.IsModeDevice
func IsModeDevice(mode os.FileMode) bool { return mode&os.ModeDevice != 0 }

// IsModeNamedPipe give the result of whether a file is a named pipe
// Deprecated see also dir.IsModePipe
func IsModeNamedPipe(mode os.FileMode) bool { return mode&os.ModeNamedPipe != 0 }

// IsModeSocket give the result of whether a file is a socket file
// Deprecated see also dir.IsModeSocket
func IsModeSocket(mode os.FileMode) bool { return mode&os.ModeSocket != 0 }

// IsModeSetuid give the result of whether a file has the setuid bit
// Deprecated see also dir.IsModeSetuid
func IsModeSetuid(mode os.FileMode) bool { return mode&os.ModeSetuid != 0 }

// IsModeSetgid give the result of whether a file has the setgid bit
// Deprecated see also dir.IsModeSetgid
func IsModeSetgid(mode os.FileMode) bool { return mode&os.ModeSetgid != 0 }

// IsModeCharDevice give the result of whether a file is a character device
// Deprecated see also dir.IsModeCharDevice
func IsModeCharDevice(mode os.FileMode) bool { return mode&os.ModeCharDevice != 0 }

// IsModeSticky give the result of whether a file is a sticky file
// Deprecated see also dir.IsModeSticky
func IsModeSticky(mode os.FileMode) bool { return mode&os.ModeSticky != 0 }

// IsModeIrregular give the result of whether a file is a non-regular file; nothing else is known about this file
// Deprecated see also dir.IsModeIrregular
func IsModeIrregular(mode os.FileMode) bool { return mode&os.ModeIrregular != 0 }

//

// IsModeExecOwner give the result of whether a file can be invoked by its unix-owner
// Deprecated see also dir.IsModeExecOwner
func IsModeExecOwner(mode os.FileMode) bool { return mode&0100 != 0 }

// IsModeExecGroup give the result of whether a file can be invoked by its unix-group
// Deprecated see also dir.IsModeExecGroup
func IsModeExecGroup(mode os.FileMode) bool { return mode&0010 != 0 }

// IsModeExecOther give the result of whether a file can be invoked by its unix-all
// Deprecated see also dir.IsModeExecOther
func IsModeExecOther(mode os.FileMode) bool { return mode&0001 != 0 }

// IsModeExecAny give the result of whether a file can be invoked by anyone
// Deprecated see also dir.IsModeExecAny
func IsModeExecAny(mode os.FileMode) bool { return mode&0111 != 0 }

// IsModeExecAll give the result of whether a file can be invoked by all users
// Deprecated see also dir.IsModeExecAll
func IsModeExecAll(mode os.FileMode) bool { return mode&0111 == 0111 }

//

// IsModeWriteOwner give the result of whether a file can be written by its unix-owner
// Deprecated see also dir.IsModeWriteOwner
func IsModeWriteOwner(mode os.FileMode) bool { return mode&0200 != 0 }

// IsModeWriteGroup give the result of whether a file can be written by its unix-group
// Deprecated see also dir.IsModeWriteGroup
func IsModeWriteGroup(mode os.FileMode) bool { return mode&0020 != 0 }

// IsModeWriteOther give the result of whether a file can be written by its unix-all
// Deprecated see also dir.IsModeWriteOther
func IsModeWriteOther(mode os.FileMode) bool { return mode&0002 != 0 }

// IsModeWriteAny give the result of whether a file can be written by anyone
// Deprecated see also dir.IsModeWriteAny
func IsModeWriteAny(mode os.FileMode) bool { return mode&0222 != 0 }

// IsModeWriteAll give the result of whether a file can be written by all users
// Deprecated see also dir.IsModeWriteAll
func IsModeWriteAll(mode os.FileMode) bool { return mode&0222 == 0222 }

//

// IsModeReadOwner give the result of whether a file can be read by its unix-owner
// Deprecated see also dir.IsModeReadOwner
func IsModeReadOwner(mode os.FileMode) bool { return mode&0400 != 0 }

// IsModeReadGroup give the result of whether a file can be read by its unix-group
// Deprecated see also dir.IsModeReadGroup
func IsModeReadGroup(mode os.FileMode) bool { return mode&0040 != 0 }

// IsModeReadOther give the result of whether a file can be read by its unix-all
// Deprecated see also dir.IsModeReadOther
func IsModeReadOther(mode os.FileMode) bool { return mode&0004 != 0 }

// IsModeReadAny give the result of whether a file can be read by anyone
// Deprecated see also dir.IsModeReadAny
func IsModeReadAny(mode os.FileMode) bool { return mode&0444 != 0 }

// IsModeReadAll give the result of whether a file can be read by all users
// Deprecated see also dir.IsModeReadAll
func IsModeReadAll(mode os.FileMode) bool { return mode&0444 == 0444 }

//

// FileExists returns the existence of an directory or file
// Deprecated see also dir.FileExist
func FileExists(filepath string) bool {
	if _, err := os.Stat(os.ExpandEnv(filepath)); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// EnsureDir checks and creates the directory.
// Deprecated see also dir.EnsureDir
func EnsureDir(dir string) (err error) {
	if len(dir) == 0 {
		return errors.New("empty directory")
	}
	if !FileExists(dir) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}

// EnsureDirEnh checks and creates the directory, via sudo if necessary.
// Deprecated see also dir.EnsureDirEnh
func EnsureDirEnh(dir string) (err error) {
	if len(dir) == 0 {
		return errors.New("empty directory")
	}
	if !FileExists(dir) {
		err = os.MkdirAll(dir, 0755)
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EACCES {
			var u *user.User
			u, err = user.Current()
			if _, _, err = Sudo("mkdir", "-p", dir); err == nil {
				_, _, err = Sudo("chown", u.Username+":", dir)
			}

			//if _, _, err = exec.Sudo("mkdir", "-p", dir); err != nil {
			//	logrus.Warnf("Failed to create directory %q, using default stderr. error is: %v", dir, err)
			//} else if _, _, err = exec.Sudo("chown", u.Username+":", dir); err != nil {
			//	logrus.Warnf("Failed to create directory %q, using default stderr. error is: %v", dir, err)
			//}
		}
	}
	return
}

// RemoveDirRecursive removes a directory and any children it contains.
// Deprecated see also dir.RemoveDirRecursive
func RemoveDirRecursive(dir string) (err error) {
	// RemoveContentsInDir(dir)
	err = os.RemoveAll(dir)
	return
}

// // RemoveContentsInDir removes all file and sub-directory in a directory
// func RemoveContentsInDir(dir string) error {
// 	d, err := os.Open(dir)
// 	if err != nil {
// 		return err
// 	}
// 	defer d.Close()
// 	names, err := d.Readdirnames(-1)
// 	if err != nil {
// 		return err
// 	}
// 	for _, name := range names {
// 		err = os.RemoveAll(filepath.Join(dir, name))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// NormalizeDir make dir name normalized
// Deprecated see also dir.NormalizeDir
func NormalizeDir(s string) string {
	return normalizeDir(s)
}

func normalizeDir(s string) string {
	p := normalizeDirBasic(s)
	p = filepath.Clean(p)
	return p
}

func normalizeDirBasic(s string) string {
	if len(s) == 0 {
		return s
	}

	s = os.Expand(s, os.Getenv)
	if s[0] == '/' {
		return s
	} else if strings.HasPrefix(s, "./") {
		return path.Join(GetCurrentDir(), s)
	} else if strings.HasPrefix(s, "../") {
		return path.Dir(path.Join(GetCurrentDir(), s))
	} else if strings.HasPrefix(s, "~/") {
		return path.Join(os.Getenv("HOME"), s[2:])
	} else {
		return s
	}
}

// AbsPath returns a clean, normalized and absolute path string for the given pathname.
// Deprecated see also dir.AbsPath
func AbsPath(pathname string) string {
	return absPath(pathname)
}

func absPath(pathname string) (abs string) {
	abs = normalizePath(pathname)
	if s, err := filepath.Abs(abs); err == nil {
		abs = s
	}
	return
}

// NormalizePath cleans up the given pathname
// Deprecated see also dir.NormalizePath
func NormalizePath(pathname string) string {
	return normalizePath(pathname)
}

func normalizePath(pathname string) string {
	p := normalizePathBasic(pathname)
	p = filepath.Clean(p)
	return p
}

func normalizePathBasic(pathname string) string {
	if len(pathname) == 0 {
		return pathname
	}

	pathname = os.Expand(pathname, os.Getenv)
	if pathname[0] == '/' {
		return pathname
	} else if strings.HasPrefix(pathname, "~/") {
		return path.Join(os.Getenv("HOME"), pathname[2:])
	} else {
		return pathname
	}
}

// ForDir walks on `root` directory and its children
// Deprecated see also dir.ForDir
func ForDir(root string, cb func(depth int, cwd string, fi os.FileInfo) (stop bool, err error)) (err error) {
	err = ForDirMax(root, 0, -1, cb)
	return
}

// ForDirMax walks on `root` directory and its children with nested levels up to `maxLength`.
//
// Example - discover folder just one level
//
//      _ = ForDirMax(dir, 0, 1, func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
//			if fi.IsDir() {
//				return
//			}
//          // ... doing something for a file,
//			return
//		})
//
// maxDepth = -1: no limit.
// initialDepth: 0 if no idea.
// Deprecated see also dir.ForDirMax
func ForDirMax(root string, initialDepth, maxDepth int, cb func(depth int, cwd string, fi os.FileInfo) (stop bool, err error)) (err error) {
	if maxDepth > 0 && initialDepth >= maxDepth {
		return
	}

	var dirs []os.FileInfo
	dirs, err = ioutil.ReadDir(os.ExpandEnv(root))
	if err != nil {
		// Logger.Fatalf("error in ForDirMax(): %v", err)
		return
	}

	var stop bool
	for _, f := range dirs {
		//Logger.Printf("  - %v", f.Name())
		if stop, err = cb(initialDepth, root, f); stop {
			return
		}
		if err != nil {
			log.NewStdLogger().Errorf("error in ForDirMax().cb: %v", err)
		} else if f.IsDir() && (maxDepth <= 0 || (maxDepth > 0 && initialDepth+1 < maxDepth)) {
			dir := path.Join(root, f.Name())
			if err = ForDirMax(dir, initialDepth+1, maxDepth, cb); err != nil {
				log.NewStdLogger().Errorf("error in ForDirMax(): %v", err)
			}
		}
	}

	return
}

// ForFile walks on `root` directory and its children
// Deprecated see also dir.ForFile
func ForFile(root string, cb func(depth int, cwd string, fi os.FileInfo) (stop bool, err error)) (err error) {
	err = ForFileMax(root, 0, -1, cb)
	return
}

// ForFileMax walks on `root` directory and its children with nested levels up to `maxLength`.
//
// Example - discover folder just one level
//
//      _ = ForFileMax(dir, 0, 1, func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
//			if fi.IsDir() {
//				return
//			}
//          // ... doing something for a file,
//			return
//		})
//
// maxDepth = -1: no limit.
// initialDepth: 0 if no idea.
//
// Deprecated see also dir.ForFileMax
func ForFileMax(root string, initialDepth, maxDepth int, cb func(depth int, cwd string, fi os.FileInfo) (stop bool, err error)) (err error) {
	if maxDepth > 0 && initialDepth >= maxDepth {
		return
	}

	var dirs []os.FileInfo
	dirs, err = ioutil.ReadDir(os.ExpandEnv(root))
	if err != nil {
		// Logger.Fatalf("error in ForFileMax(): %v", err)
		return
	}

	var stop bool
	for _, f := range dirs {
		//Logger.Printf("  - %v", f.Name())
		if err != nil {
			log.NewStdLogger().Errorf("error in ForFileMax().cb: %v", err)
		} else if f.IsDir() && (maxDepth <= 0 || (maxDepth > 0 && initialDepth+1 < maxDepth)) {
			dir := path.Join(root, f.Name())
			if err = ForFileMax(dir, initialDepth+1, maxDepth, cb); err != nil {
				log.NewStdLogger().Errorf("error in ForFileMax(): %v", err)
			}
		} else if !f.IsDir() {
			// log.Infof(" - %s", f.Name())
			if stop, err = cb(initialDepth, root, f); stop {
				return
			}
		}
	}

	return
}

// DeleteFile deletes a file if exists
// Deprecated see also dir.DeleteFile
func DeleteFile(dst string) (err error) {
	dst = os.ExpandEnv(dst)
	if FileExists(dst) {
		err = os.Remove(dst)
	}
	return
}

// CopyFileByLinkFirst copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherwise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
// Deprecated see also dir.CopyFileByLinkFirst
func CopyFileByLinkFirst(src, dst string) (err error) {
	return copyFileByLinkFirst(src, dst, true)
}

// CopyFile will make a content clone of src.
// Deprecated see also dir.CopyFile
func CopyFile(src, dst string) (err error) {
	return copyFileByLinkFirst(src, dst, false)
}

func copyFileByLinkFirst(src, dst string, linkAtFirst bool) (err error) {
	src = os.ExpandEnv(src)
	dst = os.ExpandEnv(dst)
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if linkAtFirst {
		if err = os.Link(src, dst); err == nil {
			return
		}
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
