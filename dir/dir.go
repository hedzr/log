// Copyright © 2020 Hedzr Yeh.

package dir

import (
	"errors"
	"fmt"
	"github.com/hedzr/log"
	"github.com/hedzr/log/exec"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// GetExecutableDir returns the executable file directory
func GetExecutableDir() string {
	// _ = ioutil.WriteFile("/tmp/11", []byte(strings.Join(os.Args,",")), 0644)
	// fmt.Printf("os.Args[0] = %v\n", os.Args[0])

	d, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Println(d)
	return d
}

// GetExecutablePath returns the executable file path
func GetExecutablePath() string {
	p, _ := filepath.Abs(os.Args[0])
	return p
}

// GetCurrentDir returns the current workingFlag directory
// it should be equal with os.Getenv("PWD")
func GetCurrentDir() string {
	d, _ := os.Getwd()
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Println(d)
	return d
}

// IsDirectory tests whether `path` is a directory or not
func IsDirectory(filepath string) (bool, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

// IsRegularFile tests whether `path` is a normal regular file or not
func IsRegularFile(filepath string) (bool, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	return fileInfo.Mode().IsRegular(), err
}

func timeSpecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

// FileModeIs tests the mode of 'filepath' with 'tester'. Examples:
//
//     var yes = exec.FileModeIs("/etc/passwd", exec.IsModeExecAny)
//     var yes = exec.FileModeIs("/etc/passwd", exec.IsModeDirectory)
//
func FileModeIs(filepath string, tester func(mode os.FileMode) bool) (ret bool) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return
	}
	ret = tester(fileInfo.Mode())
	return
}

// IsModeRegular give the result of whether a file is a regular file
func IsModeRegular(mode os.FileMode) bool { return mode.IsRegular() }

// IsModeDirectory give the result of whether a file is a directory
func IsModeDirectory(mode os.FileMode) bool { return mode&os.ModeDir != 0 }

// IsModeSymbolicLink give the result of whether a file is a symbolic link
func IsModeSymbolicLink(mode os.FileMode) bool { return mode&os.ModeSymlink != 0 }

// IsModeDevice give the result of whether a file is a device
func IsModeDevice(mode os.FileMode) bool { return mode&os.ModeDevice != 0 }

// IsModeNamedPipe give the result of whether a file is a named pipe
func IsModeNamedPipe(mode os.FileMode) bool { return mode&os.ModeNamedPipe != 0 }

// IsModeSocket give the result of whether a file is a socket file
func IsModeSocket(mode os.FileMode) bool { return mode&os.ModeSocket != 0 }

// IsModeSetuid give the result of whether a file has the setuid bit
func IsModeSetuid(mode os.FileMode) bool { return mode&os.ModeSetuid != 0 }

// IsModeSetgid give the result of whether a file has the setgid bit
func IsModeSetgid(mode os.FileMode) bool { return mode&os.ModeSetgid != 0 }

// IsModeCharDevice give the result of whether a file is a character device
func IsModeCharDevice(mode os.FileMode) bool { return mode&os.ModeCharDevice != 0 }

// IsModeSticky give the result of whether a file is a sticky file
func IsModeSticky(mode os.FileMode) bool { return mode&os.ModeSticky != 0 }

// IsModeIrregular give the result of whether a file is a non-regular file; nothing else is known about this file
func IsModeIrregular(mode os.FileMode) bool { return mode&os.ModeIrregular != 0 }

//

// IsModeExecOwner give the result of whether a file can be invoked by its unix-owner
func IsModeExecOwner(mode os.FileMode) bool { return mode&0100 != 0 }

// IsModeExecGroup give the result of whether a file can be invoked by its unix-group
func IsModeExecGroup(mode os.FileMode) bool { return mode&0010 != 0 }

// IsModeExecOther give the result of whether a file can be invoked by its unix-all
func IsModeExecOther(mode os.FileMode) bool { return mode&0001 != 0 }

// IsModeExecAny give the result of whether a file can be invoked by anyone
func IsModeExecAny(mode os.FileMode) bool { return mode&0111 != 0 }

// IsModeExecAll give the result of whether a file can be invoked by all users
func IsModeExecAll(mode os.FileMode) bool { return mode&0111 == 0111 }

//

// IsModeWriteOwner give the result of whether a file can be written by its unix-owner
func IsModeWriteOwner(mode os.FileMode) bool { return mode&0200 != 0 }

// IsModeWriteGroup give the result of whether a file can be written by its unix-group
func IsModeWriteGroup(mode os.FileMode) bool { return mode&0020 != 0 }

// IsModeWriteOther give the result of whether a file can be written by its unix-all
func IsModeWriteOther(mode os.FileMode) bool { return mode&0002 != 0 }

// IsModeWriteAny give the result of whether a file can be written by anyone
func IsModeWriteAny(mode os.FileMode) bool { return mode&0222 != 0 }

// IsModeWriteAll give the result of whether a file can be written by all users
func IsModeWriteAll(mode os.FileMode) bool { return mode&0222 == 0222 }

//

// IsModeReadOwner give the result of whether a file can be read by its unix-owner
func IsModeReadOwner(mode os.FileMode) bool { return mode&0400 != 0 }

// IsModeReadGroup give the result of whether a file can be read by its unix-group
func IsModeReadGroup(mode os.FileMode) bool { return mode&0040 != 0 }

// IsModeReadOther give the result of whether a file can be read by its unix-all
func IsModeReadOther(mode os.FileMode) bool { return mode&0004 != 0 }

// IsModeReadAny give the result of whether a file can be read by anyone
func IsModeReadAny(mode os.FileMode) bool { return mode&0444 != 0 }

// IsModeReadAll give the result of whether a file can be read by all users
func IsModeReadAll(mode os.FileMode) bool { return mode&0444 == 0444 }

//

// FileExists returns the existence of an directory or file
func FileExists(filepath string) bool {
	if _, err := os.Stat(os.ExpandEnv(filepath)); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// EnsureDir checks and creates the directory.
func EnsureDir(d string) (err error) {
	if len(d) == 0 {
		return errors.New("empty directory")
	}
	if !FileExists(d) {
		err = os.MkdirAll(d, 0755)
	}
	return
}

// EnsureDirEnh checks and creates the directory, via sudo if necessary.
func EnsureDirEnh(d string) (err error) {
	if len(d) == 0 {
		return errors.New("empty directory")
	}
	if !FileExists(d) {
		err = os.MkdirAll(d, 0755)
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EACCES {
			var u *user.User
			u, err = user.Current()
			if _, _, err = exec.Sudo("mkdir", "-p", d); err == nil {
				_, _, err = exec.Sudo("chown", u.Username+":", d)
			}

			//if _, _, err = exec.Sudo("mkdir", "-p", d); err != nil {
			//	logrus.Warnf("Failed to create directory %q, using default stderr. error is: %v", d, err)
			//} else if _, _, err = exec.Sudo("chown", u.Username+":", d); err != nil {
			//	logrus.Warnf("Failed to create directory %q, using default stderr. error is: %v", d, err)
			//}
		}
	}
	return
}

// RemoveDirRecursive removes a directory and any children it contains.
func RemoveDirRecursive(d string) (err error) {
	// RemoveContentsInDir(d)
	err = os.RemoveAll(d)
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

func FollowSymLink(pathname string) string {
	if t, err := filepath.EvalSymlinks(pathname); err == nil {
		return t
	}
	return pathname
}

// NormalizePath cleans up the given pathname
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
func ForDir(
	root string,
	cb func(depth int, dirname string, fi os.FileInfo) (stop bool, err error),
	excludes ...string,
) (err error) {
	err = ForDirMax(root, 0, -1, cb, excludes...)
	return
}

// ForDirMax walks on `root` directory and its children with nested levels up to `maxLength`.
//
// Example - discover folder just one level
//
//      _ = ForDirMax(dir, 0, 1, func(depth int, dirname string, fi os.FileInfo) (stop bool, err error) {
//			if fi.IsDir() {
//				return
//			}
//          // ... doing something for a file,
//			return
//		})
//
// maxDepth = -1: no limit.
// initialDepth: 0 if no idea.
func ForDirMax(
	root string,
	initialDepth int,
	maxDepth int,
	cb func(depth int, dirname string, fi os.FileInfo) (stop bool, err error),
	excludes ...string,
) (err error) {
	if maxDepth > 0 && initialDepth >= maxDepth {
		return
	}

	//rootDir := os.ExpandEnv(root)
	rootDir := path.Clean(NormalizeDir(root))

	return forDirMax(rootDir, initialDepth, maxDepth, cb, excludes...)
}

func forDirMax(
	rootDir string,
	initialDepth int,
	maxDepth int,
	cb func(depth int, dirname string, fi os.FileInfo) (stop bool, err error),
	excludes ...string,
) (err error) {
	var dirs []os.FileInfo
	dirs, err = ioutil.ReadDir(rootDir)
	if err != nil {
		// Logger.Fatalf("error in ForDirMax(): %v", err)
		return
	}

	var stop bool

	//files, err :=os.ReadDir(rootDir)
	var fi os.FileInfo
	fi, err = os.Stat(rootDir)
	if err != nil {
		return
	}
	if stop, err = cb(initialDepth, rootDir, fi); stop {
		return
	}

	stop, err = forDirMaxLoops(dirs, rootDir, initialDepth, maxDepth, cb, excludes...)
	return
}

func forDirMaxLoops(
	dirs []os.FileInfo,
	rootDir string,
	initialDepth int,
	maxDepth int,
	cb func(depth int, dirname string, fi os.FileInfo) (stop bool, err error),
	excludes ...string,
) (stop bool, err error) {

	for _, f := range dirs {
		//Logger.Printf("  - %v", f.Name())
		if err != nil {
			log.NewStdLogger().Errorf("error in ForDirMax().cb: %v", err)
			continue
		}

		if f.IsDir() && (maxDepth <= 0 || (maxDepth > 0 && initialDepth+1 < maxDepth)) {
			d := path.Join(rootDir, f.Name())
			if forFileMatched(d, excludes...) {
				continue
			}

			if stop, err = cb(initialDepth, d, f); stop {
				return
			}
			if err = ForDirMax(d, initialDepth+1, maxDepth, cb); err != nil {
				log.NewStdLogger().Errorf("error in ForDirMax(): %v", err)
			}
		}
	}

	return
}

// ForFile walks on `root` directory and its children
func ForFile(
	root string,
	cb func(depth int, dirname string, fi os.FileInfo) (stop bool, err error),
	excludes ...string,
) (err error) {
	err = ForFileMax(root, 0, -1, cb, excludes...)
	return
}

// ForFileMax walks on `root` directory and its children with nested levels up to `maxLength`.
//
// Example - discover folder just one level
//
//      _ = ForFileMax(dir, 0, 1, func(depth int, dirname string, fi os.FileInfo) (stop bool, err error) {
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
// Known issue:
// can't walk at ~/.local/share/NuGet/v3-cache/1ca707a4d90792ce8e42453d4e350886a0fdaa4d:_api.nuget.org_v3_index.json.
// workaround: use filepath.Walk
//
func ForFileMax(
	root string,
	initialDepth, maxDepth int,
	cb func(depth int, dirname string, fi os.FileInfo) (stop bool, err error),
	excludes ...string,
) (err error) {
	if maxDepth > 0 && initialDepth >= maxDepth {
		return
	}

	//rootDir := os.ExpandEnv(root)
	rootDir := path.Clean(NormalizeDir(root))

	return forFileMax(rootDir, initialDepth, maxDepth, cb, excludes...)
}

func forFileMax(
	rootDir string,
	initialDepth, maxDepth int,
	cb func(depth int, dirname string, fi os.FileInfo) (stop bool, err error),
	excludes ...string,
) (err error) {
	var dirs []os.FileInfo
	dirs, err = ioutil.ReadDir(rootDir)
	//var dirs []os.DirEntry
	//dirs, err = os.ReadDir(rootDir)
	if err != nil {
		// Logger.Fatalf("error in ForFileMax(): %v", err)
		return
	}

	var stop bool
	for _, f := range dirs {
		//Logger.Printf("  - %v", f.Name())
		if err != nil {
			log.NewStdLogger().Errorf("error in ForFileMax().cb: %v", err)
			continue
		}

		if f.IsDir() && (maxDepth <= 0 || (maxDepth > 0 && initialDepth < maxDepth)) {
			d := path.Join(rootDir, f.Name())
			if forFileMatched(d, excludes...) {
				continue
			}

			if err = ForFileMax(d, initialDepth+1, maxDepth, cb, excludes...); err != nil {
				log.NewStdLogger().Errorf("error in ForFileMax(): %v", err)
			}

			continue
		}

		if !f.IsDir() {
			d := path.Join(rootDir, f.Name())
			if forFileMatched(d, excludes...) {
				continue
			}

			// log.Infof(" - %s", f.Name())
			//fi, _ := f.Info()
			if stop, err = cb(initialDepth, rootDir, f); stop {
				return
			}
		}
	}

	return
}

//func forDirMatched(f os.DirEntry, root string, excludes ...string) (matched bool) {
//	fullName := path.Join(root, f.Name())
//	for _, ptn := range excludes {
//		if IsWildMatch(fullName, ptn) {
//			matched = true
//			break
//		}
//	}
//	return
//}

//func forFileMatched(f os.FileInfo, root string, excludes ...string) (matched bool) {
//	fullName := path.Join(root, f.Name())
//	matched = inExcludes(fullName, excludes...)
//	//if matched, _ = filepath.Match(ptn, fullName); matched {
//	//	break
//	//}
//	return
//}

func forFileMatched(name string, excludes ...string) (yes bool) {
	for _, ptn := range excludes {
		if yes = IsWildMatch(name, ptn); yes {
			break
		}
	}
	return
}

// PushDir provides a shortcut to enter a folder and restore at
// the end of your current function scope.
// PushDir returns a functor and assumes you will DEFER call it.
//
// For example:
//
//     func TestSth() {
//         defer dir.PushDir("/your/working/dir")()
//         // do sth under '/your/working/dir' ...
//     }
//
// BEWARE DON'T miss the ending brakets for defer call.
func PushDir(dirname string) (closer func()) {
	savedDir := GetCurrentDir()
	var err error
	if err = os.Chdir(dirname); err != nil {
		//err = nil //ignore path err
		return func() {}
	}
	return func() {
		if err == nil {
			_ = os.Chdir(savedDir)
		}
	}
}

// DeleteFile deletes a file if exists
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
func CopyFileByLinkFirst(src, dst string) (err error) {
	return copyFileByLinkFirst(src, dst, true)
}

// CopyFile will make a content clone of src.
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
