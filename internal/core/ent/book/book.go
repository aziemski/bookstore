// Code generated by ent, DO NOT EDIT.

package book

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldFeatured holds the string denoting the featured field in the database.
	FieldFeatured = "featured"
	// Table holds the table name of the book in the database.
	Table = "books"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldAuthor,
	FieldDescription,
	FieldCategory,
	FieldFeatured,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// AuthorValidator is a validator for the "author" field. It is called by the builders before save.
	AuthorValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	CategoryValidator func(string) error
	// DefaultFeatured holds the default value on creation for the "featured" field.
	DefaultFeatured bool
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Book queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByFeatured orders the results by the featured field.
func ByFeatured(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeatured, opts...).ToFunc()
}
