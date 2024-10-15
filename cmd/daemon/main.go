package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	i := 0

	f, err := os.OpenFile("/tmp/mylog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	}
	defer f.Close()

	for {
		_, err = f.WriteString(fmt.Sprintf("line %d", i))
		i += 1
		time.Sleep(3 * time.Second)
	}
}
