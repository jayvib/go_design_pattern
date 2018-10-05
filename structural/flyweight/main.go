package main

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type Team struct {
	ID             uint64
	Name           uint64
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date         time.Time
	VisitorID    uint64
	LocalID      uint64
	LocalScore   byte
	VisitorScore byte
	LocalShoots  uint16
	VisitorShoots uint16
}

func NewTeamFactory() *teamFlyweightFactory {
	return &teamFlyweightFactory{
		createdTeams: make(map[uint64]*Team, 0),
	}
}

// uses factory pattern for the initialization of object
type teamFlyweightFactory struct {
	createdTeams map[uint64]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamID uint64) *Team {
	if t.createdTeams[teamID] != nil { // this is not thread safe.. need to use sync.Once
		return t.createdTeams[teamID]
	}

	team := getTeamFactory(teamID) // when
	t.createdTeams[teamID] = &team // haven't yet initialize

	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(teamID uint64) Team {
	switch teamID {
	case TEAM_A:
		return Team{
			ID: 1,
			Name: TEAM_A,
		}
	default:
		return Team{
			ID: 2,
			Name: TEAM_B,
		}
	}
}
