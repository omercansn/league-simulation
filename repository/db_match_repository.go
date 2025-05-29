package repository

import (
	"league-simulation/entities"
	"log"
)

type DBMatchRepository struct{}

func (r *DBMatchRepository) GetAllMatches() []entities.Match {
	rows, err := DB.Query("SELECT id, week, hometeamid, awayteamid, homegoals, awaygoals, played FROM matches")
	if err != nil {
		log.Println("GetAllMatches error:", err)
		return nil
	}
	defer rows.Close()

	var matches []entities.Match
	for rows.Next() {
		var m entities.Match
		err = rows.Scan(&m.ID, &m.Week, &m.HomeTeamID, &m.AwayTeamID, &m.HomeGoals, &m.AwayGoals, &m.Played)
		if err == nil {
			matches = append(matches, m)
		}
	}
	return matches
}

func (r *DBMatchRepository) AddMatch(match entities.Match) error {
	_, err := DB.Exec("INSERT INTO matches (week, hometeamid, awayteamid, homegoals, awaygoals, played) VALUES ($1, $2, $3, $4, $5, $6)",
		match.Week, match.HomeTeamID, match.AwayTeamID, match.HomeGoals, match.AwayGoals, match.Played)
	return err
}

func (r *DBMatchRepository) UpdateMatch(matchToUpdate *entities.Match) error {
	_, err := DB.Exec(
		"UPDATE matches SET week=$1, hometeamid=$2, awayteamid=$3, homegoals=$4, awaygoals=$5, played=$6 WHERE id=$7",
		matchToUpdate.Week, matchToUpdate.HomeTeamID, matchToUpdate.AwayTeamID, matchToUpdate.HomeGoals, matchToUpdate.AwayGoals, matchToUpdate.Played, matchToUpdate.ID)
	return err
}

func (r *DBMatchRepository) FindMatchByID(id int) *entities.Match {
	row := DB.QueryRow("SELECT id, week, hometeamid, awayteamid, homegoals, awaygoals, played FROM matches WHERE id=$1", id)
	var m entities.Match
	err := row.Scan(&m.ID, &m.Week, &m.HomeTeamID, &m.AwayTeamID, &m.HomeGoals, &m.AwayGoals, &m.Played)
	if err != nil {
		return nil
	}
	return &m
}
