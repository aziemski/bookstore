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
		field.String("author").
			NotEmpty(),
		field.String("summary").
			NotEmpty(),
		field.String("description").
			NotEmpty(),
		field.String("category").
			NotEmpty(),
		field.String("image_link").
			NotEmpty(),
		field.Bool("featured").
			Default(false),
	}
}

func (Book) Edges() []ent.Edge {
	return nil
}
