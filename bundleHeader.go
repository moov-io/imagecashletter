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

// Errors specific to a BundleHeader Record

// BundleHeader Record
type BundleHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// recordType defines the type of record.
	recordType string
	// A code that identifies the type of bundle. It is the same value as the CollectionTypeIndicator
	// in the CashLetterHeader within which the bundle is contained, unless the CollectionTypeIndicator
	// in the CashLetterHeader is 99.
	// Values:
	// 00: Preliminary Forward Information–Used when information may change and the information is treated
	// as not final.
	// 01: Forward Presentment–For the collection and settlement of checks (demand instruments).
	// Data are treated as final.
	// 02: Forward Presentment–Same-Day Settlement–For the collection and settlement of checks (demand instruments)
	// presented under the Federal Reserve’s same day settlement amendments to Regulation CC (12CFR Part 229).
	// Data are treated as final.
	// 03: Return–For the return of check(s). Transaction carries value. Data are treated as final.
	// 04: Return Notification–For the notification of return of check(s). Transaction carries no value. The Return
	// Notification Indicator (Field 12) in the Return Record (Type 31) has to be interrogated to determine whether a
	// notice is a preliminary or final notification.
	// 05: Preliminary Return Notification–For the notification of return of check(s). Transaction carries no value.
	// Used to indicate that an item may be returned. This field supersedes the Return Notification Indicator
	// (Field 12) in the Return Record (Type 31).
	// 06: Final Return Notification–For the notification of return of check(s). Transaction carries no value. Used to
	// indicate that an item will be returned. This field supersedes the Return Notification Indicator (Field 12)
	// in the Return Record (Type 31).
	CollectionTypeIndicator string `json:"collectionTypeIndicator"`
	// DestinationRoutingNumber contains the routing and transit number of the institution that
	// receives and processes the cash letter or the bundle.  Format: TTTTAAAAC, where:
	//  TTTT Federal Reserve Prefix
	//  AAAA ABA Institution Identifier
	//  C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	DestinationRoutingNumber string `json:"destinationRoutingNumber"`
	// ECEInstitutionRoutingNumber contains the routing and transit number of the institution
	// that creates the bundle header.  Format: TTTTAAAAC, where:
	//	TTTT Federal Reserve Prefix
	//	AAAA ABA Institution Identifier
	//	C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	ECEInstitutionRoutingNumber string `json:"eceInstitutionRoutingNumber"`
	// BundleBusinessDate is the business date of the bundle.
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	BundleBusinessDate time.Time `json:"bundleBusinessDate"`
	// BundleCreationDate is the date that the bundle is created. It is Eastern Time zone format unless
	// different clearing arrangements have been made
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	BundleCreationDate time.Time `json:"bundleCreationDate"`
	// BundleID is number that identifies the bundle, assigned by the institution that creates the bundle.
	BundleID string `json:"bundleID"`
	// BundleSequenceNumber is a number assigned by the institution that creates the bundle. Usually denotes
	// the relative position of the bundle within the cash letter.  NumericBlank
	BundleSequenceNumber string `json:"BundleSequenceNumber,omitempty"`
	// CycleNumber is a code assigned by the institution that creates the bundle.  Denotes the cycle under which
	// the bundle is created.
	CycleNumber string `json:"cycleNumber"`
	// ReturnLocationRoutingNumber is a bank routing number used by some processors.
	// This will be blank in the resulting file if it is empty.
	ReturnLocationRoutingNumber string `json:"returnLocationRoutingNumber,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for imagecashletter data validation
	validator
	// converters is composed for imagecashletter to golang Converters
	converters
}

// NewBundleHeader returns a new BundleHeader with default values for non exported fields
func NewBundleHeader() *BundleHeader {
	bh := &BundleHeader{}
	bh.setRecordType()
	return bh
}

func (bh *BundleHeader) setRecordType() {
	if bh == nil {
		return
	}
	bh.recordType = "20"
}

// Parse takes the input record string and parses the BundleHeader values
func (bh *BundleHeader) Parse(record string) {
	if utf8.RuneCountInString(record) < 68 {
		return
	}

	// Character position 1-2, Always "20"
	bh.setRecordType()
	// 03-04
	bh.CollectionTypeIndicator = record[2:4]
	// 05-13
	bh.DestinationRoutingNumber = bh.parseStringField(record[4:13])
	// 14-22
	bh.ECEInstitutionRoutingNumber = bh.parseStringField(record[13:22])
	// 23-30
	bh.BundleBusinessDate = bh.parseYYYYMMDDDate(record[22:30])
	// 31-38
	bh.BundleCreationDate = bh.parseYYYYMMDDDate(record[30:38])
	// 39-48
	bh.BundleID = bh.parseStringField(record[38:48])
	// 49-52
	bh.BundleSequenceNumber = bh.parseStringField(record[48:52])
	// 53-54
	bh.CycleNumber = bh.parseStringField(record[52:54])
	// 55-63
	bh.ReturnLocationRoutingNumber = bh.parseStringField(record[54:63])
	// 64-68
	bh.UserField = bh.parseStringField(record[63:68])
	// 69-80
	bh.reserved = "            "
}

func (bh *BundleHeader) UnmarshalJSON(data []byte) error {
	type Alias BundleHeader
	aux := struct {
		*Alias
	}{
		(*Alias)(bh),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	bh.setRecordType()
	return nil
}

// String writes the BundleHeader struct to a string.
func (bh *BundleHeader) String() string {
	if bh == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(bh.recordType)
	buf.WriteString(bh.CollectionTypeIndicatorField())
	buf.WriteString(bh.DestinationRoutingNumberField())
	buf.WriteString(bh.ECEInstitutionRoutingNumberField())
	buf.WriteString(bh.BundleBusinessDateField())
	buf.WriteString(bh.BundleCreationDateField())
	buf.WriteString(bh.BundleIDField())
	buf.WriteString(bh.BundleSequenceNumberField())
	buf.WriteString(bh.CycleNumberField())
	buf.WriteString(bh.ReturnLocationRoutingNumberField())
	buf.WriteString(bh.UserFieldField())
	buf.WriteString(bh.reservedField())
	return buf.String()
}

// Validate performs imagecashletter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (bh *BundleHeader) Validate() error {
	if bh == nil {
		return nil
	}
	if err := bh.fieldInclusion(); err != nil {
		return err
	}
	if bh.recordType != "20" {
		msg := fmt.Sprintf(msgRecordType, 20)
		return &FieldError{FieldName: "recordType", Value: bh.recordType, Msg: msg}
	}
	// Mandatory
	if err := bh.isCollectionTypeIndicator(bh.CollectionTypeIndicator); err != nil {
		return &FieldError{FieldName: "CollectionTypeIndicator",
			Value: bh.CollectionTypeIndicator, Msg: err.Error()}
	}
	if err := bh.isAlphanumeric(bh.BundleID); err != nil {
		return &FieldError{FieldName: "BundleID", Value: bh.BundleID, Msg: err.Error()}
	}
	if err := bh.isAlphanumeric(bh.CycleNumber); err != nil {
		return &FieldError{FieldName: "CycleNumber", Value: bh.CycleNumber, Msg: err.Error()}
	}
	if err := bh.isAlphanumericSpecial(bh.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: bh.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (bh *BundleHeader) fieldInclusion() error {
	if bh == nil {
		return nil
	}
	if bh.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: bh.recordType,
			Msg:   msgFieldInclusion + ", did you use BundleHeader()?"}
	}
	if bh.CollectionTypeIndicator == "" {
		return &FieldError{FieldName: "CollectionTypeIndicator",
			Value: bh.CollectionTypeIndicator,
			Msg:   msgFieldInclusion + ", did you use BundleHeader()?"}
	}
	if bh.DestinationRoutingNumber == "" {
		return &FieldError{FieldName: "DestinationRoutingNumber",
			Value: bh.DestinationRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use BundleHeader()?"}
	}
	if bh.DestinationRoutingNumberField() == "000000000" {
		return &FieldError{FieldName: "DestinationRoutingNumber",
			Value: bh.DestinationRoutingNumber, Msg: msgFieldInclusion}
	}
	if bh.ECEInstitutionRoutingNumber == "" {
		return &FieldError{FieldName: "ECEInstitutionRoutingNumber",
			Value: bh.ECEInstitutionRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use BundleHeader()?"}
	}
	if bh.ECEInstitutionRoutingNumberField() == "000000000" {
		return &FieldError{FieldName: "ECEInstitutionRoutingNumber",
			Value: bh.ECEInstitutionRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use BundleHeader()?"}
	}
	if bh.BundleBusinessDate.IsZero() {
		return &FieldError{FieldName: "BundleBusinessDate",
			Value: bh.BundleBusinessDate.String(),
			Msg:   msgFieldInclusion + ", did you use BundleHeader()?"}
	}
	if bh.BundleCreationDate.IsZero() {
		return &FieldError{FieldName: "BundleCreationDate",
			Value: bh.BundleCreationDate.String(),
			Msg:   msgFieldInclusion + ", did you use BundleHeader()?"}
	}
	if bh.BundleSequenceNumberField() == "    " {
		return &FieldError{FieldName: "BundleSequenceNumber",
			Value: bh.BundleSequenceNumber,
			Msg:   msgFieldInclusion + ", did you use BundleHeader()?"}
	}
	return nil
}

// CollectionTypeIndicatorField gets the CollectionTypeIndicator field
func (bh *BundleHeader) CollectionTypeIndicatorField() string {
	return bh.alphaField(bh.CollectionTypeIndicator, 2)
}

// DestinationRoutingNumberField gets the DestinationRoutingNumber field
func (bh *BundleHeader) DestinationRoutingNumberField() string {
	return bh.stringField(bh.DestinationRoutingNumber, 9)
}

// ECEInstitutionRoutingNumberField gets the ECEInstitutionRoutingNumber field
func (bh *BundleHeader) ECEInstitutionRoutingNumberField() string {
	return bh.stringField(bh.ECEInstitutionRoutingNumber, 9)
}

// BundleBusinessDateField gets the BundleBusinessDate in YYYYMMDD format
func (bh *BundleHeader) BundleBusinessDateField() string {
	return bh.formatYYYYMMDDDate(bh.BundleBusinessDate)
}

// BundleCreationDateField gets the BundleCreationDate in YYYYMMDD format
func (bh *BundleHeader) BundleCreationDateField() string {
	return bh.formatYYYYMMDDDate(bh.BundleCreationDate)
}

// BundleIDField gets the BundleID field space padded
func (bh *BundleHeader) BundleIDField() string {
	return bh.alphaField(bh.BundleID, 10)
}

// BundleSequenceNumberField gets the BundleSequenceNumber field zero padded
func (bh *BundleHeader) BundleSequenceNumberField() string {
	return bh.alphaField(bh.BundleSequenceNumber, 4)
}

// CycleNumberField gets the CycleNumber field
func (bh *BundleHeader) CycleNumberField() string {
	return bh.alphaField(bh.CycleNumber, 2)
}

// ReturnLocationRoutingNumberField gets the ReturnLocationRoutingNumber field
// or spaces if blank
func (bh *BundleHeader) ReturnLocationRoutingNumberField() string {
	return bh.alphaField(bh.ReturnLocationRoutingNumber, 9)
}

// UserFieldField gets the UserField field
func (bh *BundleHeader) UserFieldField() string {
	return bh.alphaField(bh.UserField, 5)
}

// reservedField returns blank spaces for paddding
func (bh *BundleHeader) reservedField() string {
	return bh.alphaField(bh.reserved, 12)
}

// SetBundleSequenceNumber sets BundleSequenceNumber
func (bh *BundleHeader) SetBundleSequenceNumber(seq int) string {
	bh.BundleSequenceNumber = bh.numericField(seq, 4)
	return bh.BundleSequenceNumber
}
