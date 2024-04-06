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

func (s *Server) RegisterWith(e *echo.Echo) {
	g := e.Group("api/v1")
	RegisterHandlers(g, s)
}

func (s *Server) PostBooks(c echo.Context) error {
	ctx := c.Request().Context()

	var in Book
	if err := c.Bind(&in); err != nil {
		return errRsp(c, http.StatusBadRequest, err)
	}

	bookSpec := &core.Book{
		Title:       in.Title,
		Author:      in.Author,
		Summary:     in.Summary,
		Description: in.Description,
		ImageLink:   in.ImageLink,
		Category:    in.Category,
		Featured:    in.Featured,
	}

	if in.Id != nil {
		bookSpec.ID = *in.Id
	}

	created, err := s.repo.InsertNew(ctx, bookSpec)

	if err != nil {
		return errRsp(c, http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, apiBook(created))
}

func (s *Server) GetBooksSearch(c echo.Context, params GetBooksSearchParams) error {
	ctx := c.Request().Context()

	q := ""
	if params.Q != nil {
		q = *params.Q
	}

	found := s.repo.Query(ctx, core.QueryArgs{
		SearchTerm: q,
		Offset:     params.Offset,
		Limit:      params.Limit,
	})

	bookList := apiBookList(found)

	return c.JSON(http.StatusOK, bookList)
}

func (s *Server) DeleteBooksId(c echo.Context, id string) error {
	// TODO implement me
	panic("implement me")
}

func (s *Server) GetBooksId(c echo.Context, id string) error {
	ctx := c.Request().Context()

	b, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiBook(b))
}

func (s *Server) GetBooks(ctx echo.Context, params GetBooksParams) error {
	// TODO implement me
	panic("implement me")
}

func (s *Server) PutBooksId(c echo.Context, id string) error {
	// TODO implement me
	panic("implement me")
}

func apiBookList(found []core.Book) BookList {
	items := make([]BookItem, 0, len(found))
	for _, in := range found {
		b := BookItem{
			Author:    in.Author,
			Id:        in.ID,
			ImageLink: in.ImageLink,
			Summary:   in.Summary,
			Title:     in.Title,
		}
		items = append(items, b)
	}
	return BookList{
		Items: items,
		Total: len(items),
	}
}

func apiBook(in *core.Book) Book {
	return Book{
		Author:      in.Author,
		Category:    in.Category,
		Description: in.Description,
		Featured:    in.Featured,
		Id:          &in.ID,
		ImageLink:   in.ImageLink,
		Summary:     in.Summary,
		Title:       in.Title,
	}
}

func errRsp(c echo.Context, code int, err error) error {
	return c.JSON(code, Error{
		Code: http.StatusText(code),
		Msg:  err.Error(),
	})
}
