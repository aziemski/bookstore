// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package ports

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Book defines model for Book.
type Book struct {
	Tittle *string `json:"tittle,omitempty"`
}
