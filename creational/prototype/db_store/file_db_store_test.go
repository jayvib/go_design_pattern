package db_store

import "testing"

func TestFileDBStore(t *testing.T) {
	fileDbStore, err := NewFileDBStore("test.json")
	if err != nil {
		t.Fatal(err)
	}
	_ = fileDbStore
}
