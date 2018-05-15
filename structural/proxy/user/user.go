package user

import (
	"errors"
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
	Cache(user User)
	Add(user User)
	Len() int
}

type User struct {
	ID int32
}

type UserList []User

func (ul *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*ul); i++ {
		if (*ul)[i].ID == id {
			return (*ul)[i], nil
		}
	}
	return User{}, errors.New("user not exist")
}

func (ul *UserList) Len() int {
	l := len(*ul)
	return l
}

func (ul *UserList) Cache(user User) {
	*ul = append((*ul)[1:], user)
}

func (ul *UserList) Add(user User) {
	*ul = append(*ul, user)
}

type UserListProxy struct {
	SomeDatabase           UserFinder
	StackCache             UserFinder
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache ID: ", id)
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}
	u.addUserToStack(user)
	u.DidLastSearchUsedCache = false
	return user, nil

}

func (u *UserListProxy) addUserToStack(user User) {
	if u.StackCache.Len() == u.StackCapacity {
		u.StackCache.Cache(user) // FIFO
		return
	}
	u.StackCache.Add(user)
}

