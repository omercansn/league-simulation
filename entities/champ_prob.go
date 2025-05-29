package entities

import "time"

type ChampionProbability struct {
    ID           int
    TeamID       int
    Season       int
    Probability  float64
    CalculatedAt time.Time
}
