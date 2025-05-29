package repository

// this is used for championship prediction in order not to affect the real data.
// this repository copies the real data and is used in championship prediction service function.
import (
	"fmt"
	"league-simulation/entities"
)

type MemoryMatchRepo struct {
	matches []entities.Match
}

func NewMemoryMatchRepo(src []entities.Match) *MemoryMatchRepo {
	cpy := make([]entities.Match, len(src))
	copy(cpy, src)
	return &MemoryMatchRepo{matches: cpy}
}
func (m *MemoryMatchRepo) GetAllMatches() []entities.Match { return m.matches }
func (m *MemoryMatchRepo) FindMatchByID(id int) *entities.Match {
	for i := range m.matches { if m.matches[i].ID == id { return &m.matches[i] } }
	return nil
}
func (m *MemoryMatchRepo) AddMatch(match entities.Match) error {
	m.matches = append(m.matches, match)
	return nil
}
func (m *MemoryMatchRepo) UpdateMatch(match *entities.Match) error {
	for i := range m.matches {
		if m.matches[i].ID == match.ID { m.matches[i] = *match; return nil }
	}
	return fmt.Errorf("match not found")
}