package mod

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Phillezi/common/utils/or"
	ev "github.com/Phillezi/common/utils/path"
)

// Walks up from the current directory until it finds a go.mod file.
func FindGoModRoot(from ...string) (string, error) {
	path := or.Or(from...)
	if path != "" {
		path = ev.EvalAbsolute(path)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = dir
	}
	for {
		modPath := filepath.Join(path, "go.mod")
		if _, err := os.Stat(modPath); err == nil {
			return path, nil
		}
		parent := filepath.Dir(path)
		if parent == path {
			break // reached root without finding go.mod
		}
		path = parent
	}
	return "", fmt.Errorf("go.mod not found in any parent directory")
}
