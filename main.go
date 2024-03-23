package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/aziemski/bookstore/internal/books/domain/books"

	"github.com/aziemski/bookstore/internal/books/adapters/entgo"
	"github.com/aziemski/bookstore/internal/books/app"
	"github.com/aziemski/bookstore/internal/books/app/query"
	"github.com/aziemski/bookstore/internal/books/ports"
	"github.com/aziemski/bookstore/internal/x/xecho/xmiddleware"
	"github.com/aziemski/bookstore/internal/x/xlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log := xlog.GetLogger()

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(xmiddleware.RequestLogger(log))

	// dbClient, err := entgo.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	dbClient, err := entgo.Open("sqlite3", "file:./bookstore.db?_fk=1")
	if err != nil {
		log.Error("unexpected db.Open err", xlog.Err(err))
		panic(err)
	}

	if err = dbClient.Schema.Create(context.Background()); err != nil {
		log.Error("unexpected db.Schema.Create err", xlog.Err(err))
		panic(err)
	}

	repo := entgo.NewRepository(dbClient, log)

	a := &app.App{
		Queries: app.Queries{
			GetBookByID: query.NewGetByIDBookHandler(repo, log),
		},
	}

	httpServer := ports.NewHTTPServer(a, log)
	httpServer.RegisterWith(e)

	go createFixtures(repo, log)

	e.Logger.Fatal(e.Start("localhost:8080"))
}

func createFixtures(repo *entgo.Repository, log *slog.Logger) {
	try := func() error {
		time.Sleep(3 * time.Second)
		log.Info("Inserting book")
		_, err := repo.InsertNew(context.Background(), &books.NewBookSpec{
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
