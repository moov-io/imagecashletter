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
	EndorsingBankRoutingNumber string `json:"endorsingBankRoutingNumber"`
	// BOFDEndorsementBusinessDate is the business date the check was endorsed.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	BOFDEndorsementBusinessDate time.Time `json:"bofdEndorsementBusinessDate"`
	// EndorsingItemSequenceNumber is a number that identifies the item at the endorsing bank.
	EndorsingBankItemSequenceNumber string `json:"endorsingBankItemSequenceNumber"`
	// TruncationIndicator identifies if the institution truncated the original check item.
	// Values: Y: Yes this institution truncated this original check item and this is first endorsement
	// for the institution.
	// N: No this institution did not truncate the original check or, this is not the first endorsement for the
	// institution or, this item is an IRD not an original check item (EPC equals 4).
	TruncationIndicator string `json:"truncationIndicator"`
	// EndorsingConversionIndicator is a code that indicates the conversion within the processing institution among
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
	EndorsingBankConversionIndicator string `json:"endorsingBankConversionIndicator"`
	// EndorsingCorrectionIndicator identifies whether and how the MICR line of this item was repaired by the
	// creator of this CheckDetailAddendumC Record for fields other than Payor Bank Routing Number and Amount.
	// Values:
	// 0: No Repair
	// 1: Repaired (form of repair unknown)
	// 2: Repaired without Operator intervention
	// 3: Repaired with Operator intervention
	// 4: Undetermined if repair has been done or no
	EndorsingBankCorrectionIndicator int `json:"endorsingBankCorrectionIndicator"`
	// ReturnReason is a code that indicates the reason for non-payment.
	ReturnReason string `json:"returnReason"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	//EndorsingBankIdentifier
	// Values:
	// 0: Depository Bank (BOFD) - this value is used when the CheckDetailAddendumC Record reflects the Return
	// Processing Bank in lieu of BOFD.
	// 1: Other Collecting Bank
	// 2: Other Returning Bank
	// 3: Payor Bank
	EndorsingBankIdentifier int `json:"endorsingBankIdentifier"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for image cash letter data validation
	validator
	// converters is composed for image cash letter to golang Converters
	converters
}

// NewCheckDetailAddendumC returns a new CheckDetailAddendumC with default values for non exported fields
func NewCheckDetailAddendumC() CheckDetailAddendumC {
	cdAddendumC := CheckDetailAddendumC{}
	cdAddendumC.setRecordType()
	return cdAddendumC
}

func (cdAddendumC *CheckDetailAddendumC) setRecordType() {
	if cdAddendumC == nil {
		return
	}
	cdAddendumC.recordType = "28"
}

// Parse takes the input record string and parses the CheckDetailAddendumC values
func (cdAddendumC *CheckDetailAddendumC) Parse(record string) {
	if utf8.RuneCountInString(record) < 60 {
		return // line too short
	}

	// Character position 1-2, Always "28"
	cdAddendumC.setRecordType()
	// 03-04
	cdAddendumC.RecordNumber = cdAddendumC.parseNumField(record[2:4])
	// 05-13
	cdAddendumC.EndorsingBankRoutingNumber = cdAddendumC.parseStringField(record[4:13])
	// 14-21
	cdAddendumC.BOFDEndorsementBusinessDate = cdAddendumC.parseYYYYMMDDDate(record[13:21])
	// 22-36
	cdAddendumC.EndorsingBankItemSequenceNumber = cdAddendumC.parseStringField(record[21:36])
	// 37-37
	cdAddendumC.TruncationIndicator = cdAddendumC.parseStringField(record[36:37])
	// 38-38
	cdAddendumC.EndorsingBankConversionIndicator = cdAddendumC.parseStringField(record[37:38])
	// 39-39
	cdAddendumC.EndorsingBankCorrectionIndicator = cdAddendumC.parseNumField(record[38:39])
	// 40-40
	cdAddendumC.ReturnReason = cdAddendumC.parseStringField(record[39:40])
	// 41-59
	cdAddendumC.UserField = cdAddendumC.parseStringField(record[40:59])
	// 60-60
	cdAddendumC.EndorsingBankIdentifier = cdAddendumC.parseNumField(record[59:60])
	// 61-80
	cdAddendumC.reserved = "                    "
}

func (cdAddendumC *CheckDetailAddendumC) UnmarshalJSON(data []byte) error {
	type Alias CheckDetailAddendumC
	aux := struct {
		*Alias
	}{
		(*Alias)(cdAddendumC),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cdAddendumC.setRecordType()
	return nil
}

// String writes the CheckDetailAddendumC struct to a string.
func (cdAddendumC *CheckDetailAddendumC) String() string {
	if cdAddendumC == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(cdAddendumC.recordType)
	buf.WriteString(cdAddendumC.RecordNumberField())
	buf.WriteString(cdAddendumC.EndorsingBankRoutingNumberField())
	buf.WriteString(cdAddendumC.BOFDEndorsementBusinessDateField())
	buf.WriteString(cdAddendumC.EndorsingBankItemSequenceNumberField())
	buf.WriteString(cdAddendumC.TruncationIndicatorField())
	buf.WriteString(cdAddendumC.EndorsingBankConversionIndicatorField())
	buf.WriteString(cdAddendumC.EndorsingBankCorrectionIndicatorField())
	buf.WriteString(cdAddendumC.ReturnReasonField())
	buf.WriteString(cdAddendumC.UserFieldField())
	buf.WriteString(cdAddendumC.EndorsingBankIdentifierField())
	buf.WriteString(cdAddendumC.reservedField())
	return buf.String()
}

// Validate performs image cash letter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (cdAddendumC *CheckDetailAddendumC) Validate() error {
	if err := cdAddendumC.fieldInclusion(); err != nil {
		return err
	}
	if cdAddendumC.recordType != "28" {
		msg := fmt.Sprintf(msgRecordType, 28)
		return &FieldError{FieldName: "recordType", Value: cdAddendumC.recordType, Msg: msg}
	}
	if err := cdAddendumC.isNumeric(cdAddendumC.EndorsingBankRoutingNumber); err != nil {
		return &FieldError{FieldName: "EndorsingBankRoutingNumber",
			Value: cdAddendumC.EndorsingBankRoutingNumber, Msg: err.Error()}
	}
	// Mandatory
	if err := cdAddendumC.isTruncationIndicator(cdAddendumC.TruncationIndicator); err != nil {
		return &FieldError{FieldName: "TruncationIndicator",
			Value: cdAddendumC.TruncationIndicator, Msg: err.Error()}
	}
	// Conditional
	if cdAddendumC.EndorsingBankConversionIndicator != "" {
		if err := cdAddendumC.isConversionIndicator(cdAddendumC.EndorsingBankConversionIndicator); err != nil {
			return &FieldError{FieldName: "EndorsingBankConversionIndicator",
				Value: cdAddendumC.EndorsingBankConversionIndicator, Msg: err.Error()}
		}
	}
	// Conditional
	if cdAddendumC.EndorsingBankCorrectionIndicatorField() != "" {
		if err := cdAddendumC.isCorrectionIndicator(cdAddendumC.EndorsingBankCorrectionIndicator); err != nil {
			return &FieldError{FieldName: "EndorsingBankCorrectionIndicator",
				Value: cdAddendumC.EndorsingBankCorrectionIndicatorField(), Msg: err.Error()}
		}
	}
	if err := cdAddendumC.isAlphanumeric(cdAddendumC.ReturnReason); err != nil {
		return &FieldError{FieldName: "ReturnReason",
			Value: cdAddendumC.ReturnReason, Msg: err.Error()}
	}
	if err := cdAddendumC.isAlphanumericSpecial(cdAddendumC.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: cdAddendumC.UserField, Msg: err.Error()}
	}
	if err := cdAddendumC.isEndorsingBankIdentifier(cdAddendumC.EndorsingBankIdentifier); err != nil {
		return &FieldError{FieldName: "EndorsingBankIdentifier",
			Value: cdAddendumC.EndorsingBankIdentifierField(), Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (cdAddendumC *CheckDetailAddendumC) fieldInclusion() error {
	if cdAddendumC.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: cdAddendumC.recordType,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumC()?"}
	}
	if cdAddendumC.RecordNumber == 0 {
		return &FieldError{FieldName: "RecordNumber",
			Value: cdAddendumC.RecordNumberField(),
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumC()?"}
	}
	if cdAddendumC.EndorsingBankRoutingNumber == "" {
		return &FieldError{FieldName: "EndorsingBankRoutingNumber",
			Value: cdAddendumC.EndorsingBankRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumC()?"}
	}
	if cdAddendumC.EndorsingBankRoutingNumberField() == "000000000" && !IsFRBCompatibilityModeEnabled() {
		return &FieldError{FieldName: "EndorsingBankRoutingNumber",
			Value: cdAddendumC.EndorsingBankRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumC()?"}
	}
	if cdAddendumC.BOFDEndorsementBusinessDate.IsZero() {
		return &FieldError{FieldName: "BOFDEndorsementBusinessDate",
			Value: cdAddendumC.BOFDEndorsementBusinessDate.String(),
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumC()?"}
	}
	if !IsFRBCompatibilityModeEnabled() && cdAddendumC.EndorsingBankItemSequenceNumberField() == "               " {
		return &FieldError{FieldName: "EndorsingBankItemSequenceNumber",
			Value: cdAddendumC.EndorsingBankItemSequenceNumber,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumC()?"}
	}
	if cdAddendumC.TruncationIndicator == "" {
		return &FieldError{FieldName: "TruncationIndicator",
			Value: cdAddendumC.TruncationIndicator,
			Msg:   msgFieldInclusion + ", did you use CheckDetailAddendumC()?"}
	}
	return nil
}

// RecordNumberField gets a string of the RecordNumber field
func (cdAddendumC *CheckDetailAddendumC) RecordNumberField() string {
	return cdAddendumC.numericField(cdAddendumC.RecordNumber, 2)
}

// EndorsingBankRoutingNumberField gets a string of the EndorsingBankRoutingNumber field
func (cdAddendumC *CheckDetailAddendumC) EndorsingBankRoutingNumberField() string {
	return cdAddendumC.stringField(cdAddendumC.EndorsingBankRoutingNumber, 9)
}

// BOFDEndorsementBusinessDateField gets the BOFDEndorsementBusinessDate in YYYYMMDD format
func (cdAddendumC *CheckDetailAddendumC) BOFDEndorsementBusinessDateField() string {
	return cdAddendumC.formatYYYYMMDDDate(cdAddendumC.BOFDEndorsementBusinessDate)
}

// EndorsingBankItemSequenceNumberField gets the EndorsingBankItemSequenceNumber field
func (cdAddendumC *CheckDetailAddendumC) EndorsingBankItemSequenceNumberField() string {
	return cdAddendumC.alphaField(cdAddendumC.EndorsingBankItemSequenceNumber, 15)
}

// TruncationIndicatorField gets the TruncationIndicator field
func (cdAddendumC *CheckDetailAddendumC) TruncationIndicatorField() string {
	return cdAddendumC.alphaField(cdAddendumC.TruncationIndicator, 1)
}

// EndorsingBankConversionIndicatorField gets the EndorsingBankConversionIndicator field
func (cdAddendumC *CheckDetailAddendumC) EndorsingBankConversionIndicatorField() string {
	return cdAddendumC.alphaField(cdAddendumC.EndorsingBankConversionIndicator, 1)
}

// EndorsingBankCorrectionIndicatorField gets a string of the EndorsingBankCorrectionIndicator field
func (cdAddendumC *CheckDetailAddendumC) EndorsingBankCorrectionIndicatorField() string {
	return cdAddendumC.numericField(cdAddendumC.EndorsingBankCorrectionIndicator, 1)
}

// ReturnReasonField gets the ReturnReason field
func (cdAddendumC *CheckDetailAddendumC) ReturnReasonField() string {
	return cdAddendumC.alphaField(cdAddendumC.ReturnReason, 1)
}

// UserFieldField gets the UserField field
func (cdAddendumC *CheckDetailAddendumC) UserFieldField() string {
	return cdAddendumC.alphaField(cdAddendumC.UserField, 19)
}

// EndorsingBankIdentifierField gets the EndorsingBankIdentifier field
func (cdAddendumC *CheckDetailAddendumC) EndorsingBankIdentifierField() string {
	return cdAddendumC.numericField(cdAddendumC.EndorsingBankIdentifier, 1)
}

// reservedField gets reserved - blank space
func (cdAddendumC *CheckDetailAddendumC) reservedField() string {
	return cdAddendumC.alphaField(cdAddendumC.reserved, 20)
}

// SetEndorsingBankItemSequenceNumber sets EndorsingBankItemSequenceNumber
func (cdAddendumC *CheckDetailAddendumC) SetEndorsingBankItemSequenceNumber(seq int) string {
	itemSequence := strconv.Itoa(seq)
	cdAddendumC.EndorsingBankItemSequenceNumber = itemSequence
	return cdAddendumC.EndorsingBankItemSequenceNumber
}
