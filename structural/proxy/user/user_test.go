package user

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	someDatabase := &UserList{}

	var knownIDs []int32

	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		user := User{ ID:n }
		if len(knownIDs) < 3 {
			knownIDs = append(knownIDs, n)
		}
		someDatabase.Add(user)
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCapacity: 2,
		StackCache:    &UserList{},
	}



	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if proxy.StackCache.Len() != 1 {
			t.Error("After one successful search in an empty cache, the size it must be one")
		}

		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if proxy.StackCache.Len() != 1 {
			t.Error("Cache must not grow if we asked for an object that is stored on it")
		}

		if !proxy.DidLastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	})

	user1, err := proxy.FindUser(knownIDs[0])
	if err != nil {
		t.Fatal(err)
	}

	user2, _ := proxy.FindUser(knownIDs[1])
	if proxy.DidLastSearchUsedCache {
		t.Error("The user wasn't stored on the proxy cache yet")
	}

	user3, _ := proxy.FindUser(knownIDs[2])
	if proxy.DidLastSearchUsedCache {
		t.Error("The user wasn't stored on the proxy cache yet")
	}

	for i := 0; i < proxy.StackCache.Len(); i++ {
		if (*proxy.StackCache.(*UserList))[i].ID == user1.ID {
			t.Error("User that should be gone was found")
		}
	}

	if proxy.StackCache.Len() != 2 {
		t.Error("After inserting 3 users the cache should not grow more than two")
	}

	for _, v := range *proxy.StackCache.(*UserList) {
		if v != user2 && v != user3 {
			t.Error("A non expected user was found on the cache")
		}
	}
}
