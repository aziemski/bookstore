package ports

import (
	"log/slog"
	"net/http"

	"github.com/aziemski/bookstore/internal/books/app"

	"github.com/labstack/echo/v4"
)

var _ ServerInterface = (*HTTPServer)(nil)

type HTTPServer struct {
	app *app.App
	log *slog.Logger
}

func NewHTTPServer(app *app.App, log *slog.Logger) *HTTPServer {
	return &HTTPServer{
		app: app,
		log: log,
	}
}

func (h *HTTPServer) GetBookByID(c echo.Context, id string) error {
	ctx := c.Request().Context()

	b, err := h.app.Queries.GetBookByID.Handle(ctx, id)

	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, Book{
		Title: b.Title})
}

func (h *HTTPServer) RegisterWith(e *echo.Echo) {
	g := e.Group("api/v1")
	RegisterHandlers(g, h)
}
