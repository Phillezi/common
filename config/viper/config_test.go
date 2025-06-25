package viperconf

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	lggr "github.com/Phillezi/common/config/logger"
	"github.com/Phillezi/common/config/testutils"
	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	testutils.GenericTest(t, "InitConfig with valid config file", func(tmpDir string) error {
		t.Setenv("XDG_CONFIG_HOME", tmpDir)
		project := "myproject"
		dir := filepath.Join(tmpDir, project)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		_, err := testutils.CreateFileInDir(dir, "config.yaml", "foo: bar")
		return err
	}, func(tmpDir string) error {
		project := "myproject"
		mock, logs := lggr.NewMockLogger()
		SetLogger(mock)

		InitConfig(project)

		if len(logs.Debugs) == 0 {
			return errors.New("expected debug logs")
		}
		found := false
		for _, msg := range logs.Debugs {
			if strings.Contains(msg, "Using config file") {
				found = true
			}
		}
		if !found {
			return errors.New("expected 'Using config file' in debug logs")
		}
		return nil
	})
}

func TestInitConfig_NoConfigFile(t *testing.T) {
	testutils.GenericTest(t, "InitConfig with no config file", func(tmpDir string) error {
		t.Setenv("XDG_CONFIG_HOME", tmpDir)
		return nil // No config file created
	}, func(tmpDir string) error {
		project := "myproject"
		mock, logs := lggr.NewMockLogger()
		SetLogger(mock)

		InitConfig(project)

		if len(logs.Debugs) == 0 {
			return errors.New("expected debug logs")
		}
		found := false
		for _, msg := range logs.Debugs {
			if strings.Contains(msg, "Config file not found") {
				found = true
				break
			}
		}
		if !found {
			return errors.New("expected 'Config file not found' message in debug logs")
		}
		return nil
	})
}

func TestInitConfig_CustomFilename(t *testing.T) {
	testutils.GenericTest(t, "InitConfig with custom filename", func(tmpDir string) error {
		t.Setenv("XDG_CONFIG_HOME", tmpDir)
		project := "myproject"
		dir := filepath.Join(tmpDir, project)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		_, err := testutils.CreateFileInDir(dir, "custom.yaml", "foo: bar")
		return err
	}, func(tmpDir string) error {
		project := "myproject"
		mock, logs := lggr.NewMockLogger()
		SetLogger(mock)

		InitConfig(project, "custom")

		if len(logs.Debugs) == 0 {
			return errors.New("expected debug logs")
		}
		found := false
		for _, msg := range logs.Debugs {
			if strings.Contains(msg, "Using config file") {
				found = true
			}
		}
		if !found {
			return errors.New("expected 'Using config file' in debug logs")
		}
		return nil
	})
}

func TestInitConfig_NestedDirectory(t *testing.T) {
	testutils.GenericTest(t, "InitConfig reads config in nested dir", func(tmpDir string) error {
		t.Setenv("XDG_CONFIG_HOME", tmpDir)
		project := "kthcloud"
		dir := filepath.Join(tmpDir, project)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		_, err := testutils.CreateFileInDir(dir, "config.yaml", "foo: baz")
		return err
	}, func(tmpDir string) error {
		project := "kthcloud"
		mock, logs := lggr.NewMockLogger()
		SetLogger(mock)

		InitConfig(project)

		found := false
		for _, msg := range logs.Debugs {
			if strings.Contains(msg, "Using config file") {
				found = true
			}
		}
		if !found {
			return errors.New("expected config file to be used from nested directory")
		}
		return nil
	})
}

func TestInitConfig_EnvOverride(t *testing.T) {
	testutils.GenericTest(t, "Env variable overrides config", func(tmpDir string) error {
		t.Setenv("XDG_CONFIG_HOME", tmpDir)
		t.Setenv("FOO", "env-bar")

		project := "myproject"
		dir := filepath.Join(tmpDir, project)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		_, err := testutils.CreateFileInDir(dir, "config.yaml", "foo: bar")
		return err
	}, func(tmpDir string) error {
		project := "myproject"
		mock, _ := lggr.NewMockLogger()
		SetLogger(mock)

		InitConfig(project)
		val := viper.GetString("foo")
		if val != "env-bar" {
			return errors.New("expected env var to override config file value")
		}
		return nil
	})
}
