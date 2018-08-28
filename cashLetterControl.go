// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "time"

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
		recordType: "90",
	}
	return clc
}

// Parse takes the input record string and parses the CashLetterControl values

// String writes the CashLetterControl struct to a string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
