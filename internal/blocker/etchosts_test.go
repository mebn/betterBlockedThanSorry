package blocker

import (
	"fmt"
	"os"
	"testing"
)

func fileSetup(t *testing.T) (*os.File, string, string) {
	filename := "/tmp/bbtsignoreme4real"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Fatal("Failed to open/create file, try running with sudo.")
	}

	file.Truncate(0)
	fileContent := "some\ncontent\n\n"
	file.WriteString(fileContent)

	return file, fileContent, filename
}

func TestGenerateEtcHosts(t *testing.T) {
	blocklist := []string {"aaa.com", "www.bbb.com"}

	want := fmt.Sprintf("%s%s%s\n", startBlock, `
127.0.0.1 aaa.com
127.0.0.1 www.aaa.com
127.0.0.1 www.bbb.com
127.0.0.1 www.www.bbb.com
`, endBlock)

	got := GenerateEtcHosts(blocklist)

	if want != got {
		t.Fatalf("want != got:\n%s\n%s", want, got)
	}
}

func TestFileTamperedWith(t *testing.T) {
	file, fileContent, _ := fileSetup(t)
	defer file.Close()

	blockpart := GenerateEtcHosts([]string {"hello"})
	file.WriteString(blockpart)

	// test no modifications
	want := false
	got := FileTamperedWith(file, blockpart)

	if want != got {
		t.Fatalf("[no modifications] want != got:\n%t\n%t", want, got)
	}

	// test modifications (changes inside start-and endBlock)
	file.Truncate(0)
	file.WriteString(fileContent)
	file.WriteString(startBlock)
	file.WriteString("\n")
	file.WriteString("not 127.0.0.1 hello")
	file.WriteString(endBlock)
	file.WriteString("\n")
	
	want = true
	got = FileTamperedWith(file, blockpart)

	if want != got {
		t.Fatalf("[inside modifications] want != got:\n%t\n%t", want, got)
	}

	// test modifications (no startBlock)
	file.Truncate(0)
	file.WriteString(fileContent)
	file.WriteString("\n")
	file.WriteString("127.0.0.1 hello")
	file.WriteString("127.0.0.1 www.hello")
	file.WriteString(endBlock)
	file.WriteString("\n")
	
	want = true
	got = FileTamperedWith(file, blockpart)

	if want != got {
		t.Fatalf("[no startBlock modifications] want != got:\n%t\n%t", want, got)
	}

	// test modifications (no endBlock)
	file.Truncate(0)
	file.WriteString(fileContent)
	file.WriteString("\n")
	file.WriteString(startBlock)
	file.WriteString("127.0.0.1 hello")
	file.WriteString("127.0.0.1 www.hello")
	file.WriteString("\n")
	
	want = true
	got = FileTamperedWith(file, blockpart)

	if want != got {
		t.Fatalf("[no endBlock modifications] want != got:\n%t\n%t", want, got)
	}
}

func TestAddBlock(t *testing.T) {
	file, fileContent, filename := fileSetup(t)
	defer file.Close()

	blockpart := GenerateEtcHosts([]string {"hello"})
	AddBlock(file, blockpart)

	want := fmt.Sprintf("%s%s", fileContent, blockpart)
	gotByte, _ := os.ReadFile(filename)
	got := string(gotByte)

	if want != got {
		t.Fatalf("want != got:\n%s\n%s", want, got)
	}
}

func TestRemoveBlock(t *testing.T) {
	file, fileContent, filename := fileSetup(t)
	defer file.Close()

	blockpart := GenerateEtcHosts([]string {"hello"})
	AddBlock(file, blockpart)

	RemoveBlock(file)

	want := fileContent
	gotByte, _ := os.ReadFile(filename)
	got := string(gotByte)

	if want != got {
		t.Fatalf("want != got:\n%s\n%s", want, got)
	}
}