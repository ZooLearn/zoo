package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/ZooLearn/zoo/ent/mixins"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(100),
		field.String("email").MaxLen(100),
		field.String("password"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.TimeMixin{},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
