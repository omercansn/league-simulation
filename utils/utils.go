package utils

import (
	"league-simulation/entities"
	"math/rand"
	"time"
)

// a helper function for shuffling teams so that every match order will be changed
func ShuffleTeams(teams []entities.Team) []entities.Team {
	// a unique seed for shuffling
    rand.New(rand.NewSource(time.Now().UnixNano()))
    shuffled := make([]entities.Team, len(teams))
    copy(shuffled, teams)
    rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
    return shuffled
}


