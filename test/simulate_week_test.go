package test

import (
	"league-simulation/entities"
	"league-simulation/repository"
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)


func TestSimulateWeek(t *testing.T) {
	teamRepo := repository.NewMemoryTeamRepo(utils.TestTeams())
	matchRepo := repository.NewMemoryMatchRepo([]entities.Match{})
	service.CreateFixture(teamRepo, matchRepo)
	matches := service.SimulateWeek(1, teamRepo, matchRepo)
	if len(matches) == 0 {
		t.Errorf("SimulateWeek did not play any matches")
	}
	utils.LogTestPassed(t)
}