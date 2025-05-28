package repository

import "league-simulation/entities"

type TeamRepository interface {
	GetAllTeams() []entities.Team
	FindTeamByID(id int) *entities.Team
	FindTeamByName(name string) *entities.Team
	AddTeam(team entities.Team) error
}