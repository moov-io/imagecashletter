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

// Errors specific to a ReturnDetailAddendumD Record

// ReturnDetailAddendumD Record
type ReturnDetailAddendumD struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// RecordNumber is a number representing the order in which each ReturnDetailAddendumD was created.
	// ReturnDetailAddendumD shall be in sequential order starting with 1.  Maximum 99,
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
	// This field is optional in earlier version of the specs.
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
	// creator of this ReturnDetailAddendumD Record for fields other than Payor Bank Routing Number and Amount.
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
	// EndorsingBankIdentifier
	// Values:
	// 0: Depository Bank (BOFD) - this value is used when the ReturnDetailAddendumD Record reflects the Return
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

// NewReturnDetailAddendumD returns a new ReturnDetailAddendumD with default values for non exported fields
func NewReturnDetailAddendumD() ReturnDetailAddendumD {
	rdAddendumD := ReturnDetailAddendumD{}
	rdAddendumD.setRecordType()
	return rdAddendumD
}

func (rdAddendumD *ReturnDetailAddendumD) setRecordType() {
	if rdAddendumD == nil {
		return
	}
	rdAddendumD.recordType = "35"
}

// Parse takes the input record string and parses the ReturnDetailAddendumD values
func (rdAddendumD *ReturnDetailAddendumD) Parse(record string) {
	if utf8.RuneCountInString(record) < 60 {
		return // line too short
	}

	// Character position 1-2, Always "35"
	rdAddendumD.setRecordType()
	// 03-04
	rdAddendumD.RecordNumber = rdAddendumD.parseNumField(record[2:4])
	// 05-13
	rdAddendumD.EndorsingBankRoutingNumber = rdAddendumD.parseStringField(record[4:13])
	// 14-21
	rdAddendumD.BOFDEndorsementBusinessDate = rdAddendumD.parseYYYYMMDDDate(record[13:21])
	// 22-36
	rdAddendumD.EndorsingBankItemSequenceNumber = rdAddendumD.parseStringField(record[21:36])
	// 37-37
	rdAddendumD.TruncationIndicator = rdAddendumD.parseStringField(record[36:37])
	// 38-38
	rdAddendumD.EndorsingBankConversionIndicator = rdAddendumD.parseStringField(record[37:38])
	// 39-39
	rdAddendumD.EndorsingBankCorrectionIndicator = rdAddendumD.parseNumField(record[38:39])
	// 40-40
	rdAddendumD.ReturnReason = rdAddendumD.parseStringField(record[39:40])
	// 41-59
	rdAddendumD.UserField = rdAddendumD.parseStringField(record[40:59])
	// 60-60
	rdAddendumD.EndorsingBankIdentifier = rdAddendumD.parseNumField(record[59:60])
	// 61-80
	rdAddendumD.reserved = "                    "
}

func (rdAddendumD *ReturnDetailAddendumD) UnmarshalJSON(data []byte) error {
	type Alias ReturnDetailAddendumD
	aux := struct {
		*Alias
	}{
		(*Alias)(rdAddendumD),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rdAddendumD.setRecordType()
	return nil
}

// String writes the ReturnDetailAddendumD struct to a string.
func (rdAddendumD *ReturnDetailAddendumD) String() string {
	if rdAddendumD == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rdAddendumD.recordType)
	buf.WriteString(rdAddendumD.RecordNumberField())
	buf.WriteString(rdAddendumD.EndorsingBankRoutingNumberField())
	buf.WriteString(rdAddendumD.BOFDEndorsementBusinessDateField())
	buf.WriteString(rdAddendumD.EndorsingBankItemSequenceNumberField())
	buf.WriteString(rdAddendumD.TruncationIndicatorField())
	buf.WriteString(rdAddendumD.EndorsingBankConversionIndicatorField())
	buf.WriteString(rdAddendumD.EndorsingBankCorrectionIndicatorField())
	buf.WriteString(rdAddendumD.ReturnReasonField())
	buf.WriteString(rdAddendumD.UserFieldField())
	buf.WriteString(rdAddendumD.EndorsingBankIdentifierField())
	buf.WriteString(rdAddendumD.reservedField())
	return buf.String()
}

// Validate performs image cash letter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rdAddendumD *ReturnDetailAddendumD) Validate() error {
	if err := rdAddendumD.fieldInclusion(); err != nil {
		return err
	}
	if rdAddendumD.recordType != "35" {
		msg := fmt.Sprintf(msgRecordType, 35)
		return &FieldError{FieldName: "recordType", Value: rdAddendumD.recordType, Msg: msg}
	}
	if err := rdAddendumD.isNumeric(rdAddendumD.EndorsingBankRoutingNumber); err != nil {
		return &FieldError{FieldName: "EndorsingBankRoutingNumber",
			Value: rdAddendumD.EndorsingBankRoutingNumber, Msg: err.Error()}
	}
	if err := rdAddendumD.isNumeric(rdAddendumD.EndorsingBankItemSequenceNumber); err != nil {
		return &FieldError{FieldName: "EndorsingBankItemSequenceNumber",
			Value: rdAddendumD.EndorsingBankItemSequenceNumber, Msg: msgNumeric}
	}
	// Mandatory
	if err := rdAddendumD.isTruncationIndicator(rdAddendumD.TruncationIndicator); err != nil {
		return &FieldError{FieldName: "TruncationIndicator",
			Value: rdAddendumD.TruncationIndicator, Msg: err.Error()}
	}
	// Conditional
	if rdAddendumD.EndorsingBankConversionIndicator != "" {
		if err := rdAddendumD.isConversionIndicator(rdAddendumD.EndorsingBankConversionIndicator); err != nil {
			return &FieldError{FieldName: "EndorsingBankConversionIndicator",
				Value: rdAddendumD.EndorsingBankConversionIndicator, Msg: err.Error()}
		}
	}
	// Conditional
	if rdAddendumD.EndorsingBankCorrectionIndicatorField() != "" {
		if err := rdAddendumD.isCorrectionIndicator(rdAddendumD.EndorsingBankCorrectionIndicator); err != nil {
			return &FieldError{FieldName: "EndorsingBankCorrectionIndicator",
				Value: rdAddendumD.EndorsingBankCorrectionIndicatorField(), Msg: err.Error()}
		}
	}
	if err := rdAddendumD.isAlphanumeric(rdAddendumD.ReturnReason); err != nil {
		return &FieldError{FieldName: "ReturnReason",
			Value: rdAddendumD.ReturnReason, Msg: err.Error()}
	}
	if err := rdAddendumD.isAlphanumericSpecial(rdAddendumD.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: rdAddendumD.UserField, Msg: err.Error()}
	}
	if err := rdAddendumD.isEndorsingBankIdentifier(rdAddendumD.EndorsingBankIdentifier); err != nil {
		return &FieldError{FieldName: "EndorsingBankIdentifier",
			Value: rdAddendumD.EndorsingBankIdentifierField(), Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rdAddendumD *ReturnDetailAddendumD) fieldInclusion() error {
	if rdAddendumD.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: rdAddendumD.recordType,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumD()?"}
	}
	if rdAddendumD.RecordNumber == 0 {
		return &FieldError{FieldName: "RecordNumber",
			Value: rdAddendumD.RecordNumberField(),
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumD()?"}
	}
	if rdAddendumD.EndorsingBankRoutingNumber == "" {
		return &FieldError{FieldName: "EndorsingBankRoutingNumber",
			Value: rdAddendumD.EndorsingBankRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumD()?"}
	}
	if rdAddendumD.EndorsingBankRoutingNumberField() == "000000000" {
		return &FieldError{FieldName: "EndorsingBankRoutingNumber",
			Value: rdAddendumD.EndorsingBankRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumD()?"}
	}
	if rdAddendumD.BOFDEndorsementBusinessDate.IsZero() {
		return &FieldError{FieldName: "BOFDEndorsementBusinessDate",
			Value: rdAddendumD.BOFDEndorsementBusinessDate.String(),
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumD()?"}
	}
	if rdAddendumD.TruncationIndicator == "" {
		return &FieldError{FieldName: "TruncationIndicator",
			Value: rdAddendumD.TruncationIndicator,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumD()?"}
	}
	return nil
}

// RecordNumberField gets a string of the RecordNumber field
func (rdAddendumD *ReturnDetailAddendumD) RecordNumberField() string {
	return rdAddendumD.numericField(rdAddendumD.RecordNumber, 2)
}

// EndorsingBankRoutingNumberField gets a string of the EndorsingBankRoutingNumber field
func (rdAddendumD *ReturnDetailAddendumD) EndorsingBankRoutingNumberField() string {
	return rdAddendumD.stringField(rdAddendumD.EndorsingBankRoutingNumber, 9)
}

// BOFDEndorsementBusinessDateField gets the BOFDEndorsementBusinessDate in YYYYMMDD format
func (rdAddendumD *ReturnDetailAddendumD) BOFDEndorsementBusinessDateField() string {
	return rdAddendumD.formatYYYYMMDDDate(rdAddendumD.BOFDEndorsementBusinessDate)
}

// EndorsingBankItemSequenceNumberField gets the EndorsingBankItemSequenceNumber field
func (rdAddendumD *ReturnDetailAddendumD) EndorsingBankItemSequenceNumberField() string {
	return rdAddendumD.alphaField(rdAddendumD.EndorsingBankItemSequenceNumber, 15)
}

// TruncationIndicatorField gets the TruncationIndicator field
func (rdAddendumD *ReturnDetailAddendumD) TruncationIndicatorField() string {
	return rdAddendumD.alphaField(rdAddendumD.TruncationIndicator, 1)
}

// EndorsingBankConversionIndicatorField gets the EndorsingBankConversionIndicator field
func (rdAddendumD *ReturnDetailAddendumD) EndorsingBankConversionIndicatorField() string {
	return rdAddendumD.alphaField(rdAddendumD.EndorsingBankConversionIndicator, 1)
}

// EndorsingBankCorrectionIndicatorField gets a string of the EndorsingBankCorrectionIndicator field
func (rdAddendumD *ReturnDetailAddendumD) EndorsingBankCorrectionIndicatorField() string {
	return rdAddendumD.numericField(rdAddendumD.EndorsingBankCorrectionIndicator, 1)
}

// ReturnReasonField gets the ReturnReason field
func (rdAddendumD *ReturnDetailAddendumD) ReturnReasonField() string {
	return rdAddendumD.alphaField(rdAddendumD.ReturnReason, 1)
}

// UserFieldField gets the UserField field
func (rdAddendumD *ReturnDetailAddendumD) UserFieldField() string {
	return rdAddendumD.alphaField(rdAddendumD.UserField, 19)
}

// EndorsingBankIdentifierField gets the EndorsingBankIdentifier field
func (rdAddendumD *ReturnDetailAddendumD) EndorsingBankIdentifierField() string {
	return rdAddendumD.numericField(rdAddendumD.EndorsingBankIdentifier, 1)
}

// reservedField gets reserved - blank space
func (rdAddendumD *ReturnDetailAddendumD) reservedField() string {
	return rdAddendumD.alphaField(rdAddendumD.reserved, 20)
}

// SetEndorsingBankItemSequenceNumber sets EndorsingBankItemSequenceNumber
func (rdAddendumD *ReturnDetailAddendumD) SetEndorsingBankItemSequenceNumber(seq int) string {
	rdAddendumD.EndorsingBankItemSequenceNumber = rdAddendumD.numericField(seq, 15)
	return rdAddendumD.EndorsingBankItemSequenceNumber
}
