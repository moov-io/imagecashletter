// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"errors"
	"fmt"
)

var (
	msgCreditTotalIndicator       = "is an invalid Credit Total Indicator"
	msgDocumentationTypeIndicator = "is an invalid Documentation Type Indicator"
	msgCorrectionIndicator        = "is an invalid Correction Indicator"
	// FileHeader
	msgStandardLevel   = "is an invalid Standard Level"
	msgResendIndicator = "is an invalid Resend Indicator"
	msgTestIndicator   = "is an invalid Test Indicator"
	// CashLetterHeader
	msgCLCollectionType      = "is an invalid Collection Type Indicator"
	msgCLRecordTypeIndicator = "is an invalid Cash Letter Record Type Indicator"
	msgCLReturnsIndicator    = "is an invalid Returns Indicator"
	// CheckDetail
	msgReturnAcceptanceIndicator = "is an invalid Return Acceptance Indicator"
	msgMICRValidIndicator        = "is an invalid MICR Valid Indicator"
	msgBOFDIndicator             = "is an invalid Bank Of First Deposit Indicator"
	msgArchiveTypeIndicator      = "is an invalid Archive Type Indicator"
	// CheckDetail Addendum
	msgTruncationIndicator        = "is an invalid Truncation Indicator"
	msgConversionIndicator        = "is an invalid Conversion Indicator"
	msgImageReferenceKeyIndicator = "is an invalid Image Reference Key Indicator"
)

// validator is common validation and formatting of golang types to x9 type strings
type validator struct{}

// FieldError is returned for errors at a field level in a record
type FieldError struct {
	FieldName string // field name where error happened
	Value     string // value that cause error
	Msg       string // context of the error.
}

// Error message is constructed
// FieldName Msg Value
// Example1: BatchCount $% has none alphanumeric characters
// Example2: BatchCount 5 is out-of-balance with file count 6
func (e *FieldError) Error() string {
	return fmt.Sprintf("%s %s %s", e.FieldName, e.Value, e.Msg)
}

// isCreditTotalIndicator ensures CreditTotalIndicator of a FileControl, CashLetterControl, and BundleControl is valid
func (v *validator) isCreditTotalIndicator(code int) error {
	switch code {
	case
		// 	Credit Items are NOT included in totals
		0,
		//  Credit Items are included in totals
		1:
		return nil
	}
	return errors.New(msgCreditTotalIndicator)
}

// isDocumentationTypeIndicator ensures DocumentationTypeIndicator of a CashLetterHeader and CheckDetail is valid
func (v *validator) isDocumentationTypeIndicator(code string) error {
	switch code {
	case
		// No image provided, paper provided separately
		"A",
		// No image provided, paper provided separately, image upon request
		"B",
		// Image provided separately, no paper provided
		"C",
		// Image provided separately, no paper provided, image upon request
		"D",
		// Image and paper provided separately
		"E",
		// Image and paper provided separately, image upon request
		"F",
		// Image included, no paper provided
		"G",
		// Image included, no paper provided, image upon request
		"H",
		// Image included, paper provided separately
		"I",
		// Image included, paper provided separately, image upon request
		"J",
		// No image provided, no paper provided
		"K",
		// No image provided, no paper provided, image upon request
		"L",
		// No image provided, Electronic Check provided separately
		"M",
		// Not Same Type–Documentation associated with each item in Cash Letter will be different. The Check Detail
		// Record (Type 25) or Return Record (Type 31) has to be interrogated for further information.
		"Z":
		return nil
	}
	return errors.New(msgDocumentationTypeIndicator)
}

// ***File Header Validations***

// isStandardLevel ensures StandardLevel of a FileHeader is valid
func (v *validator) isStandardLevel(code string) error {
	switch code {
	case
		// DSTU X9.37 - 2003
		// Current Support is for 03
		"03":
		return nil
	}
	return errors.New(msgStandardLevel)
}

// isResendIndicator ensures ResendIndicator of a FileHeader is valid
func (v *validator) isResendIndicator(code string) error {
	switch code {
	case
		// The file has been previously transmitted
		"Y",
		// The file has NOT been previously transmitted
		"N":
		return nil
	}
	return errors.New(msgResendIndicator)
}

// isTestIndicator ensures TestIndicator of a FileHeader is valid
func (v *validator) isTestIndicator(code string) error {
	switch code {
	case
		// Production File
		"P",
		// Test File
		"T":
		return nil
	}
	return errors.New(msgTestIndicator)
}

// ***Cash Letter Header Validations***

// isCLCollectionType ensures CollectionTypeIndicator of a CashLetterHeader is valid
func (v *validator) isCLCollectionType(code string) error {
	switch code {
	case
		// Preliminary Forward Information–Used when information may change and the
		// information is treated as not final.
		"00",
		// Forward Presentment–For the collection and settlement of checks (demand
		// instruments). Data are treated as final.
		"01",
		// Forward Presentment–Same-Day Settlement–For the collection and settlement of
		// checks (demand instruments) presented under the Federal Reserve’s same day
		// settlement amendments to Regulation CC (12CFR Part 229). Data are treated as
		// final.
		"02",
		// 03 Return–For the return of check(s). Transaction carries value. Data are
		// treated as final.
		"03",
		// 04 Return Notification–For the notification of return of check(s). Transaction
		// carries no value. The Return Notification Indicator (Field 12) in the Return Record
		// (Type 31) has to be interrogated to determine whether a notice is a preliminary or final
		// notification.
		"04",
		// 05 Preliminary Return Notification–For the notification of return of check(s). Transaction
		// carries no value. Used to indicate that an item may be returned. This field supersedes
		// the Return Notification Indicator (Field 12) in the Return Record (Type 31).
		"05",
		// 06 Final Return Notification–For the notification of return of check(s). Transaction
		// carries no value. Used to indicate that an item will be returned. This field
		// supersedes the Return Notification Indicator (Field 12) in the Return Record (Type 31).
		"06",
		// No Detail–There are no detail records contained within the bundle or cash letter.
		// Defined Value of the Cash Letter Record Type Indicator (Field 8) shall be set to ‘N’.
		"20",
		// Bundles not the same collection type. Use of the value is only allowed by clearing
		// arrangement.
		"99":
		return nil
	}
	return errors.New(msgCLCollectionType)
}

// isCLRecordTypeIndicator ensures CashLetterRecordTypeIndicator of a CashLetterHeader is valid
func (v *validator) isCLRecordTypeIndicator(code string) error {
	switch code {
	case
		// No electronic check records or image records (Type 2x’s, 3x’s, 5x’s); e.g., an empty cash letter.
		"N",
		// Cash letter contains electronic check records with no images (Type 2x’s and 3x’s only).
		"E",
		// Cash letter contains electronic check records (Type 2x’s, 3x’s) and image records (Type 5x’s).
		"I",
		// Cash letter contains electronic check records (Type 2x’s and 3x’s) and image records (Type 5x’s)
		// that correspond to a previously sent cash letter (i.e., E file).
		"F":
		return nil
	}
	return errors.New(msgCLRecordTypeIndicator)
}

// isCLReturnsIndicator ensures ReturnsIndicator of a CashLetterHeader is valid
func (v *validator) isCLReturnsIndicator(code string) error {
	switch code {
	case
		// Blank for Forward Presentment
		"",
		// Administrative - items being returned that are handled by the bank and usually do not directly
		// affect the customer or its account.
		"E",
		// Customer–items being returned that directly affect a customer’s account.
		"R",
		// Reject Return
		"J":
		return nil
	}
	return errors.New(msgCLReturnsIndicator)
}

// CheckDetail validations

// isReturnAcceptanceIndicator ensures ReturnAcceptanceIndicator of a CheckDetail is valid
func (v *validator) isReturnAcceptanceIndicator(code string) error {
	switch code {
	case
		// Will not accept any electronic information
		"0",
		// Will accept preliminary return notifications, returns, and final return notifications
		"1",
		// Will accept preliminary return notifications and returns
		"2",
		// Will accept preliminary return notifications and final return notifications
		"3",
		// Will accept returns and final return notifications
		"4",
		// Will accept preliminary return notifications only
		"5",
		// Will accept returns only
		"6",
		// Will accept final return notifications only
		"7",
		// Will accept preliminary return notifications, returns, final return notifications, and image returns
		"8",
		// Will accept preliminary return notifications, returns and image returns
		"9",
		// Will accept preliminary return notifications, final return notifications and image returns
		"A",
		// Will accept returns, final return notifications and image returns
		"B",
		// Will accept preliminary return notifications and image returns
		"C",
		// Will accept returns and image returns
		"D",
		// Will accept final return notifications and image returns
		"E",
		// Will accept image returns only
		"F":
		return nil
	}
	return errors.New(msgReturnAcceptanceIndicator)
}

// isMICRValidIndicator ensures MICRValidIndicator of a CheckDetail is valid
func (v *validator) isMICRValidIndicator(code int) error {
	switch code {
	case
		// Good read
		1,
		// Good read, missing field
		2,
		// Read error encountered
		3,
		// Missing field and read error encountered
		4:
		return nil
	}
	return errors.New(msgMICRValidIndicator)
}

// isBOFDIndicator ensures BOFDIndicator of a CheckDetail is valid
func (v *validator) isBOFDIndicator(code int) error {
	switch code {
	case
		// Good read
		1,
		// Good read, missing field
		2,
		// Read error encountered
		3,
		// Missing field and read error encountered
		4:
		return nil
	}
	return errors.New(msgBOFDIndicator)
}

// isCorrectionIndicator ensures CorrectionIndicator of a CheckDetail is valid
func (v *validator) isCorrectionIndicator(code int) error {
	switch code {
	case
		// No Repair
		0,
		// Repaired (form of repair unknown)
		1,
		// Repaired without Operator intervention
		2,
		// Repaired with Operator intervention
		3,
		// Undetermined if repair has been done or not
		4:
		return nil
	}
	return errors.New(msgCorrectionIndicator)
}

// isArchiveTypeIndicator ensures ArchiveTypeIndicator of a CheckDetail is valid
func (v *validator) isArchiveTypeIndicator(code string) error {
	switch code {
	case
		// Microfilm
		"A",
		// B: Image
		"B",
		// C: Paper
		"C",
		// D: Microfilm and image
		"D",
		// E: Microfilm and paper
		"E",
		// F: Image and paper
		"F",
		// Microfilm, image and paper
		"G",
		// Electronic Check Instrument
		"H",
		// None
		"I":
		return nil
	}
	return errors.New(msgArchiveTypeIndicator)
}

// CheckDetail Addendum validations

// isTruncationIndicator ensures TruncationIndicator of a CheckDetailAddendumA is valid
func (v *validator) isTruncationIndicator(code string) error {
	switch code {
	case
		// This institution truncated this original check item and this is first endorsement for the institution.
		"Y",
		// This institution did not truncate the original check or, this is not the first endorsement for the
		// institution or, this item is an IRD not an original check item (EPC equals 4).
		"N":
		return nil
	}
	return errors.New(msgTruncationIndicator)
}

// isConversionIndicator ensures BOFD and Endorsing Bank ConversionIndicator of a CheckDetailAddendumA and
// CheckDetailAddendumC is valid
func (v *validator) isConversionIndicator(code string) error {
	switch code {
	case
		// Did not convert physical document
		"0",
		// Original paper converted to IRD
		"1",
		// Original paper converted to image
		"2",
		// IRD converted to another IRD
		"3",
		// IRD converted to image of IRD
		"4",
		// Image converted to an IRD
		"5",
		// Image converted to another image (e.g., transcoded)
		"6",
		// Did not convert image (e.g., same as source)
		"7",
		// Undetermined
		"8":
		return nil
	}
	return errors.New(msgConversionIndicator)
}

// isImageReferenceKeyIndicator ensures ImageReferenceKeyIndicator of a CheckDetailAddendumB is valid
func (v *validator) isImageReferenceKeyIndicator(code int) error {
	switch code {
	case
		// ImageReferenceKeyIndicator has Defined Value of 0034 and ImageReferenceKey contains the Image Reference Key.
		0,
		// ImageReferenceKeyIndicator contains a value other than Value 0034;
		// or ImageReferenceKeyIndicator contains Value 0034, which is not a Defined Value, and the content of
		// ImageReferenceKey has no special significance with regards to an Image Reference Key;
		// or ImageReferenceKeyIndicator is 0000, meaning the ImageReferenceKey is not present.
		1:
		return nil
	}
	return errors.New(msgImageReferenceKeyIndicator)
}
