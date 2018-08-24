package main

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type Team struct {
	ID             uint64
	Name           string
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

type teamFlyweightFactory struct {
	createdTeams map[uint64]*Team
}

func (t *teamFlyweightFactory) GetTeam(name uint64) *Team {
	return nil
}

func (t *teamFlyweightFactory) GetNumberOfObjects() uint64 {
	return 0
}
