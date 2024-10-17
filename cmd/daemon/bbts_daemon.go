package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	i := 0

	f, _ := os.OpenFile("/tmp/mylog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if len(os.Args) < 2 {
		f.WriteString("got no argument \n")
	} else {
		f.WriteString(fmt.Sprintf("got argument %s \n", os.Args[1]))
	}

	for {
		f.WriteString(fmt.Sprintf("line %d \n", i))
		i += 1
		time.Sleep(time.Second * 1)
	}
}
