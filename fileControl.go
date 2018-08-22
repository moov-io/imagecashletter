// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a FileControl Record

// FileControl Record
type FileControl struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record
	recordType string
	// CashLetterCount identifies the total number of cash letters within the file.
	CashLetterCount int `json:"cashLetterCount"`
	// TotalRecordCount identifies the total number of records of all types sent in the file, including the FileControl.
	TotalRecordCount int `json:"totalRecordCount"`
	// TotalItemCount identifies the total number of Items sent within the file.
	TotalItemCount int `json:"totalItemCount"`
	// FileTotalAmount identifies the total Item amount of the complete file.
	// ToDo: int64 by default on 64bit - string for 32 bit?
	FileTotalAmount int `json:"fileTotalAmount"`
	// ImmediateOriginContactName identifies contact at the institution that creates the ECE file.
	ImmediateOriginContactName string `json:"immediateOriginContactName"`
	// ImmediateOriginContactPhoneNumber is the phone number of the contact at the institution that creates the
	// file.
	ImmediateOriginContactPhoneNumber string `json:"immediateOriginContactPhoneNumber"`
	// CreditTotalIndicator isa code that indicates whether Credits Items are included in this record’s totals.
	// If so they will be included in TotalItemCount and FileTotal Amount.
	// TotalRecordCount includes all records of all types regardless of the value of this field.
	// Values:
	// 	0: Credit Items are not included in totals
	//  1: Credit Items are included in totals
	CreditTotalIndicator int `json:"creditTotalIndicator"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewFileControl returns a new FileControl with default values for non exported fields
func NewFileControl() FileControl {
	fc := FileControl{
		recordType: "99",
	}
	return fc
}

// Parse takes the input record string and parses the FileControl values

// String writes the FileControl struct to a variable length string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
