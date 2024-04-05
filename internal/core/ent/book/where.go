// Code generated by ent, DO NOT EDIT.

package book

import (
	"entgo.io/ent/dialect/sql"
	"github.com/aziemski/bookstore/internal/core/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldAuthor, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDescription, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCategory, v))
}

// Featured applies equality check predicate on the "featured" field. It's identical to FeaturedEQ.
func Featured(v bool) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldFeatured, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldTitle, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldAuthor, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldDescription, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldCategory, v))
}

// FeaturedEQ applies the EQ predicate on the "featured" field.
func FeaturedEQ(v bool) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldFeatured, v))
}

// FeaturedNEQ applies the NEQ predicate on the "featured" field.
func FeaturedNEQ(v bool) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldFeatured, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(sql.NotPredicates(p))
}
