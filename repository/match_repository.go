package repository

import "league-simulation/entities"

type MatchRepository interface {
	GetAllMatches() []entities.Match
	AddMatch(match entities.Match) error
	UpdateMatch(matchToUpdate *entities.Match) error
	FindMatchByID(id int) *entities.Match
	
}