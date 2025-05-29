package test

import (
	"league-simulation/entities"
	"league-simulation/repository"
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)


func TestCreateFixture(t *testing.T) {
	teamRepo := repository.NewMemoryTeamRepo(utils.TestTeams())
	matchRepo := repository.NewMemoryMatchRepo([]entities.Match{})
	service.CreateFixture(teamRepo, matchRepo)
	matches := matchRepo.GetAllMatches()
	if len(matches) != service.GetTotalWeeks(matchRepo)*2 { 
		t.Errorf("Fixture creation failed, got %d matches", len(matches))
	}
	utils.LogTestPassed(t)
}
