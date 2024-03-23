package query

import (
	"context"

	"github.com/aziemski/bookstore/internal/books/domain/books"
)

type GetByIDBookHandler struct {
	repo books.Repository
}

func (h GetByIDBookHandler) Handle(context context.Context, id string) (*books.Book, error) {
	return h.repo.FindByID(context, id)
}
