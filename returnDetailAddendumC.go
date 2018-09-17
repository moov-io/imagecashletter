// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
)

// Errors specific to a ReturnDetailAddendumC Record

// ReturnDetailAddendumC Record
type ReturnDetailAddendumC struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewReturnDetailAddendumC returns a new ReturnDetailAddendumC with default values for non exported fields
func NewReturnDetailAddendumC() ReturnDetailAddendumC {
	rdAddendumC := ReturnDetailAddendumC{
		recordType: "33",
	}
	return rdAddendumC
}

// Parse takes the input record string and parses the ReturnDetailAddendumC values
func (rdAddendumC *ReturnDetailAddendumC) Parse(record string) {
	// Character position 1-2, Always "33"
	rdAddendumC.recordType = "33"

}

// String writes the ReturnDetailAddendumC struct to a string.
func (rdAddendumC *ReturnDetailAddendumC) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rdAddendumC.recordType)
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rdAddendumC *ReturnDetailAddendumC) Validate() error {
	if err := rdAddendumC.fieldInclusion(); err != nil {
		return err
	}
	if rdAddendumC.recordType != "33" {
		msg := fmt.Sprintf(msgRecordType, 33)
		return &FieldError{FieldName: "recordType", Value: rdAddendumC.recordType, Msg: msg}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rdAddendumC *ReturnDetailAddendumC) fieldInclusion() error {
	if rdAddendumC.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: rdAddendumC.recordType, Msg: msgFieldInclusion}
	}
	return nil
}
