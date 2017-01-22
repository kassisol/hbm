package filedir

import (
	"os"
)

func FileExists(f string) bool {
	_, err := os.Lstat(f)
	if err != nil {
		return false
	}

	return true
}

func IsSymlink(f string) (bool, string, error) {
	t := false
	link := ""

	fi, err := os.Lstat(f)
	if err != nil {
		return t, link, err
	}

	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		t = true
		link, err = os.Readlink(f)
		if err != nil {
			return t, link, err
		}
	}

	return t, link, nil
}
