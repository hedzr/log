// Copyright © 2023 Hedzr Yeh.

package detects

import (
	"os"
	"strings"

	"github.com/hedzr/log/states"
)

// InDebugging return the status if cmdr was built with debug mode / or the app running under a debugger attached.
//
// To enable the debugger attached mode for cmdr, run `go build` with `-tags=delve` options. eg:
//
//	go run -tags=delve ./cli
//	go build -tags=delve -o my-app ./cli
//
// For Goland, you can enable this under 'Run/Debug Configurations', by adding the following into 'Go tool arguments:'
//
//	-tags=delve
//
// InDebugging() is a synonym to IsDebuggerAttached().
//
// NOTE that `isdelve` algor is from https://stackoverflow.com/questions/47879070/how-can-i-see-if-the-goland-debugger-is-running-in-the-program
//
// noinspection GoBoolExpressions
func InDebugging() bool {
	return states.Env().InDebugging() // isdelve.Enabled
}

// IsDebuggerAttached return the status if cmdr was built with debug mode / or the app running under a debugger attached.
//
// To enable the debugger attached mode for cmdr, run `go build` with `-tags=delve` options. eg:
//
//	go run -tags=delve ./cli
//	go build -tags=delve -o my-app ./cli
//
// For Goland, you can enable this under 'Run/Debug Configurations', by adding the following into 'Go tool arguments:'
//
//	-tags=delve
//
// IsDebuggerAttached() is a synonym to InDebugging().
//
// NOTE that `isdelve` algor is from https://stackoverflow.com/questions/47879070/how-can-i-see-if-the-goland-debugger-is-running-in-the-program
//
// noinspection GoBoolExpressions
func IsDebuggerAttached() bool {
	return states.Env().InDebugging() // isdelve.Enabled
}

func InTracing() bool {
	return states.Env().GetTraceMode()
}

// InTestingT detects whether is running under 'go test' mode
func InTestingT(args []string) bool {
	if !strings.HasSuffix(args[0], ".test") &&
		!strings.Contains(args[0], "/T/___Test") {

		// [0] = /var/folders/td/2475l44j4n3dcjhqbmf3p5l40000gq/T/go-build328292371/b001/exe/main
		// !strings.Contains(SavedOsArgs[0], "/T/go-build")

		for _, s := range args {
			if s == "-test.v" || s == "-test.run" {
				return true
			}
		}
		return false

	}
	return true
}

// InTesting detects whether is running under go test mode
func InTesting() bool {
	return InTestingT(os.Args)
	// if !strings.HasSuffix(tool.SavedOsArgs[0], ".test") &&
	//	!strings.Contains(tool.SavedOsArgs[0], "/T/___Test") {
	//
	//	// [0] = /var/folders/td/2475l44j4n3dcjhqbmf3p5l40000gq/T/go-build328292371/b001/exe/main
	//	// !strings.Contains(SavedOsArgs[0], "/T/go-build")
	//
	//	for _, s := range tool.SavedOsArgs {
	//		if s == "-test.v" || s == "-test.run" {
	//			return true
	//		}
	//	}
	//	return false
	//
	// }
	// return true
}

// func InTestingWith(osArgs []string) bool { return log.InTestingT(osArgs) } // test if in testing mode

// InDevelopingTime detects whether is in developing time.
//
// If the main program has been built as an executable binary, we
// would assume which is not in developing time.
//
// If log.GetDebugMode() is true, that's in developing time too.
func InDevelopingTime() (status bool) {
	return InDebugging() || InTesting()
}

// InDockerEnvSimple detects whether is running within docker
// container environment.
//
// InDockerEnvSimple finds if `/.dockerenv` exists or not.
func InDockerEnvSimple() (status bool) {
	return isRunningInDockerContainer()
}

func isRunningInDockerContainer() bool {
	// docker creates a .dockerenv file at the root
	// of the directory tree inside the container.
	// if this file exists then the viewer is running
	// from inside a container so return true

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
