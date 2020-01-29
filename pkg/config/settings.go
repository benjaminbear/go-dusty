package config

import (
	"os"
	"path/filepath"
)

func defaultConfigPath() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(cfgDir, "dusty/config.yml"), nil
}
