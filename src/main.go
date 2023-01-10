package main

import (
	"flag"
	"fmt"
	"net/http"

	"api_client/config"
	"api_client/infra/database"
	myMiddleware "api_client/infra/middleware"

	"api_client/infra/router"
	"api_client/registry"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	configPath := flag.String("configPath", ".", "path to the config file")
	flag.Parse()

	config.ReadConfig(*configPath)
	// init db
	db := database.NewDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	client := new(http.Client)
	r := registry.NewRegistry(db, client)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())
	e.HTTPErrorHandler = myMiddleware.JSONErrorHandler

	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		e.Logger.Fatal(err)
	}
}
