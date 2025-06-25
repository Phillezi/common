package testutils

import (
	"os"
	"path/filepath"
	"testing"
)

// GenericTest is a reusable testing wrapper
func GenericTest(t *testing.T, name string, setup func(tmpDir string) error, test func(tmpDir string) error) {
	t.Run(name, func(t *testing.T) {
		tmpDir := t.TempDir()

		// Setup phase
		if setup != nil {
			if err := setup(tmpDir); err != nil {
				t.Fatalf("Setup failed: %v", err)
			}
		}

		// Test execution
		if err := test(tmpDir); err != nil {
			t.Errorf("Test failed: %v", err)
		}
	})
}

// WriteFile writes content to a file path
func WriteFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// CreateFileInDir is a helper to create a file inside a dir
func CreateFileInDir(dir, filename, content string) (string, error) {
	fullPath := filepath.Join(dir, filename)
	return fullPath, WriteFile(fullPath, content)
}
