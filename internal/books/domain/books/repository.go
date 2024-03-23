package books

import "context"

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

type Repository interface {
	InsertNew(ctx context.Context, in *NewBookSpec) (*Book, error)
	FindByID(ctx context.Context, id string) (*Book, error)
}
