package test

import (
	"league-simulation/entities"
	"league-simulation/repository"
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)


func TestGetLeagueTable(t *testing.T) {
	teamRepo := repository.NewMemoryTeamRepo(utils.TestTeams())
	matchRepo := repository.NewMemoryMatchRepo([]entities.Match{})
	service.CreateFixture(teamRepo, matchRepo)
	service.SimulateWeek(1, teamRepo, matchRepo)
	table := service.GetLeagueTable(teamRepo, matchRepo)
	if len(table) != 4 {
		t.Errorf("League table should have 4 teams, got %d", len(table))
	}
	utils.LogTestPassed(t)
}