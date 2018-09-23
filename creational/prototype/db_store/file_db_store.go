package db_store

import "errors"

func NewFileDBStore(filename string) *FileDBStore {
	return nil
}

// FileDBStore is a file storage that implements the store interface
type FileDBStore struct {
	filename string
	items []*Person
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