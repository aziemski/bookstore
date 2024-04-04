package home

import (
	"github.com/aziemski/bookstore/internal/web/home/homeviews"
	"github.com/aziemski/bookstore/internal/x/xecho"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) Handle(c echo.Context) error {
	homeView := homeviews.HomeIndex("hello")

	return xecho.RenderView(c, homeView)
}
