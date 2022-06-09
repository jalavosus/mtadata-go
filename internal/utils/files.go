package utils

import (
	"os"
	"path/filepath"
)

const (
	openFileWriteFlag = os.O_RDWR | os.O_CREATE
	openFileReadFlag  = os.O_RDONLY
	openFilePerm      = 0755
)

// OpenFileWrite wraps os.OpenFile, passing a combo of
// os.O_RDWR | os.O_CREATE as the flag parameter (enabling file write _or_ creation),
// and 0755 as the file permission.
func OpenFileWrite(fp string) (*os.File, error) {
	return os.OpenFile(fp, openFileWriteFlag, openFilePerm)
}

// OpenFileRead wraps os.OpenFile, passing os.O_RDONLY
// and 0755 as the open flag and file permission.
func OpenFileRead(fp string) (*os.File, error) {
	return os.OpenFile(fp, openFileReadFlag, openFilePerm)
}

// ClearFile calls Truncate(0) and Seek(0, 0) on an os.File object,
// clearing all of its contents.
func ClearFile(f *os.File) error {
	if err := f.Truncate(0); err != nil {
		return err
	}

	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	return nil
}

// AbsolutePath wraps filepath.Abs,
// panicking if filepath.Abs returns an error.
func AbsolutePath(fp string) string {
	if filepath.IsAbs(fp) {
		return fp
	}

	absFp, err := filepath.Abs(fp)
	if err != nil {
		panic(err)
	}

	return absFp
}
