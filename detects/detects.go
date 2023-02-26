package detects

import (
	"os"
	"strings"

	"github.com/hedzr/log/buildtags"
	"github.com/hedzr/log/dir"
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
