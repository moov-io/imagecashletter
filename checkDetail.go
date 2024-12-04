// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Errors specific to a CheckDetail Record

var (
	msgDocumentationTypeIndicator = "is Invalid"
)

// CheckDetail Record
type CheckDetail struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// AuxiliaryOnUs identifies a code used on commercial checks at the discretion of the payor bank.
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
	// B: No image provided, paper provided separately, image upon request
	// C:	Image provided separately, no paper provided
	// D: Image provided separately, no paper provided, image upon request
	// E:	Image and paper provided separately
	// F: Image and paper provided separately, image upon request
	// G: Image included, no paper provided
	// H: Image included, no paper provided, image upon request
	// I:	Image included, paper provided separately
	// J: Image included, paper provided separately, image upon request
	// K:	No image provided, no paper provided
	// L: No image provided, no paper provided, image upon request
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
	// CheckDetailAddendumA, CheckDetailAddendumB and CheckDetailAddendumC types.  It matches the total number
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
	CheckDetailAddendumA []CheckDetailAddendumA `json:"checkDetailAddendumA"`
	// CheckDetailAddendumB
	CheckDetailAddendumB []CheckDetailAddendumB `json:"checkDetailAddendumB"`
	// CheckDetailAddendumC
	CheckDetailAddendumC []CheckDetailAddendumC `json:"checkDetailAddendumC"`
	// ImageViewDetail
	ImageViewDetail []ImageViewDetail `json:"imageViewDetail"`
	// ImageViewData
	ImageViewData []ImageViewData `json:"imageViewData"`
	// ImageViewAnalysis
	ImageViewAnalysis []ImageViewAnalysis `json:"imageViewAnalysis"`
	// validator is composed for imagecashletter data validation
	validator
	// converters is composed for imagecashletter to golang Converters
	converters
}

// NewCheckDetail returns a new CheckDetail with default values for non exported fields
func NewCheckDetail() *CheckDetail {
	cd := &CheckDetail{}
	cd.setRecordType()
	return cd
}

func (cd *CheckDetail) setRecordType() {
	if cd == nil {
		return
	}
	cd.recordType = "25"
}

// Parse takes the input record string and parses the CheckDetail values
func (cd *CheckDetail) Parse(record string) {
	if utf8.RuneCountInString(record) < 80 {
		return // line too short
	}

	// Character position 1-2, Always "25"
	cd.setRecordType()
	// 03-17
	cd.AuxiliaryOnUs = cd.parseStringField(record[2:17])
	// 18-18
	cd.ExternalProcessingCode = cd.parseStringField(record[17:18])
	// 19-26
	cd.PayorBankRoutingNumber = cd.parseStringField(record[18:26])
	// 27-27
	cd.PayorBankCheckDigit = cd.parseStringField(record[26:27])
	// 28-47
	cd.OnUs = cd.parseStringField(record[27:47])
	// 48-57
	cd.ItemAmount = cd.parseNumField(record[47:57])
	// 58-72
	cd.EceInstitutionItemSequenceNumber = cd.parseStringField(record[57:72])
	// 73-73
	cd.DocumentationTypeIndicator = cd.parseStringField(record[72:73])
	// 74-74
	cd.ReturnAcceptanceIndicator = cd.parseStringField(record[73:74])
	// 75-75
	cd.MICRValidIndicator = cd.parseNumField(record[74:75])
	// 76-76
	cd.BOFDIndicator = cd.parseStringField(record[75:76])
	// 77-78
	cd.AddendumCount = cd.parseNumField(record[76:78])
	// 79-79
	cd.CorrectionIndicator = cd.parseNumField(record[78:79])
	// 80-80
	cd.ArchiveTypeIndicator = cd.parseStringField(record[79:80])
}

func (cd *CheckDetail) UnmarshalJSON(data []byte) error {
	type Alias CheckDetail
	aux := struct {
		*Alias
	}{
		(*Alias)(cd),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cd.setRecordType()
	return nil
}

// String writes the CheckDetail struct to a variable length string.
func (cd *CheckDetail) String() string {
	if cd == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(cd.recordType)
	buf.WriteString(cd.AuxiliaryOnUsField())
	buf.WriteString(cd.ExternalProcessingCodeField())
	buf.WriteString(cd.PayorBankRoutingNumberField())
	buf.WriteString(cd.PayorBankCheckDigitField())
	buf.WriteString(cd.OnUsField())
	buf.WriteString(cd.ItemAmountField())
	buf.WriteString(cd.EceInstitutionItemSequenceNumberField())
	buf.WriteString(cd.DocumentationTypeIndicatorField())
	buf.WriteString(cd.ReturnAcceptanceIndicatorField())
	buf.WriteString(cd.MICRValidIndicatorField())
	buf.WriteString(cd.BOFDIndicatorField())
	buf.WriteString(cd.AddendumCountField())
	buf.WriteString(cd.CorrectionIndicatorField())
	buf.WriteString(cd.ArchiveTypeIndicatorField())
	return buf.String()
}

// Validate performs imagecashletter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (cd *CheckDetail) Validate() error {
	if err := cd.fieldInclusion(); err != nil {
		return err
	}
	if cd.recordType != "25" {
		msg := fmt.Sprintf(msgRecordType, 25)
		return &FieldError{FieldName: "recordType", Value: cd.recordType, Msg: msg}
	}
	if cd.DocumentationTypeIndicator != "" {
		// Z is valid for CashLetter DocumentationTypeIndicator only
		if cd.DocumentationTypeIndicator == "Z" {
			msg := msgDocumentationTypeIndicator
			return &FieldError{FieldName: "DocumentationTypeIndicator", Value: cd.DocumentationTypeIndicator, Msg: msg}
		}
		if err := cd.isDocumentationTypeIndicator(cd.DocumentationTypeIndicator); err != nil {
			return &FieldError{FieldName: "DocumentationTypeIndicator", Value: cd.DocumentationTypeIndicator, Msg: err.Error()}
		}
	}
	// Conditional
	if cd.ReturnAcceptanceIndicator != "" {
		if err := cd.isReturnAcceptanceIndicator(cd.ReturnAcceptanceIndicator); err != nil {
			return &FieldError{FieldName: "ReturnAcceptanceIndicator", Value: cd.ReturnAcceptanceIndicator, Msg: err.Error()}
		}
	}
	// Conditional
	if cd.MICRValidIndicatorField() != "" {
		if err := cd.isMICRValidIndicator(cd.MICRValidIndicator); err != nil {
			return &FieldError{FieldName: "MICRValidIndicator", Value: cd.MICRValidIndicatorField(), Msg: err.Error()}
		}
	}
	// Mandatory
	if err := cd.isBOFDIndicator(cd.BOFDIndicator); err != nil {
		return &FieldError{FieldName: "BOFDIndicator", Value: cd.BOFDIndicator, Msg: err.Error()}
	}
	// Conditional
	if cd.CorrectionIndicatorField() != "" {
		if err := cd.isCorrectionIndicator(cd.CorrectionIndicator); err != nil {
			return &FieldError{FieldName: "CorrectionIndicator", Value: cd.CorrectionIndicatorField(), Msg: err.Error()}
		}
	}
	// Conditional
	if cd.ArchiveTypeIndicator != "" {
		if err := cd.isArchiveTypeIndicator(cd.ArchiveTypeIndicator); err != nil {
			return &FieldError{FieldName: "ArchiveTypeIndicator", Value: cd.ArchiveTypeIndicator, Msg: err.Error()}
		}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (cd *CheckDetail) fieldInclusion() error {
	if cd.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: cd.recordType,
			Msg:   msgFieldInclusion + ", did you use CheckDetail()?"}
	}
	if cd.PayorBankRoutingNumber == "" {
		return &FieldError{FieldName: "PayorBankRoutingNumber",
			Value: cd.PayorBankRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetail()?"}
	}
	if cd.PayorBankRoutingNumberField() == "00000000" {
		return &FieldError{FieldName: "PayorBankRoutingNumber",
			Value: cd.PayorBankRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetail()?"}
	}
	if cd.PayorBankCheckDigit == "" {
		return &FieldError{FieldName: "PayorBankCheckDigit",
			Value: cd.PayorBankCheckDigit,
			Msg:   msgFieldInclusion + ", did you use CheckDetail()?"}
	}
	if cd.EceInstitutionItemSequenceNumberField() == "               " {
		return &FieldError{FieldName: "EceInstitutionItemSequenceNumber",
			Value: cd.EceInstitutionItemSequenceNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetail()?"}
	}
	if cd.BOFDIndicator == "" {
		return &FieldError{FieldName: "BOFDIndicator",
			Value: cd.BOFDIndicator,
			Msg:   msgFieldInclusion + ", did you use CheckDetail()?"}
	}
	return nil
}

// AuxiliaryOnUsField gets the AuxiliaryOnUs field
func (cd *CheckDetail) AuxiliaryOnUsField() string {
	return cd.nbsmField(cd.AuxiliaryOnUs, 15)
}

// ExternalProcessingCodeField gets the ExternalProcessingCode field - Also known as Position 44
func (cd *CheckDetail) ExternalProcessingCodeField() string {
	return cd.alphaField(cd.ExternalProcessingCode, 1)
}

// PayorBankRoutingNumberField gets the PayorBankRoutingNumber field
func (cd *CheckDetail) PayorBankRoutingNumberField() string {
	return cd.stringField(cd.PayorBankRoutingNumber, 8)
}

// PayorBankCheckDigitField gets the PayorBankCheckDigit field
func (cd *CheckDetail) PayorBankCheckDigitField() string {
	return cd.stringField(cd.PayorBankCheckDigit, 1)
}

// OnUsField gets the OnUs field
func (cd *CheckDetail) OnUsField() string {
	return cd.nbsmField(cd.OnUs, 20)
}

// ItemAmountField gets the ItemAmount right justified and zero padded
func (cd *CheckDetail) ItemAmountField() string {
	return cd.numericField(cd.ItemAmount, 10)
}

// EceInstitutionItemSequenceNumberField gets a string of the EceInstitutionItemSequenceNumber field
func (cd *CheckDetail) EceInstitutionItemSequenceNumberField() string {
	return cd.alphaField(cd.EceInstitutionItemSequenceNumber, 15)
}

// DocumentationTypeIndicatorField gets the DocumentationTypeIndicator field
func (cd *CheckDetail) DocumentationTypeIndicatorField() string {
	return cd.alphaField(cd.DocumentationTypeIndicator, 1)
}

// ReturnAcceptanceIndicatorField gets the ReturnAcceptanceIndicator field
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

// AddCheckDetailAddendumA appends an AddendumA to the CheckDetail
func (cd *CheckDetail) AddCheckDetailAddendumA(cdAddendaA CheckDetailAddendumA) []CheckDetailAddendumA {
	cd.CheckDetailAddendumA = append(cd.CheckDetailAddendumA, cdAddendaA)
	return cd.CheckDetailAddendumA
}

// GetCheckDetailAddendumA returns a slice of AddendumA for the CheckDetail
func (cd *CheckDetail) GetCheckDetailAddendumA() []CheckDetailAddendumA {
	return cd.CheckDetailAddendumA
}

// AddCheckDetailAddendumB appends an AddendumA to the CheckDetail
func (cd *CheckDetail) AddCheckDetailAddendumB(cdAddendaB CheckDetailAddendumB) []CheckDetailAddendumB {
	cd.CheckDetailAddendumB = append(cd.CheckDetailAddendumB, cdAddendaB)
	return cd.CheckDetailAddendumB
}

// GetCheckDetailAddendumB returns a slice of AddendumB for the CheckDetail
func (cd *CheckDetail) GetCheckDetailAddendumB() []CheckDetailAddendumB {
	return cd.CheckDetailAddendumB
}

// AddCheckDetailAddendumC appends an AddendumC to the CheckDetail
func (cd *CheckDetail) AddCheckDetailAddendumC(cdAddendaC CheckDetailAddendumC) []CheckDetailAddendumC {
	cd.CheckDetailAddendumC = append(cd.CheckDetailAddendumC, cdAddendaC)
	return cd.CheckDetailAddendumC
}

// GetCheckDetailAddendumC returns a slice of AddendumC for the CheckDetail
func (cd *CheckDetail) GetCheckDetailAddendumC() []CheckDetailAddendumC {
	return cd.CheckDetailAddendumC
}

// AddImageViewDetail appends an ImageViewDetail to the CheckDetail
func (cd *CheckDetail) AddImageViewDetail(ivDetail ImageViewDetail) []ImageViewDetail {
	cd.ImageViewDetail = append(cd.ImageViewDetail, ivDetail)
	return cd.ImageViewDetail
}

// GetImageViewDetail returns a slice of ImageViewDetail for the CheckDetail
func (cd *CheckDetail) GetImageViewDetail() []ImageViewDetail {
	return cd.ImageViewDetail
}

// AddImageViewData appends an ImageViewData to the CheckDetail
func (cd *CheckDetail) AddImageViewData(ivData ImageViewData) []ImageViewData {
	cd.ImageViewData = append(cd.ImageViewData, ivData)
	return cd.ImageViewData
}

// GetImageViewData returns a slice of ImageViewData for the CheckDetail
func (cd *CheckDetail) GetImageViewData() []ImageViewData {
	return cd.ImageViewData
}

// AddImageViewAnalysis appends an ImageViewAnalysis to the CheckDetail
func (cd *CheckDetail) AddImageViewAnalysis(ivAnalysis ImageViewAnalysis) []ImageViewAnalysis {
	cd.ImageViewAnalysis = append(cd.ImageViewAnalysis, ivAnalysis)
	return cd.ImageViewAnalysis
}

// GetImageViewAnalysis returns a slice of ImageViewAnalysis for the CheckDetail
func (cd *CheckDetail) GetImageViewAnalysis() []ImageViewAnalysis {
	return cd.ImageViewAnalysis
}

// SetEceInstitutionItemSequenceNumber sets EceInstitutionItemSequenceNumber
func (cd *CheckDetail) SetEceInstitutionItemSequenceNumber(seq int) string {
	cd.EceInstitutionItemSequenceNumber = cd.numericField(seq, 15)
	return cd.EceInstitutionItemSequenceNumber
}
