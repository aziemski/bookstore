package search

import (
	"log/slog"

	"github.com/aziemski/bookstore/internal/core"
	"github.com/aziemski/bookstore/internal/web/search/searchviews"
	"github.com/aziemski/bookstore/internal/x/xecho"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	r *core.Repository
}

func NewHandler(r *core.Repository) *Handler {
	return &Handler{r: r}
}

func (h *Handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	q := ""
	l := 1
	if err := echo.QueryParamsBinder(c).
		String("q", &q).
		Int("l", &l).
		BindError(); err != nil {
		slog.Error("bind error", "err", err)
	}

	limit := 20
	books := h.r.Query(ctx, core.QueryArgs{SearchTerm: q, Limit: &limit})

	sv := searchviews.SearchResult(q, books)

	return xecho.RenderView(c, sv)
}
