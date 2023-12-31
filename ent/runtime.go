// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ZooLearn/zoo/ent/schema"
	"github.com/ZooLearn/zoo/ent/tag"
	"github.com/ZooLearn/zoo/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	tagMixin := schema.Tag{}.Mixin()
	tagMixinFields1 := tagMixin[1].Fields()
	_ = tagMixinFields1
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagMixinFields1[0].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagMixinFields1[1].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tag.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tag.UpdateDefaultUpdatedAt = tagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tagDescKey is the schema descriptor for key field.
	tagDescKey := tagFields[0].Descriptor()
	// tag.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	tag.KeyValidator = tagDescKey.Validators[0].(func(string) error)
	// tagDescAlias is the schema descriptor for alias field.
	tagDescAlias := tagFields[1].Descriptor()
	// tag.AliasValidator is a validator for the "alias" field. It is called by the builders before save.
	tag.AliasValidator = tagDescAlias.Validators[0].(func(string) error)
	// tagDescColor is the schema descriptor for color field.
	tagDescColor := tagFields[2].Descriptor()
	// tag.ColorValidator is a validator for the "color" field. It is called by the builders before save.
	tag.ColorValidator = tagDescColor.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields1[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields1[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
