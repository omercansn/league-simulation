package repository

import "league-simulation/entities"

var teams = []entities.Team{
	{ID: 1, Name: "Team A", Strength: 80, MatchesPlayed: 0, MatchesWon: 0, MatchesDrawn: 0, MatchesLost: 0, GoalsFor: 0, GoalsAgainst: 0, GoalDifference: 0, Points: 0},
	{ID: 2, Name: "Team B", Strength: 65, MatchesPlayed: 0, MatchesWon: 0, MatchesDrawn: 0, MatchesLost: 0, GoalsFor: 0, GoalsAgainst: 0, GoalDifference: 0, Points: 0},
	{ID: 3, Name: "Team C", Strength: 71, MatchesPlayed: 0, MatchesWon: 0, MatchesDrawn: 0, MatchesLost: 0, GoalsFor: 0, GoalsAgainst: 0, GoalDifference: 0, Points: 0},
	{ID: 4, Name: "Team D", Strength: 59, MatchesPlayed: 0, MatchesWon: 0, MatchesDrawn: 0, MatchesLost: 0, GoalsFor: 0, GoalsAgainst: 0, GoalDifference: 0, Points: 0},
}

func GetAllTeams() []entities.Team {
	return teams
}

func FindTeamByID(id int) *entities.Team {
	for _, team := range teams {
		if team.ID == id {
			return &team
		}
	}
	return nil
}

func FindTeamByName(name string) *entities.Team {
	for _, team := range teams {
		if team.Name == name {
			return &team
		}
	}
	return nil
}
