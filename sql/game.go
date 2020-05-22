package garysql

import (
	"log"

	"github.com/garymcbay/garyapiNEW"
	"github.com/go-gorp/gorp"
)

type GameStore struct {
	db *gorp.DbMap
}

func NewGameStore(db *gorp.DbMap) *GameStore {
	return &GameStore{
		db: db,
	}
}

func (g *GameStore) Game(id int64) (*garyapiNEW.Game, error) {
	var game *garyapiNEW.Game
	if err := g.db.SelectOne(&game, `
	SELECT * FROM games 
	WHERE id=?`, id); err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameStore) Games() ([]garyapiNEW.Game, error) {
	var games []garyapiNEW.Game

	if _, err := g.db.Select(&games, `SELECT * FROM games ORDER BY id`); err != nil {
		return nil, err
	}
	return games, nil
}

func (g *GameStore) Create(req garyapiNEW.GameCreate) (*garyapiNEW.Game, error) {
	game := &garyapiNEW.Game{
		Publisher: req.Publisher,
		Developer: req.Developer,
		Title:     req.Title,
	}
	err := g.db.Insert(game)
	if err != nil {
		checkErr(err, "Insert row error")
	}
	return game, nil
}

func (g *GameStore) Delete(id int64) error {

	_, err := g.db.Exec(`
	DELETE FROM games 
	WHERE id = ?;`, id)
	if err != nil {
		checkErr(err, "Delete row error")
	}
	return nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
