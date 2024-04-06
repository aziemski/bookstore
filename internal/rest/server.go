package rest

import (
	"net/http"

	"github.com/aziemski/bookstore/internal/core"
	"github.com/labstack/echo/v4"
)

type Server struct {
	repo *core.Repository
}

func NewServer(repo *core.Repository) *Server {
	return &Server{
		repo: repo,
	}
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

func (s *Server) GetBooks(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostBooks(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetBooksSearch(ctx echo.Context, params GetBooksSearchParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteBooksId(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetBooksId(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PutBooksId(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func core2api(m *core.Book) Book {
	return Book{
		Title: m.Title,
	}
}
