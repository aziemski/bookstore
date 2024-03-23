package app

import "github.com/aziemski/bookstore/internal/books/app/query"

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	GetBookByID query.GetByIDBookHandler
}
