package entgo

import (
	"context"
	"log/slog"

	"github.com/aziemski/bookstore/internal/books/domain/books"
)

var _ books.Repository = (*Repository)(nil)

type Repository struct {
	client *Client

	log *slog.Logger
}

func NewRepository(client *Client, log *slog.Logger) *Repository {
	return &Repository{
		client: client,
		log:    log,
	}
}

func (r *Repository) InsertNew(ctx context.Context, in *books.NewBookSpec) (*books.Book, error) {
	b, err := r.client.Book.Create().
		SetID(in.ID).
		SetTitle(in.Title).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &books.Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: "",
	}, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*books.Book, error) {
	b, err := r.client.Book.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &books.Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: "",
	}, nil
}
