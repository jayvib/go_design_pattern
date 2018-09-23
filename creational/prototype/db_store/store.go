package db_store

const (
	FILE storage = 1
	MYSQL storage = 2
	MONGO  storage = 3
)


var globalStore Store

type Store interface {
	Save(*Person) error
	FindByFirstName(string) (*Person, error)
	Clone() Store
}

func NewStore(s storage) (Store, error) {
	// Factory Design Pattern
	switch s {
	case FILE:
		return nil, nil
	case MYSQL:
		return nil, nil
	case MONGO:
		return nil, nil
	default:
		return nil, nil
	}
}



