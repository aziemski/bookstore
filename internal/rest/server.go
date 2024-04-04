package rest

import (
	"net/http"

	"github.com/aziemski/bookstore/internal/core"
	"github.com/labstack/echo/v4"
)

type Server struct {
	repo *core.Repository
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetBookByID(c echo.Context, id string) error {
	ctx := c.Request().Context()

	b, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, core2api(b))
}

func (s *Server) RegisterWith(e *echo.Echo) {
	g := e.Group("api/v1")
	RegisterHandlers(g, s)
}

func core2api(m *core.Book) Book {
	return Book{
		Title: m.Title,
	}
}
