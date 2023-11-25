// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"net/http"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldT holds the string denoting the t field in the database.
	FieldT = "t"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldURLs holds the string denoting the urls field in the database.
	FieldURLs = "urls"
	// FieldRaw holds the string denoting the raw field in the database.
	FieldRaw = "raw"
	// FieldDirs holds the string denoting the dirs field in the database.
	FieldDirs = "dirs"
	// FieldInts holds the string denoting the ints field in the database.
	FieldInts = "ints"
	// FieldFloats holds the string denoting the floats field in the database.
	FieldFloats = "floats"
	// FieldStrings holds the string denoting the strings field in the database.
	FieldStrings = "strings"
	// FieldIntsValidate holds the string denoting the ints_validate field in the database.
	FieldIntsValidate = "ints_validate"
	// FieldFloatsValidate holds the string denoting the floats_validate field in the database.
	FieldFloatsValidate = "floats_validate"
	// FieldStringsValidate holds the string denoting the strings_validate field in the database.
	FieldStringsValidate = "strings_validate"
	// FieldAddr holds the string denoting the addr field in the database.
	FieldAddr = "addr"
	// FieldUnknown holds the string denoting the unknown field in the database.
	FieldUnknown = "unknown"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldT,
	FieldURL,
	FieldURLs,
	FieldRaw,
	FieldDirs,
	FieldInts,
	FieldFloats,
	FieldStrings,
	FieldIntsValidate,
	FieldFloatsValidate,
	FieldStringsValidate,
	FieldAddr,
	FieldUnknown,
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
	// DefaultDirs holds the default value on creation for the "dirs" field.
	DefaultDirs func() []http.Dir
	// DefaultInts holds the default value on creation for the "ints" field.
	DefaultInts []int
	// IntsValidateValidator is a validator for the "ints_validate" field. It is called by the builders before save.
	IntsValidateValidator func([]int) error
	// FloatsValidateValidator is a validator for the "floats_validate" field. It is called by the builders before save.
	FloatsValidateValidator func([]float64) error
	// StringsValidateValidator is a validator for the "strings_validate" field. It is called by the builders before save.
	StringsValidateValidator func([]string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}
