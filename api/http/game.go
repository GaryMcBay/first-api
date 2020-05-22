package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/gommon/log"

	"github.com/garymcbay/garyapiNEW"
	"github.com/labstack/echo"
)

type GameHandler struct {
	svc garyapiNEW.GameService
}

func NewGameHandler(svc garyapiNEW.GameService) *GameHandler {
	return &GameHandler{
		svc: svc,
	}
}

// add echo handlers
func (h *GameHandler) Routes(e *echo.Echo) {
	e.POST("/games", h.create)
	e.GET("/games/:id", h.game)
	e.GET("/games", h.games)
	e.DELETE("/games/:id", h.delete)

}

func (h GameHandler) create(c echo.Context) error {
	var req garyapiNEW.GameCreate
	err := c.Bind(&req)
	if err != nil {
		log.Printf("Failed to create game", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	resp, err := h.svc.Create(req)
	if err != nil {
		log.Printf("Failed unmarshaling in createGame: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h GameHandler) game(c echo.Context) error {

	gameID := c.Param("id")
	id, err := strconv.Atoi(gameID)
	if err != nil {
		log.Printf("Failed to get gameID", err)
		return c.String(http.StatusBadRequest, "")
	}
	resp, err := h.svc.Game(int64(id))
	if err != nil {
		log.Printf("Failed to get game", err)
		return c.String(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, resp)
}

func (h GameHandler) games(c echo.Context) error {

	games, err := h.svc.Games()

	if err != nil {
		log.Printf("Failed to get game", err)
		return c.String(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, games)
}

func (h GameHandler) delete(c echo.Context) error {
	gameID := c.Param("id")
	id, err := strconv.Atoi(gameID)
	if err != nil {
		log.Printf("Failed to get game", err)
		return c.String(http.StatusBadRequest, "")
	}
	err = h.svc.Delete(int64(id))
	if err != nil {
		log.Printf("Failed to delete game", err)
		return c.String(http.StatusInternalServerError, "")
	}
	return c.NoContent(http.StatusNoContent)
}
