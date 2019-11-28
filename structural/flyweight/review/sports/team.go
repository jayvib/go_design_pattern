package sports

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type team int

const (
	TEAMA team = iota
	TEAMB
)

type Team struct {
	ID             team
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

// What's the gatcha here?? ...its memory synchronization
type teamFlyweightFactory struct {
	// RWMutex will block the RLock() if someone called
	// Lock().....
	// Multiple RLock() call will not block.
	mu           sync.RWMutex
	createdTeams map[team]*Team // caching the pool
}

func (t *teamFlyweightFactory) GetTeam(tm team) *Team {
	// check if the team already exist

	logrus.Println("Getting the cached...")
	t.mu.RLock()
	cachedTeam, ok := t.createdTeams[tm]
	if ok {
		logrus.Println("Returning from cached")
		return cachedTeam
	}
	t.mu.RUnlock()

	logrus.Println("opppsss.. not cached yet")
	// Create an instance
	t.mu.Lock()
	// simulate initialization
	logrus.Println("Initializing....")
	logrus.Println("blocking the reader")
	time.Sleep(5 * time.Second)
	newTeam := getTeamFactory(tm)
	t.createdTeams[tm] = newTeam
	t.mu.Unlock()
	return newTeam
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(tm team) *Team {
	switch tm {
	case TEAMA:
		return &Team{
			ID:   tm,
			Name: "Team A",
		}
	default:
		return &Team{
			ID:   tm,
			Name: "Team B",
		}
	}
}
