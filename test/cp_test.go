package test

import (
	"league-simulation/entities"
	"league-simulation/repository"
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)

func TestChampionProbabilities(t *testing.T) {
	teamRepo := repository.NewMemoryTeamRepo(utils.TestTeams())
	matchRepo := repository.NewMemoryMatchRepo([]entities.Match{})
	service.CreateFixture(teamRepo, matchRepo)
	service.SimulateWeek(1, teamRepo, matchRepo)
	curr := service.GetCurrentWeek(matchRepo)
	total := service.GetTotalWeeks(matchRepo)
	probs := service.ChampionProbabilities(curr, total, teamRepo, matchRepo)
	if len(probs) != 4 {
		t.Errorf("Probabilities should be calculated for 4 teams, got %d", len(probs))
	}
	utils.LogTestPassed(t)
}
