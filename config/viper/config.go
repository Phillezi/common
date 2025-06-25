package viperconf

import (
	"fmt"
	"os"
	"path"

	"github.com/Phillezi/common/utils/or"
	"github.com/spf13/viper"
)

// InitConfig initializes configuration
func InitConfig(projectName string, filenames ...string) {
	if len(filenames) > 0 {
		filenames = append(filenames, "config")
	} else {
		filenames = []string{"config"}
	}
	viper.SetConfigName(or.Or(filenames...)) // Name of the config file (without extension)
	viper.SetConfigType("yaml")              // File format (yaml)
	viper.AddConfigPath(GetConfigPath(projectName))
	viper.AutomaticEnv() // Read environment variables

	// Load config file
	if err := viper.ReadInConfig(); err != nil {
		logger.Debug("Config file not found, using defaults or environment variables.")
	} else {
		logger.Debug("Using config file: ", viper.ConfigFileUsed())
	}

}

func getConfigPath(projectName string) (string, error) {
	basePath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configPath := path.Join(basePath, projectName)
	fileDescr, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			return "", err
		}
		fileDescr, err = os.Stat(configPath)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	if !fileDescr.IsDir() {
		return "", fmt.Errorf("default config dir is file")
	}
	return configPath, nil
}

func GetConfigPath(projectName string) string {
	configPath, err := getConfigPath(projectName)
	if err != nil {
		logger.Warn("error getting config path", err)
		logger.Info("defaulting to:", func() string {
			wd, err := os.Getwd()
			if err != nil {
				return "."
			}
			return wd
		}())
		configPath = "."
	}
	return configPath
}
