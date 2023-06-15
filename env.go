// Copyright © 2023 Hedzr Yeh.

package log

import (
	"github.com/hedzr/log/detects"
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
func InDebugging() bool { return detects.InDebugging() }

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
func IsDebuggerAttached() bool { return detects.InDebugging() }

func InTracing() bool { return detects.InTracing() }

// InTestingT detects whether is running under 'go test' mode
func InTestingT(args []string) bool { return detects.InTestingT(args) }

// InTesting detects whether is running under go test mode
func InTesting() bool { return detects.InTesting() }

// func InTestingWith(osArgs []string) bool { return log.InTestingT(osArgs) } // test if in testing mode

// InDevelopingTime detects whether is in developing time.
//
// If the main program has been built as an executable binary, we
// would assume which is not in developing time.
//
// If log.GetDebugMode() is true, that's in developing time too.
func InDevelopingTime() (status bool) { return detects.InDevelopingTime() }

// InDockerEnvSimple detects whether is running within docker
// container environment.
//
// InDockerEnvSimple finds if `/.dockerenv` exists or not.
func InDockerEnvSimple() (status bool) { return detects.InDockerEnvSimple() }
