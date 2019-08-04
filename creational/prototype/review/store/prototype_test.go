package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBCloner(t *testing.T) {
	t.Run("Memory DB", func(t *testing.T) {
		memDb := DBCloner.GetClone(MemoryDB)

		// Check if the pointer is not the same.
		// Expected: Pointer should not be the same.
		t.Logf("%p %p\n", memDb, MemoryDatabase)
		assert.False(t, memDb == MemoryDatabase, "Pointer should not be equal")

		// Check if the value is from memory database
		// Expected: Return has a word "memory"
		result := memDb.Get()
		t.Log(result)
		assert.Contains(t, result, "memory", "Result expected to have a word memory")

		memDb.SetState("good")
		t.Log(memDb.GetState())
		memDb2 := DBCloner.GetClone(MemoryDB)
		assert.NotEqual(t, memDb.GetState(), memDb2.GetState(), "State should be equal got %s want %s",
			memDb2.GetState(), memDb.GetState())
	})
}
