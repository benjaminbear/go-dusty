package rpcserver

import (
	"os"
	"path/filepath"
)

func defaultConfigPath() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(cfgDir, "dusty"), nil
}

func defaultReposPath() (string, error) {
	cfgDir, err := defaultConfigPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(cfgDir, "repos"), nil
}
