package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/ZooLearn/zoo/ent/mixins"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").MaxLen(100),
		field.String("alias").MaxLen(100),
		field.String("color").MaxLen(10),
		field.String("imageUrl"),
	}
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.TimeMixin{},
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{}
}
