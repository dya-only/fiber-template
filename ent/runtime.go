// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-template/ent/schema"
	"go-template/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateAt is the schema descriptor for create_at field.
	userDescCreateAt := userFields[2].Descriptor()
	// user.DefaultCreateAt holds the default value on creation for the create_at field.
	user.DefaultCreateAt = userDescCreateAt.Default.(func() time.Time)
	// userDescUpdateAt is the schema descriptor for update_at field.
	userDescUpdateAt := userFields[3].Descriptor()
	// user.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	user.UpdateDefaultUpdateAt = userDescUpdateAt.UpdateDefault.(func() time.Time)
}