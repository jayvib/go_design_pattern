package store

type from int

const (
	MemoryDB from = iota
	RedisDB
	MysqlDB
)

var DBCloner = dbcloner{}

type dbcloner struct{}

func (dbcloner) GetClone(f from) GetStater {
	switch f {
	case MemoryDB:
		newDB := *MemoryDatabase
		return &newDB
	case RedisDB:
		return &(*RedisDatabase)
	case MysqlDB:
		return &(*MysqlDatabase)
	}
	return nil
}
