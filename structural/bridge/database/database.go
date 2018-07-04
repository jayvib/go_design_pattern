package database

type cruder interface {
	Put(key, value string) error
	Post(key, value string) error
	Get(key string) (value string, err error)
	Delete(key string) error
}

type Database struct {
	db cruder
}

func (d Database) PutData(key, value string) error {
	return d.db.Put(key, value)
}

func (d Database) PostData(key, value string) error {
	return d.db.Post(key, value)
}

func (d Database) GetData(key string) (value string, err error) {
	return d.db.Get(key)
}

func (d Database) DeleteData(key string) error {
	return d.db.Delete(key)
}

