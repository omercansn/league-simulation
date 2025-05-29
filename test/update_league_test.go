package test

import (
	"league-simulation/entities"
	"league-simulation/repository"
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)

func TestUpdateLeagueStatistics(t *testing.T) {
	teamRepo := repository.NewMemoryTeamRepo(utils.TestTeams())
	teamA := &entities.Team{ID: 1, Name: "A", Strength: 80}
	teamB := &entities.Team{ID: 2, Name: "B", Strength: 70}
	match := &entities.Match{ID: 1, HomeGoals: 2, AwayGoals: 1, Played: true}
	service.UpdateLeagueStatistics(match, teamA, teamB, teamRepo)
	if teamA.Points != 3 || teamB.Points != 0 {
		t.Errorf("Points not updated correctly: got %d and %d", teamA.Points, teamB.Points)
	}
	utils.LogTestPassed(t)
}
