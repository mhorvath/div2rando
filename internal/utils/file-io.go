package utils

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
)

func GetUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, "div2rando"), nil
}

func LoadJSON(filename string, v interface{}) error {
	homeDir, err := GetUserHomeDir()
	if err != nil {
		return err
	}

	filePath := filepath.Join(homeDir, filename)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, v)
	if err != nil {
		return err
	}

	return nil
}

func SaveJSON(filename string, v interface{}) error {
	homeDir, err := GetUserHomeDir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(homeDir, 0755)
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, filename)
	data, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
