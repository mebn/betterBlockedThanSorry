package env

import (
	"path/filepath"
	"runtime"
)

func getCurrentFilePath() string {
	_, file, _, ok := runtime.Caller(1)

	if !ok {
		panic("Unable to get current file info")
	}

	return filepath.Dir(file)
}
