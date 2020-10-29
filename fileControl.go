// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

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
	// validator is composed for image cash letter data validation
	validator
	// converters is composed for image cash letter to golang Converters
	converters
}

// NewFileControl returns a new FileControl with default values for non exported fields
func NewFileControl() FileControl {
	fc := FileControl{}
	fc.setRecordType()
	return fc
}

func (fc *FileControl) setRecordType() {
	if fc == nil {
		return
	}

	fc.recordType = "99"
	fc.reserved = "               "
}

// Parse takes the input record string and parses the FileControl values
func (fc *FileControl) Parse(record string) {
	if utf8.RuneCountInString(record) < 65 {
		return
	}
	// Character position 1-2, Always "99"
	fc.setRecordType()
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

func (fc *FileControl) UnmarshalJSON(data []byte) error {
	type Alias FileControl
	aux := struct {
		*Alias
	}{
		(*Alias)(fc),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fc.setRecordType()
	return nil
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
	buf.WriteString(fc.CreditTotalIndicatorField())
	buf.WriteString(fc.reservedField())
	return buf.String()
}

// Validate performs image cash letter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (fc *FileControl) Validate() error {
	if err := fc.fieldInclusion(); err != nil {
		return err
	}
	if fc.recordType != "99" {
		msg := fmt.Sprintf(msgRecordType, 99)
		return &FieldError{FieldName: "recordType", Value: fc.recordType, Msg: msg}
	}
	if err := fc.isAlphanumericSpecial(fc.ImmediateOriginContactName); err != nil {
		return &FieldError{FieldName: "ImmediateOriginContactName",
			Value: fc.ImmediateOriginContactName, Msg: err.Error()}
	}
	if err := fc.isNumeric(fc.ImmediateOriginContactPhoneNumber); err != nil {
		return &FieldError{FieldName: "ImmediateOriginContactPhoneNumber",
			Value: fc.ImmediateOriginContactPhoneNumber, Msg: err.Error()}
	}
	// Conditional
	if fc.CreditTotalIndicatorField() != "" {
		if err := fc.isCreditTotalIndicator(fc.CreditTotalIndicator); err != nil {
			return &FieldError{FieldName: "CreditTotalIndicator", Value: fc.CreditTotalIndicatorField(), Msg: err.Error()}
		}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (fc *FileControl) fieldInclusion() error {
	if fc.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: fc.recordType,
			Msg:   msgFieldInclusion + ", did you use FileControl()?"}
	}
	if fc.CashLetterCount == 0 {
		return &FieldError{FieldName: "CashLetterCount",
			Value: fc.CashLetterCountField(),
			Msg:   msgFieldInclusion + ", did you use FileControl()?"}
	}
	if fc.TotalRecordCount == 0 {
		return &FieldError{FieldName: "TotalRecordCount",
			Value: fc.TotalRecordCountField(),
			Msg:   msgFieldInclusion + ", did you use FileControl()?"}
	}
	if fc.TotalItemCount == 0 {
		return &FieldError{FieldName: "TotalItemCount",
			Value: fc.TotalItemCountField(),
			Msg:   msgFieldInclusion + ", did you use FileControl()?"}
	}
	if fc.FileTotalAmount == 0 {
		return &FieldError{FieldName: "FileTotalAmount",
			Value: fc.FileTotalAmountField(),
			Msg:   msgFieldInclusion + ", did you use FileControl()?"}
	}
	return nil
}

// CashLetterCountField gets a string of the CashLetterCount zero padded
func (fc *FileControl) CashLetterCountField() string {
	return fc.numericField(fc.CashLetterCount, 6)
}

// TotalRecordCountField gets a string of the TotalRecordCount zero padded
func (fc *FileControl) TotalRecordCountField() string {
	return fc.numericField(fc.TotalRecordCount, 8)
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

// CreditTotalIndicatorField gets a string of the CreditTotalIndicator field
func (fc *FileControl) CreditTotalIndicatorField() string {
	return fc.numericField(fc.CreditTotalIndicator, 1)
}

// reservedField gets reserved - blank space
func (fc *FileControl) reservedField() string {
	return fc.alphaField(fc.reserved, 15)
}
