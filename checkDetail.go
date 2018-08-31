// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
)

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a CheckDetail Record

// CheckDetail Record
type CheckDetail struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	//AuxiliaryOnUs identifies a code used on commercial checks at the discretion of the payor bank.
	AuxiliaryOnUs string `json:"auxiliaryOnUs"`
	// ExternalProcessingCode identifies a code used for special purposes as authorized by the Accredited
	// Standards Committee X9. Also known as Position 44.
	ExternalProcessingCode string `json:"externalProcessingCode"`
	// PayorBankRoutingNumber identifies a number that identifies the institution by or through which the item is
	// payable. Must be a valid routing and transit number issued by the ABA’s Routing Number Registrar. Shall
	// represent the first 8 digits of a 9-digit routing number or 8 numeric digits of a 4 dash 4 routing number.
	// A valid routing number consists of 2 fields: the eight- digit Payor Bank Routing Number  and the
	// one-digit Payor Bank Routing Number Check Digit.
	// Format: TTTTAAAA, where:
	// TTTT: Federal Reserve Prefix
	// AAAA: ABA Institution Identifier
	PayorBankRoutingNumber string `json:"payorBankRoutingNumber"`
	// PayorBankCheckDigit identifies a digit representing the routing number check digit. The combination of Payor
	// Bank Routing Number and payor Bank Routing Number Check Digit  must be a mod-checked routing number with a
	// valid check digit.
	PayorBankCheckDigit string `json:"payorBankCheckDigit"`
	// OnUs identifies data specified by the payor bank. On-Us data usually consists of the payor’s account number,
	// a serial number or transaction code, or both.
	OnUs string `json:"onUs"`
	// Amount identifies the amount of the check.  All amounts fields have two implied decimal points.
	// e.g., 100000 is $1,000.00
	ItemAmount int `json:"itemAmount"`
	// EceInstitutionItemSequenceNumber identifies a number assigned by the institution that creates the CheckDetail.
	// Field must contain a numeric value. It cannot be all blanks.
	EceInstitutionItemSequenceNumber string `json:"eceInstitutionItemSequenceNumber"`
	// ToDo: CashLetterHeader.CashLetterDocumentation = "Z", CheckDetail.DocumentationTypeIndicator cannot be Z.
	// ToDo: CheckDetail.DocumentationTypeIndicator is defined CashLetterHeader.CashLetterDocumentation = "Z" should
	// ToDo: Z, and this value supersedes.
	// DocumentationTypeIndicator identifies a code that indicates the type of documentation that supports the check
	// record.
	// This field is superseded by the Cash Letter Documentation Type Indicator in the Cash Letter Header
	// Record (Type 10) for all Defined Values except ‘Z’ Not Same Type. In the case of Defined Value of ‘Z’, the
	// Documentation Type Indicator in this record takes precedent.
	//
	// Shall be present when Cash Letter Documentation Type Indicator (Field 9) in the Cash Letter Header Record
	// (Type 10) is Defined Value of ‘Z’.
	//
	// Values:
	// A: No image provided, paper provided separately
	// B: No image provided, paper provided separately, image upon request ‘C’	Image provided separately, no paper
	// provided
	// D: Image provided separately, no paper provided, image upon request ‘E’	Image and paper provided separately
	// F: Image and paper provided separately, image upon request
	// G: Image included, no paper provided
	// H: Image included, no paper provided, image upon request ‘I’	Image included, paper provided separately
	// J: Image included, paper provided separately, image upon request ‘K’	No image provided, no paper provided
	// L: No image provided, no paper provided, image upon request ‘M’	No image provided, Electronic Check provided
	// separately
	// M: No image provided, Electronic Check provided separately
	DocumentationTypeIndicator string `json:"documentationTypeIndicator"`
	// ReturnAcceptanceIndicator is a code that indicates whether the institution that creates the CheckDetail
	// will or will not support electronic return processing.
	// Values:
	// 0:	Will not accept any electronic information
	// 1:	Will accept preliminary return notifications, returns, and final return notifications
	// 2:	Will accept preliminary return notifications and returns
	// 3:	Will accept preliminary return notifications and final return notifications
	// 4:	Will accept returns and final return notifications
	// 5:	Will accept preliminary return notifications only
	// 6:	Will accept returns only
	// 7:	Will accept final return notifications only
	// 8:	Will accept preliminary return notifications, returns, final return notifications, and image returns
	// 9:	Will accept preliminary return notifications, returns and image returns
	// A:	Will accept preliminary return notifications, final return notifications and image returns
	// B:	Will accept returns, final return notifications and image returns
	// C:	Will accept preliminary return notifications and image returns
	// D:	Will accept returns and image returns
	// E:	Will accept final return notifications and image returns
	// F:	Will accept image returns only
	ReturnAcceptanceIndicator string `json:"returnAcceptanceIndicator"`
	// MICRValidIndicator is a code that indicates whether any character in the Magnetic Ink Character Recognition
	// (MICR) property is unreadable, or the OnUs property is missing from the CheckDetail.
	// 1: Good read
	// 2: Good read, missing field
	// 3: Read error encountered
	// 4: Missing field and read error encountered
	MICRValidIndicator int `json:"micrValidIndicator"`
	// BOFDIndicator is a code that indicates whether the ECE institution indicated on the Bundle Header Record (Type 20)
	// is the Bank of First Deposit (BOFD). This field shall be consistent with values contained in the Check Detail
	// Addendum A Record (Type 26) and Check Detail Addendum C Record (Type 28).
	// Values:
	// Y: ECE institution is BOFD
	// N: ECE institution is not BOFD
	// U: ECE institution relationship to BOFD is undetermined
	BOFDIndicator string `json:"bofdIndicator"`
	// AddendumCount is a number of Check Detail Record Addenda to follow. This represents the number of
	// CheckDetailAddendumA, CheckDetailAddendumB Record and CheckDetailAddendumC types.  It matches the total number
	// of addendum records associated with this item. The standard supports up to 99 addendum records.
	AddendumCount int `json:"addendumCount"`
	// CorrectionIndicator identifies whether and how the MICR line was repaired, for fields other than Payor Bank
	// Routing Number and Amount.
	// Values:
	// 0: No Repair
	// 1: Repaired (form of repair unknown)
	// 2: Repaired without Operator intervention
	// 3: Repaired with Operator intervention
	// 4: Undetermined if repair has been done or not
	CorrectionIndicator int `json:"correctionIndicator"`
	// ArchiveTypeIndicator is a code that indicates the type of archive that supports this CheckDetail.
	// Access method, availability and time-frames shall be defined by clearing arrangements.
	// Values:
	// A: Microfilm
	// B: Image
	// C: Paper
	// D: Microfilm and image
	// E: Microfilm and paper
	// F: Image and paper
	// G: Microfilm, image and paper
	// H: Electronic Check Instrument
	// I: None
	ArchiveTypeIndicator string `json:"archiveTypeIndicator"`
	// CheckDetailAddendumA
	CheckDetailAddendumA CheckDetailAddendumA `json:"checkDetailAddendumA"`
	// CheckDetailAddendumB
	CheckDetailAddendumB CheckDetailAddendumB `json:"checkDetailAddendumB"`
	// CheckDetailAddendumC
	CheckDetailAddendumC CheckDetailAddendumC `json:"checkDetailAddendumC"`
	// ImageView
	ImageView ImageView `json:"imageView"`

	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewCheckDetail returns a new CheckDetail with default values for non exported fields
func NewCheckDetail() *CheckDetail {
	cd := &CheckDetail{
		recordType: "25",
	}
	return cd
}

// Parse takes the input record string and parses the CheckDetail values
func (cd *CheckDetail) Parse(record string) {
}

// String writes the CheckDetail struct to a variable length string.
func (cd *CheckDetail) String() string {
	var buf strings.Builder
	buf.Grow(80)
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (cd *CheckDetail) Validate() error {
	if err := cd.fieldInclusion(); err != nil {
		return err
	}
	if cd.recordType != "25" {
		msg := fmt.Sprintf(msgRecordType, 25)
		return &FieldError{FieldName: "recordType", Value: cd.recordType, Msg: msg}
	}
	if err := cd.isDocumentationTypeIndicator(cd.DocumentationTypeIndicator); err != nil {
		return &FieldError{FieldName: "DocumentationTypeIndicator", Value: cd.DocumentationTypeIndicator, Msg: err.Error()}
	}
	if err := cd.isReturnAcceptanceIndicator(cd.ReturnAcceptanceIndicator); err != nil {
		return &FieldError{FieldName: "ReturnAcceptanceIndicator", Value: cd.ReturnAcceptanceIndicator, Msg: err.Error()}
	}
	if err := cd.isMICRValidIndicator(cd.MICRValidIndicator); err != nil {
		return &FieldError{FieldName: "MICRValidIndicator", Value: cd.MICRValidIndicatorField(), Msg: err.Error()}
	}
	if err := cd.isBOFDIndicator(cd.BOFDIndicator); err != nil {
		return &FieldError{FieldName: "BOFDIndicator", Value: cd.BOFDIndicator, Msg: err.Error()}
	}
	if err := cd.isCorrectionIndicator(cd.CorrectionIndicator); err != nil {
		return &FieldError{FieldName: "CorrectionIndicator", Value: cd.CorrectionIndicatorField(), Msg: err.Error()}
	}
	if err := cd.isArchiveTypeIndicator(cd.ArchiveTypeIndicator); err != nil {
		return &FieldError{FieldName: "ArchiveTypeIndicator", Value: cd.ArchiveTypeIndicator, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (cd *CheckDetail) fieldInclusion() error {
	if cd.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: cd.recordType, Msg: msgFieldInclusion}
	}
	if cd.PayorBankRoutingNumber == "" {
		return &FieldError{FieldName: "PayorBankRoutingNumber",
		Value: cd.PayorBankRoutingNumber, Msg: msgFieldInclusion}
	}
	if cd.PayorBankCheckDigit == "" {
		return &FieldError{FieldName: "PayorBankCheckDigit", Value: cd.PayorBankCheckDigit, Msg: msgFieldInclusion}
	}
	if cd.EceInstitutionItemSequenceNumber == "" {
		return &FieldError{FieldName: "EceInstitutionItemSequenceNumber",
		Value: cd.EceInstitutionItemSequenceNumber, Msg: msgFieldInclusion}
	}
	if cd.BOFDIndicator == "" {
		return &FieldError{FieldName: "BOFDIndicator", Value: cd.BOFDIndicator, Msg: msgFieldInclusion}
	}
	return nil
}

// ItemAmountField gets the ItemAmount right justified and zero padded
func (cd *CheckDetail) ItemAmountField() string {
	return cd.numericField(cd.ItemAmount, 1)
}

// EceInstitutionItemSequenceNumberField gets the EceInstitutionItemSequenceNumber
func (cd *CheckDetail) EceInstitutionItemSequenceNumberField() string {
	return cd.alphaField(cd.EceInstitutionItemSequenceNumber, 1)
}

// DocumentationTypeIndicatorField gets the DocumentationTypeIndicator
func (cd *CheckDetail) DocumentationTypeIndicatorField() string {
	return cd.alphaField(cd.DocumentationTypeIndicator, 1)
}

// ReturnAcceptanceIndicatorField gets the ReturnAcceptanceIndicator
func (cd *CheckDetail) ReturnAcceptanceIndicatorField() string {
	return cd.alphaField(cd.ReturnAcceptanceIndicator, 1)
}

// MICRValidIndicatorField gets a string of the MICRValidIndicator field
func (cd *CheckDetail) MICRValidIndicatorField() string {
	return cd.numericField(cd.MICRValidIndicator, 1)
}

// BOFDIndicatorField gets the BOFDIndicator field
func (cd *CheckDetail) BOFDIndicatorField() string {
	return cd.alphaField(cd.BOFDIndicator, 1)
}

// AddendumCountField gets a string of the AddendumCount field
func (cd *CheckDetail) AddendumCountField() string {
	return cd.numericField(cd.AddendumCount, 2)
}

// CorrectionIndicatorField gets a string of the CorrectionIndicator field
func (cd *CheckDetail) CorrectionIndicatorField() string {
	return cd.numericField(cd.CorrectionIndicator, 1)
}

// ArchiveTypeIndicatorField gets the ArchiveTypeIndicator field
func (cd *CheckDetail) ArchiveTypeIndicatorField() string {
	return cd.alphaField(cd.ArchiveTypeIndicator, 1)
}