package controller

import (
	"league-simulation/repository"
	"league-simulation/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var teamRepo repository.TeamRepository = &repository.DBTeamRepository{}
var matchRepo repository.MatchRepository = &repository.DBMatchRepository{}

// GET /league-table
func GetLeagueTableHandler(c *gin.Context) {
	table := service.GetLeagueTable(teamRepo, matchRepo)
	c.JSON(http.StatusOK, table)
}

// POST /simulate-week/:week
func SimulateWeekHandler(c *gin.Context) {
	weekString := c.Param("week")
	week, err := strconv.Atoi(weekString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week number"})
		return
	}
	results := service.SimulateWeek(week, teamRepo, matchRepo)
	c.JSON(http.StatusOK, results)
}

// POST /create-fixture
func CreateFixtureHandler(c *gin.Context) {
	service.CreateFixture(teamRepo, matchRepo)
	c.JSON(http.StatusCreated, gin.H{"message": "Fixture created!"})
}

// GET /champion-probabilities
func GetChampionProbabilitiesHandler(c *gin.Context) {
	currentWeek := service.GetCurrentWeek(matchRepo)
	totalWeeks := service.GetTotalWeeks(matchRepo)
	probs := service.ChampionProbabilities(currentWeek, totalWeeks, teamRepo, matchRepo)
	c.JSON(http.StatusOK, probs)
}
