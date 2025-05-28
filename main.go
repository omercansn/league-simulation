package main

import (
	"league-simulation/api-controller"
	"league-simulation/repository"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    repository.InitDB("host=localhost port=5432 user=postgres password=6915 dbname=insider_premierleaguesimulation sslmode=disable")
    r.GET("/league-table", controller.GetLeagueTableHandler)
    r.POST("/simulate-week/:week", controller.SimulateWeekHandler)
    r.POST("/create-fixture", controller.CreateFixtureHandler)

    // Eğer başka endpointler ekleyeceksen burada devam edebilirsin.

    r.Run(":8080") // Sunucu 8080 portundan başlar
}

