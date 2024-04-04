package search

import (
	"github.com/aziemski/bookstore/internal/web/search/searchviews"
	"github.com/aziemski/bookstore/internal/x/xecho"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) Handle(c echo.Context) error {
	sv := searchviews.SearchResult("test")

	return xecho.RenderView(c, sv)
}
