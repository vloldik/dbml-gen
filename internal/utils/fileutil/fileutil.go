package fileutil

import (
	"io/fs"
	"os"
)

func EnsureFolderExists(location string, mode fs.FileMode) error {
	if err := os.MkdirAll(location, mode); err != nil {
		if err == os.ErrExist {
			return nil
		}
		return err
	}
	return nil
}
