package database

import (
	"os"
	"reflect"
	"testing"
)

func TestEndtime(t *testing.T) {
	path := "/tmp/testingdatabase"

	//setup
	db, err := NewDB(path)
	if err != nil {
		t.Fatal("error:", err)
	}
	defer db.CloseDB()

	// testing
	want := int64(0)

	got, err := db.GetEndtime()
	if err != nil {
		t.Fatal("getendtime failed: ", err)
	}

	if got != want {
		t.Fatal(got, want)
	}

	want = int64(1337)

	err = db.SetEndtime(want)
	if err != nil {
		t.Fatal("SetEndtime failed: ", err)
	}

	got, err = db.GetEndtime()
	if err != nil {
		t.Fatal("getendtime failed: ", err)
	}

	if got != want {
		t.Fatal(got, want)
	}

	// cleanup
	db.CloseDB()
	os.RemoveAll(path)
}

func TestBlocklist(t *testing.T) {
	path := "/tmp/testingdatabase"

	//setup
	db, err := NewDB(path)
	if err != nil {
		t.Fatal("error:", err)
	}
	defer db.CloseDB()

	// testing
	want := []string{}

	got, err := db.GetBlocklist()
	if err != nil {
		t.Fatal("getBlocklist failed: ", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatal(got, want)
	}

	want = []string{"a", "b", "c"}

	err = db.SetBlocklist(want)
	if err != nil {
		t.Fatal("SetBlocklist failed: ", err)
	}

	got, err = db.GetBlocklist()
	if err != nil {
		t.Fatal("getBlocklist failed: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal(got, want)
	}

	// cleanup
	db.CloseDB()
	os.RemoveAll(path)
}
