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

// Errors specific to a CreditItem Record

// Current Implementation: CreditItem(s) Precede CheckDetail(s) - CreditItem(s) outside the leading Bundle
// and Within the First Cash Letter.  Please adjust reader and writer for your specific clearing arrangement
// implementation or contact MOOV for your particular implementation.
//
// FileHeader
// CashLetterHeader Record
// CreditItem
// BundleHeader Record
// 1st CheckDetail
// 2nd CheckDetail
// N* CheckDetail
// Last CheckDetail
// BundleControl
// BundleHeader
// 1st CheckDetail
// 2nd CheckDetail
// N* CheckDetail
// Last CheckDetail
// BundleControl
// CashLetterControl
// FileControl

// CreditItem Record
type CreditItem struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// AuxiliaryOnUs identifies a code used on commercial checks at the discretion of the payor bank.
	AuxiliaryOnUs string `json:"auxiliaryOnUs"`
	// ExternalProcessingCode identifies a code used for special purposes as authorized by the Accredited
	// Standards Committee X9. Also known as Position 44.
	ExternalProcessingCode string `json:"externalProcessingCode"`
	// PostingBankRoutingNumber is a routing number assigned by the posting bank to identify this credit.
	// Format: TTTTAAAA, where:
	// TTTT: Federal Reserve Prefix
	// AAAA: ABA Institution Identifier
	PostingBankRoutingNumber string `json:"postingBankRoutingNumber"`
	// OnUs identifies data specified by the payor bank. On-Us data usually consists of the payor’s account number,
	// a serial number or transaction code, or both.
	OnUs string `json:"onUs"`
	// Amount identifies the amount of the check.  All amounts fields have two implied decimal points.
	// e.g., 100000 is $1,000.00
	ItemAmount int `json:"itemAmount"`
	// CreditItemSequenceNumber identifies a number assigned by the institution that creates the CreditItem
	CreditItemSequenceNumber string `json:"creditItemSequenceNumber"`
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
	// C: Image provided separately, no paper provided
	// D: Image provided separately, no paper provided, image upon request
	// E: Image and paper provided separately
	// F: Image and paper provided separately, image upon request
	// G: Image included, no paper provided
	// H: Image included, no paper provided, image upon request
	// I: Image included, paper provided separately
	// J: Image included, paper provided separately, image upon request
	// K: No image provided, no paper provided
	// L: No image provided, no paper provided, image upon request
	DocumentationTypeIndicator string `json:"documentationTypeIndicator"`
	// AccountTypeCode is a code that indicates the type of account to which this CreditItem is associated.
	// Values:
	// 0: Unknown
	// 1: DDA account
	// 2: General Ledger account
	// 3: Savings account
	// 4: Money Market account
	// 5: Other account
	AccountTypeCode string `json:"accountTypeCode"`
	// SourceWorkCode is a code used to identify the source of the work associated with this CreditItem.
	// Values:
	// 00: Unknown
	// 01: Internal–ATM
	// 02: Internal–Branch
	// 03: Internal–Other
	// 04: External–Bank to Bank (Correspondent)
	// 05: External–Business to Bank (Customer)
	// 06: External–Business to Bank Remote Capture
	// 07: External–Processor to Bank
	// 08: External–Bank to Processor
	// 09: Lockbox
	// 10: International–Internal
	// 11: International–External
	// 21–50: User Defined
	SourceWorkCode string `json:"sourceWorkCode"`
	// UserField is a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	validator
	// converters is composed for imagecashletter to golang Converters
	converters
}

// NewCreditItem returns a new CreditItem with default values for non exported fields
func NewCreditItem() *CreditItem {
	ci := &CreditItem{}
	ci.setRecordType()
	return ci
}

func (ci *CreditItem) setRecordType() {
	if ci == nil {
		return
	}
	ci.recordType = "62"
}

// Parse takes the input record string and parses the CreditItem values
func (ci *CreditItem) Parse(record string) {
	if utf8.RuneCountInString(record) < 96 {
		return // line is too short
	}
	// Character position 1-2, Always "62"
	ci.setRecordType()
	// 03-17
	ci.AuxiliaryOnUs = ci.parseStringField(record[2:17])
	// 18-18
	ci.ExternalProcessingCode = ci.parseStringField(record[17:18])
	// 19-27
	ci.PostingBankRoutingNumber = ci.parseStringField(record[18:27])
	// 28-47
	ci.OnUs = ci.parseStringField(record[27:47])
	// 48-61
	ci.ItemAmount = ci.parseNumField(record[47:61])
	// 62-76
	ci.CreditItemSequenceNumber = ci.parseStringField(record[61:76])
	// 77-77
	ci.DocumentationTypeIndicator = ci.parseStringField(record[76:77])
	// 78-78
	ci.AccountTypeCode = ci.parseStringField(record[77:78])
	// 79-80
	ci.SourceWorkCode = ci.parseStringField(record[78:80])
	// 81-96
	ci.UserField = ci.parseStringField(record[80:96])
	// 97-100
	ci.reserved = "    "
}

func (ci *CreditItem) UnmarshalJSON(data []byte) error {
	type Alias CreditItem
	aux := struct {
		*Alias
	}{
		(*Alias)(ci),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ci.setRecordType()
	return nil
}

// String writes the CreditItem struct to a variable length string.
func (ci *CreditItem) String() string {
	if ci == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(100)
	buf.WriteString(ci.recordType)
	buf.WriteString(ci.AuxiliaryOnUsField())
	buf.WriteString(ci.ExternalProcessingCodeField())
	buf.WriteString(ci.PostingBankRoutingNumberField())
	buf.WriteString(ci.OnUsField())
	buf.WriteString(ci.ItemAmountField())
	buf.WriteString(ci.CreditItemSequenceNumberField())
	buf.WriteString(ci.DocumentationTypeIndicatorField())
	buf.WriteString(ci.AccountTypeCodeField())
	buf.WriteString(ci.SourceWorkCodeField())
	buf.WriteString(ci.UserFieldField())
	buf.WriteString(ci.reservedField())
	return buf.String()
}

// Validate performs imagecashletter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (ci *CreditItem) Validate() error {
	if err := ci.fieldInclusion(); err != nil {
		return err
	}
	if ci.recordType != "62" {
		msg := fmt.Sprintf(msgRecordType, 62)
		return &FieldError{FieldName: "recordType", Value: ci.recordType, Msg: msg}
	}
	if ci.DocumentationTypeIndicator != "" {
		// Z is valid for CashLetter DocumentationTypeIndicator only
		if ci.DocumentationTypeIndicator == "Z" {
			msg := msgDocumentationTypeIndicator
			return &FieldError{FieldName: "DocumentationTypeIndicator", Value: ci.DocumentationTypeIndicator, Msg: msg}
		}
		// M is not valid for CreditItem DocumentationTypeIndicator
		if ci.DocumentationTypeIndicator == "M" {
			msg := msgDocumentationTypeIndicator
			return &FieldError{FieldName: "DocumentationTypeIndicator", Value: ci.DocumentationTypeIndicator, Msg: msg}
		}
		if err := ci.isDocumentationTypeIndicator(ci.DocumentationTypeIndicator); err != nil {
			return &FieldError{FieldName: "DocumentationTypeIndicator", Value: ci.DocumentationTypeIndicator, Msg: err.Error()}
		}
	}
	if err := ci.isAccountTypeCode(ci.AccountTypeCode); err != nil {
		return &FieldError{FieldName: "AccountTypeCode", Value: ci.AccountTypeCode, Msg: err.Error()}
	}
	if err := ci.isSourceWorkCode(ci.SourceWorkCode); err != nil {
		return &FieldError{FieldName: "SourceWorkCode", Value: ci.SourceWorkCode, Msg: err.Error()}
	}
	if err := ci.isAlphanumericSpecial(ci.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: ci.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (ci *CreditItem) fieldInclusion() error {
	if ci.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: ci.recordType,
			Msg:   msgFieldInclusion + ", did you use CreditItem()?"}
	}
	if ci.PostingBankRoutingNumber == "" {
		return &FieldError{FieldName: "PostingBankRoutingNumber",
			Value: ci.PostingBankRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use CreditItem()?"}
	}
	if ci.PostingBankRoutingNumberField() == "000000000" {
		return &FieldError{FieldName: "PostingBankRoutingNumber",
			Value: ci.PostingBankRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use CreditItem()?"}
	}
	if ci.CreditItemSequenceNumberField() == "               " {
		return &FieldError{FieldName: "CreditItemSequenceNumber",
			Value: ci.CreditItemSequenceNumber,
			Msg:   msgFieldInclusion + ", did you use CreditItem()?"}
	}
	return nil
}

// AuxiliaryOnUsField gets the AuxiliaryOnUs field
func (ci *CreditItem) AuxiliaryOnUsField() string {
	return ci.nbsmField(ci.AuxiliaryOnUs, 15)
}

// ExternalProcessingCodeField gets the ExternalProcessingCode field
func (ci *CreditItem) ExternalProcessingCodeField() string {
	return ci.alphaField(ci.ExternalProcessingCode, 1)
}

// PostingBankRoutingNumberField gets the PostingBankRoutingNumber field
func (ci *CreditItem) PostingBankRoutingNumberField() string {
	return ci.stringField(ci.PostingBankRoutingNumber, 9)
}

// OnUsField gets the OnUs field
func (ci *CreditItem) OnUsField() string {
	return ci.nbsmField(ci.OnUs, 20)
}

// ItemAmountField gets the temAmount field
func (ci *CreditItem) ItemAmountField() string {
	return ci.numericField(ci.ItemAmount, 14)
}

// CreditItemSequenceNumberField gets the CreditItemSequenceNumber field
func (ci *CreditItem) CreditItemSequenceNumberField() string {
	return ci.alphaField(ci.CreditItemSequenceNumber, 15)
}

// DocumentationTypeIndicatorField gets the DocumentationTypeIndicator field
func (ci *CreditItem) DocumentationTypeIndicatorField() string {
	return ci.alphaField(ci.DocumentationTypeIndicator, 1)
}

// AccountTypeCodeField gets the AccountTypeCode field
func (ci *CreditItem) AccountTypeCodeField() string {
	return ci.alphaField(ci.AccountTypeCode, 1)
}

// SourceWorkCodeField gets the SourceWorkCode field
func (ci *CreditItem) SourceWorkCodeField() string {
	return ci.alphaField(ci.SourceWorkCode, 2)
}

// UserFieldField gets the UserField field
func (ci *CreditItem) UserFieldField() string {
	return ci.alphaField(ci.UserField, 16)
}

// reservedField gets reserved - blank space
func (ci *CreditItem) reservedField() string {
	return ci.alphaField(ci.reserved, 4)
}
