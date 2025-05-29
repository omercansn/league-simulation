
package repository

import (
	"league-simulation/entities"
	"log"
)

type DBTeamRepository struct{}

func (r *DBTeamRepository) GetAllTeams() []entities.Team {
	rows, err := DB.Query("SELECT id, name, strength, matchesplayed, matcheswon, matchesdrawn, matcheslost, goalsfor, goalsagainst, goaldifference, points FROM teams")
	if err != nil {
		log.Println("GetAllTeams error:", err)
		return nil
	}
	defer rows.Close()

	var teams []entities.Team
	for rows.Next() {
		var t entities.Team
		err = rows.Scan(&t.ID, &t.Name, &t.Strength, &t.MatchesPlayed, &t.MatchesWon, &t.MatchesDrawn, &t.MatchesLost, &t.GoalsFor, &t.GoalsAgainst, &t.GoalDifference, &t.Points)
		if err == nil {
			teams = append(teams, t)
		}
	}
	log.Printf("DEBUG: GetAllTeams found %d teams", len(teams))
	return teams
}

func (r *DBTeamRepository) FindTeamByID(id int) *entities.Team {
	row := DB.QueryRow("SELECT id, name, strength, matchesplayed, matcheswon, matchesdrawn, matcheslost, goalsfor, goalsagainst, goaldifference, points FROM teams WHERE id = $1", id)
	var t entities.Team
	err := row.Scan(&t.ID, &t.Name, &t.Strength, &t.MatchesPlayed, &t.MatchesWon, &t.MatchesDrawn, &t.MatchesLost, &t.GoalsFor, &t.GoalsAgainst, &t.GoalDifference, &t.Points)
	if err != nil {
		return nil
	}
	return &t
}

func (r *DBTeamRepository) FindTeamByName(name string) *entities.Team {
	row := DB.QueryRow("SELECT id, name, strength, matchesplayed, matcheswon, matchesdrawn, matcheslost, goalsfor, goalsagainst, goaldifference, points FROM teams WHERE name = $1", name)
	var t entities.Team
	err := row.Scan(&t.ID, &t.Name, &t.Strength, &t.MatchesPlayed, &t.MatchesWon, &t.MatchesDrawn, &t.MatchesLost, &t.GoalsFor, &t.GoalsAgainst, &t.GoalDifference, &t.Points)
	if err != nil {
		return nil
	}
	return &t
}

func (r *DBTeamRepository) AddTeam(team entities.Team) error {
	_, err := DB.Exec("INSERT INTO team (name, strength) VALUES ($1, $2)", team.Name, team.Strength)
	return err
}

func (r *DBTeamRepository) UpdateTeam(team *entities.Team) error {
	_, err := DB.Exec(
		`UPDATE teams SET 
		 matchesplayed = $1, matcheswon = $2, matchesdrawn = $3, matcheslost = $4,
		 goalsfor = $5, goalsagainst = $6, goaldifference = $7, points = $8
		 WHERE id = $9`,
		team.MatchesPlayed, team.MatchesWon, team.MatchesDrawn, team.MatchesLost,
		team.GoalsFor, team.GoalsAgainst, team.GoalDifference, team.Points, team.ID)
	return err
}