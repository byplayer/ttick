package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func Exists(path string) (bool, error) {
	_, fileErr := os.Stat(path)
	if fileErr == nil {
		return true, nil
	}
	if os.IsNotExist(fileErr) {
		return false, nil
	}
	return true, nil
}

func AssureExists(filePath string) error {
	path := filepath.Dir(filePath)
	exists, err := Exists(path)
	if err != nil {
		return err
	}
	if !exists {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("couldn't create path: %s", path)
		}
	}
	return nil
}
