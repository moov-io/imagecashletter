// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
)

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

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
func (fc *FileControl) Parse(record string) {
	// Character position 1-2, Always "99"
	fc.recordType = "99"
	// 03-08
	fc.CashLetterCount = fc.parseNumField(record[2:8])
	// 09-16
	fc.TotalRecordCount = fc.parseNumField(record[8:16])
	// 17-24
	fc.TotalItemCount = fc.parseNumField(record[16:24])
	// 25-40
	fc.FileTotalAmount = fc.parseNumField(record[24:40])
	// 41-54
	fc.ImmediateOriginContactName = fc.parseStringField(record[40:54])
	// 55-64
	fc.ImmediateOriginContactPhoneNumber = fc.parseStringField(record[54:64])
	// 65-65
	fc.CreditTotalIndicator = fc.parseNumField(record[64:65])
	// 66-80 reserved - Leave blank
	fc.reserved = "               "
}

// String writes the FileControl struct to a string.
func (fc *FileControl) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(fc.recordType)
	buf.WriteString(fc.CashLetterCountField())
	buf.WriteString(fc.TotalRecordCountField())
	buf.WriteString(fc.TotalItemCountField())
	buf.WriteString(fc.FileTotalAmountField())
	buf.WriteString(fc.ImmediateOriginContactNameField())
	buf.WriteString(fc.ImmediateOriginContactPhoneNumberField())
	buf.WriteString(fmt.Sprintf("%v", fc.CreditTotalIndicator))
	buf.WriteString(fc.reservedField())
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (fc *FileControl) Validate() error {
	if err := fc.fieldInclusion(); err != nil {
		return err
	}
	if fc.recordType != "99" {
		msg := fmt.Sprintf(msgRecordType, 99)
		return &FieldError{FieldName: "recordType", Value: fc.recordType, Msg: msg}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (fc *FileControl) fieldInclusion() error {
	if fc.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: fc.recordType, Msg: msgFieldInclusion}
	}
	if fc.CashLetterCount == 0 {
		return &FieldError{FieldName: "CashLetterCount", Value: fc.CashLetterCountField(), Msg: msgFieldInclusion}
	}
	if fc.TotalRecordCount == 0 {
		return &FieldError{FieldName: "TotalRecordCount", Value: fc.TotalRecordCountField(), Msg: msgFieldInclusion}
	}
	if fc.TotalItemCount == 0 {
		return &FieldError{FieldName: "TotalItemCount", Value: fc.TotalItemCountField(), Msg: msgFieldInclusion}
	}
	if fc.FileTotalAmount == 0 {
		return &FieldError{FieldName: "FileTotalAmount ", Value: fc.FileTotalAmountField(), Msg: msgFieldInclusion}
	}
	return nil
}

// CashLetterCountField gets a string of the CashLetterCount zero padded
func (fc *FileControl) CashLetterCountField() string {
	return fc.numericField(fc.CashLetterCount, 6)
}

// TotalRecordCountField gets a string of the TotalRecordCount zero padded
func (fc *FileControl) TotalRecordCountField() string {
	return fc.numericField(fc.CashLetterCount, 8)
}

// TotalItemCountField gets a string of TotalItemCount zero padded
func (fc *FileControl) TotalItemCountField() string {
	return fc.numericField(fc.TotalItemCount, 8)
}

// FileTotalAmountField gets a string of FileTotalAmount zero padded
func (fc *FileControl) FileTotalAmountField() string {
	return fc.numericField(fc.FileTotalAmount, 16)
}

// ImmediateOriginContactNameField gets the ImmediateOriginContactName field padded
func (fc *FileControl) ImmediateOriginContactNameField() string {
	return fc.alphaField(fc.ImmediateOriginContactName, 14)
}

// ImmediateOriginContactPhoneNumberField gets the ImmediateOriginContactPhoneNumber field padded
func (fc *FileControl) ImmediateOriginContactPhoneNumberField() string {
	return fc.alphaField(fc.ImmediateOriginContactPhoneNumber, 10)
}

// reservedField gets reserved - blank space
func (fc *FileControl) reservedField() string {
	return fc.alphaField(fc.reserved, 15)
}
