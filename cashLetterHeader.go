// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "time"

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a CashLetterHeader Record

// CashLetterHeader Record is mandatory.
type CashLetterHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record
	// Value: 10
	recordType string
	// CollectionTypeIndicator is a code that identifies the type of cash letter.
	// Values:
	// 00: Preliminary Forward Information–Used when information may change and the
	// information is treated as not final.
	// 01: Forward Presentment–For the collection and settlement of checks (demand
	// instruments). Data are treated as final.
	// 02: Forward Presentment–Same-Day Settlement–For the collection and settlement of
	// checks (demand instruments) presented under the Federal Reserve’s same day
	// settlement amendments to Regulation CC (12CFR Part 229). Data are treated as
	// final.
	// 03: Return–For the return of check(s). Transaction carries value. Data are
	// treated as final.
	// 04: Return Notification–For the notification of return of check(s). Transaction
	// carries no value. The Return Notification Indicator (Field 12) in the Return Record
	// (Type 31) has to be interrogated to determine whether a notice is a preliminary or final
	// notification.
	// 05: Preliminary Return Notification–For the notification of return of check(s). Transaction
	// carries no value. Used to indicate that an item may be returned. This field supersedes
	// the Return Notification Indicator (Field 12) in the Return Record (Type 31).
	// 06: Final Return Notification–For the notification of return of check(s). Transaction
	// carries no value. Used to indicate that an item will be returned. This field
	// supersedes the Return Notification Indicator (Field 12) in the Return Record (Type 31).
	// 20: No Detail–There are no detail records contained within the bundle or cash letter.
	// Defined Value of the Cash Letter Record Type Indicator (Field 8) shall be set to ‘N’.
	// 99: Bundles not the same collection type. Use of the value is only allowed by clearing
	// arrangement.
	CollectionTypeIndicator string `json:"collectionTypeIndicator"`
	// DestinationRoutingNumber contains the routing and transit number of the institution that
	// receives and processes the cash letter or the bundle.  Format: TTTTAAAAC, where:
	//  TTTT Federal Reserve Prefix
	//  AAAA ABA Institution Identifier
	//  C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	DestinationRoutingNumber string `json:"destinationRoutingNumber"`
	// ECEInstitutionRoutingNumber contains the routing and transit number of the institution that
	// that creates the Cash Letter Header Record.  Format: TTTTAAAAC, where:
	//  TTTT Federal Reserve Prefix
	//  AAAA ABA Institution Identifier
	//  C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	ECEInstitutionRoutingNumber string `json:"eceInstitutionRoutingNumber"`
	// CashLetterBusinessDate is the business date of the cash letter.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	CashLetterBusinessDate time.Time `json:"cashLetterBusinessDate"`
	// CashLetterCreationDate is the date that the cash letter is created which shall be in Eastern
	// Time zone format. Other time zones may be used under clearing arrangements.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	CashLetterCreationDate time.Time `json:"cashLetterCreationDate"`
	// CashLetterCreationTime is the time that the cash letter is created.  Default time shall be in
	// Eastern Time zone format. Other time zones may be used under clearing arrangements.
	// Format: hhmm, where: hh hour, mm minute
	// Values:
	// hh '00' through '23'
	// mm '00' through '59'
	CashLetterCreationTime time.Time `json:"cashLetterCreationTime"`
	// CashLetterRecordTypeIndicator is a code that indicates the presence of records or the type of
	// records contained in the cash letter.   If an image is associated with any Check Detail Record
	// (Type 25) or Return Record (Type 31), the cash letter must have a Cash Letter Record Type Indicator
	// of I or F.
	// Values:
	// N: No electronic check records or image records (Type 2x’s, 3x’s, 5x’s); e.g., an empty cash letter.
	// E: Cash letter contains electronic check records with no images (Type 2x’s and 3x’s only).
	// I: Cash letter contains electronic check records (Type 2x’s, 3x’s) and image records (Type 5x’s).
	// F: Cash letter contains electronic check records (Type 2x’s and 3x’s) and image records (Type 5x’s)
	// that correspond to a previously sent cash letter (i.e., E file).
	//
	// The fields in this file that contain posting data shall not be changed from the previously sent CashLetter
	// with CollectionTypeIndicator values of 01, 02 or 03. ItemsCount and TotalAmount of the CashLetterControl with
	// a RecordTypeIndicator value of F must equal the corresponding fields in a CashLetter with a RecordTypeIndicator
	// value of E.
	CashLetterRecordTypeIndicator string `json:"cashLetterRecordTypeIndicator"`
	// CashLetterDocumentationTypeIndicator is a code that indicates the type of documentation that supports
	// all check records in the cash letter
	// Values:
	// A: No image provided, paper provided separately
	// B: No image provided, paper provided separately, image upon request
	// C: Image provided separately, no paper provided
	// D: Image provided separately, no paper provided, image upon request
	// E: Image and paper provided separately
	// F: Image and paper provided separately, image upon request
	// G: Image included, no paper provided
	// H: Image included, no paper provided, image upon request
	// I: Image included, paper provided separately
	// J: Image included, paper provided separately, image upon request
	// K: No image provided, no paper provided
	// L: No image provided, no paper provided, image upon request
	// M: No image provided, Electronic Check provided separately
	// Z: Not Same Type–Documentation associated with each item in Cash Letter will be different. The Check Detail
	// Record (Type 25) or Return Record (Type 31) has to be interrogated for further information.
	CashLetterDocumentationTypeIndicator string `json:"cashLetterdocumentationTypeIndicator"`
	// OriginatorContactName is the name of contact at the institution that creates the cash letter.
	OriginatorContactName string `json:"originatorContactName"`
	// OriginatorContactPhoneNumber is the phone number of the contact at the institution that creates
	// the cash letter.
	OriginatorContactPhoneNumber string `json:"originatorContactPhoneNumber"`
	// FedWorkType is any valid codes specified by the Federal Reserve Bank.
	FedWorkType string `json:"fedWorkType"`
	// ReturnsIndicator identifies type pf returns.
	// Values:
	// "": Blank for Forward Presentment
	// E: Administrative - items being returned that are handled by the bank and usually do not directly
	// affect the customer or its account.
	// R: Customer–items being returned that directly affect a customer’s account.
	// J: Reject Return
	ReturnsIndicator string `json:"returnsIndicator"`
	// UserField is a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewCashLetterHeader returns a new CashLetterHeader with default values for non exported fields
func NewCashLetterHeader() *CashLetterHeader {
	clh := &CashLetterHeader{
		recordType: "10",
	}
	return clh
}

// Parse takes the input record string and parses the CashLetterHeader values

// String writes the CashLetterHeader struct to a variable length string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
