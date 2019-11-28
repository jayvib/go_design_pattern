package user

import "errors"

// UserFinder is the interface that the database and the Proxy implement.
type UserFinder interface {
	FindUser(id int) (User, error)
}

type User struct {
	ID int
}

type UserList []User

func (l *UserList) FindUser(id int) (User, error) {
	return User{}, errors.New("not implemented yet")
}

type UserListProxy struct {
	SomeDatabase           *UserList
	StackCache             *UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int) (User, error) {
	return User{}, errors.New("not implemented yet")
}
