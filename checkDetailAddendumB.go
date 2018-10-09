// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
)

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a CheckDetailAddendumB Record

// CheckDetailAddendumB Record
type CheckDetailAddendumB struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// ImageReferenceKeyIndicator identifies whether ImageReferenceKeyLength contains a variable value within the
	// allowable range, or contains a defined value and the content is ItemReferenceKey.
	// Values:
	// 0: ImageReferenceKeyIndicator has Defined Value of 0034 and ImageReferenceKey contains the Image Reference Key.
	// 1: ImageReferenceKeyIndicator contains a value other than Value 0034;
	// or ImageReferenceKeyIndicator contains Value 0034, which is not a Defined Value, and the content of
	// ImageReferenceKey has no special significance with regards to an Image Reference Key;
	// or ImageReferenceKeyIndicator is 0000, meaning the ImageReferenceKey is not present.
	ImageReferenceKeyIndicator int `json:"imageReferenceKeyIndicator"`
	// MicrofilmArchiveSequenceNumber A number that identifies the item in the microfilm archive system;
	// it may be different than the Check Detail.ECEInstitutionItemSequenceNumber and from the ImageReferenceKey.
	MicrofilmArchiveSequenceNumber string `json:"microfilmArchiveSequenceNumber"`
	// ImageReferenceKeyLength is the number of characters in the ImageReferenceKey
	// Values:
	// ToDo: Verify defined value meaning see Annex H of the spec, add validator
	// 0034: ImageReferenceKey contains the ImageReferenceKey (ImageReferenceKeyIndicator is 0).
	// 0000: ImageReferenceKey not present (ImageReferenceKeyIndicator is 1).
	// 0001 - 9999: May include Value 0034, and ImageReferenceKey has no special significance to
	// Image Reference Key (ImageReferenceKey is 1).
	LengthImageReferenceKey string `json:"imageReferenceKeyLength"`
	// ImageReferenceKey  is used to find the image of the item in the image data system.
	ImageReferenceKey string `json:"imageReferenceKey"`
	//Description describes the transaction
	Description string `json:"description"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewCheckDetailAddendumB returns a new CheckDetailAddendumB with default values for non exported fields
func NewCheckDetailAddendumB() CheckDetailAddendumB {
	cdAddendumB := CheckDetailAddendumB{
		recordType: "27",
	}
	return cdAddendumB
}

// Parse takes the input record string and parses the CheckDetailAddendumB values
func (cdAddendumB *CheckDetailAddendumB) Parse(record string) {
	// Character position 1-2, Always "27"
	cdAddendumB.recordType = "27"
	// 03-03
	cdAddendumB.ImageReferenceKeyIndicator = cdAddendumB.parseNumField(record[2:3])
	// 04-18
	cdAddendumB.MicrofilmArchiveSequenceNumber = cdAddendumB.parseStringField(record[3:18])
	// 19-22
	cdAddendumB.LengthImageReferenceKey = cdAddendumB.parseStringField(record[18:22])
	// 23 (22+X)
	cdAddendumB.ImageReferenceKey = cdAddendumB.parseStringField(record[22:cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey)])
	// 23+X - 37+X
	cdAddendumB.Description = cdAddendumB.parseStringField(record[22+cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey) : 37+cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey)])
	// 38+X - 41+X
	cdAddendumB.UserField = cdAddendumB.parseStringField(record[37+cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey) : 41+cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey)])
	// 42+X - 46+X
	cdAddendumB.reserved = cdAddendumB.parseStringField(record[41+cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey) : 46+cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey)])
}

// String writes the CheckDetailAddendumB struct to a string.
func (cdAddendumB *CheckDetailAddendumB) String() string {
	var buf strings.Builder
	buf.Grow(46)
	buf.WriteString(cdAddendumB.recordType)
	buf.WriteString(cdAddendumB.ImageReferenceKeyIndicatorField())
	buf.WriteString(cdAddendumB.MicrofilmArchiveSequenceNumberField())
	buf.WriteString(cdAddendumB.LengthImageReferenceKeyField())
	buf.Grow(cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey))
	buf.WriteString(cdAddendumB.ImageReferenceKeyField())
	buf.WriteString(cdAddendumB.DescriptionField())
	buf.WriteString(cdAddendumB.UserFieldField())
	buf.WriteString(cdAddendumB.reservedField())
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (cdAddendumB *CheckDetailAddendumB) Validate() error {
	if err := cdAddendumB.fieldInclusion(); err != nil {
		return err
	}
	if cdAddendumB.recordType != "27" {
		msg := fmt.Sprintf(msgRecordType, 27)
		return &FieldError{FieldName: "recordType", Value: cdAddendumB.recordType, Msg: msg}
	}
	// Mandatory
	if err := cdAddendumB.isImageReferenceKeyIndicator(cdAddendumB.ImageReferenceKeyIndicator); err != nil {
		return &FieldError{FieldName: "ImageReferenceKeyIndicator",
			Value: cdAddendumB.ImageReferenceKeyIndicatorField(), Msg: err.Error()}
	}
	if err := cdAddendumB.isAlphanumericSpecial(cdAddendumB.ImageReferenceKey); err != nil {
		return &FieldError{FieldName: "ImageReferenceKey", Value: cdAddendumB.ImageReferenceKey, Msg: err.Error()}
	}
	if err := cdAddendumB.isAlphanumericSpecial(cdAddendumB.Description); err != nil {
		return &FieldError{FieldName: "Description", Value: cdAddendumB.Description, Msg: err.Error()}
	}
	if err := cdAddendumB.isAlphanumericSpecial(cdAddendumB.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: cdAddendumB.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (cdAddendumB *CheckDetailAddendumB) fieldInclusion() error {
	if cdAddendumB.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: cdAddendumB.recordType, Msg: msgFieldInclusion}
	}
	if cdAddendumB.MicrofilmArchiveSequenceNumberField() == "               " {
		return &FieldError{FieldName: "MicrofilmArchiveSequenceNumber",
			Value: cdAddendumB.MicrofilmArchiveSequenceNumber, Msg: msgFieldInclusion}
	}
	return nil
}

// ImageReferenceKeyIndicatorField gets a string of the ImageReferenceKeyIndicator field
func (cdAddendumB *CheckDetailAddendumB) ImageReferenceKeyIndicatorField() string {
	return cdAddendumB.numericField(cdAddendumB.ImageReferenceKeyIndicator, 1)
}

// MicrofilmArchiveSequenceNumberField gets the MicrofilmArchiveSequenceNumber field
func (cdAddendumB *CheckDetailAddendumB) MicrofilmArchiveSequenceNumberField() string {
	return cdAddendumB.alphaField(cdAddendumB.MicrofilmArchiveSequenceNumber, 15)
}

// LengthImageReferenceKeyField gets the LengthImageReferenceKey field
func (cdAddendumB *CheckDetailAddendumB) LengthImageReferenceKeyField() string {
	return cdAddendumB.stringField(cdAddendumB.LengthImageReferenceKey, 4)
}

// ImageReferenceKeyField gets the ImageReferenceKey field
func (cdAddendumB *CheckDetailAddendumB) ImageReferenceKeyField() string {
	return cdAddendumB.alphaField(cdAddendumB.ImageReferenceKey, uint(cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey)))
}

// DescriptionField gets the Description field
func (cdAddendumB *CheckDetailAddendumB) DescriptionField() string {
	return cdAddendumB.alphaField(cdAddendumB.Description, 15)
}

// UserFieldField gets the UserField field
func (cdAddendumB *CheckDetailAddendumB) UserFieldField() string {
	return cdAddendumB.alphaField(cdAddendumB.UserField, 4)
}

// reservedField gets reserved - blank space
func (cdAddendumB *CheckDetailAddendumB) reservedField() string {
	return cdAddendumB.alphaField(cdAddendumB.reserved, 5)
}
