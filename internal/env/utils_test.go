package env

// import (
// 	"testing"
// )

// func TestSafePath(t *testing.T) {
// 	path := SafePath(folder, "testing", "folder")
// 	println(path)

// 	t.Fatal()
// }

// func TestJoin(t *testing.T) {
// 	fullPath := filepath.Join("a", "path", "to", "a", "file")
// 	println(fullPath)
// 	t.Fatal()
// }

// func TestSafePath(t *testing.T) {
// 	got := safePath(home(), ".bbtsDEVTEST", "afolder", "afile.txt")
// 	want := home() + "/.bbtsDEVTEST/afolder/afile.txt"

// 	if got != want {
// 		t.Fatal(got, want)
// 	}

// 	os.WriteFile(got, []byte("hello world"), 0644)

// 	b, _ := os.ReadFile(got)
// 	got = string(b)
// 	want = "hello world"

// 	if got != want {
// 		t.Fatal(got, want)
// 	}

// 	// cleanup
// 	err := os.RemoveAll(".bbtsDEVTEST")
// 	if err != nil {
// 		println(err)
// 		t.Fatal()
// 	}
// }
