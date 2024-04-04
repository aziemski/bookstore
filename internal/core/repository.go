package core

import (
	"context"

	"github.com/aziemski/bookstore/internal/books/domain/books"
	"github.com/aziemski/bookstore/internal/core/db"
)

type Book struct {
	ID          string
	Title       string
	Description string
}

type NewBookSpec struct {
	ID          string
	Title       string
	Description string
}

type Repository struct {
	client *db.Client
}

func NewRepository(client *db.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) InsertNew(ctx context.Context, in *NewBookSpec) (*books.Book, error) {
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

func (r *Repository) FindByID(ctx context.Context, id string) (*Book, error) {
	b, err := r.client.Book.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: "",
	}, nil
}
