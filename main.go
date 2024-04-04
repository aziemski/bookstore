package main

import (
	"context"
	"flag"
	"log/slog"
	"time"

	"github.com/aziemski/bookstore/internal/core"
	"github.com/aziemski/bookstore/internal/core/ent"
	"github.com/aziemski/bookstore/internal/rest"
	"github.com/aziemski/bookstore/internal/web"
	"github.com/aziemski/bookstore/internal/x/xecho/xmiddleware"
	"github.com/aziemski/bookstore/internal/x/xlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	address := flag.String("a", "localhost:8080", "server:port")
	flag.Parse()

	log := xlog.GetLogger()

	e := echo.New()
	e.Static("/", "assets")

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

	restServer := rest.NewServer()
	restServer.RegisterWith(e)

	go createFixtures(repo, log)

	web.SetupRoutes(e)

	e.Logger.Fatal(e.Start(*address))
}

func createFixtures(repo *core.Repository, log *slog.Logger) {
	try := func() error {
		time.Sleep(3 * time.Second)
		log.Info("Inserting book")
		_, err := repo.InsertNew(context.Background(), &core.NewBookSpec{
			ID:          "123",
			Title:       "Some title",
			Description: "Description",
		})

		if err != nil {
			log.Warn("Cannot save book", xlog.Err(err))
		}

		return err
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			log.Warn("Received cancellation signal due to deadline.")
			return
		default:
			if err := try(); err == nil {
				return
			}
		}
	}
}
