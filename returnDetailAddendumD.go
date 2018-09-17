// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
)

// Errors specific to a ReturnDetailAddendumD Record

// ReturnDetailAddendumD Record
type ReturnDetailAddendumD struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewReturnDetailAddendumD returns a new ReturnDetailAddendumD with default values for non exported fields
func NewReturnDetailAddendumD() ReturnDetailAddendumD {
	rdAddendumD := ReturnDetailAddendumD{
		recordType: "35",
	}
	return rdAddendumD
}

// Parse takes the input record string and parses the ReturnDetailAddendumD values
func (rdAddendumD *ReturnDetailAddendumD) Parse(record string) {
	// Character position 1-2, Always "35"
	rdAddendumD.recordType = "35"

}

// String writes the ReturnDetailAddendumD struct to a string.
func (rdAddendumD *ReturnDetailAddendumD) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rdAddendumD.recordType)
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rdAddendumD *ReturnDetailAddendumD) Validate() error {
	if err := rdAddendumD.fieldInclusion(); err != nil {
		return err
	}
	if rdAddendumD.recordType != "35" {
		msg := fmt.Sprintf(msgRecordType, 35)
		return &FieldError{FieldName: "recordType", Value: rdAddendumD.recordType, Msg: msg}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rdAddendumD *ReturnDetailAddendumD) fieldInclusion() error {
	if rdAddendumD.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: rdAddendumD.recordType, Msg: msgFieldInclusion}
	}
	return nil
}
