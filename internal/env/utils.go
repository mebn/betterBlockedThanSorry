package env

import (
	"os"
	"path/filepath"
	"runtime"
)

func currentDir() string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Dir(file)
}

func safePath(args ...string) string {
	fullPath := filepath.Join(args...) // dir, dir, filename
	dir := filepath.Dir(fullPath)
	os.MkdirAll(dir, 0755)
	return fullPath
}

func MoveProgram() error {
	_, err := os.Stat(FirstProgramPath)
	if os.IsNotExist(err) {
		// source file does not exist
		// TODO: download new daemon
		return err
	}
	if err != nil {
		// failed to check source file
		return err
	}

	err = os.Rename(FirstProgramPath, ProgramPath)
	return err
}
