package main

import "testing"

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := teamFlyweightFactory{}
	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	//teamA2 := factory.GetN
}