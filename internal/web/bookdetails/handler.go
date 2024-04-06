package bookdetails

import (
	"github.com/aziemski/bookstore/internal/core"
	"github.com/aziemski/bookstore/internal/web/bookdetails/bookdetailviews"
	"github.com/aziemski/bookstore/internal/x/xecho"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) Handle(c echo.Context) error {
	dv := bookdetailviews.BookDetails("", core.Book{
		ID:          "id",
		Title:       "title",
		Author:      "authordd",
		Summary:     "summary",
		Description: "desc",
		ImageLink:   "https://http.cat/208.jpg",
		Category:    "category",
		Featured:    false,
	})
	return xecho.RenderView(c, dv)
}
