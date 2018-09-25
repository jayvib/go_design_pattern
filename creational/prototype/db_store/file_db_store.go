package db_store

import (
	"encoding/json"
	"errors"
	"os"
)

func NewFileDBStore(filename string) (*FileDBStore, error) {
	store := &FileDBStore{
		filename: filename,
		Items:    make([]*Person, 0),
	}
	f, err := os.Open(filename)
	if err != nil {
		if !os.IsExist(err) {
			return store, nil
		}
		return nil, err
	}

	err = json.NewDecoder(f).Decode(store)
	if err != nil {
		return nil, err
	}
	return store, nil
}

// FileDBStore is a file storage that implements the store interface
type FileDBStore struct {
	filename string
	Items    []*Person
}

func (f *FileDBStore) Save(p *Person) error {
	return errors.New("not implemented yet")
}

func (f *FileDBStore) FindByFirstName(firstname string) (*Person, error) {
	return nil, errors.New("not implemented yet")
}

func (f *FileDBStore) Clone() Store {
	return nil
}