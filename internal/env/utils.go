package env

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func currentDir() string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Dir(file)
}

func Home() string {
	currentUser, _ := user.Current()
	return currentUser.HomeDir
}

func SafePath(args ...string) string {
	path := filepath.Join(args...)
	_ = os.MkdirAll(path, 0755)
	return path
}

func SafeFile(args ...string) string {
	path := filepath.Join(args...) // dir, dir, filename
	dir := filepath.Dir(path)
	_ = SafePath(dir)
	return path
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
