package repository

import (
    "league-simulation/entities"
    "log"
)

// GetAllTeams: Fetch all teams from the database
func GetAllTeams() []entities.Team {
    rows, err := DB.Query("SELECT id, name, strength, matches_played, matches_won, matches_drawn, matches_lost, goals_for, goals_against, goal_difference, points FROM team")
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
    return teams
}

// FindTeamByID: Find a team by ID
func FindTeamByID(id int) *entities.Team {
    row := DB.QueryRow("SELECT id, name, strength, matches_played, matches_won, matches_drawn, matches_lost, goals_for, goals_against, goal_difference, points FROM team WHERE id = $1", id)
    var t entities.Team
    err := row.Scan(&t.ID, &t.Name, &t.Strength, &t.MatchesPlayed, &t.MatchesWon, &t.MatchesDrawn, &t.MatchesLost, &t.GoalsFor, &t.GoalsAgainst, &t.GoalDifference, &t.Points)
    if err != nil {
        return nil
    }
    return &t
}

// FindTeamByName: Find a team by Name
func FindTeamByName(name string) *entities.Team {
    row := DB.QueryRow("SELECT id, name, strength, matches_played, matches_won, matches_drawn, matches_lost, goals_for, goals_against, goal_difference, points FROM team WHERE name = $1", name)
    var t entities.Team
    err := row.Scan(&t.ID, &t.Name, &t.Strength, &t.MatchesPlayed, &t.MatchesWon, &t.MatchesDrawn, &t.MatchesLost, &t.GoalsFor, &t.GoalsAgainst, &t.GoalDifference, &t.Points)
    if err != nil {
        return nil
    }
    return &t
}

// AddTeam: Insert new team
func AddTeam(team entities.Team) error {
    _, err := DB.Exec("INSERT INTO team (name, strength) VALUES ($1, $2)", team.Name, team.Strength)
    return err
}
