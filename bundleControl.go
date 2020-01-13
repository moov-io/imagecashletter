// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Errors specific to a BundleControl Record

// BundleControl Record
type BundleControl struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// BundleItemsCount identifies the total number of items within the bundle.
	BundleItemsCount int `json:"bundleitemsCount"`
	// BundleTotalAmount identifies the total amount of item amounts within the bundle.
	BundleTotalAmount int `json:"bundleTotalAmount"`
	// MICRValidTotalAmount identifies the total amount of all CheckDetail Records within the bundle which
	// contains 1 in the MICRValidIndicator .
	MICRValidTotalAmount int `json:"micrValidTotalAmount"`
	// BundleImagesCount identifies the total number of Image ViewDetail Records  within the bundle.
	BundleImagesCount int `json:"bundleImagesCount"`
	// UserField is used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// CreditTotalIndicator identifies a code that indicates whether Credits Items are included in the totals.
	// If so they will be included in Items CashLetterItemsCount, CashLetterTotalAmount and
	// CashLetterImagesCount.
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

// NewBundleControl returns a new BundleControl with default values for non exported fields
func NewBundleControl() *BundleControl {
	bc := &BundleControl{
		recordType: "70",
		reserved:   "                        ",
	}
	return bc
}

// Parse takes the input record string and parses the BundleControl values
func (bc *BundleControl) Parse(record string) {
	if utf8.RuneCountInString(record) < 56 {
		return
	}

	// Character position 1-2, Always "70"
	bc.recordType = "70"
	// 03-06
	bc.BundleItemsCount = bc.parseNumField(record[2:6])
	// 07-18
	bc.BundleTotalAmount = bc.parseNumField(record[6:18])
	// 19-30
	bc.MICRValidTotalAmount = bc.parseNumField(record[18:30])
	// 31-35
	bc.BundleImagesCount = bc.parseNumField(record[30:35])
	// 36-55
	bc.UserField = bc.parseStringField(record[35:55])
	// 56-56
	bc.CreditTotalIndicator = bc.parseNumField(record[55:56])
	// 57-80
	bc.reserved = "                        "

}

// String writes the BundleControl struct to a string.
func (bc *BundleControl) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(bc.recordType)
	buf.WriteString(bc.BundleItemsCountField())
	buf.WriteString(bc.BundleTotalAmountField())
	buf.WriteString(bc.MICRValidTotalAmountField())
	buf.WriteString(bc.BundleImagesCountField())
	buf.WriteString(bc.UserFieldField())
	buf.WriteString(bc.CreditTotalIndicatorField())
	buf.WriteString(bc.reservedField())
	return buf.String()
}

// Validate performs image cash letter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (bc *BundleControl) Validate() error {
	if err := bc.fieldInclusion(); err != nil {
		return err
	}
	if bc.recordType != "70" {
		msg := fmt.Sprintf(msgRecordType, 70)
		return &FieldError{FieldName: "recordType", Value: bc.recordType, Msg: msg}
	}
	if err := bc.isAlphanumericSpecial(bc.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: bc.UserField, Msg: err.Error()}
	}
	if bc.CreditTotalIndicatorField() != "" {
		if err := bc.isCreditTotalIndicator(bc.CreditTotalIndicator); err != nil {
			return &FieldError{FieldName: "CreditTotalIndicator", Value: bc.CreditTotalIndicatorField(), Msg: err.Error()}
		}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (bc *BundleControl) fieldInclusion() error {
	if bc.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: bc.recordType,
			Msg:   msgFieldInclusion + ", did you use BundleControl()?"}
	}
	if bc.BundleItemsCount == 0 {
		return &FieldError{FieldName: "BundleItemsCount",
			Value: bc.BundleItemsCountField(),
			Msg:   msgFieldInclusion + ", did you use BundleControl()?"}
	}
	if bc.BundleTotalAmount == 0 {
		return &FieldError{FieldName: "BundleTotalAmount",
			Value: bc.BundleTotalAmountField(),
			Msg:   msgFieldInclusion + ", did you use BundleControl()?"}
	}
	return nil
}

//BundleItemsCountField gets a string of the BundleItemsCount zero padded
func (bc *BundleControl) BundleItemsCountField() string {
	return bc.numericField(bc.BundleItemsCount, 4)
}

// BundleTotalAmountField gets a string of the BundleTotalAmount zero padded
func (bc *BundleControl) BundleTotalAmountField() string {
	return bc.numericField(bc.BundleTotalAmount, 12)
}

// MICRValidTotalAmountField gets a string of the MICRValidTotalAmount zero padded
func (bc *BundleControl) MICRValidTotalAmountField() string {
	return bc.numericField(bc.MICRValidTotalAmount, 12)
}

// BundleImagesCountField gets a string of the BundleImagesCount zero padded
func (bc *BundleControl) BundleImagesCountField() string {
	return bc.numericField(bc.BundleImagesCount, 5)
}

// UserFieldField gets the UserField field
func (bc *BundleControl) UserFieldField() string {
	return bc.alphaField(bc.UserField, 20)
}

// CreditTotalIndicatorField gets a string of the CreditTotalIndicator field
func (bc *BundleControl) CreditTotalIndicatorField() string {
	return bc.numericField(bc.CreditTotalIndicator, 1)
}

// reservedField gets reserved - blank space
func (bc *BundleControl) reservedField() string {
	return bc.alphaField(bc.reserved, 24)
}
