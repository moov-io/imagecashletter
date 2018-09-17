// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
)

// Errors specific to a ReturnDetail Record

// ReturnDetail Record
type ReturnDetail struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// ReturnDetailAddendumA
	ReturnDetailAddendumA []ReturnDetailAddendumA `json:"returnDetailAddendumA"`
	// ReturnDetailAddendumB
	ReturnDetailAddendumB []ReturnDetailAddendumB `json:"returnDetailAddendumB"`
	// ReturnDetailAddendumC
	ReturnDetailAddendumC []ReturnDetailAddendumC `json:"returnDetailAddendumC"`
	// ReturnDetailAddendumD
	ReturnDetailAddendumD []ReturnDetailAddendumD `json:"returnDetailAddendumD"`
	// ImageViewDetail
	ImageViewDetail []ImageViewDetail `json:"imageViewDetail"`
	// ImageViewData
	ImageViewData []ImageViewData `json:"imageViewData"`
	// ImageViewAnalysis
	ImageViewAnalysis []ImageViewAnalysis `json:"imageViewAnalysis"`
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewReturnDetail returns a new ReturnDetail with default values for non exported fields
func NewReturnDetail() *ReturnDetail {
	rd := &ReturnDetail{
		recordType: "31",
	}
	return rd
}

// Parse takes the input record string and parses the ReturnDetail values
func (rd *ReturnDetail) Parse(record string) {
	// Character position 1-2, Always "31"
	rd.recordType = "31"
}

// String writes the ReturnDetail struct to a variable length string.
func (rd *ReturnDetail) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rd.recordType)
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rd *ReturnDetail) Validate() error {
	if err := rd.fieldInclusion(); err != nil {
		return err
	}
	if rd.recordType != "31" {
		msg := fmt.Sprintf(msgRecordType, 31)
		return &FieldError{FieldName: "recordType", Value: rd.recordType, Msg: msg}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rd *ReturnDetail) fieldInclusion() error {
	if rd.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: rd.recordType, Msg: msgFieldInclusion}
	}
	return nil
}
