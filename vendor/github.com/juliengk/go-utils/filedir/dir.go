package filedir

import (
	"os"
)

func CreateIfNotExist(path string, perm os.FileMode) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, perm)
		if err != nil {
			return err
		}
	}

	return nil
}
