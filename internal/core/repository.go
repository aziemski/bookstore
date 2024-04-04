package core

import (
	"context"

	"github.com/aziemski/bookstore/internal/core/ent"
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
	db *ent.Client
}

func NewRepository(db *ent.Client) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertNew(ctx context.Context, in *NewBookSpec) (*Book, error) {
	b, err := r.db.Book.Create().
		SetID(in.ID).
		SetTitle(in.Title).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: "",
	}, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Book, error) {
	b, err := r.db.Book.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: "",
	}, nil
}

func (r *Repository) GetTotalCount(ctx context.Context) (int, error) {
	return r.db.Book.Query().Count(ctx)
}
