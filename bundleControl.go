// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a BundleControl Record

// BundleControl Record
type BundleControl struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// BundleItemsCount identifies the total number of items within the bundle.
	BundleItemsCount int `json:"bundleitemsCount"`
	// ToDo: int64 by default on 64bit - string for 32 bit?
	// BundleTotalAmount identifies the total amount of item amounts within the bundle.
	BundleTotalAmount int `json:"bundleTotalAmount"`
	// ToDo: int64 by default on 64bit - string for 32 bit?
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
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewBundleControl returns a new BundleControl with default values for non exported fields
func NewBundleControl() *BundleControl {
	bc := &BundleControl{
		recordType: "70",
	}
	return bc
}

// Parse takes the input record string and parses the BundleControl values

// String writes the BundleControl struct to a variable length string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
