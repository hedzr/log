// Package detects collects all detector, buildtags, environment settings in one place.
//
// By using detects package, you have no needs to call `buildtags`, `states` and other similar packages.
package detects

import (
	"os"
	"strings"

	"github.com/hedzr/log/buildtags"
	"github.com/hedzr/log/dir"
	"github.com/hedzr/log/isdelve"
	"github.com/hedzr/log/states"
	"github.com/hedzr/log/trace"
)

// InK8s detects if the service is running under k8s environment.
func InK8s() bool {
	return os.Getenv("KUBERNETES_SERVICE_HOST") != "" || buildtags.IsK8sBuild()
}

// InK8sYN is yet another DetectInK8s impl
func InK8sYN() bool {
	return dir.FileExists("/var/run/secrets/kubernetes.io") || buildtags.IsK8sBuild()
}

// InIstio detects if the service is running under istio injected.
//
// # IMPORTANT
//
// To make this detector work properly, you must mount a DownwordAPI
// volume to your container/pod. See also:
// https://kubernetes.io/en/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/
func InIstio() bool {
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		kf := "/etc/podinfo/labels"
		if dir.FileExists(kf) {
			if data, err := dir.ReadFile(kf); err == nil {
				// lines:=strings.Split(string(data),"\n")
				if strings.Contains(string(data), "service.istio.io/canonical-name") {
					return true
				}
			}
		}
	}
	return buildtags.IsIstioBuild()
}

// InDocker detects if the service is running under docker environment.
//
// InDocker test these two conditions:
// 1. find if `/.dockerenv` exists or not.
// 2. `docker` in buildtags
func InDocker() bool {
	if dir.FileExists("/.dockerenv") {
		return true
	}
	return buildtags.IsDockerBuild()
}

func IsDockerBuild() bool { return buildtags.IsDockerBuild() }
func IsK8sBuild() bool    { return buildtags.IsK8sBuild() }
func IsIstioBuild() bool  { return buildtags.IsIstioBuild() }

func IsVerboseBuild() bool       { return buildtags.VerboseEnabled } // is verbose build
func IsVerboseModeEnabled() bool { return Env().IsVerboseMode() }    // is verbose build, or is CLI Verbose mode enabled (by `--verbose`)?
func GetVerboseLevel() int       { return Env().CountOfVerbose() }   //
func IsQuietModeEnabled() bool   { return Env().IsQuietMode() }      // is quiet build, or is CLI Quiet mode enabled (by `--verbose`)?
func GetQuietLevel() int         { return Env().CountOfQuiet() }     //
func IsNoColorModeEnabled() bool { return Env().IsNoColorMode() }    //
func GetNoColorLevel() int       { return Env().CountOfNoColor() }   //
func IsDebugBuild() bool         { return isdelve.Enabled }          // is debug build?
func IsDebugModeEnabled() bool   { return Env().GetDebugMode() }     // is debug build, or is CLI debug mode enabled (by `--debug`)?
func GetDebugLevel() int         { return Env().GetDebugLevel() }    //
func IsTracingEnabled() bool     { return trace.IsEnabled() }        // is tracing-flag true in trace package
func IsTraceModeEnabled() bool   { return Env().GetTraceMode() }     // is CLI trace mode enabled (by `--trace`)? or is tracing-flag true in trace package
func GetTraceLevel() int         { return Env().GetTraceLevel() }    //

// States or Env returns a minimal environment settings for a typical CLI app.
//
// See also states.CmdrMinimal.
func States() states.CmdrMinimal { return states.Env() }
func Env() states.CmdrMinimal    { return states.Env() } // States or Env returns a minimal environment settings for a typical CLI app.
