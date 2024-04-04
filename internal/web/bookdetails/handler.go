package bookdetails

import (
	"github.com/aziemski/bookstore/internal/web/bookdetails/bookdetailviews"
	"github.com/aziemski/bookstore/internal/x/xecho"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) Handle(c echo.Context) error {
	dv := bookdetailviews.BookDetailsIndex("test")
	return xecho.RenderView(c, dv)
}
