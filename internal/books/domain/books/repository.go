package books

import "context"

type Book struct {
	ID          string
	Title       string
	Description string
}

type Repository interface {
	FindByID(ctx context.Context, id string) (*Book, error)
}
