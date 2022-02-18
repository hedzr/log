package log

import (
	"runtime"
	"strings"
	"sync"
)

// CalcStackFrames _
func CalcStackFrames(skipFramesAtFirst int) (skipped int) {

	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, 2)
		_ = runtime.Callers(0, pcs)
		logPackage = "github.com/hedzr/log"
		logrusPackage = "github.com/sirupsen/logrus"
		knownPackages = append(knownPackages, "runtime", "reflect", logPackage, logrusPackage)
		knownPathes = []string{"/usr/local/go/src/"}
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(skipFramesAtFirst+2 /*minimumCallerDepth*/ /*+skipFrames*/, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	skipped = 0
	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)
		if !contains(knownPackages, pkg) && !containsPartialsOnly(knownPathes, f.File) {
			break
		}
		skipped++
	}

	return
}

var (
	knownPackages             []string
	knownPathes               []string
	logPackage, logrusPackage string
	callerInitOnce            sync.Once
)

const (
	maximumCallerDepth int = 25
)

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func contains(names []string, name string) bool {
	for _, n := range names {
		if strings.EqualFold(n, name) {
			return true
		}
	}
	return false
}

func containsPartialsOnly(partialNames []string, testedString string) (contains bool) {
	for _, n := range partialNames {
		if strings.Contains(testedString, n) {
			return true
		}
	}
	return
}
