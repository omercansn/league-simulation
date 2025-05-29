package repository

// this is used for championship prediction in order not to affect the real data.
// this repository copies the real data and is used in championship prediction service function.

import (
	"fmt"
	"league-simulation/entities"
)

type MemoryTeamRepo struct {
	teams []entities.Team
}

func NewMemoryTeamRepo(src []entities.Team) *MemoryTeamRepo {
	cpy := make([]entities.Team, len(src))
	copy(cpy, src)
	return &MemoryTeamRepo{teams: cpy}
}

// returns all teams in memory
func (m *MemoryTeamRepo) GetAllTeams() []entities.Team {
	return m.teams
}

// finds a team by its id
func (m *MemoryTeamRepo) FindTeamByID(id int) *entities.Team {
	for i := range m.teams {
		if m.teams[i].ID == id {
			return &m.teams[i]
		}
	}
	return nil
}

// finds a team by its name
func (m *MemoryTeamRepo) FindTeamByName(name string) *entities.Team {
	for i := range m.teams {
		if m.teams[i].Name == name {
			return &m.teams[i]
		}
	}
	return nil
}

// adds a new team to memory
func (m *MemoryTeamRepo) AddTeam(team entities.Team) error {
	for _, t := range m.teams {
		if t.ID == team.ID {
			return fmt.Errorf("team already exists")
		}
	}
	m.teams = append(m.teams, team)
	return nil
}
// updates the scores andthe other variables
func (m *MemoryTeamRepo) UpdateTeam(team *entities.Team) error {
	for i := range m.teams {
		if m.teams[i].ID == team.ID {
			m.teams[i] = *team
			return nil
		}
	}
	return fmt.Errorf("team not found")
}
