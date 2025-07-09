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
	// 0034: ImageReferenceKey contains the ImageReferenceKey (ImageReferenceKeyIndicator is 0).
	// 0000: ImageReferenceKey not present (ImageReferenceKeyIndicator is 1).
	// 0001 - 9999: May include Value 0034, and ImageReferenceKey has no special significance to
	// Image Reference Key (ImageReferenceKey is 1).
	LengthImageReferenceKey string `json:"imageReferenceKeyLength"`
	// ImageReferenceKey  is used to find the image of the item in the image data system.
	ImageReferenceKey string `json:"imageReferenceKey"`
	// Description describes the transaction
	Description string `json:"description"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for image cash letter data validation
	validator
	// converters is composed for image cash letter to golang Converters
	converters
}

// NewCheckDetailAddendumB returns a new CheckDetailAddendumB with default values for non exported fields
func NewCheckDetailAddendumB() CheckDetailAddendumB {
	cdAddendumB := CheckDetailAddendumB{}
	cdAddendumB.setRecordType()
	return cdAddendumB
}

func (cdAddendumB *CheckDetailAddendumB) setRecordType() {
	if cdAddendumB == nil {
		return
	}
	cdAddendumB.recordType = "27"
}

// Parse takes the input record string and parses the CheckDetailAddendumB values
func (cdAddendumB *CheckDetailAddendumB) Parse(record string) {
	if utf8.RuneCountInString(record) < 22 {
		return // line too short
	}

	// Character position 1-2, Always "27"
	cdAddendumB.setRecordType()
	// 03-03
	cdAddendumB.ImageReferenceKeyIndicator = cdAddendumB.parseNumField(record[2:3])
	// 04-18
	cdAddendumB.MicrofilmArchiveSequenceNumber = cdAddendumB.parseStringField(record[3:18])
	// 19-22
	cdAddendumB.LengthImageReferenceKey = cdAddendumB.parseStringField(record[18:22])

	imageRefLength := cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey)
	if imageRefLength <= 0 || utf8.RuneCountInString(record) < 46+imageRefLength {
		return // line too short
	}

	// 23 (22+X)
	cdAddendumB.ImageReferenceKey = cdAddendumB.parseStringField(record[22 : 22+imageRefLength])
	// 23+X - 37+X
	cdAddendumB.Description = cdAddendumB.parseStringField(record[22+imageRefLength : 37+imageRefLength])
	// 38+X - 41+X
	cdAddendumB.UserField = cdAddendumB.parseStringField(record[37+imageRefLength : 41+imageRefLength])
	// 42+X - 46+X
	cdAddendumB.reserved = cdAddendumB.parseStringField(record[41+imageRefLength : 46+imageRefLength])
}

func (cdAddendumB *CheckDetailAddendumB) UnmarshalJSON(data []byte) error {
	type Alias CheckDetailAddendumB
	aux := struct {
		*Alias
	}{
		(*Alias)(cdAddendumB),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cdAddendumB.setRecordType()
	return nil
}

// String writes the CheckDetailAddendumB struct to a string.
func (cdAddendumB *CheckDetailAddendumB) String() string {
	if cdAddendumB == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(46)
	buf.WriteString(cdAddendumB.recordType)
	buf.WriteString(cdAddendumB.ImageReferenceKeyIndicatorField())
	buf.WriteString(cdAddendumB.MicrofilmArchiveSequenceNumberField())
	buf.WriteString(cdAddendumB.LengthImageReferenceKeyField())
	if size := cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey); validSizeInt(size) {
		buf.Grow(size)
	}
	buf.WriteString(cdAddendumB.ImageReferenceKeyField())
	buf.WriteString(cdAddendumB.DescriptionField())
	buf.WriteString(cdAddendumB.UserFieldField())
	buf.WriteString(cdAddendumB.reservedField())
	return buf.String()
}

// Validate performs image cash letter format rule checks on the record and returns an error if not Validated
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
		return &FieldError{FieldName: "recordType",
			Value: cdAddendumB.recordType,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumB()?"}
	}
	if cdAddendumB.MicrofilmArchiveSequenceNumberField() == "               " {
		return &FieldError{FieldName: "MicrofilmArchiveSequenceNumber",
			Value: cdAddendumB.MicrofilmArchiveSequenceNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumB()?"}
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
	max := cdAddendumB.parseNumField(cdAddendumB.LengthImageReferenceKey)
	if !validSizeInt(max) {
		return ""
	}
	return cdAddendumB.alphaField(cdAddendumB.ImageReferenceKey, uint(max)) //nolint:gosec
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
