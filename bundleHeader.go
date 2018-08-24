// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "time"

// Errors specific to a BundleHeader Record

// BundleHeader Record
type BundleHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// recordType defines the type of record.
	recordType string
	// A code that identifies the type of bundle. It is the same value as the CollectionTypeIndicator
	// in the CashLetterHeader within which the bundle is contained, unless the CollectionTypeIndicator
	// in the CashLetterHeader is 99.
	// Values:
	// 00: Preliminary Forward Information–Used when information may change and the information is treated
	// as not final.
	// 01: Forward Presentment–For the collection and settlement of checks (demand instruments).
	// Data are treated as final.
	// 02: Forward Presentment–Same-Day Settlement–For the collection and settlement of checks (demand instruments)
	// presented under the Federal Reserve’s same day settlement amendments to Regulation CC (12CFR Part 229).
	// Data are treated as final.
	// 03: Return–For the return of check(s). Transaction carries value. Data are treated as final.
	// 04: Return Notification–For the notification of return of check(s). Transaction carries no value. The Return
	// Notification Indicator (Field 12) in the Return Record (Type 31) has to be interrogated to determine whether a
	// notice is a preliminary or final notification.
	// 05: Preliminary Return Notification–For the notification of return of check(s). Transaction carries no value.
	// Used to indicate that an item may be returned. This field supersedes the Return Notification Indicator
	// (Field 12) in the Return Record (Type 31).
	// 06: Final Return Notification–For the notification of return of check(s). Transaction carries no value. Used to
	// indicate that an item will be returned. This field supersedes the Return Notification Indicator (Field 12)
	// in the Return Record (Type 31).
	CollectionTypeIndicator string `json:"collectionTypeIndicator"`
	// DestinationRoutingNumber contains the routing and transit number of the institution that
	// receives and processes the cash letter or the bundle.  Format: TTTTAAAAC, where:
	//  TTTT Federal Reserve Prefix
	//  AAAA ABA Institution Identifier
	//  C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	DestinationRoutingNumber string `json:"destinationRoutingNumber"`
	// ECEInstitutionRoutingNumber contains the routing and transit number of the institution that
	// that creates the bundle header.  Format: TTTTAAAAC, where:
	//	TTTT Federal Reserve Prefix
	//	AAAA ABA Institution Identifier
	//	C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	ECEInstitutionRoutingNumber string `json:"eceInstitutionRoutingNumber"`
	// BundleBusinessDate is the business date of the bundle.
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	BundleBusinessDate time.Time `json:"bundleBusinessDate"`
	// BundleCreationDate is the date that the bundle is created. It is Eastern Time zone format unless
	// different clearing arrangements have been made
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	BundleCreationDate time.Time `json:"bundleCreationDate"`
	// BundleID is number that identifies the bundle, assigned by the institution that creates the bundle.
	BundleID string `json:"bundleID"`
	// BundleSequenceNumber is a number assigned by the institution that creates the bundle. Usually denotes
	// the relative position of the bundle within the cash letter.  NumericBlank
	BundleSequenceNumber int `json:"BundleSequenceNumber,omitempty"`
	// CycleNumber is a code assigned by the institution that creates the bundle.  Denotes the cycle under which
	// the bundle is created.
	CycleNumber string `json:"cycleNumber"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reservedTwo is a field reserved for future use.  Reserved should be blank.
	reservedTwo string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewBundleHeader returns a new BundleHeader with default values for non exported fields
func NewBundleHeader() *BundleHeader {
	bh := &BundleHeader{
		recordType: "20",
	}
	return bh
}

// Parse takes the input record string and parses the BundleHeader values

// String writes the BundleHeader struct to a variable length string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
