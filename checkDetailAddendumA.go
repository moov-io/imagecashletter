// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "time"

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a CheckDetailAddendumA Record

// CheckDetailAddendumA Record
type CheckDetailAddendumA struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// recordNumber is a number representing the order in which each CheckDetailAddendumA was created.
	// CheckDetailAddendumA shall be in sequential order starting with 1. Maximum 99.
	RecordNumber int `json:"recordNumber"`
	// RoutingNumber (Return Location Routing Number) is valid routing and transit number indicating where returns,
	// final return notifications, and preliminary return notifications are sent, usually the BOFD.
	// Format: TTTTAAAAC, where:
	// TTTT Federal Reserve Prefix
	// AAAA ABA Institution Identifier
	// C Check Digit
	// For a number that identifies a non-financial institution: NNNNNNNNN
	RoutingNumber string `json:"routingNumber"`
	// BOFDEndorsementDate is the date of endorsement.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	BOFDEndorsementDate time.Time `json:"bofdEndorsementDate"`
	// BOFDItemSequenceNumber is a number that identifies the item in the CheckDetailAddendumA.
	BOFDItemSequenceNumber string `json:"bofdItemSequenceNumber"`
	// BOFDAccountNumber is a number that identifies the depository account at the Bank of First Deposit.
	BOFDAccountNumber string `json:"bofdAccountNumber"`
	// BOFDBranchCode is a code that identifies the branch at the Bank of First Deposit.
	BOFDBranchCode string `json:"bofdBranchCode"`
	// PayeeName is the name of the payee from the check.
	PayeeName string `json:"payeeName"`
	// TruncationIndicator identifies if the institution truncated the original check item.
	// Values: Y: Yes this institution truncated this original check item and this is first endorsement
	// for the institution.
	// N: No this institution did not truncate the original check or, this is not the first endorsement for the
	// institution or, this item is an IRD not an original check item (EPC equals 4).
	TruncationIndicator string `json:"truncationIndicator"`
	// BOFDConversionIndicator is a code that indicates the conversion within the processing institution between
	// original paper check, image and IRD. The indicator is specific to the action of institution that created
	// this record.
	//Values:
	// 0: Did not convert physical document
	// 1: Original paper converted to IRD
	// 2: Original paper converted to image
	// 3: IRD converted to another IRD
	// 4: IRD converted to image of IRD
	// 5: Image converted to an IRD
	// 6: Image converted to another image (e.g., transcoded)
	// 7: Did not convert image (e.g., same as source)
	// 8: Undetermined
	BOFDConversionIndicator string `json:"BOFDConversionIndicator"`
	// BOFDCorrectionIndicator identifies whether and how the MICR line of this item was repaired by the
	// creator of this CheckDetailAddendumA Record for fields other than Payor Bank Routing Number and Amount.
	// Values:
	// 0: No Repair
	// 1: Repaired (form of repair unknown)
	// 2: Repaired without Operator intervention
	// 3: Repaired with Operator intervention
	// 4: Undetermined if repair has been done or not
	BOFDCorrectionIndicator int `json:"BOFDCorrectionIndicator"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewCheckDetailAddendumA returns a new CheckDetailAddendumA with default values for non exported fields
func NewCheckDetailAddendumA() *CheckDetailAddendumA {
	checkAddendumA := &CheckDetailAddendumA{
		recordType: "26",
	}
	return checkAddendumA
}

// Parse takes the input record string and parses the CheckDetailAddendumA values

// String writes the CheckDetailAddendumA struct to a string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
