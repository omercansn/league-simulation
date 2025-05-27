package main

import (
    "github.com/gin-gonic/gin"
    "league-simulation/api-controller"
)

func main() {
    r := gin.Default()

    r.GET("/league-table", controller.GetLeagueTableHandler)
    r.POST("/simulate-week/:week", controller.SimulateWeekHandler)
    r.POST("/create-fixture", controller.CreateFixtureHandler)

    // Eğer başka endpointler ekleyeceksen burada devam edebilirsin.

    r.Run(":8080") // Sunucu 8080 portundan başlar
}

