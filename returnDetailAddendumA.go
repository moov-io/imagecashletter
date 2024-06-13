// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Errors specific to a ReturnDetailAddendumA Record

// ReturnDetailAddendumA Record
type ReturnDetailAddendumA struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// RecordNumber represents the chronological order (oldest to newest) in which each ReturnDetailAddendumA was
	// created. The ReturnDetailAddendumA shall be in sequential order according to this field. ReturnDetailAddendumA
	// RecordNumber9s) shall be in sequential order starting with 1, indicating the oldest addendum, and incrementing
	// by 1 for each subsequent addendum.
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
	// According to spec version of x9.100-187-2016, this field is conditional. It is required if creating
	// a new Check Detail Addendum A Record, otherwise conditional.
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

// NewReturnDetailAddendumA returns a new ReturnDetailAddendumA with default values for non exported fields
func NewReturnDetailAddendumA() ReturnDetailAddendumA {
	rdAddendumA := ReturnDetailAddendumA{}
	rdAddendumA.setRecordType()
	return rdAddendumA
}

func (rdAddendumA *ReturnDetailAddendumA) setRecordType() {
	if rdAddendumA == nil {
		return
	}
	rdAddendumA.recordType = "32"
}

// Parse takes the input record string and parses the ReturnDetailAddendumA values
func (rdAddendumA *ReturnDetailAddendumA) Parse(record string) {
	if utf8.RuneCountInString(record) < 77 {
		return // line too short
	}

	// Character position 1-2, Always "32"
	rdAddendumA.setRecordType()
	// 03-03
	rdAddendumA.RecordNumber = rdAddendumA.parseNumField(record[2:3])
	// 04-12
	rdAddendumA.ReturnLocationRoutingNumber = rdAddendumA.parseStringField(record[3:12])
	// 13-20
	rdAddendumA.BOFDEndorsementDate = rdAddendumA.parseYYYYMMDDDate(record[12:20])
	// 21-35
	rdAddendumA.BOFDItemSequenceNumber = rdAddendumA.parseStringField(record[20:35])
	// 36-53
	rdAddendumA.BOFDAccountNumber = rdAddendumA.parseStringField(record[35:53])
	// 54-58
	rdAddendumA.BOFDBranchCode = rdAddendumA.parseStringField(record[53:58])
	// 59-73
	rdAddendumA.PayeeName = rdAddendumA.parseStringField(record[58:73])
	// 74-74
	rdAddendumA.TruncationIndicator = rdAddendumA.parseStringField(record[73:74])
	// 75-75
	rdAddendumA.BOFDConversionIndicator = rdAddendumA.parseStringField(record[74:75])
	// 76-76
	rdAddendumA.BOFDCorrectionIndicator = rdAddendumA.parseNumField(record[75:76])
	// 77-77
	rdAddendumA.UserField = rdAddendumA.parseStringField(record[76:77])
	// 78-80
	rdAddendumA.reserved = "   "

}

func (rdAddendumA *ReturnDetailAddendumA) UnmarshalJSON(data []byte) error {
	type Alias ReturnDetailAddendumA
	aux := struct {
		*Alias
	}{
		(*Alias)(rdAddendumA),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rdAddendumA.setRecordType()
	return nil
}

// String writes the ReturnDetailAddendumA struct to a string.
func (rdAddendumA *ReturnDetailAddendumA) String() string {
	if rdAddendumA == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rdAddendumA.recordType)
	buf.WriteString(rdAddendumA.RecordNumberField())
	buf.WriteString(rdAddendumA.ReturnLocationRoutingNumberField())
	buf.WriteString(rdAddendumA.BOFDEndorsementDateField())
	buf.WriteString(rdAddendumA.BOFDItemSequenceNumberField())
	buf.WriteString(rdAddendumA.BOFDAccountNumberField())
	buf.WriteString(rdAddendumA.BOFDBranchCodeField())
	buf.WriteString(rdAddendumA.PayeeNameField())
	buf.WriteString(rdAddendumA.TruncationIndicatorField())
	buf.WriteString(rdAddendumA.BOFDConversionIndicatorField())
	buf.WriteString(rdAddendumA.BOFDCorrectionIndicatorField())
	buf.WriteString(rdAddendumA.UserFieldField())
	buf.WriteString(rdAddendumA.reservedField())
	return buf.String()
}

// Validate performs image cash letter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rdAddendumA *ReturnDetailAddendumA) Validate() error {
	if err := rdAddendumA.fieldInclusion(); err != nil {
		return err
	}
	if rdAddendumA.recordType != "32" {
		msg := fmt.Sprintf(msgRecordType, 32)
		return &FieldError{FieldName: "recordType", Value: rdAddendumA.recordType, Msg: msg}
	}
	if err := rdAddendumA.isNumeric(rdAddendumA.ReturnLocationRoutingNumber); err != nil {
		return &FieldError{FieldName: "ReturnLocationRoutingNumber",
			Value: rdAddendumA.ReturnLocationRoutingNumber, Msg: err.Error()}
	}
	if err := rdAddendumA.isAlphanumericSpecial(rdAddendumA.BOFDAccountNumber); err != nil {
		return &FieldError{FieldName: "BOFDAccountNumber",
			Value: rdAddendumA.BOFDAccountNumber, Msg: err.Error()}
	}
	if err := rdAddendumA.isAlphanumericSpecial(rdAddendumA.BOFDBranchCode); err != nil {
		return &FieldError{FieldName: "BOFDBranchCode",
			Value: rdAddendumA.BOFDBranchCode, Msg: err.Error()}
	}
	if err := rdAddendumA.isAlphanumericSpecial(rdAddendumA.PayeeName); err != nil {
		return &FieldError{FieldName: "PayeeName",
			Value: rdAddendumA.PayeeName, Msg: err.Error()}
	}
	// Mandatory
	if err := rdAddendumA.isTruncationIndicator(rdAddendumA.TruncationIndicator); err != nil {
		return &FieldError{FieldName: "TruncationIndicator",
			Value: rdAddendumA.TruncationIndicator, Msg: err.Error()}
	}
	// Conditional
	if rdAddendumA.BOFDConversionIndicator != "" {
		if err := rdAddendumA.isConversionIndicator(rdAddendumA.BOFDConversionIndicator); err != nil {
			return &FieldError{FieldName: "BOFDConversionIndicator",
				Value: rdAddendumA.BOFDConversionIndicator, Msg: err.Error()}
		}
	}
	// Conditional
	if rdAddendumA.BOFDCorrectionIndicatorField() != "" {
		if err := rdAddendumA.isCorrectionIndicator(rdAddendumA.BOFDCorrectionIndicator); err != nil {
			return &FieldError{FieldName: "BOFDCorrectionIndicator",
				Value: rdAddendumA.BOFDCorrectionIndicatorField(), Msg: err.Error()}
		}
	}
	if err := rdAddendumA.isAlphanumericSpecial(rdAddendumA.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: rdAddendumA.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rdAddendumA *ReturnDetailAddendumA) fieldInclusion() error {
	if rdAddendumA.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: rdAddendumA.recordType,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumA()?"}
	}
	if rdAddendumA.RecordNumber == 0 {
		return &FieldError{FieldName: "RecordNumber",
			Value: rdAddendumA.RecordNumberField(),
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumA()?"}
	}
	if rdAddendumA.ReturnLocationRoutingNumber == "" {
		return &FieldError{FieldName: "ReturnLocationRoutingNumber",
			Value: rdAddendumA.ReturnLocationRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumA()?"}
	}
	if rdAddendumA.ReturnLocationRoutingNumberField() == "000000000" {
		return &FieldError{FieldName: "ReturnLocationRoutingNumber",
			Value: rdAddendumA.ReturnLocationRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumA()?"}
	}
	if rdAddendumA.BOFDEndorsementDate.IsZero() && !IsFRBCompatibilityModeEnabled() {
		return &FieldError{FieldName: "BOFDEndorsementDate",
			Value: rdAddendumA.BOFDEndorsementDate.String(),
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumA()?"}
	}
	if rdAddendumA.TruncationIndicator == "" {
		return &FieldError{FieldName: "TruncationIndicator",
			Value: rdAddendumA.TruncationIndicator,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumA()?"}
	}
	return nil
}

// RecordNumberField gets a string of the RecordNumber field
func (rdAddendumA *ReturnDetailAddendumA) RecordNumberField() string {
	return rdAddendumA.numericField(rdAddendumA.RecordNumber, 1)
}

// ReturnLocationRoutingNumberField gets a string of the ReturnLocationRoutingNumber field
func (rdAddendumA *ReturnDetailAddendumA) ReturnLocationRoutingNumberField() string {
	return rdAddendumA.stringField(rdAddendumA.ReturnLocationRoutingNumber, 9)
}

// BOFDEndorsementDateField gets the BOFDEndorsementDate in YYYYMMDD format
func (rdAddendumA *ReturnDetailAddendumA) BOFDEndorsementDateField() string {
	return rdAddendumA.formatYYYYMMDDDate(rdAddendumA.BOFDEndorsementDate)
}

// BOFDItemSequenceNumberField gets a string of the BOFDItemSequenceNumber field zero padded
func (rdAddendumA *ReturnDetailAddendumA) BOFDItemSequenceNumberField() string {
	return rdAddendumA.alphaField(rdAddendumA.BOFDItemSequenceNumber, 15)
}

// BOFDAccountNumberField gets the BOFDAccountNumber field
func (rdAddendumA *ReturnDetailAddendumA) BOFDAccountNumberField() string {
	return rdAddendumA.alphaField(rdAddendumA.BOFDAccountNumber, 18)
}

// BOFDBranchCodeField gets the BOFDBranchCode field
func (rdAddendumA *ReturnDetailAddendumA) BOFDBranchCodeField() string {
	return rdAddendumA.alphaField(rdAddendumA.BOFDBranchCode, 5)
}

// PayeeNameField gets the PayeeName field
func (rdAddendumA *ReturnDetailAddendumA) PayeeNameField() string {
	return rdAddendumA.alphaField(rdAddendumA.PayeeName, 15)
}

// TruncationIndicatorField gets the TruncationIndicator field
func (rdAddendumA *ReturnDetailAddendumA) TruncationIndicatorField() string {
	return rdAddendumA.alphaField(rdAddendumA.TruncationIndicator, 1)
}

// BOFDConversionIndicatorField gets the BOFDConversionIndicator field
func (rdAddendumA *ReturnDetailAddendumA) BOFDConversionIndicatorField() string {
	return rdAddendumA.alphaField(rdAddendumA.BOFDConversionIndicator, 1)
}

// BOFDCorrectionIndicatorField gets a string of the BOFDCorrectionIndicator field
func (rdAddendumA *ReturnDetailAddendumA) BOFDCorrectionIndicatorField() string {
	return rdAddendumA.numericField(rdAddendumA.BOFDCorrectionIndicator, 1)
}

// UserFieldField gets the UserField field
func (rdAddendumA *ReturnDetailAddendumA) UserFieldField() string {
	return rdAddendumA.alphaField(rdAddendumA.UserField, 1)
}

// reservedField gets reserved - blank space
func (rdAddendumA *ReturnDetailAddendumA) reservedField() string {
	return rdAddendumA.alphaField(rdAddendumA.reserved, 3)
}

// SetBOFDItemSequenceNumber sets BOFDItemSequenceNumber
func (rdAddendumA *ReturnDetailAddendumA) SetBOFDItemSequenceNumber(seq int) string {
	itemSequence := strconv.Itoa(seq)
	rdAddendumA.BOFDItemSequenceNumber = itemSequence
	return rdAddendumA.BOFDItemSequenceNumber
}
