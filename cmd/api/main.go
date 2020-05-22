package main

import (
	"database/sql"
	"log"

	garyapi "github.com/garymcbay/garyapiNEW"

	garysql "github.com/garymcbay/garyapiNEW/sql"

	garyhttp "github.com/garymcbay/garyapiNEW/api/http"
	"github.com/garymcbay/garyapiNEW/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"fmt"
	"net/http"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("yaaaaaaas")
	// setup db
	dbmap := initDB()
	defer dbmap.Db.Close()

	//Echo middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	}))

	//Handlers
	garyhttp.NewGameHandler(service.NewGameService(garysql.NewGameStore(dbmap))).Routes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "You've connected to games_db @localhost")
	})
	e.Start(":8000")
}

func initDB() *gorp.DbMap {
	//DB config
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/games_db")
	checkErr(err, "sql.Open failed")

	//construct dbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{
		Engine:   "InnoDB",
		Encoding: "utf8mb4",
	}}

	//Add table called Games

	dbmap.AddTableWithName(garyapi.Game{}, "games").SetKeys(true, "ID")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Creation of table failed")
	return dbmap

}

func checkErr(err error, msg string) {
	//Error check method
	if err != nil {
		log.Fatalln(msg, err)
	}
}
