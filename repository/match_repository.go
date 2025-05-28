package repository

import (
    "league-simulation/entities"
    "log"
)

// GetAllMatches: Fetch all matches from the database
func GetAllMatches() []entities.Match {
    rows, err := DB.Query("SELECT id, week, home_team_id, away_team_id, home_goals, away_goals, played FROM match")
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

// AddMatch: Insert a new match
func AddMatch(match entities.Match) error {
    _, err := DB.Exec("INSERT INTO match (week, home_team_id, away_team_id, home_goals, away_goals, played) VALUES ($1, $2, $3, $4, $5, $6)",
        match.Week, match.HomeTeamID, match.AwayTeamID, match.HomeGoals, match.AwayGoals, match.Played)
    return err
}

// UpdateMatch: Update an existing match
func UpdateMatch(matchToUpdate *entities.Match) error {
    _, err := DB.Exec(
        "UPDATE match SET week=$1, home_team_id=$2, away_team_id=$3, home_goals=$4, away_goals=$5, played=$6 WHERE id=$7",
        matchToUpdate.Week, matchToUpdate.HomeTeamID, matchToUpdate.AwayTeamID, matchToUpdate.HomeGoals, matchToUpdate.AwayGoals, matchToUpdate.Played, matchToUpdate.ID)
    return err
}

// FindMatchByID: Find a match by ID
func FindMatchByID(id int) *entities.Match {
    row := DB.QueryRow("SELECT id, week, home_team_id, away_team_id, home_goals, away_goals, played FROM match WHERE id=$1", id)
    var m entities.Match
    err := row.Scan(&m.ID, &m.Week, &m.HomeTeamID, &m.AwayTeamID, &m.HomeGoals, &m.AwayGoals, &m.Played)
    if err != nil {
        return nil
    }
    return &m
}
