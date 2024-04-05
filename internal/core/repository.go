package core

import (
	"context"
	"log/slog"

	"github.com/aziemski/bookstore/internal/core/ent"
	"github.com/aziemski/bookstore/internal/core/ent/book"
	"github.com/aziemski/bookstore/internal/x/xlog"
)

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	Category    string `json:"category"`
	Featured    bool   `json:"featured"`
}

type Repository struct {
	db *ent.Client
}

func NewRepository(db *ent.Client) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertNew(ctx context.Context, in *Book) (*Book, error) {
	b, err := r.db.Book.Create().
		SetID(in.ID).
		SetTitle(in.Title).
		SetAuthor(in.Author).
		SetDescription(in.Description).
		SetSummary(in.Summary).
		SetImageLink(in.ImageLink).
		SetCategory(in.Category).
		SetFeatured(in.Featured).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	out := ent2core(b)
	return &out, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Book, error) {
	in, err := r.db.Book.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	out := ent2core(in)
	return &out, nil
}

func (r *Repository) GetTotalCount(ctx context.Context) (int, error) {
	return r.db.Book.Query().Count(ctx)
}

func (r *Repository) FindFeaturedBooks(ctx context.Context) []Book {
	var result []Book

	found, err := r.db.Book.Query().
		Where(book.FeaturedEQ(true)).
		All(ctx)

	if err != nil {
		slog.Error("sql, unexpected FindFeaturedBooks err", xlog.Err(err))
		return result
	}

	for _, in := range found {
		out := ent2core(in)
		result = append(result, out)
	}

	return result
}

func (r *Repository) Find(ctx context.Context, q string) []Book {
	var result []Book

	found, err := r.db.Book.Query().
		Where(
			book.Or(
				book.TitleContains(q),
				book.AuthorContains(q),
				book.SummaryContains(q),
				book.DescriptionContains(q),
				book.CategoryContains(q),
			)).
		All(ctx)

	if err != nil {
		slog.Error("sql, unexpected FindFeaturedBooks err", xlog.Err(err))
		return result
	}

	for _, in := range found {
		out := ent2core(in)
		result = append(result, out)
	}

	return result
}

func ent2core(in *ent.Book) Book {
	return Book{
		ID:          in.ID,
		Title:       in.Title,
		Author:      in.Author,
		Summary:     in.Summary,
		Description: in.Description,
		Category:    in.Category,
		Featured:    in.Featured,
		ImageLink:   in.ImageLink,
	}
}
