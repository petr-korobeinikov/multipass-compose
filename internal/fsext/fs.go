package fsext

import (
	"errors"
	"io/fs"
	"os"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return false
	}

	return true
}
