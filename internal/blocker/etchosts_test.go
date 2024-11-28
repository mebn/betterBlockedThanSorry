package blocker

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func initFile(t *testing.T, path string, data ...string) {
	file, err := os.Create(path) // truncate file
	if err != nil {
		t.Fatal("creating file failed")
	}
	if len(data) != 0 {
		file.WriteString(data[0])
	}
	file.Close()
}

func TestGenerateBlocklist(t *testing.T) {
	want := []string{
		"127.0.0.1 a.com",
		"127.0.0.1 www.a.com",
		"::1 a.com",
		"::1 www.a.com",
		"127.0.0.1 www.b.com",
		"127.0.0.1 www.www.b.com",
		"::1 www.b.com",
		"::1 www.www.b.com",
	}

	blocklist := []string{"a.com", "www.b.com"}

	got := generateBlocklist(blocklist)

	if !reflect.DeepEqual(want, got) {
		t.Fatal(want, got)
	}
}

func TestAddBlock(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal("Failed to create temp dir:", err)
	}
	defer os.RemoveAll(tempDir)
	path := filepath.Join(tempDir, "file")

	t.Run("add block", func(t *testing.T) {
		initFile(t, path)

		etcHosts := NewEtcHosts(path, []string{"a.com", "www.b.com"})
		etcHosts.AddBlock()

		gotB, _ := os.ReadFile(path)
		got := string(gotB)

		want := `
127.0.0.1 a.com
127.0.0.1 www.a.com
::1 a.com
::1 www.a.com
127.0.0.1 www.b.com
127.0.0.1 www.www.b.com
::1 www.b.com
::1 www.www.b.com`

		if want != got {
			t.Fatal(want, got)
		}
	})

	t.Run("should not override data", func(t *testing.T) {
		initFile(t, path, `
some
data`)

		etcHosts := NewEtcHosts(path, []string{"a.com", "www.b.com"})
		etcHosts.AddBlock()

		gotB, _ := os.ReadFile(path)
		got := string(gotB)

		want := `
some
data
127.0.0.1 a.com
127.0.0.1 www.a.com
::1 a.com
::1 www.a.com
127.0.0.1 www.b.com
127.0.0.1 www.www.b.com
::1 www.b.com
::1 www.www.b.com`

		if want != got {
			t.Fatal(want, got)
		}
	})
}

func TestDeleteBlock(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal("Failed to create temp dir:", err)
	}
	defer os.RemoveAll(tempDir)
	path := filepath.Join(tempDir, "file")

	t.Run("remove a block", func(t *testing.T) {
		initFile(t, path, `
some
data`)

		etcHosts := NewEtcHosts(path, []string{"a.com", "www.b.com"})
		etcHosts.AddBlock()
		etcHosts.RemoveBlock()

		gotB, _ := os.ReadFile(path)
		got := string(gotB)

		want := `
some
data`

		if want != got {
			t.Fatal(want, got)
		}
	})
}

func TestIsTamperedWith(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal("Failed to create temp dir:", err)
	}
	defer os.RemoveAll(tempDir)
	path := filepath.Join(tempDir, "file")

	t.Run("no modifications", func(t *testing.T) {
		initFile(t, path)

		etcHosts := NewEtcHosts(path, []string{"a.com", "www.b.com"})
		etcHosts.AddBlock()

		got := etcHosts.IsTamperedWith()
		want := false

		if want != got {
			t.Fatal(want, got)
		}
	})

	t.Run("modifications", func(t *testing.T) {
		initFile(t, path)

		etcHosts := NewEtcHosts(path, []string{"a.com", "www.b.com"})
		etcHosts.AddBlock()

		initFile(t, path)

		got := etcHosts.IsTamperedWith()
		want := true

		if want != got {
			t.Fatal(want, got)
		}
	})
}
