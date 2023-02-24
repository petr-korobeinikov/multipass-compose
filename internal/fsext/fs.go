package fsext

import (
	"errors"
	"io/fs"
	"os"
)

const ModeFile = 0o644

func FileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return false
	}

	return true
}
