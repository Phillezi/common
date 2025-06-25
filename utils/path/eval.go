package path

import (
	"os"
	"path/filepath"
	"strings"
)

// Eval expands ~ to the user's home directory if present.
func Eval(path string) string {
	if strings.HasPrefix(path, "~") {
		if home, err := os.UserHomeDir(); err == nil {
			return filepath.Join(home, strings.TrimPrefix(path, "~"))
		}
	}
	return path
}

// EvalAbsolute expands ~ and returns the absolute path.
func EvalAbsolute(path string) string {
	expanded := Eval(path)
	abs, err := filepath.Abs(expanded)
	if err != nil {
		// If we can't resolve to absolute, return the expanded path as a fallback
		return expanded
	}
	return abs
}
