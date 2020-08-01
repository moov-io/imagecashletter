// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// RoutingNumberSummary Record
type RoutingNumberSummary struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// CashLetterRoutingNumber
	CashLetterRoutingNumber string `json:"cashLetterRoutingNumber"`
	// RoutingNumberTotalAmount
	RoutingNumberTotalAmount int `json:"routingNumberTotalAmount"`
	// RoutingNumberItemCount
	RoutingNumberItemCount int `json:"routingNumberItemCount"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for imagecashletter data validation
	validator
	// converters is composed for imagecashletter to golang Converters
	converters
}

// NewRoutingNumberSummary returns a new RoutingNumberSummary with default values for non exported fields
func NewRoutingNumberSummary() *RoutingNumberSummary {
	rns := &RoutingNumberSummary{}
	rns.setRecordType()
	return rns
}

func (rns *RoutingNumberSummary) setRecordType() {
	if rns == nil {
		return
	}
	rns.recordType = "85"
}

// Parse takes the input record string and parses the ImageViewDetail values
func (rns *RoutingNumberSummary) Parse(record string) {
	if utf8.RuneCountInString(record) < 55 {
		return // line too short
	}

	// Character position 1-2, Always "85"
	rns.setRecordType()
	// 03-11
	rns.CashLetterRoutingNumber = rns.parseStringField(record[2:11])
	// 12-25
	rns.RoutingNumberTotalAmount = rns.parseNumField(record[11:25])
	// 26-31
	rns.RoutingNumberItemCount = rns.parseNumField(record[26:31])
	// 32-55
	rns.UserField = rns.parseStringField(record[31:55])
	// 56-80
	rns.reserved = "                         "
}

// String writes the ImageViewDetail struct to a string.
func (rns *RoutingNumberSummary) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rns.recordType)
	buf.WriteString(rns.CashLetterRoutingNumberField())
	buf.WriteString(rns.RoutingNumberTotalAmountField())
	buf.WriteString(rns.RoutingNumberItemCountField())
	buf.WriteString(rns.UserFieldField())
	buf.WriteString(rns.reservedField())
	return buf.String()
}

// Validate performs imagecashletter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rns *RoutingNumberSummary) Validate() error {
	if err := rns.fieldInclusion(); err != nil {
		return err
	}
	if rns.recordType != "85" {
		msg := fmt.Sprintf(msgRecordType, 85)
		return &FieldError{FieldName: "recordType", Value: rns.recordType, Msg: msg}
	}
	if err := rns.isAlphanumericSpecial(rns.UserField); err != nil {
		return &FieldError{FieldName: "UserField",
			Value: rns.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rns *RoutingNumberSummary) fieldInclusion() error {
	if rns.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: rns.recordType,
			Msg:   msgFieldInclusion + ", did you use RoutingNumberSummary()?"}
	}
	if rns.CashLetterRoutingNumber == "" {
		return &FieldError{FieldName: "CashLetterRoutingNumber",
			Value: rns.CashLetterRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use RoutingNumberSummary()?"}
	}
	return nil
}

// CashLetterRoutingNumberField gets the CashLetterRoutingNumber routing number field
func (rns *RoutingNumberSummary) CashLetterRoutingNumberField() string {
	return rns.stringField(rns.CashLetterRoutingNumber, 9)
}

// RoutingNumberTotalAmountField gets a string of RoutingNumberTotalAmount zero padded
func (rns *RoutingNumberSummary) RoutingNumberTotalAmountField() string {
	return rns.numericField(rns.RoutingNumberTotalAmount, 14)
}

// RoutingNumberItemCountField gets a string of RoutingNumberItemCount zero padded
func (rns *RoutingNumberSummary) RoutingNumberItemCountField() string {
	return rns.numericField(rns.RoutingNumberItemCount, 6)
}

// UserFieldField gets the UserField field
func (rns *RoutingNumberSummary) UserFieldField() string {
	return rns.alphaField(rns.UserField, 24)
}

// reservedField gets the reserved field
func (rns *RoutingNumberSummary) reservedField() string {
	return rns.alphaField(rns.reserved, 25)
}
