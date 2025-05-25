package repository

import "league-simulation/entities"

var matches []entities.Match

func GetAllMatches() []entities.Match {
	return matches
}

func AddMatch(match entities.Match){
	matches = append(matches, match)
}

func UpdateMatch(matchToUpdate *entities.Match){
	for i, m := range matches {
		if m.ID == matchToUpdate.ID {
			matches[i] = *matchToUpdate
			return
		}
	}
}

func FindMatchByID(id int) *entities.Match {
	for i := range matches {
		if matches[i].ID == id {
			return &matches[i]
		}
	}
	return nil
}

