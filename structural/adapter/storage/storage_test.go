package storage

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type anotherMockDB struct {
	mock.Mock
}

func (m *anotherMockDB) Find(text string) string {
	args := m.Called(text)
	return args.String(0)
}

func TestMockDB(t *testing.T) {
	testObj := new(anotherMockDB)
	testObj.On("Find", "test").Return("test")
	testObj.AssertExpectations(t)
}

func TestDB(t *testing.T) {
	var database DB
	t.Run("Old Database", func(t *testing.T) {
		mockDB := newMockDB()
		database = mockDB
		err := database.Put("test")
		if err != nil {
			t.Fatal(err.Error())
		}

		txt := database.Find("test")
		if txt == "" {
			t.Fatal("empty result")
		}

		if txt != "test" {
			t.Fatal("result from find doesn't matched")
		}
	})

	t.Run("New Database", func(t *testing.T) {

		latestStorage := newLatestStorage(newMockDB())
		err := latestStorage.Store("test1")
		if err != nil {
			t.Fatal(err.Error())
		}

		res := latestStorage.Search("test1")
		if res == "" {
			t.Fatal("empty result")
		}

		err = saveText(latestStorage, "test2")
		if err != nil {
			t.Fatal(err.Error())
		}

		res = findText(latestStorage, "test2")
		if res == "" {
			t.Fatal("empty result")
		}
	})

}
