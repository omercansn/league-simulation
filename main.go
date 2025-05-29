package main

import (
	"league-simulation/api-controller"
	"league-simulation/repository"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    repository.InitDB("host=localhost port=5432 user=postgres password=6915 dbname=insider_premierleaguesimulation sslmode=disable")
    r.GET("/teams", controller.GetAllTeamsHandler)
    r.GET("/fixture", controller.GetFullFixtureHandler)
    r.GET("/league-table", controller.GetLeagueTableHandler)
    r.POST("/simulate-week/:week", controller.SimulateWeekHandler)
    r.POST("/create-fixture", controller.CreateFixtureHandler)
    r.GET("/champion-probabilities", controller.GetChampionProbabilitiesHandler)
    r.POST("/reset-teams", controller.ResetTeamsHandler)
    r.Run(":8080")
}

