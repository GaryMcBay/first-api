package garyapiNEW

type Game struct {
	ID        int64  `json:"id"			db:"game_id"`
	Title     string `json:"title"		db:"title"`
	Developer string `json:"developer" 	db:"developer"`
	Publisher string `json:"publisher" 	db:"publisher"`
}

type GameCreate struct {
	Title     string `json:	"title, omitempty" `
	Developer string `json:	"developer, omitempty"`
	Publisher string `json:	"publisher, omitempty"`
}

type GameService interface {
	//Enforce rules to Game operations
	Game(id int64) (*Game, error)
	Games() ([]Game, error)
	Create(req GameCreate) (*Game, error)
	Delete(id int64) error
}

type GameStore interface {
	GameService
}
