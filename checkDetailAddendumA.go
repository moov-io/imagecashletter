// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

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
	ReturnLocationRoutingNumber string `json:"returnLocationRoutingNumber"`
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
	BOFDConversionIndicator string `json:"bofdConversionIndicator"`
	// BOFDCorrectionIndicator identifies whether and how the MICR line of this item was repaired by the
	// creator of this CheckDetailAddendumA Record for fields other than Payor Bank Routing Number and Amount.
	// Values:
	// 0: No Repair
	// 1: Repaired (form of repair unknown)
	// 2: Repaired without Operator intervention
	// 3: Repaired with Operator intervention
	// 4: Undetermined if repair has been done or not
	BOFDCorrectionIndicator int `json:"bofdCorrectionIndicator"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for image cash letter data validation
	validator
	// converters is composed for image cash letter to golang Converters
	converters
}

// NewCheckDetailAddendumA returns a new CheckDetailAddendumA with default values for non exported fields
func NewCheckDetailAddendumA() CheckDetailAddendumA {
	cdAddendumA := CheckDetailAddendumA{}
	cdAddendumA.setRecordType()
	return cdAddendumA
}

func (cdAddendumA *CheckDetailAddendumA) setRecordType() {
	if cdAddendumA == nil {
		return
	}
	cdAddendumA.recordType = "26"
}

// Parse takes the input record string and parses the CheckDetailAddendumA values
func (cdAddendumA *CheckDetailAddendumA) Parse(record string) {
	if utf8.RuneCountInString(record) < 77 {
		return // line too short
	}

	// Character position 1-2, Always "26"
	cdAddendumA.setRecordType()
	// 03-03
	cdAddendumA.RecordNumber = cdAddendumA.parseNumField(record[2:3])
	// 04-12
	cdAddendumA.ReturnLocationRoutingNumber = cdAddendumA.parseStringField(record[3:12])
	// 13-20
	cdAddendumA.BOFDEndorsementDate = cdAddendumA.parseYYYYMMDDDate(record[12:20])
	// 21-35
	cdAddendumA.BOFDItemSequenceNumber = cdAddendumA.parseStringField(record[20:35])
	// 36-53
	cdAddendumA.BOFDAccountNumber = cdAddendumA.parseStringField(record[35:53])
	// 54-58
	cdAddendumA.BOFDBranchCode = cdAddendumA.parseStringField(record[53:58])
	// 59-73
	cdAddendumA.PayeeName = cdAddendumA.parseStringField(record[58:73])
	// 74-74
	cdAddendumA.TruncationIndicator = cdAddendumA.parseStringField(record[73:74])
	// 75-75
	cdAddendumA.BOFDConversionIndicator = cdAddendumA.parseStringField(record[74:75])
	// 76-76
	cdAddendumA.BOFDCorrectionIndicator = cdAddendumA.parseNumField(record[75:76])
	// 77-77
	cdAddendumA.UserField = cdAddendumA.parseStringField(record[76:77])
	// 78-80
	cdAddendumA.reserved = "   "
}

func (cdAddendumA *CheckDetailAddendumA) UnmarshalJSON(data []byte) error {
	type Alias CheckDetailAddendumA
	aux := struct {
		*Alias
	}{
		(*Alias)(cdAddendumA),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cdAddendumA.setRecordType()
	return nil
}

// String writes the CheckDetailAddendumA struct to a string.
func (cdAddendumA *CheckDetailAddendumA) String() string {
	if cdAddendumA == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(cdAddendumA.recordType)
	buf.WriteString(cdAddendumA.RecordNumberField())
	buf.WriteString(cdAddendumA.ReturnLocationRoutingNumberField())
	buf.WriteString(cdAddendumA.BOFDEndorsementDateField())
	buf.WriteString(cdAddendumA.BOFDItemSequenceNumberField())
	buf.WriteString(cdAddendumA.BOFDAccountNumberField())
	buf.WriteString(cdAddendumA.BOFDBranchCodeField())
	buf.WriteString(cdAddendumA.PayeeNameField())
	buf.WriteString(cdAddendumA.TruncationIndicatorField())
	buf.WriteString(cdAddendumA.BOFDConversionIndicatorField())
	buf.WriteString(cdAddendumA.BOFDCorrectionIndicatorField())
	buf.WriteString(cdAddendumA.UserFieldField())
	buf.WriteString(cdAddendumA.reservedField())
	return buf.String()
}

// Validate performs image cash letter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (cdAddendumA *CheckDetailAddendumA) Validate() error {
	if err := cdAddendumA.fieldInclusion(); err != nil {
		return err
	}
	if cdAddendumA.recordType != "26" {
		msg := fmt.Sprintf(msgRecordType, 26)
		return &FieldError{FieldName: "recordType", Value: cdAddendumA.recordType, Msg: msg}
	}
	if err := cdAddendumA.isNumeric(cdAddendumA.ReturnLocationRoutingNumber); err != nil {
		return &FieldError{FieldName: "ReturnLocationRoutingNumber",
			Value: cdAddendumA.ReturnLocationRoutingNumber, Msg: err.Error()}
	}
	if err := cdAddendumA.isAlphanumericSpecial(cdAddendumA.BOFDAccountNumber); err != nil {
		return &FieldError{FieldName: "BOFDAccountNumber",
			Value: cdAddendumA.BOFDAccountNumber, Msg: err.Error()}
	}
	if err := cdAddendumA.isAlphanumericSpecial(cdAddendumA.BOFDBranchCode); err != nil {
		return &FieldError{FieldName: "BOFDBranchCode",
			Value: cdAddendumA.BOFDBranchCode, Msg: err.Error()}
	}
	if err := cdAddendumA.isAlphanumericSpecial(cdAddendumA.PayeeName); err != nil {
		return &FieldError{FieldName: "PayeeName",
			Value: cdAddendumA.PayeeName, Msg: err.Error()}
	}
	// Mandatory
	if err := cdAddendumA.isTruncationIndicator(cdAddendumA.TruncationIndicator); err != nil {
		return &FieldError{FieldName: "TruncationIndicator",
			Value: cdAddendumA.TruncationIndicator, Msg: err.Error()}
	}
	// Conditional
	if cdAddendumA.BOFDConversionIndicator != "" {
		if err := cdAddendumA.isConversionIndicator(cdAddendumA.BOFDConversionIndicator); err != nil {
			return &FieldError{FieldName: "BOFDConversionIndicator",
				Value: cdAddendumA.BOFDConversionIndicator, Msg: err.Error()}
		}
	}
	// Conditional
	if cdAddendumA.BOFDCorrectionIndicatorField() != "" {
		if err := cdAddendumA.isCorrectionIndicator(cdAddendumA.BOFDCorrectionIndicator); err != nil {
			return &FieldError{FieldName: "BOFDCorrectionIndicator",
				Value: cdAddendumA.BOFDCorrectionIndicatorField(), Msg: err.Error()}
		}
	}
	if err := cdAddendumA.isAlphanumericSpecial(cdAddendumA.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: cdAddendumA.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (cdAddendumA *CheckDetailAddendumA) fieldInclusion() error {
	if cdAddendumA.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: cdAddendumA.recordType,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumA()?"}
	}
	if cdAddendumA.RecordNumber == 0 {
		return &FieldError{FieldName: "RecordNumber",
			Value: cdAddendumA.RecordNumberField(),
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumA()?"}
	}
	if cdAddendumA.ReturnLocationRoutingNumber == "" {
		return &FieldError{FieldName: "ReturnLocationRoutingNumber",
			Value: cdAddendumA.ReturnLocationRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumA()?"}
	}
	if !IsFRBCompatibilityModeEnabled() {
		if cdAddendumA.ReturnLocationRoutingNumberField() == "000000000" {
			return &FieldError{FieldName: "ReturnLocationRoutingNumber",
				Value: cdAddendumA.ReturnLocationRoutingNumber,
				Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumA()?"}
		}
	}
	if cdAddendumA.BOFDEndorsementDate.IsZero() {
		return &FieldError{FieldName: "BOFDEndorsementDate",
			Value: cdAddendumA.BOFDEndorsementDate.String(),
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumA()?"}
	}
	if cdAddendumA.BOFDItemSequenceNumber == "               " {
		return &FieldError{FieldName: "BOFDItemSequenceNumber",
			Value: cdAddendumA.BOFDItemSequenceNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumA()?"}
	}
	if cdAddendumA.TruncationIndicator == "" {
		if IsFRBCompatibilityModeEnabled() {
			cdAddendumA.TruncationIndicator = "N"
		} else {
			return &FieldError{FieldName: "TruncationIndicator",
				Value: cdAddendumA.TruncationIndicator,
				Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumA()?"}
		}
	}
	return nil
}

// RecordNumberField gets a string of the RecordNumber field
func (cdAddendumA *CheckDetailAddendumA) RecordNumberField() string {
	return cdAddendumA.numericField(cdAddendumA.RecordNumber, 1)
}

// ReturnLocationRoutingNumberField gets a string of the ReturnLocationRoutingNumber field
func (cdAddendumA *CheckDetailAddendumA) ReturnLocationRoutingNumberField() string {
	return cdAddendumA.stringField(cdAddendumA.ReturnLocationRoutingNumber, 9)
}

// BOFDEndorsementDateField gets the BOFDEndorsementDate in YYYYMMDD format
func (cdAddendumA *CheckDetailAddendumA) BOFDEndorsementDateField() string {
	return cdAddendumA.formatYYYYMMDDDate(cdAddendumA.BOFDEndorsementDate)
}

// BOFDItemSequenceNumberField gets a string of the BOFDItemSequenceNumber field zero padded
func (cdAddendumA *CheckDetailAddendumA) BOFDItemSequenceNumberField() string {
	return cdAddendumA.alphaField(cdAddendumA.BOFDItemSequenceNumber, 15)
}

// BOFDAccountNumberField gets the BOFDAccountNumber field
func (cdAddendumA *CheckDetailAddendumA) BOFDAccountNumberField() string {
	return cdAddendumA.alphaField(cdAddendumA.BOFDAccountNumber, 18)
}

// BOFDBranchCodeField gets the BOFDBranchCode field
func (cdAddendumA *CheckDetailAddendumA) BOFDBranchCodeField() string {
	return cdAddendumA.alphaField(cdAddendumA.BOFDBranchCode, 5)
}

// PayeeNameField gets the PayeeName field
func (cdAddendumA *CheckDetailAddendumA) PayeeNameField() string {
	return cdAddendumA.alphaField(cdAddendumA.PayeeName, 15)
}

// TruncationIndicatorField gets the TruncationIndicator field
func (cdAddendumA *CheckDetailAddendumA) TruncationIndicatorField() string {
	return cdAddendumA.alphaField(cdAddendumA.TruncationIndicator, 1)
}

// BOFDConversionIndicatorField gets the BOFDConversionIndicator field
func (cdAddendumA *CheckDetailAddendumA) BOFDConversionIndicatorField() string {
	return cdAddendumA.alphaField(cdAddendumA.BOFDConversionIndicator, 1)
}

// BOFDCorrectionIndicatorField gets a string of the BOFDCorrectionIndicator field
func (cdAddendumA *CheckDetailAddendumA) BOFDCorrectionIndicatorField() string {
	return cdAddendumA.numericField(cdAddendumA.BOFDCorrectionIndicator, 1)
}

// UserFieldField gets the UserField field
func (cdAddendumA *CheckDetailAddendumA) UserFieldField() string {
	return cdAddendumA.alphaField(cdAddendumA.UserField, 1)
}

// reservedField gets reserved - blank space
func (cdAddendumA *CheckDetailAddendumA) reservedField() string {
	return cdAddendumA.alphaField(cdAddendumA.reserved, 3)
}

// SetBOFDItemSequenceNumber sets BOFDItemSequenceNumber
func (cdAddendumA *CheckDetailAddendumA) SetBOFDItemSequenceNumber(seq int) string {
	cdAddendumA.BOFDItemSequenceNumber = cdAddendumA.numericField(seq, 15)
	return cdAddendumA.BOFDItemSequenceNumber
}
