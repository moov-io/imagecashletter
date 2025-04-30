// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	// upperAlphanumericRegex    = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	alphanumericRegex        = regexp.MustCompile(`[^ a-zA-Z0-9]`)
	alphanumericRegexSpecial = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	numericRegex             = regexp.MustCompile(`[^ 0-9]`)
	msgAlphanumeric          = "has non alphanumeric characters"
	msgAlphanumericSpecial   = "has non alphanumeric or special characters"
	// msgUpperAlpha             = "is not uppercase A-Z or 0-9"
	msgNumeric        = "is not 0-9"
	msgFieldInclusion = "is a mandatory field and has a default value"
	// msgValidFieldLength    = "is not length %d"
	msgInvalid     = "is invalid"
	msgInvalidDate = "is not a valid date"
)

const (
	// maxBufferGrowth is the high limit for growing string builders and byte buffers.
	//
	// 1e8/1024/1024 is ~95MB which should be enough for anybody
	maxBufferGrowth = 1e8
)

func validSizeInt(n int) bool {
	return n > 0 && n < maxBufferGrowth
}

func validSizeUint(n uint) bool {
	return n < maxBufferGrowth
}

// validator is common validation and formatting of golang types to imagecashletter type strings
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
	return errors.New(msgInvalid)
}

// isDocumentationTypeIndicator ensures DocumentationTypeIndicator of a CashLetterHeader and CheckDetail is valid
func (v *validator) isDocumentationTypeIndicator(code string) error {
	switch code {
	case
		// Conditional value, blank/space indicates no DocumentationTypeIndicator
		"",
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
		// Not Same Type–Documentation associated with each item in Cash Letter will be different. The CheckDetail
		// Record or ReturnRecord has to be interrogated for further information.
		// Z is only valid for CashLetter
		"Z":
		return nil
	}
	return errors.New(msgInvalid)
}

// ***File Header Validations***

// isStandardLevel ensures StandardLevel of a FileHeader is valid
func (v *validator) isStandardLevel(code string) error {
	switch code {
	case
		// 03: DSTU X9.37 - 2003
		"03",
		// 30: X9.100-187-2008
		"30",
		// 35: X9.100-187-2013 and 2016
		"35":
		return nil
	}
	return errors.New(msgInvalid)
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
	return errors.New(msgInvalid)
}

// isTestFileIndicator ensures TestFileIndicator of a FileHeader is valid
func (v *validator) isTestFileIndicator(code string) error {
	switch code {
	case
		// Production File
		"P",
		// Test File
		"T":
		return nil
	}
	return errors.New(msgInvalid)
}

// isCompanionDocumentIndicatorUS ensures CompanionDocumentIndicatorUS of a FileHeader is valid
func (v *validator) isCompanionDocumentIndicatorUS(code string) error {
	switch code {
	case
		// Conditional value, blank/space indicates no CompanionDocumentIndicator
		"",
		// 0–9 Reserved for United States use
		"0", "1", "2", "3", "4", "5", "6", "7":
		// Other - as defined by clearing arrangements. - Not implemented
		return nil
	}
	return errors.New(msgInvalid)
}

// isCompanionDocumentIndicatorCA ensures CompanionDocumentIndicatorCA of a FileHeader is valid
func (v *validator) isCompanionDocumentIndicatorCA(code string) error {
	switch code {
	case
		// Conditional value, blank/space indicates no CompanionDocumentIndicator
		"",
		// A-J Reserved for Canadian use
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J":
		// Other - as defined by clearing arrangements. - Not implemented
		return nil
	}
	return errors.New(msgInvalid)
}

// ***Cash Letter Header Validations***

// isCollectionTypeIndicator ensures CollectionTypeIndicator of a CashLetterHeader is valid
func (v *validator) isCollectionTypeIndicator(code string) error {
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
	return errors.New(msgInvalid)
}

// isRecordTypeIndicator ensures CashLetterRecordTypeIndicator of a CashLetterHeader is valid
func (v *validator) isRecordTypeIndicator(code string) error {
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
	return errors.New(msgInvalid)
}

// isReturnsIndicator ensures ReturnsIndicator of a CashLetterHeader is valid
func (v *validator) isReturnsIndicator(code string) error {
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
		"J",
		// Altered/Fictitious Item/Suspected Counterfeit/Counterfeit
		"N":
		return nil
	}
	return errors.New(msgInvalid)
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
	return errors.New(msgInvalid)
}

// isMICRValidIndicator ensures MICRValidIndicator of a CheckDetail is valid
func (v *validator) isMICRValidIndicator(code int) error {
	switch code {
	case
		// Not specified (this field is conditional "Shall be present only under clearing arrangements")
		0,
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
	return errors.New(msgInvalid)
}

// isBOFDIndicator ensures BOFDIndicator of a CheckDetail is valid
func (v *validator) isBOFDIndicator(code string) error {
	switch code {
	case
		// ECE institution is BOFD
		"Y",
		// ECE institution is not BOFD
		"N",
		// ECE institution relationship to BOFD is undetermined
		"U":
		return nil
	}
	return errors.New(msgInvalid)
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
	return errors.New(msgInvalid)
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
	return errors.New(msgInvalid)
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
	return errors.New(msgInvalid)
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
	return errors.New(msgInvalid)
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
	return errors.New(msgInvalid)
}

// isEndorsingBankIdentifier ensures EndorsingBankIdentifier of a CheckDetailAddendumC is valid
func (v *validator) isEndorsingBankIdentifier(code int) error {
	switch code {
	case
		// Depository Bank (BOFD) - this value is used when the CheckDetailAddendumC Record reflects the Return
		// Processing Bank in lieu of BOFD.
		0,
		// Other Collecting Bank
		1,
		// Other Returning Bank
		2,
		// Payor Bank
		3:
		return nil
	}
	return errors.New(msgInvalid)
}

// ImageView

// isImageIndicator ensures ImageIndicator of a ImageViewDetail is valid
func (v *validator) isImageIndicator(code int) error {
	switch code {
	case
		// Image view not present
		0,
		// Image view present, actual check
		1,
		// Image view present, not actual check
		2,
		// Image view present, unable to determine if value is 1 or 2
		3:
		return nil
	}
	return errors.New(msgInvalid)
}

// isImageViewFormatIndicator ensures ImageViewFormatIndicator of a ImageViewDetail is valid
func (v *validator) isImageViewFormatIndicator(code string) error {
	switch code {
	case
		// Agreement not required:
		// 00: TIFF 6; Extension: TIF
		"00",
		// Agreement required:
		// 01: IOCA FS 11; Extension: ICA
		"01",
		// 20: PNG (Portable Network Graphics); Extension: PNG ‘21’	JFIF (JPEG File Interchange Format); Extension: JPG
		"20",
		// 22: SPIFF (Still Picture Interchange File Format) (ITU-T Rec. T.84 Annex F); Extension: SPF
		"22",
		// 23: JBIG data stream (ITU-T Rec. T.82/ISO/IEC 11544:1993);
		// Extension: JBG ‘24’	JPEG 2000 (ISO/IEC 15444-1:2000);
		// Extension: JP2
		"23":
		return nil
	}
	return errors.New(msgInvalid)
}

// isImageViewCompressionAlgorithm ensures ImageViewCompressionAlgorithm of a ImageViewDetail is valid
func (v *validator) isImageViewCompressionAlgorithm(code string) error {
	switch code {
	case
		// Agreement not required:
		// 00: Group 4 facsimile compression (ITU-T Rec. T.563/CCITT Rec. T.6)
		"00",
		// Agreement required:
		// 01: JPEG Baseline (JPEG Interchange Format) (ITU-T Rec. T.81/ISO/IEC 10918)
		"01",
		// 02: ABIC
		"02",
		// 21: PNG (Portable Network Graphics)
		"21",
		// 22: JBIG (ITU-T Rec. T.82/ISO/IEC 11544:1993)
		"22",
		// 23: JPEG 2000 (ISO/IEC 15444–1:2000)
		"23":
		return nil
	}
	return errors.New(msgInvalid)
}

// isViewSideIndicator ensures ViewSideIndicator of a ImageViewDetail is valid
func (v *validator) isViewSideIndicator(code int) error {
	switch code {
	case
		// Front image view
		0,
		// Rear image view
		1:
		return nil
	}
	return errors.New(msgInvalid)
}

// isViewDescriptor ensures ViewDescriptor of a ImageViewDetail is valid
func (v *validator) isViewDescriptor(code string) error {
	switch code {
	case
		// Full view
		"00",
		// Partial view–unspecified Area of Interest
		"01",
		// Partial view–date Area of Interest
		"02",
		// Partial view–payee Area of Interest
		"03",
		// Partial view–convenience amount Area of Interest
		"04",
		// Partial view–amount in words (legal amount) Area of Interest
		"05",
		// Partial view–signature Area(s) of Interest
		"06",
		// Partial view–payor name and address Area of Interest
		"07",
		// Partial view–MICR line Area of Interest
		"08",
		// Partial view–memo line Area of Interest
		"09",
		// Partial view–payor bank name and address Area of Interest
		"10",
		// Partial view–payee endorsement Area of Interest
		"11",
		// Partial view–Bank Of First Deposit (BOFD) endorsement Area of Interest
		"12",
		// Partial view–transit endorsement Area of Interest
		"13",
		// 14 - 99 Reserved for X9
		"14":
		return nil
	}
	return errors.New(msgInvalid)
}

// isDigitalSignatureIndicator ensures DigitalSignatureIndicator of a ImageViewDetail is valid
func (v *validator) isDigitalSignatureIndicator(code int) error {
	switch code {
	case
		// Digital Signature is not present
		0,
		// Digital Signature is present
		1:
		return nil
	}
	return errors.New(msgInvalid)
}

// isDigitalSignatureMethod ensures DigitalSignatureMethod of a ImageViewDetail is valid
func (v *validator) isDigitalSignatureMethod(code string) error {
	switch code {
	case
		// 00: Digital Signature Algorithm (DSA) with SHA1 (ANSI X9.30)
		"00",
		// 01: RSA with MD5 (ANSI X9.31)
		"01",
		// 02: RSA with MDC2 (ANSI X9.31)
		"02",
		// 03: RSA with SHA1 (ANSI X9.31)
		"03",
		// 04: Elliptic Curve DSA (ECDSA) with SHA1 (ANSI X9.62)
		"04",
		// 05 - 99: Reserved for emerging cryptographic algorithms.
		"05":
		return nil
	}
	return errors.New(msgInvalid)
}

// isImageRecreateIndicator ensures ImageRecreateIndicator of a ImageViewDetail is valid
func (v *validator) isImageRecreateIndicator(code int) error {
	switch code {
	case
		// Sender can recreate the image view for the duration of the agreed upon retention time frames.
		0,
		// Sender cannot recreate image view.
		1:
		return nil
	}
	return errors.New(msgInvalid)
}

// isOverrideIndicator ensures OverrideIndicator of a ImageViewDetail is valid
func (v *validator) isOverrideIndicator(code string) error {
	switch code {
	case
		// blank/space indicates no observed image test failure present
		"",
		// No override information for this view or not applicable
		"0",
		// Imperfect image
		"1",
		// IQA Fail–Image view reviewed and deemed usable—no alternate format
		"A",
		// B: IQA Fail–Image view reviewed and deemed usable—alternate format included in this file
		"B",
		// C: IQA Fail–Image view reviewed and deemed usable–alternate format included in this file and original
		// document available
		"C",
		// D: IQA Fail–Image view reviewed and deemed usable–alternate format available
		"D",
		// E: IQA Fail–Image view reviewed and deemed usable–original document available
		"E",
		// F: IQA Fail–Image view reviewed and deemed usable–original document and alternate format available
		"F",
		// G: IQA Fail–Image view reviewed and deemed unusable–no alternate format
		"G",
		// H: IQA Fail–Image view reviewed and deemed unusable–alternate format included in this file
		"H",
		// I: IQA Fail–Image view reviewed and deemed unusable–alternate format included in this file and original
		// document available
		"I",
		// J: IQA Fail–Image view reviewed and deemed unusable–alternate format available
		"J",
		// K: IQA Fail–Image view reviewed and deemed unusable–original document available
		"K",
		// L: IQA Fail–Image view reviewed and deemed unusable–original document and alternate format available
		"L",
		// M: IQA Fail–Image view not reviewed–no alternate format
		"M",
		// N: IQA Fail–Image view not reviewed–alternate format included in this file
		"N",
		// IQA Fail–Image view not reviewed–alternate format included in this file and original
		"O":
		return nil
	}
	return errors.New(msgInvalid)
}

// isImageViewAnalysisValid ensures generic properties of imageViewAnalysis are valid
func (v *validator) isImageViewAnalysisValid(code string) error {
	switch code {
	case
		"",
		// Refer to ImageViewAnalysis property
		"0",
		// Refer to ImageViewAnalysis property
		"1",
		// Refer to ImageViewAnalysis property
		"2":
		return nil
	}
	return errors.New(msgInvalid)
}

// Returns

// isReturnNotificationIndicator ensures ReturnNotificationIndicator of ReturnDetail is valid
func (v *validator) isReturnNotificationIndicator(code string) error {
	switch code {
	case
		// Preliminary notification
		"1",
		// Final notification
		"2":
		return nil
	}
	return errors.New(msgInvalid)
}

// isTimesReturned ensures TimeReturned of ReturnDetail is valid
func (v *validator) isTimesReturned(code int) error {
	switch code {
	case
		// The item has been returned an unknown number of times
		0,
		// The item has been returned once
		1,
		// The item has been returned twice
		2,
		// The item has been returned three times
		3:
		return nil
	}
	return errors.New(msgInvalid)
}

// isAccountTypeCode ensures AccountTypeCode of CheckItem is valid
func (v *validator) isAccountTypeCode(code string) error {
	switch code {
	case
		// Unknown
		"0",
		// DDA account
		"1",
		// General Ledger account
		"2",
		// Savings account
		"3",
		// Money Market account
		"4",
		// Other account
		"5",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J":
		return nil
	}
	return errors.New(msgInvalid)
}

// isSourceWorkCode ensures SourceWorkCode of CheckItem is valid
func (v *validator) isSourceWorkCode(code string) error {
	switch code {
	case
		// Unknown
		"00",
		// Internal–ATM
		"01",
		// Internal–Branch
		"02",
		// Internal–Other
		"03",
		// External–Bank to Bank (Correspondent)
		"04",
		// External–Business to Bank (Customer)
		"05",
		// External–Business to Bank Remote Capture
		"06",
		// External–Processor to Bank
		"07",
		// External–Bank to Processor
		"08",
		// Lockbox
		"09",
		// International–Internal
		"10",
		// International–External
		"11",
		// User Defined
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
		"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
		"41", "42", "43", "44", "45", "46", "47", "48", "49", "50":
		return nil
	}
	return errors.New(msgInvalid)
}

// isOwnerIdentifierIndicator ensures OwnerIdentifierIndicator of User* is valid
func (v *validator) isOwnerIdentifierIndicator(code int) error {
	switch code {
	case
		// Not Used
		0,
		// Routing Number
		1,
		// DUNS Number
		2,
		// Federal Tax Identification Number
		3,
		// X9 Assignment
		4,
		// Other
		5:
		return nil
	}
	return errors.New(msgInvalid)
}

// isEndorsementIndicator ensures EndorsementIndicator of UserPayeeEndorsement is valid
func (v *validator) isEndorsementIndicator(code int) error {
	switch code {
	case
		// Endorsed in Blank–Instrument becomes payable to bearer
		0,
		// For Deposit Only
		1,
		// For Collection Only
		2,
		// Anomalous Endorsement–Endorsement made by person who is not holder of instrument
		3,
		// Restrictive Endorsement–Limiting to a particular person or situation
		4,
		// Guaranteed Endorsement–Deposit to the account of within named payee absence of endorsement guaranteed by
		// the bank whose Routing Number appears in BankRoutingNumber
		5,
		// Other
		9:
		return nil
	}
	return errors.New(msgInvalid)
}

/*// isUpperAlphanumeric checks if string only contains ASCII alphanumeric upper case characters
func (v *validator) isUpperAlphanumeric(s string) error {
	if upperAlphanumericRegex.MatchString(s) {
		return errors.New(msgUpperAlpha)
	}
	return nil
}
*/
// isAlphanumeric checks if a string only contains ASCII alphanumeric characters
func (v *validator) isAlphanumeric(s string) error {
	if alphanumericRegex.MatchString(s) {
		// ^[ A-Za-z0-9_@./#&+-]*$/
		return errors.New(msgAlphanumeric)
	}
	return nil
}

// isAlphanumericSpecial checks if a string only contains ASCII alphanumeric or special characters
func (v *validator) isAlphanumericSpecial(s string) error {
	if alphanumericRegexSpecial.MatchString(s) {
		// ^[ A-Za-z0-9_@./#&+-]*$/
		return errors.New(msgAlphanumericSpecial)
	}
	return nil
}

// isNumeric checks if a string only contains ASCII numeric (0-9) characters
func (v *validator) isNumeric(s string) error {
	if numericRegex.MatchString(s) {
		// [^ 0-9]
		return errors.New(msgNumeric)
	}
	return nil
}
