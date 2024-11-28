package database

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestEndtime(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal("Failed to create temp dir:", err)
	}
	defer os.RemoveAll(tempDir)
	path := filepath.Join(tempDir, "db.db")

	db, err := NewDB(path)
	if err != nil {
		t.Fatal("error:", err)
	}
	defer db.CloseDB()

	t.Run("default value", func(t *testing.T) {
		got, err := db.GetEndtime()
		if err != nil {
			t.Fatal("getendtime failed: ", err)
		}

		want := int64(0)

		if got != want {
			t.Fatal(got, want)
		}
	})

	t.Run("custom value", func(t *testing.T) {
		want := int64(1337)

		err = db.SetEndtime(want)
		if err != nil {
			t.Fatal("SetEndtime failed: ", err)
		}

		got, err := db.GetEndtime()
		if err != nil {
			t.Fatal("getendtime failed: ", err)
		}

		if got != want {
			t.Fatal(got, want)
		}
	})
}

func TestBlocklist(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal("Failed to create temp dir:", err)
	}
	defer os.RemoveAll(tempDir)
	path := filepath.Join(tempDir, "db.db")

	//setup
	db, err := NewDB(path)
	if err != nil {
		t.Fatal("error:", err)
	}
	defer db.CloseDB()

	t.Run("default value", func(t *testing.T) {
		got, err := db.GetBlocklist()
		if err != nil {
			t.Fatal("getBlocklist failed: ", err)
		}

		want := []string{}

		if !reflect.DeepEqual(want, got) {
			t.Fatal(got, want)
		}
	})

	t.Run("custom value", func(t *testing.T) {
		want := []string{"a", "b", "c"}

		err = db.SetBlocklist(want)
		if err != nil {
			t.Fatal("SetBlocklist failed: ", err)
		}

		got, err := db.GetBlocklist()
		if err != nil {
			t.Fatal("getBlocklist failed: ", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatal(got, want)
		}
	})
}
