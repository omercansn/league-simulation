package test

import (
	"league-simulation/entities"
	"league-simulation/repository"
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)


func TestGetCurrentWeekAndTotalWeeks(t *testing.T) {
	matches := []entities.Match{
		{ID: 1, Week: 1, Played: true},
		{ID: 2, Week: 2, Played: true},
		{ID: 3, Week: 3, Played: false},
	}
	matchRepo := repository.NewMemoryMatchRepo(matches)
	curr := service.GetCurrentWeek(matchRepo)
	total := service.GetTotalWeeks(matchRepo)
	if curr != 2 || total != 3 {
		t.Errorf("Week calculation error: got curr=%d total=%d", curr, total)
	}
	utils.LogTestPassed(t)
}