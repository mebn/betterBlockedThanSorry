package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestGetEndTime(t *testing.T) {
	file, _ := os.OpenFile("/tmp/bbtstestfile456745.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	time := time.Now().Unix()
	file.WriteString(fmt.Sprintf("%d", time))
	file.Close()

	contentB, _ := os.ReadFile("/tmp/bbtstestfile456745.log")
	content := string(contentB)
	endtime, _ := strconv.ParseInt(content, 10, 64)

	if endtime != time {
		t.Fatal(endtime, time)
	}

	os.Remove("/tmp/bbtstestfile456745.log")

}
