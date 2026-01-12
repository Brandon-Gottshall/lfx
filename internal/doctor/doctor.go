package doctor

import (
	"fmt"
	"os"
)

type Result struct {
	Issues []string
}

func Check(registryRoot string, lfConfigDir string) Result {
	var issues []string

	if registryRoot == "" {
		issues = append(issues, "registry root not provided; run from the repo root or set LFX_REGISTRY_ROOT")
	} else if _, err := os.Stat(registryRoot); err != nil {
		issues = append(issues, fmt.Sprintf("registry not found at %s; ensure that directory exists", registryRoot))
	}

	if lfConfigDir == "" {
		issues = append(issues, "lf config path not provided; set XDG_CONFIG_HOME or define ~/.config/lf")
	} else if _, err := os.Stat(lfConfigDir); err != nil {
		issues = append(issues, fmt.Sprintf("lf config dir not found at %s; create it before running lf or lfx", lfConfigDir))
	}

	return Result{Issues: issues}
}
