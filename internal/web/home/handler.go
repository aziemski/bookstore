package home

import (
	"github.com/aziemski/bookstore/internal/core"
	"github.com/aziemski/bookstore/internal/web/home/homeviews"
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

	featuredBooks := h.r.FindFeaturedBooks(ctx)
	homeView := homeviews.HomeIndex(featuredBooks)

	return xecho.RenderView(c, homeView)
}
