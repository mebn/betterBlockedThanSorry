package blocker

import (
	"os"
	"reflect"
	"testing"
)

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

func TestAddAndDeleteBLock(t *testing.T) {
	// setup
	filename := "/tmp/bbtsanothertempfile346782"
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	file.Truncate(0)
	file.WriteString(`hello
my
name
is`)
	file.Close()

	blocklist := []string{"a.com", "www.b.com"}
	etcHosts := NewEtcHosts(filename, blocklist)

	// test add
	etcHosts.AddBlock()
	want := `hello
my
name
is
127.0.0.1 a.com
127.0.0.1 www.a.com
::1 a.com
::1 www.a.com
127.0.0.1 www.b.com
127.0.0.1 www.www.b.com
::1 www.b.com
::1 www.www.b.com`
	gotB, _ := os.ReadFile(filename)
	got := string(gotB)

	if want != got {
		t.Fatal(want, got)
	}

	// test remove
	etcHosts.RemoveBlock()
	want = `hello
my
name
is`
	gotB, _ = os.ReadFile(filename)
	got = string(gotB)

	if want != got {
		t.Fatal(want, got)
	}

	// cleanup
	os.Remove(filename)
}

func TestIsTamperedWith(t *testing.T) {
	// setup
	filename := "/tmp/bbtsanothertempfile1238246"
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	file.Truncate(0)
	file.WriteString(`hello
my
name
is`)
	file.Close()

	blocklist := []string{"a.com", "www.b.com"}
	etcHosts := NewEtcHosts(filename, blocklist)

	etcHosts.AddBlock()

	// test no modifications
	want := false
	got := etcHosts.IsTamperedWith()

	if want != got {
		t.Fatal(want, got)
	}

	// test modification (a.com removed)
	file, _ = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	file.Truncate(0)
	file.WriteString(`hello
my
name
is
127.0.0.1 www.a.com
::1 a.com
::1 www.a.com
127.0.0.1 www.b.com
127.0.0.1 www.www.b.com
::1 www.b.com
::1 www.www.b.com`)

	file.Close()

	want = true
	got = etcHosts.IsTamperedWith()

	if want != got {
		t.Fatal(want, got)
	}

	// cleanup
	os.Remove(filename)
}
