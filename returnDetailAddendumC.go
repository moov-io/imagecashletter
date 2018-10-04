// Copyright 2018 The Moov Authors
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
	// ToDo: Add validator
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

// NewReturnDetailAddendumC returns a new ReturnDetailAddendumC with default values for non exported fields
func NewReturnDetailAddendumC() ReturnDetailAddendumC {
	rdAddendumC := ReturnDetailAddendumC{
		recordType: "34",
	}
	return rdAddendumC
}

// Parse takes the input record string and parses the ReturnDetailAddendumC values
func (rdAddendumC *ReturnDetailAddendumC) Parse(record string) {
	// Character position 1-2, Always "34"
	rdAddendumC.recordType = "34"
	// 03-03
	rdAddendumC.ImageReferenceKeyIndicator = rdAddendumC.parseNumField(record[2:3])
	// 04-18
	rdAddendumC.MicrofilmArchiveSequenceNumber = rdAddendumC.parseStringField(record[3:18])
	// 19-22
	rdAddendumC.LengthImageReferenceKey = rdAddendumC.parseStringField(record[18:22])
	// 23 (22+X)
	rdAddendumC.ImageReferenceKey = rdAddendumC.parseStringField(record[22:rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey)])
	// 23+X - 37+X
	rdAddendumC.Description = rdAddendumC.parseStringField(record[22+rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey) : 37+rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey)])
	// 38+X - 41+X
	rdAddendumC.UserField = rdAddendumC.parseStringField(record[37+rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey) : 41+rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey)])
	// 42+X - 46+X
	rdAddendumC.reserved = rdAddendumC.parseStringField(record[41+rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey) : 46+rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey)])

}

// String writes the ReturnDetailAddendumC struct to a string.
func (rdAddendumC *ReturnDetailAddendumC) String() string {
	var buf strings.Builder
	buf.Grow(22)
	buf.WriteString(rdAddendumC.recordType)
	buf.WriteString(rdAddendumC.ImageReferenceKeyIndicatorField())
	buf.WriteString(rdAddendumC.MicrofilmArchiveSequenceNumberField())
	buf.WriteString(rdAddendumC.LengthImageReferenceKeyField())
	buf.Grow(rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey))
	buf.WriteString(rdAddendumC.ImageReferenceKeyField())
	buf.WriteString(rdAddendumC.DescriptionField())
	buf.WriteString(rdAddendumC.UserFieldField())
	buf.WriteString(rdAddendumC.reservedField())
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rdAddendumC *ReturnDetailAddendumC) Validate() error {
	if err := rdAddendumC.fieldInclusion(); err != nil {
		return err
	}
	if rdAddendumC.recordType != "34" {
		msg := fmt.Sprintf(msgRecordType, 34)
		return &FieldError{FieldName: "recordType", Value: rdAddendumC.recordType, Msg: msg}
	}
	// Mandatory
	if err := rdAddendumC.isImageReferenceKeyIndicator(rdAddendumC.ImageReferenceKeyIndicator); err != nil {
		return &FieldError{FieldName: "ImageReferenceKeyIndicator",
			Value: rdAddendumC.ImageReferenceKeyIndicatorField(), Msg: err.Error()}
	}
	if err := rdAddendumC.isAlphanumericSpecial(rdAddendumC.ImageReferenceKey); err != nil {
		return &FieldError{FieldName: "ImageReferenceKey", Value: rdAddendumC.ImageReferenceKey, Msg: err.Error()}
	}
	if err := rdAddendumC.isAlphanumericSpecial(rdAddendumC.Description); err != nil {
		return &FieldError{FieldName: "Description", Value: rdAddendumC.Description, Msg: err.Error()}
	}
	if err := rdAddendumC.isAlphanumericSpecial(rdAddendumC.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: rdAddendumC.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rdAddendumC *ReturnDetailAddendumC) fieldInclusion() error {
	if rdAddendumC.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: rdAddendumC.recordType, Msg: msgFieldInclusion}
	}
	if rdAddendumC.ImageReferenceKeyIndicatorField() == "" {
		return &FieldError{FieldName: "ImageReferenceKeyIndicator",
			Value: rdAddendumC.ImageReferenceKeyIndicatorField(), Msg: msgFieldInclusion}
	}
	if rdAddendumC.LengthImageReferenceKeyField() == "" {
		return &FieldError{FieldName: "LengthImageReferenceKey",
			Value: rdAddendumC.LengthImageReferenceKeyField(), Msg: msgFieldInclusion}
	}
	return nil
}

// ImageReferenceKeyIndicatorField gets a string of the ImageReferenceKeyIndicator field
func (rdAddendumC *ReturnDetailAddendumC) ImageReferenceKeyIndicatorField() string {
	return rdAddendumC.numericField(rdAddendumC.ImageReferenceKeyIndicator, 1)
}

// MicrofilmArchiveSequenceNumberField gets the MicrofilmArchiveSequenceNumber field
func (rdAddendumC *ReturnDetailAddendumC) MicrofilmArchiveSequenceNumberField() string {
	return rdAddendumC.alphaField(rdAddendumC.MicrofilmArchiveSequenceNumber, 15)
}

// LengthImageReferenceKeyField gets the LengthImageReferenceKey field
func (rdAddendumC *ReturnDetailAddendumC) LengthImageReferenceKeyField() string {
	return rdAddendumC.stringField(rdAddendumC.LengthImageReferenceKey, 4)
}

// ImageReferenceKeyField gets the ImageReferenceKey field
func (rdAddendumC *ReturnDetailAddendumC) ImageReferenceKeyField() string {
	return rdAddendumC.alphaField(rdAddendumC.ImageReferenceKey, uint(rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey)))
}

// DescriptionField gets the Description field
func (rdAddendumC *ReturnDetailAddendumC) DescriptionField() string {
	return rdAddendumC.alphaField(rdAddendumC.Description, 15)
}

// UserFieldField gets the UserField field
func (rdAddendumC *ReturnDetailAddendumC) UserFieldField() string {
	return rdAddendumC.alphaField(rdAddendumC.UserField, 4)
}

// reservedField gets reserved - blank space
func (rdAddendumC *ReturnDetailAddendumC) reservedField() string {
	return rdAddendumC.alphaField(rdAddendumC.reserved, 5)
}
