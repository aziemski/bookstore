package main

import (
	"github.com/aziemski/bookstore/internal/books/ports"
	"github.com/aziemski/bookstore/internal/x/xecho/xmiddleware"
	"github.com/aziemski/bookstore/internal/x/xlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log := xlog.GetLogger()

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(xmiddleware.RequestLogger(log))

	httpServer := &ports.HTTPServer{}
	httpServer.RegisterWith(e)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
