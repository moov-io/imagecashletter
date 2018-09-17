// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
)

// Errors specific to a ReturnDetailAddendumA Record

// ReturnDetailAddendumA Record
type ReturnDetailAddendumA struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewReturnDetailAddendumA returns a new ReturnDetailAddendumA with default values for non exported fields
func NewReturnDetailAddendumA() ReturnDetailAddendumA {
	rdAddendumA := ReturnDetailAddendumA{
		recordType: "32",
	}
	return rdAddendumA
}

// Parse takes the input record string and parses the ReturnDetailAddendumA values
func (rdAddendumA *ReturnDetailAddendumA) Parse(record string) {
	// Character position 1-2, Always "32"
	rdAddendumA.recordType = "32"

}

// String writes the ReturnDetailAddendumA struct to a string.
func (rdAddendumA *ReturnDetailAddendumA) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rdAddendumA.recordType)
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rdAddendumA *ReturnDetailAddendumA) Validate() error {
	if err := rdAddendumA.fieldInclusion(); err != nil {
		return err
	}
	if rdAddendumA.recordType != "32" {
		msg := fmt.Sprintf(msgRecordType, 32)
		return &FieldError{FieldName: "recordType", Value: rdAddendumA.recordType, Msg: msg}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rdAddendumA *ReturnDetailAddendumA) fieldInclusion() error {
	if rdAddendumA.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: rdAddendumA.recordType, Msg: msgFieldInclusion}
	}
	return nil
}
