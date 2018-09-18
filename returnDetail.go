// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
	"time"
)

// Errors specific to a ReturnDetail Record

// ReturnDetail Record
type ReturnDetail struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
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
	// ReturnReason is a code that indicates the reason for non-payment.
	ReturnReason string `json:"returnReason"`
	// AddendumCount is a number of Return Record Addenda to follow. This represents the number of
	// ReturnDetailAddendumA, ReturnDetailAddendumB, ReturnDetailAddendumC and ReturnDetailAddendumD types.
	// It matches the total number of addendum records associated with this item. The standard supports up to 99
	// addendum records.
	AddendumCount int `json:"addendumCount"`
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
	// ForwardBundleDate represents for electronic check exchange items, the year, month, and day that designates the
	// business date of the original forward bundle. This data is transferred from the BundleHeader.BundleBusinessDate.
	// For items presented in paper cash letters, the year, month, and day that the cash letter was created.
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	ForwardBundleDate time.Time `json:"bundleBusinessDate"`
	// EceInstitutionItemSequenceNumber identifies a number assigned by the institution that creates the Return.
	// Field must contain a numeric value. It cannot be all blanks.
	EceInstitutionItemSequenceNumber int `json:"eceInstitutionItemSequenceNumber"`
	// ExternalProcessingCode identifies a code used for special purposes as authorized by the Accredited
	// Standards Committee X9. Also known as Position 44.
	ExternalProcessingCode string `json:"externalProcessingCode"`
	// ReturnNotificationIndicator is a A code that identifies the type of notification. The
	// CashLetterHeader.CollectionTypeIndicator and the BundleHeader.CollectionTypeIndicator when equal 05 or 06
	// takes precedence over this field.
	// Values:
	// 1: Preliminary notification
	// 2: Final notification
	ReturnNotificationIndicator int `json:"returnNotificationIndicator"`
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
	// TimesReturned is code used to indicate the number of times the paying bank has returned this item.
	//Values:
	// 0: The item has been returned an unknown number of times
	// 1: The item has been returned once
	// 2: The item has been returned twice
	// 3: The item has been returned three times
	TimesReturned int `json:"timesReturned"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// ReturnDetailAddendumA
	ReturnDetailAddendumA []ReturnDetailAddendumA `json:"returnDetailAddendumA"`
	// ReturnDetailAddendumB
	ReturnDetailAddendumB []ReturnDetailAddendumB `json:"returnDetailAddendumB"`
	// ReturnDetailAddendumC
	ReturnDetailAddendumC []ReturnDetailAddendumC `json:"returnDetailAddendumC"`
	// ReturnDetailAddendumD
	ReturnDetailAddendumD []ReturnDetailAddendumD `json:"returnDetailAddendumD"`
	// ImageViewDetail
	ImageViewDetail []ImageViewDetail `json:"imageViewDetail"`
	// ImageViewData
	ImageViewData []ImageViewData `json:"imageViewData"`
	// ImageViewAnalysis
	ImageViewAnalysis []ImageViewAnalysis `json:"imageViewAnalysis"`
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewReturnDetail returns a new ReturnDetail with default values for non exported fields
func NewReturnDetail() *ReturnDetail {
	rd := &ReturnDetail{
		recordType: "31",
	}
	return rd
}

// Parse takes the input record string and parses the ReturnDetail values
func (rd *ReturnDetail) Parse(record string) {
	// Character position 1-2, Always "31"
	rd.recordType = "31"
	// 03-10
	rd.PayorBankRoutingNumber = rd.parseStringField(record[2:10])
	// 11-11
	rd.PayorBankCheckDigit = rd.parseStringField(record[10:11])
	// 12-31
	rd.OnUs = rd.parseStringField(record[11:31])
	// 32-41
	rd.ItemAmount = rd.parseNumField(record[31:41])
	// 42-42
	rd.ReturnReason = rd.parseStringField(record[41:42])
	// 43-44
	rd.AddendumCount = rd.parseNumField(record[42:44])
	// 45-45
	rd.DocumentationTypeIndicator = rd.parseStringField(record[44:45])
	// 46-53
	rd.ForwardBundleDate = rd.parseYYYYMMDDDate(record[45:53])
	// 54-68
	rd.EceInstitutionItemSequenceNumber = rd.parseNumField(record[53:68])
	// 69-69
	rd.ExternalProcessingCode = rd.parseStringField(record[68:69])
	// 70-70
	rd.ReturnNotificationIndicator = rd.parseNumField(record[69:70])
	// 71-71
	rd.ArchiveTypeIndicator = rd.parseStringField(record[70:71])
	// 72-72
	rd.TimesReturned = rd.parseNumField(record[71:72])
	// 73-80
	rd.reserved = "        "
}

// String writes the ReturnDetail struct to a variable length string.
func (rd *ReturnDetail) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rd.recordType)
	buf.WriteString(rd.PayorBankRoutingNumberField())
	buf.WriteString(rd.PayorBankCheckDigitField())
	buf.WriteString(rd.OnUsField())
	buf.WriteString(rd.ItemAmountField())
	buf.WriteString(rd.ReturnReasonField())
	buf.WriteString(rd.AddendumCountField())
	buf.WriteString(rd.DocumentationTypeIndicatorField())
	buf.WriteString(rd.ForwardBundleDateField())
	buf.WriteString(rd.EceInstitutionItemSequenceNumberField())
	buf.WriteString(rd.ExternalProcessingCodeField())
	buf.WriteString(rd.ReturnNotificationIndicatorField())
	buf.WriteString(rd.ArchiveTypeIndicatorField())
	buf.WriteString(rd.TimesReturnedField())
	buf.WriteString(rd.reservedField())
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rd *ReturnDetail) Validate() error {
	if err := rd.fieldInclusion(); err != nil {
		return err
	}
	if rd.recordType != "31" {
		msg := fmt.Sprintf(msgRecordType, 31)
		return &FieldError{FieldName: "recordType", Value: rd.recordType, Msg: msg}
	}
	if rd.DocumentationTypeIndicator != "" {
		if err := rd.isDocumentationTypeIndicator(rd.DocumentationTypeIndicator); err != nil {
			return &FieldError{FieldName: "DocumentationTypeIndicator", Value: rd.DocumentationTypeIndicator, Msg: err.Error()}
		}
	}
	if rd.ReturnNotificationIndicatorField() != "" {
		if err := rd.isReturnNotificationIndicator(rd.ReturnNotificationIndicator); err != nil {
			return &FieldError{FieldName: "ReturnNotificationIndicator", Value: rd.ReturnNotificationIndicatorField(), Msg: err.Error()}
		}
	}
	if rd.ArchiveTypeIndicatorField() != "" {
		if err := rd.isArchiveTypeIndicator(rd.ArchiveTypeIndicator); err != nil {
			return &FieldError{FieldName: "ArchiveTypeIndicator", Value: rd.ArchiveTypeIndicatorField(), Msg: err.Error()}
		}
	}
	if rd.TimesReturnedField() != "" {
		if err := rd.isTimesReturned(rd.TimesReturned); err != nil {
			return &FieldError{FieldName: "TimesReturned", Value: rd.TimesReturnedField(), Msg: err.Error()}
		}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rd *ReturnDetail) fieldInclusion() error {
	if rd.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: rd.recordType, Msg: msgFieldInclusion}
	}
	if rd.PayorBankRoutingNumber == "" {
		return &FieldError{FieldName: "PayorBankRoutingNumber",
			Value: rd.PayorBankRoutingNumber, Msg: msgFieldInclusion}
	}
	if rd.PayorBankCheckDigit == "" {
		return &FieldError{FieldName: "PayorBankCheckDigit", Value: rd.PayorBankCheckDigit, Msg: msgFieldInclusion}
	}
	if rd.ReturnReason == "" {
		return &FieldError{FieldName: "ReturnReason", Value: rd.ReturnReason, Msg: msgFieldInclusion}
	}
	return nil
}

// PayorBankRoutingNumberField gets the PayorBankRoutingNumber field
func (rd *ReturnDetail) PayorBankRoutingNumberField() string {
	return rd.stringField(rd.PayorBankRoutingNumber, 8)
}

// PayorBankCheckDigitField gets the PayorBankCheckDigit field
func (rd *ReturnDetail) PayorBankCheckDigitField() string {
	return rd.stringField(rd.PayorBankCheckDigit, 1)
}

// OnUsField gets the OnUs field
func (rd *ReturnDetail) OnUsField() string {
	return rd.nbsmField(rd.OnUs, 20)
}

// ItemAmountField gets the ItemAmount right justified and zero padded
func (rd *ReturnDetail) ItemAmountField() string {
	return rd.numericField(rd.ItemAmount, 10)
}

// ReturnReasonField gets the ReturnReason field
func (rd *ReturnDetail) ReturnReasonField() string {
	return rd.alphaField(rd.ReturnReason, 1)
}

// AddendumCountField gets a string of the AddendumCount field
func (rd *ReturnDetail) AddendumCountField() string {
	return rd.numericField(rd.AddendumCount, 2)
}

// DocumentationTypeIndicatorField gets the DocumentationTypeIndicator field
func (rd *ReturnDetail) DocumentationTypeIndicatorField() string {
	return rd.alphaField(rd.DocumentationTypeIndicator, 1)
}

// ForwardBundleDateField gets the ForwardBundleDate in YYYYMMDD format
func (rd *ReturnDetail) ForwardBundleDateField() string {
	return rd.formatYYYYMMDDDate(rd.ForwardBundleDate)
}

// EceInstitutionItemSequenceNumberField gets a string of the EceInstitutionItemSequenceNumber field
func (rd *ReturnDetail) EceInstitutionItemSequenceNumberField() string {
	return rd.numericField(rd.EceInstitutionItemSequenceNumber, 15)
}

// ExternalProcessingCodeField gets the ExternalProcessingCode field - Also known as Position 44
func (rd *ReturnDetail) ExternalProcessingCodeField() string {
	return rd.alphaField(rd.ExternalProcessingCode, 1)
}

// ReturnNotificationIndicatorField gets a string of the ReturnNotificationIndicator field
func (rd *ReturnDetail) ReturnNotificationIndicatorField() string {
	return rd.numericField(rd.ReturnNotificationIndicator, 1)
}

// ArchiveTypeIndicatorField gets the ArchiveTypeIndicator field
func (rd *ReturnDetail) ArchiveTypeIndicatorField() string {
	return rd.alphaField(rd.ArchiveTypeIndicator, 1)
}

// TimesReturnedField gets a string of the TimesReturned field
func (rd *ReturnDetail) TimesReturnedField() string {
	return rd.numericField(rd.TimesReturned, 1)
}

// reservedField gets reserved - blank space
func (rd *ReturnDetail) reservedField() string {
	return rd.alphaField(rd.reserved, 8)
}

// AddReturnDetailAddendumA appends an AddendumA to the ReturnDetail
func (rd *ReturnDetail) AddReturnDetailAddendumA(rdAddendaA ReturnDetailAddendumA) []ReturnDetailAddendumA {
	rd.ReturnDetailAddendumA = append(rd.ReturnDetailAddendumA, rdAddendaA)
	return rd.ReturnDetailAddendumA
}

// GetReturnDetailAddendumA returns a slice of AddendumA for the ReturnDetail
func (rd *ReturnDetail) GetReturnDetailAddendumA() []ReturnDetailAddendumA {
	return rd.ReturnDetailAddendumA
}

// AddReturnDetailAddendumB appends an AddendumA to the ReturnDetail
func (rd *ReturnDetail) AddReturnDetailAddendumB(rdAddendaB ReturnDetailAddendumB) []ReturnDetailAddendumB {
	rd.ReturnDetailAddendumB = append(rd.ReturnDetailAddendumB, rdAddendaB)
	return rd.ReturnDetailAddendumB
}

// GetReturnDetailAddendumB returns a slice of AddendumB for the ReturnDetail
func (rd *ReturnDetail) GetReturnDetailAddendumB() []ReturnDetailAddendumB {
	return rd.ReturnDetailAddendumB
}

// AddReturnDetailAddendumC appends an AddendumC to the ReturnDetail
func (rd *ReturnDetail) AddReturnDetailAddendumC(rdAddendaC ReturnDetailAddendumC) []ReturnDetailAddendumC {
	rd.ReturnDetailAddendumC = append(rd.ReturnDetailAddendumC, rdAddendaC)
	return rd.ReturnDetailAddendumC
}

// GetReturnDetailAddendumC returns a slice of AddendumC for the ReturnDetail
func (rd *ReturnDetail) GetReturnDetailAddendumC() []ReturnDetailAddendumC {
	return rd.ReturnDetailAddendumC
}

// AddReturnDetailAddendumD appends an AddendumD to the ReturnDetail
func (rd *ReturnDetail) AddReturnDetailAddendumD(rdAddendaD ReturnDetailAddendumD) []ReturnDetailAddendumD {
	rd.ReturnDetailAddendumD = append(rd.ReturnDetailAddendumD, rdAddendaD)
	return rd.ReturnDetailAddendumD
}

// GetReturnDetailAddendumD returns a slice of AddendumD for the ReturnDetail
func (rd *ReturnDetail) GetReturnDetailAddendumD() []ReturnDetailAddendumD {
	return rd.ReturnDetailAddendumD
}

// AddReturnDetailImageViewDetail appends an ImageViewDetail to the ReturnDetail
func (rd *ReturnDetail) AddReturnDetailImageViewDetail(ivDetail ImageViewDetail) []ImageViewDetail {
	rd.ImageViewDetail = append(rd.ImageViewDetail, ivDetail)
	return rd.ImageViewDetail
}

// GetReturnDetailImageViewDetail returns a slice of ImageViewDetail for the ReturnDetail
func (rd *ReturnDetail) GetReturnDetailImageViewDetail() []ImageViewDetail {
	return rd.ImageViewDetail
}

// AddReturnDetailImageViewData appends an ImageViewData to the ReturnDetail
func (rd *ReturnDetail) AddReturnDetailImageViewData(ivData ImageViewData) []ImageViewData {
	rd.ImageViewData = append(rd.ImageViewData, ivData)
	return rd.ImageViewData
}

// GetReturnDetailImageViewData returns a slice of ImageViewData for the ReturnDetail
func (rd *ReturnDetail) GetReturnDetailImageViewData() []ImageViewData {
	return rd.ImageViewData
}

// AddReturnDetailImageViewAnalysis appends an ImageViewAnalysis  to the ReturnDetail
func (rd *ReturnDetail) AddReturnDetailImageViewAnalysis(ivAnalysis ImageViewAnalysis) []ImageViewAnalysis {
	rd.ImageViewAnalysis = append(rd.ImageViewAnalysis, ivAnalysis)
	return rd.ImageViewAnalysis
}

// GetReturnDetailImageViewAnalysis returns a slice of ImageViewAnalysis for the ReturnDetail
func (rd *ReturnDetail) GetReturnDetailImageViewAnalysis() []ImageViewAnalysis {
	return rd.ImageViewAnalysis
}
