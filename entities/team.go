package entities

type Team struct {
	ID      int
	Name    string
	Strength int
	MatchesPlayed int
	MatchesWon int
	MatchesDrawn int
	MatchesLost int
	GoalsFor int
	GoalsAgainst int
	GoalDifference int
	Points  int
}

