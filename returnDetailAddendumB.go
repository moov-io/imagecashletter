// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
)

// Errors specific to a ReturnDetailAddendumB Record

// ReturnDetailAddendumB Record
type ReturnDetailAddendumB struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewReturnDetailAddendumB returns a new ReturnDetailAddendumB with default values for non exported fields
func NewReturnDetailAddendumB() ReturnDetailAddendumB {
	rdAddendumB := ReturnDetailAddendumB{
		recordType: "33",
	}
	return rdAddendumB
}

// Parse takes the input record string and parses the ReturnDetailAddendumB values
func (rdAddendumB *ReturnDetailAddendumB) Parse(record string) {
	// Character position 1-2, Always "33"
	rdAddendumB.recordType = "33"

}

// String writes the ReturnDetailAddendumB struct to a string.
func (rdAddendumB *ReturnDetailAddendumB) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rdAddendumB.recordType)
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rdAddendumB *ReturnDetailAddendumB) Validate() error {
	if err := rdAddendumB.fieldInclusion(); err != nil {
		return err
	}
	if rdAddendumB.recordType != "33" {
		msg := fmt.Sprintf(msgRecordType, 33)
		return &FieldError{FieldName: "recordType", Value: rdAddendumB.recordType, Msg: msg}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rdAddendumB *ReturnDetailAddendumB) fieldInclusion() error {
	if rdAddendumB.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: rdAddendumB.recordType, Msg: msgFieldInclusion}
	}
	return nil
}
