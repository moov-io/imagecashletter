// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
	"time"
)

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format
// ToDo: ASCII vs EBCDIC

// Errors specific to a File Header Record
var (
	msgRecordType = "received expecting %d"
)

// FileHeader Record is mandatory
type FileHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// recordType defines the type of record.
	recordType string
	// standardLevel identifies the standard level of the file.
	// Values: 03, 30, 35
	// 03: DSTU X9.37 - 2003
	// 30: X9.100-187-2008
	// 35: X9.100-187-2013 and 2016
	StandardLevel string `json:"standardLevel"`
	// TestFileIndicator identifies whether the file is a test or production file.
	// Values:
	// T: Test File
	// P: Production File
	TestFileIndicator string `json:"testIndicator"`
	// ImmediateDestination contains the routing and transit number of the Federal Reserve Bank
	// (FRB) or receiver to which the file is being sent.  Format: TTTTAAAAC, where:
	//  TTTT Federal Reserve Prefix
	//  AAAA ABA Institution Identifier
	//  C Check Digit
	//  For a number that identifies a non-financial institution: NNNNNNNNN
	ImmediateDestination string `json:"immediateDestination"`
	// ImmediateOrigin contains the routing and transit number of the Federal Reserve Bank
	// (FRB) or receiver from which the file is being sent.  Format: TTTTAAAAC, where:
	// TTTT Federal Reserve Prefix
	// AAAA ABA Institution Identifier
	// C Check Digit
	// For a number that identifies a non-financial institution: NNNNNNNNN
	ImmediateOrigin string `json:"immediateOrigin"`
	// FileCreationDate is the date that the immediate origin institution creates the file.  Default time shall be in
	// Eastern Time zone format. Other time zones may be used under clearing arrangements.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	FileCreationDate time.Time `json:"fileCreationDate"`
	// FileCreationTime is the time the immediate origin institution creates the file. Default time shall be in
	// Eastern Time zone format. Other time zones may be used under clearing arrangements.
	// Format: hhmm, where: hh hour, mm minute
	// Values:
	// hh '00' through '23'
	// mm '00' through '59'
	FileCreationTime time.Time `json:"fileCreationTime"`
	// ResendIndicator indicates whether the file has been previously transmitted.
	// Values:
	// Y: Yes
	// N: No
	ResendIndicator string `json:"ResendIndicator"`
	// ImmediateDestinationName identifies the short name of the institution that receives the file.
	ImmediateDestinationName string `json:"immediateDestinationName"`
	// ImmediateOriginName identifies the short name of the institution that sends the file.
	ImmediateOriginName string `json:"ImmediateOriginName"`
	// FileIDModifier is a code that permits multiple files, created on the same date, same time and between the
	// same institutions, to be distinguished one from another. If all of the following fields in a previous file are
	// equal to the same fields in this file: FileHeader ImmediateDestination, ImmediateOrigin, FileCreationDate, and
	// FileCreationTime, it must be defined.
	FileIDModifier string `json:"fileIDModifier"`
	// CountryCode is a 2-character code as approved by the International Organization for Standardization (ISO) used
	// to identify the country in which the payer bank is located.
	// Example: US = United States
	// Values for other countries can be found on the International Organization for Standardization
	// website: www.iso.org.
	CountryCode string `json:"countryCode"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// CompanionDocumentIndicator identifies a field used to indicate the Companion Document being used.
	// Shall be present only under clearing arrangements. Companion Document usage and values
	// defined by clearing arrangements.
	// Values:
	// 0–9 Reserved for United States use
	// A–J Reserved for Canadian use
	// Other - as defined by clearing arrangements.
	CompanionDocumentIndicator string `json:"companionDocumentIndicator"`
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewFileHeader returns a new FileHeader with default values for non exported fields
func NewFileHeader() FileHeader {
	fh := FileHeader{
		recordType:    "01",
		StandardLevel: "35",
	}
	return fh
}

// Parse takes the input record string and parses the FileHeader values
func (fh *FileHeader) Parse(record string) {
	// Character position 1-2, Always "01"
	fh.recordType = "01"
	// 03-04
	fh.StandardLevel = fh.parseStringField(record[2:4])
	// 05-05
	fh.TestFileIndicator = fh.parseStringField(record[4:5])
	// 06-14
	fh.ImmediateDestination = fh.parseStringField(record[5:14])
	// 15-23
	fh.ImmediateOrigin = fh.parseStringField(record[14:23])
	// 24-31
	fh.FileCreationDate = fh.parseYYYYMMDDDate(record[23:31])
	// 32-35
	fh.FileCreationTime = fh.parseSimpleTime(record[31:35])
	// 36-36
	fh.ResendIndicator = fh.parseStringField(record[35:36])
	// 37-54
	fh.ImmediateDestinationName = fh.parseStringField(record[36:54])
	// 55-72
	fh.ImmediateOriginName = fh.parseStringField(record[54:72])
	// 73-73
	fh.FileIDModifier = fh.parseStringField(record[72:73])
	// 74-75
	fh.CountryCode = fh.parseStringField(record[73:75])
	// 76-79
	fh.UserField = fh.parseStringField(record[75:79])
	// 80-80
	fh.CompanionDocumentIndicator = fh.parseStringField(record[79:80])
}

// String writes the FileHeader struct to a string.
func (fh *FileHeader) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(fh.recordType)
	buf.WriteString(fh.StandardLevelField())
	buf.WriteString(fh.TestFileIndicatorField())
	buf.WriteString(fh.ImmediateDestinationField())
	buf.WriteString(fh.ImmediateOriginField())
	buf.WriteString(fh.FileCreationDateField())
	buf.WriteString(fh.FileCreationTimeField())
	buf.WriteString(fh.ResendIndicatorField())
	buf.WriteString(fh.ImmediateDestinationNameField())
	buf.WriteString(fh.ImmediateOriginNameField())
	buf.WriteString(fh.FileIDModifierField())
	buf.WriteString(fh.CountryCodeField())
	buf.WriteString(fh.UserFieldField())
	buf.WriteString(fh.CompanionDocumentIndicatorField())
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (fh *FileHeader) Validate() error {
	if err := fh.fieldInclusion(); err != nil {
		return err
	}
	if fh.recordType != "01" {
		msg := fmt.Sprintf(msgRecordType, 01)
		return &FieldError{FieldName: "recordType", Value: fh.recordType, Msg: msg}
	}
	if err := fh.isStandardLevel(fh.StandardLevel); err != nil {
		return &FieldError{FieldName: "StandardLevel", Value: fh.StandardLevel, Msg: err.Error()}
	}
	// Mandatory
	if err := fh.isTestFileIndicator(fh.TestFileIndicator); err != nil {
		return &FieldError{FieldName: "TestFileIndicator", Value: fh.TestFileIndicator, Msg: err.Error()}
	}
	// Mandatory
	if err := fh.isResendIndicator(fh.ResendIndicator); err != nil {
		return &FieldError{FieldName: "ResendIndicator", Value: fh.ResendIndicator, Msg: err.Error()}
	}
	if err := fh.isAlphanumericSpecial(fh.ImmediateDestinationName); err != nil {
		return &FieldError{FieldName: "ImmediateDestinationName", Value: fh.ImmediateDestinationName, Msg: err.Error()}
	}
	if err := fh.isAlphanumericSpecial(fh.ImmediateOriginName); err != nil {
		return &FieldError{FieldName: "ImmediateOriginName", Value: fh.ImmediateOriginName, Msg: err.Error()}
	}
	if err := fh.isAlphanumeric(fh.FileIDModifier); err != nil {
		return &FieldError{FieldName: "FileIDModifier", Value: fh.FileIDModifier, Msg: err.Error()}
	}
	// Conditional
	if fh.CountryCode == "US" {
		if err := fh.isCompanionDocumentIndicatorUS(fh.CompanionDocumentIndicator); err != nil {
			return &FieldError{FieldName: "CompanionDocumentIndicator", Value: fh.CompanionDocumentIndicator, Msg: err.Error()}
		}
	}
	// Conditional
	if fh.CountryCode == "CA" {
		if err := fh.isCompanionDocumentIndicatorCA(fh.CompanionDocumentIndicator); err != nil {
			return &FieldError{FieldName: "CompanionDocumentIndicator", Value: fh.CompanionDocumentIndicator, Msg: err.Error()}
		}
	}
	if err := fh.isAlphanumericSpecial(fh.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: fh.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (fh *FileHeader) fieldInclusion() error {
	if fh.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: fh.recordType, Msg: msgFieldInclusion}
	}
	if fh.StandardLevel == "" {
		return &FieldError{FieldName: "StandardLevel", Value: fh.StandardLevel, Msg: msgFieldInclusion}
	}
	if fh.TestFileIndicator == "" {
		return &FieldError{FieldName: "TestFileIndicator", Value: fh.TestFileIndicator, Msg: msgFieldInclusion}
	}
	if fh.ResendIndicator == "" {
		return &FieldError{FieldName: "ResendIndicator", Value: fh.ResendIndicator, Msg: msgFieldInclusion}
	}
	if fh.ImmediateDestination == "" {
		return &FieldError{FieldName: "ImmediateDestination", Value: fh.ImmediateDestination, Msg: msgFieldInclusion}
	}
	if fh.ImmediateDestination == "000000000" {
		return &FieldError{FieldName: "ImmediateDestination", Value: fh.ImmediateDestination, Msg: msgFieldInclusion}
	}
	if fh.ImmediateOrigin == "" {
		return &FieldError{FieldName: "ImmediateOrigin", Value: fh.ImmediateOrigin, Msg: msgFieldInclusion}
	}
	if fh.ImmediateOrigin == "000000000" {
		return &FieldError{FieldName: "ImmediateOrigin", Value: fh.ImmediateOrigin, Msg: msgFieldInclusion}
	}
	if fh.FileCreationDate.IsZero() {
		return &FieldError{FieldName: "FileCreationDate", Value: fh.FileCreationDate.String(), Msg: msgFieldInclusion}
	}
	if fh.FileCreationTime.IsZero() {
		return &FieldError{FieldName: "FileCreationTime", Value: fh.FileCreationTime.String(), Msg: msgFieldInclusion}
	}
	return nil

}

// StandardLevelField gets the StandardLevel field
func (fh *FileHeader) StandardLevelField() string {
	return fh.alphaField(fh.StandardLevel, 2)
}

// TestFileIndicatorField gets the TestFileIndicator field
func (fh *FileHeader) TestFileIndicatorField() string {
	return fh.alphaField(fh.TestFileIndicator, 1)
}

// ImmediateDestinationField gets the ImmediateDestination routing number field
func (fh *FileHeader) ImmediateDestinationField() string {
	return fh.stringField(fh.ImmediateDestination, 9)
}

// ImmediateOriginField gets the ImmediateOrigin routing number field
func (fh *FileHeader) ImmediateOriginField() string {
	return fh.stringField(fh.ImmediateOrigin, 9)
}

// FileCreationDateField gets the FileCreationDate field in YYYYMMDD format
func (fh *FileHeader) FileCreationDateField() string {
	return fh.formatYYYYMMDDDate(fh.FileCreationDate)
}

// FileCreationTimeField gets the FileCreationTime in HHMM format
func (fh *FileHeader) FileCreationTimeField() string {
	return fh.formatSimpleTime(fh.FileCreationTime)
}

// ResendIndicatorField gets the TestFileIndicator field
func (fh *FileHeader) ResendIndicatorField() string {
	return fh.alphaField(fh.ResendIndicator, 1)
}

// ImmediateDestinationNameField gets the ImmediateDestinationName field padded
func (fh *FileHeader) ImmediateDestinationNameField() string {
	return fh.alphaField(fh.ImmediateDestinationName, 18)
}

// ImmediateOriginNameField gets the ImmediateOriginName field padded
func (fh *FileHeader) ImmediateOriginNameField() string {
	return fh.alphaField(fh.ImmediateOriginName, 18)
}

// CountryCodeField gets the CountryCode field
func (fh *FileHeader) CountryCodeField() string {
	return fh.alphaField(fh.CountryCode, 2)
}

// UserFieldField gets the UserField field
func (fh *FileHeader) UserFieldField() string {
	return fh.alphaField(fh.UserField, 4)
}

// FileIDModifierField gets the FileIDModifier field
func (fh *FileHeader) FileIDModifierField() string {
	return fh.alphaField(fh.FileIDModifier, 1)
}

// CompanionDocumentIndicatorField gets the CompanionDocumentIndicator field
func (fh *FileHeader) CompanionDocumentIndicatorField() string {
	return fh.alphaField(fh.CompanionDocumentIndicator, 1)
}
