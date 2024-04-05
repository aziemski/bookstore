package main

import (
	"context"
	"embed"
	"flag"
	"github.com/aziemski/bookstore/internal/core"
	"github.com/aziemski/bookstore/internal/core/ent"
	"github.com/aziemski/bookstore/internal/core/fixures"
	"github.com/aziemski/bookstore/internal/rest"
	"github.com/aziemski/bookstore/internal/web"
	"github.com/aziemski/bookstore/internal/x/xecho/xmiddleware"
	"github.com/aziemski/bookstore/internal/x/xlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed static/css/*.css
var cssDir embed.FS

func main() {
	address := flag.String("a", "localhost:8080", "server:port")
	flag.Parse()

	log := xlog.GetLogger()

	e := echo.New()
	e.StaticFS("/", cssDir)

	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(xmiddleware.RequestLogger(log))

	// dbClient, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	dbClient, err := ent.Open("sqlite3", "file:./bookstore.db?_fk=1")
	if err != nil {
		log.Error("unexpected db.Open err", xlog.Err(err))
		panic(err)
	}

	if err = dbClient.Schema.Create(context.Background()); err != nil {
		log.Error("unexpected db.Schema.Create err", xlog.Err(err))
		panic(err)
	}

	repo := core.NewRepository(dbClient)

	restServer := rest.NewServer(repo)
	restServer.RegisterWith(e)

	go fixures.Populate(repo, log)

	web.SetupRoutes(e, repo)

	e.Logger.Fatal(e.Start(*address))
}
