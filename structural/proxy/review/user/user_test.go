package user

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2342342)

	for i := 0; i < 100000; i++ {
		someDatabase = append(someDatabase, User{ ID: rand.Intn(31)})
	}

	proxy := UserListProxy{
		SomeDatabase: &someDatabase,
		StackCapacity: 2,
		StackCache: new(UserList),
	}

	knownIDs := [3]int{
		someDatabase[3].ID,
		someDatabase[4].ID,
		someDatabase[5].ID,
	}

	t.Run("FindUser - Empty cache", func(t *testing.T){
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Errorf("Expecting returned user id %d but got %d\n", knownIDs[0], user.ID)
		}

		if len(*proxy.StackCache) != 1 {
			t.Error("Afer successful search in an empty cache, the size of it must be one")
		}

		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T){
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Errorf("Expecting returned user id %d but got %d\n", knownIDs[0], user.ID)
		}

		if len(*proxy.StackCache) != 1 {
			t.Error("Afer successful search in an empty cache, the size of it must be one")
		}

		if !proxy.DidLastSearchUsedCache {
			t.Error("Cache must not grow if we asked for an object that is stored.")
		}
	})
}
