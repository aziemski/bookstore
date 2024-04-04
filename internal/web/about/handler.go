package about

import (
	"github.com/aziemski/bookstore/internal/web/about/aboutviews"
	"github.com/aziemski/bookstore/internal/x/xecho"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) Handle(c echo.Context) error {
	dv := aboutviews.AboutIndex("test")
	return xecho.RenderView(c, dv)
}
