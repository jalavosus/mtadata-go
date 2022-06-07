package utils

import (
	"os"
)

const (
	openFileWriteFlag = os.O_RDWR | os.O_CREATE
	openFileReadFlag  = os.O_RDONLY
	openFilePerm      = 0755
)

func OpenFileWrite(fp string) (*os.File, error) {
	return os.OpenFile(fp, openFileWriteFlag, openFilePerm)
}

func OpenFileRead(fp string) (*os.File, error) {
	return os.OpenFile(fp, openFileReadFlag, openFilePerm)
}

func ClearFile(f *os.File) error {
	if err := f.Truncate(0); err != nil {
		return err
	}

	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	return nil
}
