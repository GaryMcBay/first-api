package service

import (
	"errors"
	"strings"

	"github.com/garymcbay/garyapiNEW"
)

type GameService struct {
	repo garyapiNEW.GameStore
}

func NewGameService(repo garyapiNEW.GameStore) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (g *GameService) Create(req garyapiNEW.GameCreate) (*garyapiNEW.Game, error) {
	// validate request & return error if any of the properties are empty
	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("title cannot be empty")
	}
	if strings.TrimSpace(req.Developer) == "" {
		return nil, errors.New("developer cannot be empty")
	}
	if strings.TrimSpace(req.Publisher) == "" {
		return nil, errors.New("publisher cannot be empty")
	}
	return g.repo.Create(req)
}

func (g *GameService) Delete(id int64) error {
	return g.repo.Delete(id)
}

func (g *GameService) Game(id int64) (*garyapiNEW.Game, error) {
	return g.repo.Game(id)
}

func (g *GameService) Games() ([]garyapiNEW.Game, error) {
	return g.repo.Games()
}
