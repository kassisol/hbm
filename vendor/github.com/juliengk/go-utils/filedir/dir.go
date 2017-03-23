package filedir

import (
	"os"
)

func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateDirIfNotExist(path string, parent bool, perm os.FileMode) error {
	if !DirExists(path) {
		if parent {
			if err := os.MkdirAll(path, perm); err != nil {
				return err
			}
		} else {
			if err := os.Mkdir(path, perm); err != nil {
				return err
			}
		}
	}

	return nil
}
