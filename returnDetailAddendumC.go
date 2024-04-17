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

// Errors specific to a ReturnDetailAddendumC Record

// ReturnDetailAddendumC Record
type ReturnDetailAddendumC struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// validator is composed for imagecashletter data validation
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
	// validator is composed for imagecashletter data validation
	validator
	// converters is composed for imagecashletter to golang Converters
	converters
}

// NewReturnDetailAddendumC returns a new ReturnDetailAddendumC with default values for non exported fields
func NewReturnDetailAddendumC() ReturnDetailAddendumC {
	rdAddendumC := ReturnDetailAddendumC{}
	rdAddendumC.setRecordType()
	return rdAddendumC
}

func (rdAddendumC *ReturnDetailAddendumC) setRecordType() {
	if rdAddendumC == nil {
		return
	}
	rdAddendumC.recordType = "34"
}

// Parse takes the input record string and parses the ReturnDetailAddendumC values
func (rdAddendumC *ReturnDetailAddendumC) Parse(record string) {
	if utf8.RuneCountInString(record) < 22 {
		return // line too short
	}

	// Character position 1-2, Always "34"
	rdAddendumC.setRecordType()
	// 03-03
	rdAddendumC.ImageReferenceKeyIndicator = rdAddendumC.parseNumField(record[2:3])
	// 04-18
	rdAddendumC.MicrofilmArchiveSequenceNumber = rdAddendumC.parseStringField(record[3:18])
	// 19-22
	rdAddendumC.LengthImageReferenceKey = rdAddendumC.parseStringField(record[18:22])

	imageRefKeyLength := rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey)
	if imageRefKeyLength <= 0 || utf8.RuneCountInString(record) < 46+imageRefKeyLength {
		return // line too short
	}

	// 23 (22+X)
	rdAddendumC.ImageReferenceKey = rdAddendumC.parseStringField(record[22 : 22+imageRefKeyLength])
	// 23+X - 37+X
	rdAddendumC.Description = rdAddendumC.parseStringField(record[22+imageRefKeyLength : 37+imageRefKeyLength])
	// 38+X - 41+X
	rdAddendumC.UserField = rdAddendumC.parseStringField(record[37+imageRefKeyLength : 41+imageRefKeyLength])
	// 42+X - 46+X
	rdAddendumC.reserved = rdAddendumC.parseStringField(record[41+imageRefKeyLength : 46+imageRefKeyLength])

}

func (rdAddendumC *ReturnDetailAddendumC) UnmarshalJSON(data []byte) error {
	type Alias ReturnDetailAddendumC
	aux := struct {
		*Alias
	}{
		(*Alias)(rdAddendumC),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rdAddendumC.setRecordType()
	return nil
}

// String writes the ReturnDetailAddendumC struct to a string.
func (rdAddendumC *ReturnDetailAddendumC) String() string {
	if rdAddendumC == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(22)
	buf.WriteString(rdAddendumC.recordType)
	buf.WriteString(rdAddendumC.ImageReferenceKeyIndicatorField())
	buf.WriteString(rdAddendumC.MicrofilmArchiveSequenceNumberField())
	buf.WriteString(rdAddendumC.LengthImageReferenceKeyField())
	if size := rdAddendumC.parseNumField(rdAddendumC.LengthImageReferenceKey); validSize(size) {
		buf.Grow(size)
	}
	buf.WriteString(rdAddendumC.ImageReferenceKeyField())
	buf.WriteString(rdAddendumC.DescriptionField())
	buf.WriteString(rdAddendumC.UserFieldField())
	buf.WriteString(rdAddendumC.reservedField())
	return buf.String()
}

// Validate performs imagecashletter format rule checks on the record and returns an error if not Validated
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
		return &FieldError{FieldName: "recordType",
			Value: rdAddendumC.recordType,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumC()?"}
	}
	if rdAddendumC.MicrofilmArchiveSequenceNumberField() == "               " {
		return &FieldError{FieldName: "MicrofilmArchiveSequenceNumber",
			Value: rdAddendumC.MicrofilmArchiveSequenceNumber,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumC()?"}
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
