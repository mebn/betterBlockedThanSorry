package env

import (
	"fmt"
	"testing"
)

func TestGetCurrentFilePath(t *testing.T) {
	fmt.Println(getCurrentFilePath())
	fmt.Println(ProgramPath)

	// t.Fatal("")
}
