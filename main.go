package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
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

	restServer := rest.NewServer(repo)
	restServer.RegisterWith(e)

	go createFixtures(repo, log)

	web.SetupRoutes(e)

	e.Logger.Fatal(e.Start(*address))
}

type FakeBook struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	Category    string `json:"category"`
	Featured    bool   `json:"featured"`
}

func readFakeBooks(log *slog.Logger) []FakeBook {
	var books []FakeBook

	file, err := os.Open("assets/fake_books.json")
	if err != nil {
		log.Error("unexpected os.Open() err", xlog.Err(err))
		return books
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&books); err != nil {
		log.Error("unexpected decoder.Decode() err", xlog.Err(err))
		return nil
	}

	return books
}

func createFixtures(repo *core.Repository, log *slog.Logger) {
	try := func() error {
		if c, err := repo.GetTotalCount(context.Background()); c > 0 && err == nil {
			log.Info("book records found ", "count", c)
			return nil
		}
		time.Sleep(3 * time.Second)

		fakeBooks := readFakeBooks(log)
		for i, fb := range fakeBooks {
			log.Info("Inserting book")
			_, err := repo.InsertNew(context.Background(), &core.NewBookSpec{
				ID:          fmt.Sprintf("b%d", i),
				Title:       fb.Title,
				Description: fb.Description,
			})
			if err != nil {
				log.Warn("Cannot save book", xlog.Err(err))
				return err
			}
		}

		return nil
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
