package search

import (
	"github.com/aziemski/bookstore/internal/web/search/searchviews"
	"github.com/aziemski/bookstore/internal/x/xecho"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type Handler struct {
}

func (h *Handler) Handle(c echo.Context) error {
	q := ""
	l := 1
	if err := echo.QueryParamsBinder(c).
		String("q", &q).
		Int("l", &l).
		BindError(); err != nil {
		slog.Error("bind error", "err", err)
	}

	sv := searchviews.SearchResult(q)

	return xecho.RenderView(c, sv)
}
