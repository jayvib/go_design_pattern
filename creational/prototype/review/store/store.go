package store

var MemoryDatabase = &Database{Getter: &memoryGetter{}}
var RedisDatabase = &Database{Getter: &redisGetter{}}
var MysqlDatabase = &Database{Getter: &mysqlGetter{}}

type Getter interface {
	Get() string
}

type Stater interface {
	SetState(str string)
	GetState() string
}

type GetStater interface {
	Getter
	Stater
}

type Database struct {
	Getter
	state string // just something to modify
}

func (db *Database) SetState(str string) {
	db.state = str
}

func (db *Database) GetState() string {
	return db.state
}

type memoryGetter struct{}

func (m *memoryGetter) Get() string {
	return ""
}

type redisGetter struct{}

func (r *redisGetter) Get() string {
	return ""
}

type mysqlGetter struct{}

func (r *mysqlGetter) Get() string {
	return ""
}
