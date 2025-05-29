package utils

import (
	"fmt"
	"league-simulation/entities"
	"math/rand"
	"path/filepath"
	"runtime"
	"testing"
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

// a helper function for using in the test functions
func TestTeams() []entities.Team {
	return []entities.Team{
		{ID: 1, Name: "Team A", Strength: 80},
		{ID: 2, Name: "Team B", Strength: 70},
		{ID: 3, Name: "Team C", Strength: 60},
		{ID: 4, Name: "Team D", Strength: 50},
	}
}

// a helper function for finding the test file name and print it if passed
func LogTestPassed(t *testing.T) {
    pc, file, _, _ := runtime.Caller(1)
    funcName := runtime.FuncForPC(pc).Name()
    fmt.Printf("[TEST] %s - %s: successfully passed\n", filepath.Base(file), filepath.Base(funcName))
}
