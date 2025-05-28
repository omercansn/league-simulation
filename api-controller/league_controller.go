package controller

import (
    "github.com/gin-gonic/gin"
    "league-simulation/service"
    "strconv"
    "net/http"
)

// GET /league-table
func GetLeagueTableHandler(c *gin.Context) {
    table := service.GetLeagueTable()
    c.JSON(http.StatusOK, table)
}

// POST /simulate-week/:week
func SimulateWeekHandler(c *gin.Context) {
    weekString := c.Param("week")
    week, error := strconv.Atoi(weekString)
    if error!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week number"})
        return
    }
    results := service.SimulateWeek(week)
    c.JSON(http.StatusOK, results)
}

// POST /create-fixture
func CreateFixtureHandler(c *gin.Context) {
    service.CreateFixture()
    c.JSON(http.StatusCreated, gin.H{"message": "Fixture created!"})
}
