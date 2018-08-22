// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "time"

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a CheckDetailAddendumC Record

// CheckDetailAddendumC Record
type CheckDetailAddendumC struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// RecordNumber is a number representing the order in which each CheckDetailAddendumC was created.
	// CheckDetailAddendumC shall be in sequential order starting with 1.  Maximum 99,
	RecordNumber int `json:"recordNumber"`
	// RoutingNumber (Endorsing Bank Routing Number) is valid routing and transit number indicating the bank that
	// endorsed the check.
	// Format: TTTTAAAAC, where:
	// TTTT Federal Reserve Prefix
	// AAAA ABA Institution Identifier
	// C Check Digit
	// For a number that identifies a non-financial institution: NNNNNNNNN
	RoutingNumber string `json:"routingNumber"`
	// BOFDEndorsementBusinessDate is the business date the check was endorsed.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	BOFDEndorsementBusinessDate time.Time `json:"bofdEndorsementBusinessDate"`
	// EndorsingItemSequenceNumber is a number that identifies the item at the endorsing bank.
	EndorsingItemSequenceNumber string `json:"endorsingItemSequenceNumber"`
	// TruncationIndicator identifies if the institution truncated the original check item.
	// Values: Y: Yes this institution truncated this original check item and this is first endorsement
	// for the institution.
	// N: No this institution did not truncate the original check or, this is not the first endorsement for the
	// institution or, this item is an IRD not an original check item (EPC equals 4).
	TruncationIndicator string `json:"truncationIndicator"`
	// BOFDConversionIndicator is a code that indicates the conversion within the processing institution among
	// original paper check, image and IRD. The indicator is specific to the action of institution identified in the
	// Endorsing Bank RoutingNumber.
	// Values:
	// 0: Did not convert physical document
	// 1: Original paper converted to IRD
	// 2: Original paper converted to image
	// 3: IRD converted to another IRD
	// 4: IRD converted to image of IRD
	// 5: Image converted to an IRD
	// 6: Image converted to another image (e.g., transcoded)
	// 7: Did not convert image (e.g., same as source)
	// 8: Undetermined
	BOFDConversionIndicator string `json:"endorsingConversionIndicator"`
	// EndorsingCorrectionIndicator identifies whether and how the MICR line of this item was repaired by the
	// creator of this CheckDetailAddendumA Record for fields other than Payor Bank Routing Number and Amount.
	// Values:
	// 0: No Repair
	// 1: Repaired (form of repair unknown)
	// 2: Repaired without Operator intervention
	// 3: Repaired with Operator intervention
	// 4: Undetermined if repair has been done or no
	EndorsingCorrectionIndicator int `json:"endorsingCorrectionIndicator"`
	// ReturnReason is a code that indicates the reason for non-payment.
	ReturnReason string `json:"returnReason"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	//EndorsingBankIdentifier
	// Values:
	// 0: Depository Bank (BOFD) - this value is used when the CheckDetailAddendumA Record reflects the Return
	// Processing Bank in lieu of BOFD.
	// 1: Other Collecting Bank
	// 2: Other Returning Bank
	// 3: Payor Bank
	EndorsingBankIdentifier string `json:"endorsingBankIdentifier"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewCheckDetailAddendumC returns a new CheckDetailAddendumC with default values for non exported fields
func NewCheckDetailAddendumC() *CheckDetailAddendumC {
	checkAddendumC := &CheckDetailAddendumC{
		recordType: "28",
	}
	return checkAddendumC
}

// Parse takes the input record string and parses the CheckDetailAddendumC values

// String writes the CheckDetailAddendumC struct to a variable length string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
