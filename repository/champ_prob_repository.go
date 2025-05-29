// repository/champion_probability_repository.go
package repository

import (
    "league-simulation/entities"

)

type ChampionProbabilityRepository interface {
    SaveChampionProbability(prob entities.ChampionProbability) error
    GetChampionProbabilities(season int) ([]entities.ChampionProbability, error)
    ResetChampionProbabilities(season int) error
}

type DBChampionProbabilityRepository struct{}

func (r *DBChampionProbabilityRepository) SaveChampionProbability(prob entities.ChampionProbability) error {
    _, err := DB.Exec(
        "INSERT INTO champion_probabilities (team_id, season, probability, calculated_at) VALUES ($1, $2, $3, $4)",
        prob.TeamID, prob.Season, prob.Probability, prob.CalculatedAt)
    return err
}

func (r *DBChampionProbabilityRepository) GetChampionProbabilities(season int) ([]entities.ChampionProbability, error) {
    rows, err := DB.Query("SELECT id, team_id, season, probability, calculated_at FROM champion_probabilities WHERE season = $1", season)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var probs []entities.ChampionProbability
    for rows.Next() {
        var prob entities.ChampionProbability
        err = rows.Scan(&prob.ID, &prob.TeamID, &prob.Season, &prob.Probability, &prob.CalculatedAt)
        if err == nil {
            probs = append(probs, prob)
        }
    }
    return probs, nil
}

func (r *DBChampionProbabilityRepository) ResetChampionProbabilities(season int) error {
    _, err := DB.Exec("DELETE FROM champion_probabilities WHERE season = $1", season)
    return err
}
