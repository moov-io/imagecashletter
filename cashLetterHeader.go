// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"strings"
	"time"
)

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a CashLetterHeader Record

// CashLetterHeader Record is mandatory.
type CashLetterHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record
	// Value: 10
	recordType string
	// CollectionTypeIndicator is a code that identifies the type of cash letter.
	// Values:
	// 00: Preliminary Forward Information–Used when information may change and the
	// information is treated as not final.
	// 01: Forward Presentment–For the collection and settlement of checks (demand
	// instruments). Data are treated as final.
	// 02: Forward Presentment–Same-Day Settlement–For the collection and settlement of
	// checks (demand instruments) presented under the Federal Reserve’s same day
	// settlement amendments to Regulation CC (12CFR Part 229). Data are treated as
	// final.
	// 03: Return–For the return of check(s). Transaction carries value. Data are
	// treated as final.
	// 04: Return Notification–For the notification of return of check(s). Transaction
	// carries no value. The Return Notification Indicator (Field 12) in the Return Record
	// (Type 31) has to be interrogated to determine whether a notice is a preliminary or final
	// notification.
	// 05: Preliminary Return Notification–For the notification of return of check(s). Transaction
	// carries no value. Used to indicate that an item may be returned. This field supersedes
	// the Return Notification Indicator (Field 12) in the Return Record (Type 31).
	// 06: Final Return Notification–For the notification of return of check(s). Transaction
	// carries no value. Used to indicate that an item will be returned. This field
	// supersedes the Return Notification Indicator (Field 12) in the Return Record (Type 31).
	// 20: No Detail–There are no detail records contained within the bundle or cash letter.
	// Defined Value of the Cash Letter Record Type Indicator (Field 8) shall be set to ‘N’.
	// 99: Bundles not the same collection type. Use of the value is only allowed by clearing
	// arrangement.
	CollectionTypeIndicator string `json:"collectionTypeIndicator"`
	// DestinationRoutingNumber contains the routing and transit number of the institution that
	// receives and processes the cash letter or the bundle.  Format: TTTTAAAAC, where:
	//  TTTT Federal Reserve Prefix
	//  AAAA ABA Institution Identifier
	//  C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	DestinationRoutingNumber string `json:"destinationRoutingNumber"`
	// ECEInstitutionRoutingNumber contains the routing and transit number of the institution that
	// that creates the Cash Letter Header Record.  Format: TTTTAAAAC, where:
	//  TTTT Federal Reserve Prefix
	//  AAAA ABA Institution Identifier
	//  C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	ECEInstitutionRoutingNumber string `json:"eceInstitutionRoutingNumber"`
	// CashLetterBusinessDate is the business date of the cash letter.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	CashLetterBusinessDate time.Time `json:"cashLetterBusinessDate"`
	// CashLetterCreationDate is the date that the cash letter is created which shall be in Eastern
	// Time zone format. Other time zones may be used under clearing arrangements.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	CashLetterCreationDate time.Time `json:"cashLetterCreationDate"`
	// CashLetterCreationTime is the time that the cash letter is created.  Default time shall be in
	// Eastern Time zone format. Other time zones may be used under clearing arrangements.
	// Format: hhmm, where: hh hour, mm minute
	// Values:
	// hh '00' through '23'
	// mm '00' through '59'
	CashLetterCreationTime time.Time `json:"cashLetterCreationTime"`
	// CashLetterRecordTypeIndicator is a code that indicates the presence of records or the type of records contained
	// in the cash letter.   If an image is associated with any CheckDetail or Return, the cash letter must have a
	// CashLetter.RecordTypeIndicator of I or F.
	// Values:
	// N: No electronic check records or image records (Type 2x’s, 3x’s, 5x’s); e.g., an empty cash letter.
	// E: Cash letter contains electronic check records with no images (Type 2x’s and 3x’s only).
	// I: Cash letter contains electronic check records (Type 2x’s, 3x’s) and image records (Type 5x’s).
	// F: Cash letter contains electronic check records (Type 2x’s and 3x’s) and image records (Type 5x’s)
	// that correspond to a previously sent cash letter (i.e., E file).
	//
	// The fields in this file that contain posting data shall not be changed from the previously sent CashLetter
	// with CollectionTypeIndicator values of 01, 02 or 03. ItemsCount and TotalAmount of the CashLetterControl with
	// a RecordTypeIndicator value of F must equal the corresponding fields in a CashLetter with a RecordTypeIndicator
	// value of E.
	CashLetterRecordTypeIndicator string `json:"cashLetterRecordTypeIndicator"`
	// CashLetterDocumentationTypeIndicator is a code that indicates the type of documentation that supports
	// all check records in the cash letter
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
	// M: No image provided, Electronic Check provided separately
	// Z: Not Same Type–Documentation associated with each item in Cash Letter will be different. The Check Detail
	// Record (Type 25) or Return Record (Type 31) has to be interrogated for further information.
	CashLetterDocumentationTypeIndicator string `json:"cashLetterDocumentationTypeIndicator"`
	// CashLetterID uniquely identifies the cash letter. It is assigned by the institution that creates the cash
	// letter and must be unique within a Cash Letter Business Date.
	CashLetterID string `json:"cashLetterID"`
	// OriginatorContactName is the name of contact at the institution that creates the cash letter.
	OriginatorContactName string `json:"originatorContactName"`
	// OriginatorContactPhoneNumber is the phone number of the contact at the institution that creates
	// the cash letter.
	OriginatorContactPhoneNumber string `json:"originatorContactPhoneNumber"`
	// FedWorkType is any valid codes specified by the Federal Reserve Bank.
	FedWorkType string `json:"fedWorkType"`
	// ReturnsIndicator identifies type pf returns.
	// Values:
	// "": Blank for Forward Presentment
	// E: Administrative - items being returned that are handled by the bank and usually do not directly
	// affect the customer or its account.
	// R: Customer–items being returned that directly affect a customer’s account.
	// J: Reject Return
	ReturnsIndicator string `json:"returnsIndicator"`
	// UserField is a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewCashLetterHeader returns a new CashLetterHeader with default values for non exported fields
func NewCashLetterHeader() *CashLetterHeader {
	clh := &CashLetterHeader{
		recordType: "10",
	}
	return clh
}

// Parse takes the input record string and parses the CashLetterHeader values
func (clh *CashLetterHeader) Parse(record string) {
	// Character position 1-2, Always "10"
	clh.recordType = "10"
	// 03-04
	clh.CollectionTypeIndicator = clh.parseStringField(record[2:4])
	// 05-13
	clh.DestinationRoutingNumber = clh.parseStringField(record[4:13])
	// 14-22
	clh.ECEInstitutionRoutingNumber = clh.parseStringField(record[13:22])
	// 23-30
	clh.CashLetterBusinessDate = clh.parseYYYYMMDDDate(record[22:30])
	// 31-38
	clh.CashLetterCreationDate = clh.parseYYYYMMDDDate(record[30:38])
	// 39-42
	clh.CashLetterCreationTime = clh.parseSimpleTime(record[38:42])
	// 43-43
	clh.CashLetterRecordTypeIndicator = clh.parseStringField(record[42:43])
	// 44-44
	clh.CashLetterDocumentationTypeIndicator = clh.parseStringField(record[43:44])
	// 45-52
	clh.CashLetterID = clh.parseStringField(record[44:52])
	// 53-66
	clh.OriginatorContactName = clh.parseStringField(record[52:66])
	// 67-76
	clh.OriginatorContactPhoneNumber = clh.parseStringField(record[66:76])
	// 77-77
	clh.FedWorkType = clh.parseStringField(record[76:77])
	// 78-78
	clh.ReturnsIndicator = clh.parseStringField(record[77:78])
	// 79-79
	clh.UserField = clh.parseStringField(record[78:79])
	// 80-80
	clh.reserved = " "
}

// String writes the CashLetterHeader struct to a string.
func (clh *CashLetterHeader) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(clh.recordType)
	buf.WriteString(clh.CollectionTypeIndicatorField())
	buf.WriteString(clh.DestinationRoutingNumberField())
	buf.WriteString(clh.ECEInstitutionRoutingNumberField())
	buf.WriteString(clh.CashLetterBusinessDateField())
	buf.WriteString(clh.CashLetterCreationDateField())
	buf.WriteString(clh.CashLetterCreationTimeField())
	buf.WriteString(clh.CashLetterRecordTypeIndicatorField())
	buf.WriteString(clh.CashLetterDocumentationTypeIndicatorField())
	buf.WriteString(clh.CashLetterIDField())
	buf.WriteString(clh.OriginatorContactNameField())
	buf.WriteString(clh.OriginatorContactPhoneNumberField())
	buf.WriteString(clh.FedWorkTypeField())
	buf.WriteString(clh.ReturnsIndicatorField())
	buf.WriteString(clh.UserFieldField())
	buf.WriteString(clh.reservedField())
	return buf.String()
}

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (clh *CashLetterHeader) Validate() error {
	if err := clh.fieldInclusion(); err != nil {
		return err
	}
	if clh.recordType != "10" {
		msg := fmt.Sprintf(msgRecordType, 10)
		return &FieldError{FieldName: "recordType", Value: clh.recordType, Msg: msg}
	}
	if err := clh.isCollectionTypeIndicator(clh.CollectionTypeIndicator); err != nil {
		return &FieldError{FieldName: "CollectionTypeIndicator",
			Value: clh.CollectionTypeIndicator, Msg: err.Error()}
	}
	if err := clh.isCashLetterRecordTypeIndicator(clh.CashLetterRecordTypeIndicator); err != nil {
		return &FieldError{FieldName: "CashLetterRecordTypeIndicator",
			Value: clh.CashLetterRecordTypeIndicator, Msg: err.Error()}
	}
	if err := clh.isDocumentationTypeIndicator(clh.CashLetterDocumentationTypeIndicator); err != nil {
		return &FieldError{FieldName: "CashLetterDocumentationTypeIndicator",
			Value: clh.CashLetterDocumentationTypeIndicator, Msg: err.Error()}
	}
	if err := clh.isAlphanumeric(clh.CashLetterID); err != nil {
		return &FieldError{FieldName: "CashLetterID", Value: clh.CashLetterID, Msg: err.Error()}
	}
	if err := clh.isAlphanumericSpecial(clh.OriginatorContactName); err != nil {
		return &FieldError{FieldName: "OriginatorContactName", Value: clh.OriginatorContactName, Msg: err.Error()}
	}
	if err := clh.isNumeric(clh.OriginatorContactPhoneNumber); err != nil {
		return &FieldError{FieldName: "OriginatorContactPhoneNumber", Value: clh.OriginatorContactPhoneNumber, Msg: err.Error()}
	}
	if err := clh.isAlphanumeric(clh.FedWorkType); err != nil {
		return &FieldError{FieldName: "FedWorkType", Value: clh.FedWorkType, Msg: err.Error()}
	}
	if err := clh.isReturnsIndicator(clh.ReturnsIndicator); err != nil {
		return &FieldError{FieldName: "ReturnsIndicator", Value: clh.ReturnsIndicator, Msg: err.Error()}
	}
	if err := clh.isAlphanumericSpecial(clh.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: clh.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (clh *CashLetterHeader) fieldInclusion() error {
	if clh.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: clh.recordType, Msg: msgFieldInclusion}
	}
	if clh.CollectionTypeIndicator == "" {
		return &FieldError{FieldName: "CollectionTypeIndicator",
			Value: clh.CollectionTypeIndicator, Msg: msgFieldInclusion}
	}
	if clh.DestinationRoutingNumber == "" {
		return &FieldError{FieldName: "DestinationRoutingNumber",
			Value: clh.DestinationRoutingNumber, Msg: msgFieldInclusion}
	}
	if clh.ECEInstitutionRoutingNumber == "" {
		return &FieldError{FieldName: "ECEInstitutionRoutingNumber",
			Value: clh.ECEInstitutionRoutingNumber, Msg: msgFieldInclusion}
	}
	if clh.CashLetterBusinessDate.IsZero() {
		return &FieldError{FieldName: "CashLetterBusinessDate",
			Value: clh.CashLetterBusinessDate.String(), Msg: msgFieldInclusion}
	}
	if clh.CashLetterCreationDate.IsZero() {
		return &FieldError{FieldName: "CashLetterCreationDate",
			Value: clh.CashLetterCreationDate.String(), Msg: msgFieldInclusion}
	}
	if clh.CashLetterCreationTime.IsZero() {
		return &FieldError{FieldName: "CashLetterCreationTime",
			Value: clh.CashLetterCreationTime.String(), Msg: msgFieldInclusion}
	}
	if clh.CashLetterRecordTypeIndicator == "" {
		return &FieldError{FieldName: "CashLetterRecordTypeIndicator",
			Value: clh.CashLetterRecordTypeIndicator, Msg: msgFieldInclusion}
	}
	if clh.CashLetterID == "" {
		return &FieldError{FieldName: "CashLetterID", Value: clh.CashLetterID, Msg: msgFieldInclusion}
	}
	return nil
}

// CollectionTypeIndicatorField gets the CollectionTypeIndicator field
func (clh *CashLetterHeader) CollectionTypeIndicatorField() string {
	return clh.stringField(clh.CollectionTypeIndicator, 2)
}

// DestinationRoutingNumberField gets the DestinationRoutingNumber field
func (clh *CashLetterHeader) DestinationRoutingNumberField() string {
	return clh.stringField(clh.DestinationRoutingNumber, 9)
}

// ECEInstitutionRoutingNumberField gets the ECEInstitutionRoutingNumber field
func (clh *CashLetterHeader) ECEInstitutionRoutingNumberField() string {
	return clh.stringField(clh.ECEInstitutionRoutingNumber, 9)
}

// CashLetterBusinessDateField gets the CashLetterBusinessDate in YYYYMMDD format
func (clh *CashLetterHeader) CashLetterBusinessDateField() string {
	return clh.formatYYYYMMDDDate(clh.CashLetterBusinessDate)
}

// CashLetterCreationDateField gets the CashLetterCreationDate in YYYYMMDD format
func (clh *CashLetterHeader) CashLetterCreationDateField() string {
	return clh.formatYYYYMMDDDate(clh.CashLetterCreationDate)
}

// CashLetterCreationTimeField gets the CashLetterCreationTime in HHMM format
func (clh *CashLetterHeader) CashLetterCreationTimeField() string {
	return clh.formatSimpleTime(clh.CashLetterCreationTime)
}

// CashLetterRecordTypeIndicatorField gets the CashLetterRecordTypeIndicator field
func (clh *CashLetterHeader) CashLetterRecordTypeIndicatorField() string {
	return clh.alphaField(clh.CashLetterRecordTypeIndicator, 1)
}

// CashLetterDocumentationTypeIndicatorField gets the CashLetterDocumentationTypeIndicator field
func (clh *CashLetterHeader) CashLetterDocumentationTypeIndicatorField() string {
	return clh.alphaField(clh.CashLetterDocumentationTypeIndicator, 1)
}

// CashLetterIDField gets the CashLetterID field
func (clh *CashLetterHeader) CashLetterIDField() string {
	return clh.alphaField(clh.CashLetterID, 8)
}

// OriginatorContactNameField gets the OriginatorContactName field
func (clh *CashLetterHeader) OriginatorContactNameField() string {
	return clh.alphaField(clh.OriginatorContactName, 14)
}

// OriginatorContactPhoneNumberField gets the OriginatorContactPhoneNumber field
func (clh *CashLetterHeader) OriginatorContactPhoneNumberField() string {
	return clh.alphaField(clh.OriginatorContactPhoneNumber, 10)
}

// FedWorkTypeField gets the FedWorkType field
func (clh *CashLetterHeader) FedWorkTypeField() string {
	return clh.alphaField(clh.FedWorkType, 1)
}

// ReturnsIndicatorField gets the ReturnsIndicator field
func (clh *CashLetterHeader) ReturnsIndicatorField() string {
	return clh.alphaField(clh.ReturnsIndicator, 1)
}

// UserFieldField gets the UserField field
func (clh *CashLetterHeader) UserFieldField() string {
	return clh.alphaField(clh.UserField, 1)
}

// reservedField gets reserved - blank space
func (clh *CashLetterHeader) reservedField() string {
	return clh.alphaField(clh.reserved, 1)
}
