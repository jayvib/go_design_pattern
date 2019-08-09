package sports

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamFlyweightFactory(t *testing.T) {

	t.Run("GetTeam", func(t *testing.T) {
		t.SkipNow()
		factory := &teamFlyweightFactory{
			createdTeams: make(map[team]*Team),
		}
		teamA1 := factory.GetTeam(TEAMA)
		assert.False(t, teamA1 == nil, "team instance should not be nil")

		teamA2 := factory.GetTeam(TEAMA)
		assert.False(t, teamA2 == nil, "team instance should not be nill")

		assert.Equal(t, teamA1, teamA2, "They should be equal pointer")
		assert.Equal(t, 1, factory.GetNumberOfObjects(), "expecting number of object to be 1")
	})

	t.Run("GetTeam - High Volume", func(t *testing.T) {
		factory := &teamFlyweightFactory{
			createdTeams: make(map[team]*Team),
		}

		initializer := func(wg *sync.WaitGroup) {
			defer wg.Done()
			factory.GetTeam(TEAMA)
		}

		var wg sync.WaitGroup
		num := 500
		for i := 0; i < num; i++ {
			wg.Add(1)
			go initializer(&wg)
		}

		wg.Wait()
		assert.Equal(t, 1, factory.GetNumberOfObjects(), "expecting number of object to be 1")
	})

}
