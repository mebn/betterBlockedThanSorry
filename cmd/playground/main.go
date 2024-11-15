package main

import (
	"fmt"

	"github.com/mebn/betterBlockedThanSorry/internal/env"
)

func main() {
	fmt.Println("playground")

	logPath := env.ProgramPath
	fmt.Println(logPath)
}

/*
go run -tags=prod cmd/playground/main.go
*/
