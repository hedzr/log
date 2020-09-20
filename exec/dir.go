// Copyright © 2020 Hedzr Yeh.

package exec

import (
	"errors"
	"github.com/hedzr/log"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"syscall"
)

// GetExecutableDir returns the executable file directory
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
func GetExecutablePath() string {
	p, _ := filepath.Abs(os.Args[0])
	return p
}

// GetCurrentDir returns the current workingFlag directory
// it should be equal with os.Getenv("PWD")
func GetCurrentDir() string {
	dir, _ := os.Getwd()
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Println(dir)
	return dir
}

// IsDirectory tests whether `path` is a directory or not
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

// IsRegularFile tests whether `path` is a normal regular file or not
func IsRegularFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.Mode().IsRegular(), err
}

// FileExists returns the existence of an directory or file
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// EnsureDir checks and creates the directory.
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
func NormalizeDir(s string) string {
	return normalizeDir(s)
}

func normalizeDir(s string) string {
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

// ForDir walks on `root` directory and its children
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
func ForDirMax(root string, initialDepth, maxDepth int, cb func(depth int, cwd string, fi os.FileInfo) (stop bool, err error)) (err error) {
	if maxDepth > 0 && initialDepth >= maxDepth {
		return
	}

	var files []os.FileInfo
	files, err = ioutil.ReadDir(os.ExpandEnv(root))
	if err != nil {
		// Logger.Fatalf("error in ForDirMax(): %v", err)
		return
	}

	var stop bool
	for _, f := range files {
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

// IsExecOwner give the result of whether a file can be invoked by its unix-owner
func IsExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}

// IsExecOwner give the result of whether a file can be invoked by its unix-group
func IsExecGroup(mode os.FileMode) bool {
	return mode&0010 != 0
}

// IsExecOwner give the result of whether a file can be invoked by its unix-all
func IsExecOther(mode os.FileMode) bool {
	return mode&0001 != 0
}

// IsExecOwner give the result of whether a file can be invoked by anyone
func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

// IsExecOwner give the result of whether a file can be invoked by all users
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}
