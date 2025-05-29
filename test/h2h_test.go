package test

import (
	"league-simulation/entities"
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)


func TestHeadToHead(t *testing.T) {
	a := entities.Team{ID: 1, Name: "A"}
	b := entities.Team{ID: 2, Name: "B"}
	matches := []entities.Match{
		{ID: 1, HomeTeamID: 1, AwayTeamID: 2, HomeGoals: 2, AwayGoals: 1, Played: true},
		{ID: 2, HomeTeamID: 2, AwayTeamID: 1, HomeGoals: 1, AwayGoals: 1, Played: true},
	}
	h2h := service.HeadToHead(a, b, matches)
	if h2h == 0 {
		t.Errorf("Head-to-head calculation seems incorrect")
	}
	utils.LogTestPassed(t)
}