package test

import (
	"league-simulation/service"
	"league-simulation/utils"
	"testing"
)


func TestSimulateGoals(t *testing.T) {
	g := service.SimulateGoals(80)
	if g < 0 {
		t.Errorf("SimulateGoals returned negative value: %d", g)
	}
	utils.LogTestPassed(t)
}
