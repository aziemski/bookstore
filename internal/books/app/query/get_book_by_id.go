package query

import (
	"context"
	"log/slog"

	"github.com/aziemski/bookstore/internal/books/adapters/entgo"

	"github.com/aziemski/bookstore/internal/books/domain/books"
)

type GetByIDBookHandler struct {
	repo books.Repository

	log *slog.Logger
}

func (h *GetByIDBookHandler) Handle(context context.Context, id string) (*books.Book, error) {
	return h.repo.FindByID(context, id)
}

func NewGetByIDBookHandler(repo *entgo.Repository, log *slog.Logger) *GetByIDBookHandler {
	return &GetByIDBookHandler{repo: repo, log: log}
}
