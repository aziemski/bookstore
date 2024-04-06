package web

import (
	"github.com/aziemski/bookstore/internal/core"
	"github.com/aziemski/bookstore/internal/web/about"
	"github.com/aziemski/bookstore/internal/web/bookdetails"
	"github.com/aziemski/bookstore/internal/web/home"
	"github.com/aziemski/bookstore/internal/web/search"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, repo *core.Repository) {
	ah := &about.Handler{}
	hh := home.NewHandler(repo)
	bdh := bookdetails.NewHandler(repo)
	sh := search.NewHandler(repo)

	e.GET("/", hh.Handle)
	e.GET("/details/:id", bdh.Handle)
	e.GET("/search", sh.Handle)
	e.GET("/about", ah.Handle)
}
