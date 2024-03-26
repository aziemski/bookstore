package command

import (
	"context"
	"log/slog"

	"github.com/aziemski/bookstore/internal/books/adapters/entgo"
	"github.com/aziemski/bookstore/internal/books/domain/books"
)

type AddBookHandler struct {
	repo books.Repository

	log *slog.Logger
}

func NewAddBoolHandler(repo *entgo.Repository, log *slog.Logger) *AddBookHandler {
	return &AddBookHandler{repo: repo, log: log}
}

func (h *AddBookHandler) Handle(ctx context.Context, in *books.NewBookSpec) (*books.Book, error) {
	return h.repo.InsertNew(ctx, in)
}
