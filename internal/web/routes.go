package web

import (
	"github.com/aziemski/bookstore/internal/web/bookdetails"
	"github.com/aziemski/bookstore/internal/web/home"
	"github.com/aziemski/bookstore/internal/web/search"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	hh := &home.Handler{}
	bdh := &bookdetails.Handler{}
	sh := &search.Handler{}

	e.GET("/", hh.Handle)
	e.GET("/details/:id", bdh.Handle)
	e.GET("/search", sh.Handle)
}
