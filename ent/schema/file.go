package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/ZooLearn/zoo/ent/mixins"
)

type File struct {
	ent.Schema
}

func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("name"),
		field.Uint8("file_type"),
		field.Uint64("size"),
		field.String("path").Optional(),
		field.String("temp_path"),
		field.String("user_id"),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.TimeMixin{},
	}
}

func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("file_type"),
	}
}

func (File) Edges() []ent.Edge {
	return []ent.Edge{}
}
