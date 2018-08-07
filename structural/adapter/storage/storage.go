package storage

type Storage interface {
	Store(text string) error
	Search(text string) string
}

// existing storage interface
type DB interface {
	Put(text string) error
	Find(text string) string
}

func newMockDB() *mockDB {
	return &mockDB{
		data: make(map[string]string),
	}
}

// type that implement the DB
type mockDB struct {
	data map[string]string
}

func (db *mockDB) Put(text string) error {
	db.data[text] = text
	return nil
}

func (db *mockDB) Find(text string) string {
	text, ok := db.data[text]
	if !ok {
		return ""
	}
	return text
}


func newLatestStorage(db DB) *LatestStorage {
	return &LatestStorage{
		db: db,
	}
}

// LatestStorage uses adapter design pattern to make the old interface satisfy with the new interface for storage.
type LatestStorage struct {
	db DB
}

func (s *LatestStorage) Store(text string) error {
	return s.db.Put(text)
}

func (s *LatestStorage) Search(text string) string {
	return s.db.Find(text)
}

func saveText(store Storage, text string) error {
	return store.Store(text)
}

func findText(store Storage, text string) string {
	return store.Search(text)
}