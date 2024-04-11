// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Errors specific to a Credit Record

// Credit Record
type Credit struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// AuxiliaryOnUs identifies a code used on commercial checks at the discretion of the payor bank.
	AuxiliaryOnUs string `json:"auxiliaryOnUs,omitempty"`
	// ExternalProcessingCode identifies a code used for special purposes as authorized by the AS Committee X9.
	ExternalProcessingCode string `json:"externalProcessingCode,omitempty"`
	// PayorBankRoutingNumber identifies a number that identifies the institution by or through
	// which the item is payable.
	PayorBankRoutingNumber string `json:"payorBankRoutingNumber"`
	// CreditAccountNumberOnUs identifies data specified by the payor bank.
	// Usually an account number, serial number or transaction code or both.
	CreditAccountNumberOnUs string `json:"creditAccountNumberOnUs"`
	// ItemAmount identifies amount of the credit in U.S. dollars.
	ItemAmount int `json:"itemAmount"`
	// InstitutionItemSequenceNumber identifies sequence number assigned by the ECE company/institution.
	ECEInstitutionItemSequenceNumber string `json:"eceInstitutionItemSequenceNumber,omitempty"`
	// DocumentationTypeIndicator identifies a code that indicates the type of documentation
	// that supports the check record.
	DocumentationTypeIndicator string `json:"documentationTypeIndicator,omitempty"`
	// AccountTypeCode identifies a code to designate account type.
	AccountTypeCode string `json:"accountTypeCode,omitempty"`
	// SourceWorkCode identifies a code that identifies the incoming work.
	SourceWorkCode string `json:"sourceWorkCode,omitempty"`
	// WorkType identifies a code that identifies the type of work.
	WorkType string `json:"workType,omitempty"`
	// InstitutionItemSequenceNumber identifies a code that identifies whether this record represents
	// a debit or credit adjustment.
	DebitCreditIndicator string `json:"debitCreditIndicator,omitempty"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string

	ImageViewDetail []ImageViewDetail `json:"imageViewDetail"`
	ImageViewData   []ImageViewData   `json:"imageViewData"`

	// validator is composed for image cash letter data validation
	validator
	// converters is composed for image cash letter to golang Converters
	converters
}

// NewCredit returns a new credit with default values for non exported fields
func NewCredit() *Credit {
	cr := &Credit{}
	cr.setRecordType()
	return cr
}

func (cr *Credit) setRecordType() {
	if cr == nil {
		return
	}
	cr.recordType = "61"
	cr.reserved = "   "
}

// Parse takes the input record string and parses the BundleControl values
func (cr *Credit) Parse(record string) {
	if utf8.RuneCountInString(record) < 77 {
		return
	}

	// Character position 1-2, Always "61"
	cr.setRecordType()
	// 03–17
	cr.AuxiliaryOnUs = cr.parseStringField(record[2:17])
	// 18
	cr.ExternalProcessingCode = cr.parseStringField(record[17:18])
	// 19–27
	cr.PayorBankRoutingNumber = cr.parseStringField(record[18:27])
	// 28–47
	cr.CreditAccountNumberOnUs = cr.parseStringField(record[27:47])
	// 48–57
	cr.ItemAmount = cr.parseNumField(record[47:57])
	// 58–72
	cr.ECEInstitutionItemSequenceNumber = record[57:72]
	// 73
	cr.DocumentationTypeIndicator = cr.parseStringField(record[72:73])
	// 74
	cr.AccountTypeCode = cr.parseStringField(record[73:74])
	// 75
	cr.SourceWorkCode = cr.parseStringField(record[74:75])
	// 76
	cr.WorkType = cr.parseStringField(record[75:76])
	// 77
	cr.DebitCreditIndicator = cr.parseStringField(record[76:77])
	// 78–80
	cr.reserved = "   "

}

func (cr *Credit) UnmarshalJSON(data []byte) error {
	type Alias Credit
	aux := struct {
		*Alias
	}{
		(*Alias)(cr),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cr.setRecordType()
	return nil
}

// String writes the BundleControl struct to a string.
func (cr *Credit) String() string {
	if cr == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(cr.recordType)
	buf.WriteString(cr.AuxiliaryOnUsField())
	buf.WriteString(cr.ExternalProcessingCodeField())
	buf.WriteString(cr.PayorBankRoutingNumberField())
	buf.WriteString(cr.CreditAccountNumberOnUsField())
	buf.WriteString(cr.ItemAmountField())
	buf.WriteString(cr.ECEInstitutionItemSequenceNumberField())
	buf.WriteString(cr.DocumentationTypeIndicatorField())
	buf.WriteString(cr.AccountTypeCodeField())
	buf.WriteString(cr.SourceWorkCodeField())
	buf.WriteString(cr.WorkTypeField())
	buf.WriteString(cr.DebitCreditIndicatorField())
	buf.WriteString(cr.reservedField())
	return buf.String()
}

// Validate performs image cash letter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (cr *Credit) Validate() error {
	if err := cr.fieldInclusion(); err != nil {
		return err
	}
	if cr.recordType != "61" {
		msg := fmt.Sprintf(msgRecordType, 61)
		return &FieldError{FieldName: "recordType", Value: cr.recordType, Msg: msg}
	}
	if cr.SourceWorkCode != "" {
		if cr.SourceWorkCode != "3" {
			return &FieldError{FieldName: "SourceWorkCode", Value: cr.SourceWorkCode, Msg: msgInvalid}
		}
	}
	if cr.AccountTypeCode != "" {
		if err := cr.isAccountTypeCode(cr.AccountTypeCode); err != nil {
			return &FieldError{FieldName: "AccountTypeCode", Value: cr.AccountTypeCode, Msg: err.Error()}
		}
	}
	if cr.DocumentationTypeIndicator != "" {
		if cr.DocumentationTypeIndicator != "G" {
			return &FieldError{FieldName: "DocumentationTypeIndicator", Value: cr.DocumentationTypeIndicator, Msg: msgInvalid}
		}
	}
	if cr.ECEInstitutionItemSequenceNumber != "" {
		if err := cr.isNumeric(cr.ECEInstitutionItemSequenceNumber); err != nil {
			return &FieldError{FieldName: "ECEInstitutionItemSequenceNumber", Value: cr.ECEInstitutionItemSequenceNumber, Msg: msgInvalid}
		}
	}
	if err := cr.isNumeric(cr.PayorBankRoutingNumber); err != nil {
		return &FieldError{FieldName: "PayorBankRoutingNumber",
			Value: cr.PayorBankRoutingNumber, Msg: err.Error()}
	}
	// Should not contain forward or back slashes
	if strings.Contains(cr.AuxiliaryOnUs, `\`) || strings.Contains(cr.AuxiliaryOnUs, `/`) {
		return &FieldError{FieldName: "AuxiliaryOnUs", Value: cr.AuxiliaryOnUsField(), Msg: msgInvalid}
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (cr *Credit) fieldInclusion() error {
	if cr.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: cr.recordType,
			Msg:   msgFieldInclusion + ", did you use Credit()?"}
	}
	if cr.PayorBankRoutingNumberField() == "000000000" {
		return &FieldError{FieldName: "PayorBankRoutingNumber",
			Value: cr.PayorBankRoutingNumberField(),
			Msg:   msgFieldInclusion + ", did you use Credit()?"}
	}
	if cr.CreditAccountNumberOnUs == "" {
		return &FieldError{FieldName: "CreditAccountNumberOnUs",
			Value: cr.CreditAccountNumberOnUsField(),
			Msg:   msgFieldInclusion + ", did you use Credit()?"}
	}
	if cr.ItemAmount == 0 {
		return &FieldError{FieldName: "ItemAmount",
			Value: cr.ItemAmountField(),
			Msg:   msgFieldInclusion + ", did you use Credit()?"}
	}
	return nil
}

// AuxiliaryOnUsField gets a string of the AuxiliaryOnUs
func (cr *Credit) AuxiliaryOnUsField() string {
	return cr.alphaField(cr.AuxiliaryOnUs, 15)
}

// ExternalProcessingCodeField gets a string of the ExternalProcessingCode
func (cr *Credit) ExternalProcessingCodeField() string {
	return cr.alphaField(cr.ExternalProcessingCode, 1)
}

// PayorBankRoutingNumberField gets a string of the PayorBankRoutingNumber zero padded
func (cr *Credit) PayorBankRoutingNumberField() string {
	return cr.alphaField(cr.PayorBankRoutingNumber, 9)
}

// CreditAccountNumberOnUsField gets a string of the CreditAccountNumberOnUs
func (cr *Credit) CreditAccountNumberOnUsField() string {
	return cr.alphaField(cr.CreditAccountNumberOnUs, 20)
}

// ItemAmountField gets a string of the ItemAmount zero padded
func (cr *Credit) ItemAmountField() string {
	return cr.numericField(cr.ItemAmount, 10)
}

// ECEInstitutionItemSequenceNumberField gets a string of the ECEInstitutionItemSequenceNumber
func (cr *Credit) ECEInstitutionItemSequenceNumberField() string {
	return cr.alphaField(cr.ECEInstitutionItemSequenceNumber, 15)
}

// DocumentationTypeIndicatorField gets a string of the DocumentationTypeIndicator
func (cr *Credit) DocumentationTypeIndicatorField() string {
	return cr.alphaField(cr.DocumentationTypeIndicator, 1)
}

// AccountTypeCodeField gets a string of the AccountTypeCode
func (cr *Credit) AccountTypeCodeField() string {
	return cr.alphaField(cr.AccountTypeCode, 1)
}

// SourceWorkCodeField gets a string of the SourceOfWorkCode
func (cr *Credit) SourceWorkCodeField() string {
	return cr.alphaField(cr.SourceWorkCode, 1)
}

// WorkTypeField gets a string of the WorkType
func (cr *Credit) WorkTypeField() string {
	return cr.alphaField(cr.WorkType, 1)
}

// DebitCreditIndicatorField gets a string of the DebitCreditIndicator
func (cr *Credit) DebitCreditIndicatorField() string {
	return cr.alphaField(cr.DebitCreditIndicator, 1)
}

// reservedField gets reserved - blank space
func (cr *Credit) reservedField() string {
	return cr.alphaField(cr.reserved, 3)
}

func (cr *Credit) AddImageViewData(ivData ImageViewData) {
	cr.ImageViewData = append(cr.ImageViewData, ivData)
}

func (cr *Credit) AddImageViewDetail(ivDetail ImageViewDetail) {
	cr.ImageViewDetail = append(cr.ImageViewDetail, ivDetail)
}
