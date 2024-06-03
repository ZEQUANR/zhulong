// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ZEQUANR/zhulong/ent/operationlog"
	"github.com/ZEQUANR/zhulong/ent/reviews"
	"github.com/ZEQUANR/zhulong/ent/schema"
	"github.com/ZEQUANR/zhulong/ent/thesis"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	operationlogFields := schema.OperationLog{}.Fields()
	_ = operationlogFields
	// operationlogDescTime is the schema descriptor for time field.
	operationlogDescTime := operationlogFields[3].Descriptor()
	// operationlog.DefaultTime holds the default value on creation for the time field.
	operationlog.DefaultTime = operationlogDescTime.Default.(func() time.Time)
	reviewsFields := schema.Reviews{}.Fields()
	_ = reviewsFields
	// reviewsDescCreateTime is the schema descriptor for create_time field.
	reviewsDescCreateTime := reviewsFields[2].Descriptor()
	// reviews.DefaultCreateTime holds the default value on creation for the create_time field.
	reviews.DefaultCreateTime = reviewsDescCreateTime.Default.(func() time.Time)
	thesisFields := schema.Thesis{}.Fields()
	_ = thesisFields
	// thesisDescCreateTime is the schema descriptor for create_time field.
	thesisDescCreateTime := thesisFields[12].Descriptor()
	// thesis.DefaultCreateTime holds the default value on creation for the create_time field.
	thesis.DefaultCreateTime = thesisDescCreateTime.Default.(func() time.Time)
}
