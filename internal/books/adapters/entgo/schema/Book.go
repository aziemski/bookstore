package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Book struct {
	ent.Schema
}

func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable().
			MaxLen(32), //nolint:gomnd // magic number is ok here
		field.String("title").
			NotEmpty(),
	}
}

func (Book) Edges() []ent.Edge {
	return nil
}
