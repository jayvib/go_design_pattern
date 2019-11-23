package redis

import (
	"github.com/jayvib/golog"
	"github.com/stretchr/testify/assert"
	"testing"
)

type person struct {
	FirstName string `redis:"firstname"`
	LastName  string `redis:"lastname"`
	Age       int8   `redis:"age"`
	Price float64 `redis:"price"`
}

func TestObjectMapper(t *testing.T) {
	golog.SetLevel(golog.DebugLevel)
	t.Run("string field", func(t *testing.T) {
		obj := struct {
			FirstName string `redis:"firstname"`
			LastName  string `redis:"lastname"`
		}{
			FirstName: "Luffy",
			LastName:  "Monkey",
		}

		want := map[string]interface{}{
			"firstname": "Luffy",
			"lastname":  "Monkey",
		}

		got := ObjectMapper(obj)
		assert.Equal(t, want, got)
	})

	t.Run("int8 type in the field", func(t *testing.T) {
		obj := struct {
			FirstName string `redis:"firstname"`
			LastName  string `redis:"lastname"`
			Age       int8   `redis:"age"`
		}{
			FirstName: "Luffy",
			LastName:  "Monkey",
			Age:       20,
		}

		want := map[string]interface{}{
			"firstname": "Luffy",
			"lastname":  "Monkey",
			"age":       "20",
		}

		got := ObjectMapper(obj)
		assert.Equal(t, want, got)
	})


	t.Run("when an object is a pointer", func(t *testing.T) {
		obj := &person{
			FirstName: "Luffy",
			LastName:  "Monkey",
			Age:       20,
			Price: 1,
		}

		want := map[string]interface{}{
			"firstname": "Luffy",
			"lastname":  "Monkey",
			"age":       "20",
			"price": "1.00",
		}

		got := ObjectMapper(obj)
		assert.Equal(t, want, got)
	})

	t.Run("float field", func(t *testing.T) {
		obj := struct {
			FirstName string  `redis:"firstname"`
			LastName  string  `redis:"lastname"`
			Age       int8    `redis:"age"`
			Price     float64 `redis:"price"`
		}{
			FirstName: "Luffy",
			LastName:  "Monkey",
			Age:       20,
			Price:     20.01,
		}

		want := map[string]interface{}{
			"firstname": "Luffy",
			"lastname":  "Monkey",
			"age":       "20",
			"price":     "20.01",
		}

		got := ObjectMapper(obj)
		assert.Equal(t, want, got)
	})
}

func TestMapParser(t *testing.T) {
	input := map[string]string{
		"firstname": "Luffy",
		"lastname":  "Monkey",
		"age":       "20",
		"price":     "20.01",
	}

	var p person
	MapParser(input, &p)
	want := person{
		FirstName: "Luffy",
		LastName:  "Monkey",
		Age:       20,
		Price: 20.01,
	}

	assert.Equal(t, want, p)
	t.Logf(p.FirstName)
}
