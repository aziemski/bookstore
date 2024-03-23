package main

import (
	"github.com/aziemski/bookstore/internal/books/ports"
	"log/slog"
	"os"

	"github.com/aziemski/bookstore/internal/x/xecho/xmiddleware"
	"github.com/labstack/echo/v4"
)

func main() {
	log := slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		}))
	slog.SetDefault(log)

	e := echo.New()
	e.HideBanner = true

	e.Use(xmiddleware.RequestLogger(log))

	httpServer := &ports.HttpServer{}
	httpServer.RegisterWith(e)

	//apiGroup := e.Group("api")

	e.Logger.Fatal(e.Start("localhost:8080"))

}
