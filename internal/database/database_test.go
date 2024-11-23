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
	want := "1337"

	got, err := db.GetEndtime()
	if err == nil {
		t.Fatal("We should get an error here. Database should be empty.", got, err)
	}

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
	want := []string{"a", "b", "c"}

	got, err := db.GetBlocklist()
	if err == nil {
		t.Fatal("We should get an error here. Database should be empty.", got, err)
	}

	err = db.SetBlocklist(want)
	if err != nil {
		t.Fatal("SetBlocklist failed: ", err)
	}

	got, err = db.GetBlocklist()
	if err != nil {
		t.Fatal("getendtime failed: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal(got, want)
	}

	// cleanup
	db.CloseDB()
	os.RemoveAll(path)
}
