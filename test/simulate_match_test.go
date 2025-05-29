package test

import (
	"league-simulation/entities"
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)

func TestSimulateMatch(t *testing.T) {
	teamA := entities.Team{ID: 1, Name: "A", Strength: 80}
	teamB := entities.Team{ID: 2, Name: "B", Strength: 70}
	match := entities.Match{ID: 1}
	service.SimulateMatch(&match, &teamA, &teamB)
	if !match.Played {
		t.Errorf("Match should be marked as played")
	}
	utils.LogTestPassed(t)
}

