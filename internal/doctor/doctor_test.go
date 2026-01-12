package doctor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheck_MissingRegistry(t *testing.T) {
	temp := t.TempDir()

	registryRoot := filepath.Join(temp, "registry")
	result := Check(registryRoot, temp)

	if len(result.Issues) == 0 {
		t.Fatalf("expected missing registry issue")
	}
	found := false
	for _, issue := range result.Issues {
		if issue == "registry not found at "+registryRoot+"; ensure that directory exists" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected registry issue, got %v", result.Issues)
	}
}

func TestCheck_MissingLfConfig(t *testing.T) {
	temp := t.TempDir()
	registryDir := filepath.Join(temp, "registry")
	if err := os.MkdirAll(registryDir, 0o755); err != nil {
		t.Fatalf("failed to create registry dir: %v", err)
	}

	lfConfigDir := filepath.Join(temp, "lf")
	result := Check(registryDir, lfConfigDir)

	if len(result.Issues) == 0 {
		t.Fatalf("expected missing lf config issue")
	}
	found := false
	for _, issue := range result.Issues {
		if issue == "lf config dir not found at "+lfConfigDir+"; create it before running lf or lfx" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected lf config issue, got %v", result.Issues)
	}
}
