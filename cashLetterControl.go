// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
	"time"
)

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a CashLetterControl Record

//CashLetterControl Record
type CashLetterControl struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// CashLetterBundleCount identifies the total number of bundles within the cash letter.
	CashLetterBundleCount int `json:"cashLetterBundleCount"`
	// CashLetterItemsCount identifies the total number of items within the cash letter.
	CashLetterItemsCount int `json:"cashLetterItemsCount"`
	// ToDo: int64 by default on 64bit - string for 32 bit?
	// CashLetterTotalAmount identifies the total dollar value of all item amounts within the cash letter.
	CashLetterTotalAmount int `json:"cashLetterTotalAmount"`
	// CashLetterImagesCount identifies the total number of ImageViewDetail(s) within the CashLetter.
	CashLetterImagesCount int `json:"cashLetterImagesCount"`
	// ECEInstitutionName identifies the short name of the institution that creates the CashLetterControl.
	ECEInstitutionName string `json:"eceInstitutionName"`
	// SettlementDate identifies the date that the institution that creates the cash letter expects settlement.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	SettlementDate time.Time `json:"settlementDate"`
	// CreditTotalIndicator identifies a code that indicates whether Credits Items are included in the totals.
	// If so they will be included in Items CashLetterItemsCount, CashLetterTotalAmount and CashLetterImagesCount.
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

// NewCashLetterControl returns a new CashLetterControl with default values for non exported fields
func NewCashLetterControl() *CashLetterControl {
	clc := &CashLetterControl{
		recordType:     "90",
		SettlementDate: time.Now(),
		reserved:       "              ",
	}
	return clc
}

// Parse takes the input record string and parses the CashLetterControl values
func (clc *CashLetterControl) Parse(record string) {
	// Character position 1-2, Always "90"
	clc.recordType = "90"
	// 03-08
	clc.CashLetterBundleCount = clc.parseNumField(record[2:8])
	// 09-16
	clc.CashLetterItemsCount = clc.parseNumField(record[8:16])
	// 17-30
	clc.CashLetterTotalAmount = clc.parseNumField(record[16:30])
	// 31-39
	clc.CashLetterImagesCount = clc.parseNumField(record[30:39])
	// 40-57
	clc.ECEInstitutionName = clc.parseStringField(record[39:57])
	// 58-65
	clc.SettlementDate = clc.parseYYYYMMDDDate(record[57:65])
	// 66-66
	clc.CreditTotalIndicator = clc.parseNumField(record[65:66])
	// 67-80
	clc.reserved = "              "
}

// String writes the CashLetterControl struct to a string.
func (clc *CashLetterControl) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(clc.recordType)
	buf.WriteString(clc.CashLetterBundleCountField())
	buf.WriteString(clc.CashLetterItemsCountField())
	buf.WriteString(clc.CashLetterTotalAmountField())
	buf.WriteString(clc.CashLetterImagesCountField())
	buf.WriteString(clc.ECEInstitutionNameField())
	buf.WriteString(clc.SettlementDateField())
	buf.WriteString(clc.CreditTotalIndicatorField())
	buf.WriteString(clc.reservedField())
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (clc *CashLetterControl) Validate() error {
	if err := clc.fieldInclusion(); err != nil {
		return err
	}
	if clc.recordType != "90" {
		msg := fmt.Sprintf(msgRecordType, 90)
		return &FieldError{FieldName: "recordType", Value: clc.recordType, Msg: msg}
	}
	if err := clc.isAlphanumericSpecial(clc.ECEInstitutionName); err != nil {
		return &FieldError{FieldName: "ECEInstitutionName", Value: clc.ECEInstitutionName, Msg: err.Error()}
	}
	if clc.CreditTotalIndicatorField() != "" {
		if err := clc.isCreditTotalIndicator(clc.CreditTotalIndicator); err != nil {
			return &FieldError{FieldName: "CreditTotalIndicator", Value: clc.CreditTotalIndicatorField(), Msg: err.Error()}
		}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (clc *CashLetterControl) fieldInclusion() error {
	if clc.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: clc.recordType, Msg: msgFieldInclusion}
	}
	if clc.CashLetterItemsCount == 0 {
		return &FieldError{FieldName: "CashLetterItemsCount",
			Value: clc.CashLetterItemsCountField(), Msg: msgFieldInclusion}
	}
	if clc.CashLetterTotalAmount == 0 {
		return &FieldError{FieldName: "CashLetterTotalAmount",
			Value: clc.CashLetterTotalAmountField(), Msg: msgFieldInclusion}
	}
	if clc.SettlementDate.IsZero() {
		return &FieldError{FieldName: "SettlementDate",
			Value: clc.SettlementDate.String(), Msg: msgFieldInclusion}
	}
	return nil
}

// CashLetterBundleCountField gets a string of the CashLetterBundleCount zero padded
func (clc *CashLetterControl) CashLetterBundleCountField() string {
	return clc.numericField(clc.CashLetterBundleCount, 6)
}

//CashLetterItemsCountField gets a string of the CashLetterItemsCount zero padded
func (clc *CashLetterControl) CashLetterItemsCountField() string {
	return clc.numericField(clc.CashLetterItemsCount, 8)
}

// CashLetterTotalAmountField gets a string of the CashLetterTotalAmount zero padded
func (clc *CashLetterControl) CashLetterTotalAmountField() string {
	return clc.numericField(clc.CashLetterTotalAmount, 14)
}

// CashLetterImagesCountField gets a string of the CashLetterImagesCount zero padded
func (clc *CashLetterControl) CashLetterImagesCountField() string {
	return clc.numericField(clc.CashLetterImagesCount, 9)
}

// ECEInstitutionNameField gets the ECEInstitutionName field
func (clc *CashLetterControl) ECEInstitutionNameField() string {
	return clc.alphaField(clc.ECEInstitutionName, 18)
}

// SettlementDateField gets the SettlementDate in YYYYMMDD format
func (clc *CashLetterControl) SettlementDateField() string {
	return clc.formatYYYYMMDDDate(clc.SettlementDate)
}

// CreditTotalIndicatorField gets a string of the CreditTotalIndicator field
func (clc *CashLetterControl) CreditTotalIndicatorField() string {
	return clc.numericField(clc.CreditTotalIndicator, 1)
}

// reservedField gets reserved - blank space
func (clc *CashLetterControl) reservedField() string {
	return clc.alphaField(clc.reserved, 14)
}
