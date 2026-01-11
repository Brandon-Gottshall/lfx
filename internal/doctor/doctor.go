package doctor

import (
	"os"
)

type Result struct {
	Issues []string
}

func Check(registryRoot string, lfConfigDir string) Result {
	var issues []string

	if registryRoot == "" {
		issues = append(issues, "registry root not provided")
	} else if _, err := os.Stat(registryRoot); err != nil {
		issues = append(issues, "registry not found: "+registryRoot)
	}

	if lfConfigDir == "" {
		issues = append(issues, "lf config path not provided")
	} else if _, err := os.Stat(lfConfigDir); err != nil {
		issues = append(issues, "lf config dir not found: "+lfConfigDir)
	}

	return Result{Issues: issues}
}
