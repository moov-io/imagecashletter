// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// mockUserPayeeEndorsement creates a UserPayeeEndorsement
func mockUserPayeeEndorsement() *UserPayeeEndorsement {
	upe := NewUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 3
	upe.OwnerIdentifier = "230918276"
	upe.OwnerIdentifierModifier = "ZZ1"
	upe.UserRecordFormatType = "001"
	upe.FormatTypeVersionLevel = "1"
	upe.LengthUserData = "0000290"
	upe.PayeeName = "Payee Name"

	upe.EndorsementDate = time.Now()
	upe.BankRoutingNumber = "121042882"
	upe.BankAccountNumber = "123456888"
	upe.CustomerIdentifier = "A234A"
	upe.CustomerContactInformation = "Home"
	upe.StoreMerchantProcessingSiteNumber = "12345678"
	upe.InternalControlSequenceNumber = "ZB17262ZB"
	upe.Time = time.Now()
	upe.OperatorName = "ZJK"
	upe.OperatorNumber = "12345"
	upe.ManagerName = "ZBK"
	upe.ManagerNumber = "12345"
	upe.EquipmentNumber = "123456789012345"
	upe.EndorsementIndicator = 1
	upe.UserField = ""
	return upe
}

func TestUserPayeeEndorsementParseErr(t *testing.T) {
	var upe UserPayeeEndorsement
	upe.Parse("asjsahfakja")
	require.Equal(t, 0, upe.OwnerIdentifierIndicator)
}

// TestMockUserPayeeEndorsement creates a UserPayeeEndorsement
func TestMockUserPayeeEndorsement(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	require.NoError(t, upe.Validate())
	require.Equal(t, "68", upe.recordType)
	require.Equal(t, 3, upe.OwnerIdentifierIndicator)
	require.Equal(t, "230918276", upe.OwnerIdentifier)
	require.Equal(t, "ZZ1", upe.OwnerIdentifierModifier)
	require.Equal(t, "001", upe.UserRecordFormatType)
	require.Equal(t, "1", upe.FormatTypeVersionLevel)
	require.Equal(t, "0000290", upe.LengthUserData)
	require.Equal(t, "Payee Name", upe.PayeeName)

	_ = additionalUPEFields(upe, t)
}

func additionalUPEFields(upe *UserPayeeEndorsement, t *testing.T) string {
	require.Equal(t, "121042882", upe.BankRoutingNumber)
	require.Equal(t, "123456888", upe.BankAccountNumber)
	require.Equal(t, "A234A", upe.CustomerIdentifier)
	require.Equal(t, "Home", upe.CustomerContactInformation)
	require.Equal(t, "12345678", upe.StoreMerchantProcessingSiteNumber)
	require.Equal(t, "ZB17262ZB", upe.InternalControlSequenceNumber)
	require.Equal(t, "ZJK", upe.OperatorName)
	require.Equal(t, "12345", upe.OperatorNumber)
	require.Equal(t, "ZBK", upe.ManagerName)
	require.Equal(t, "12345", upe.ManagerNumber)
	require.Equal(t, "123456789012345", upe.EquipmentNumber)
	require.Equal(t, 1, upe.EndorsementIndicator)
	require.Equal(t, "", upe.UserField)
	return ""
}

// TestUPEString validation
func TestUPEString(t *testing.T) {
	line := "683230918276ZZ1                 0011  0000290Payee Name                                        20181015121042882123456888           A234A               Home                                              12345678ZB17262ZB                2222ZJK                           12345ZBK                           123451234567890123451          "
	upe := NewUserPayeeEndorsement()
	r := NewReader(strings.NewReader(line))
	r.line = line
	upe.Parse(r.line)

	require.Equal(t, line, upe.String())
}

// TestUPEParse validation
func TestUPEParse(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	require.NoError(t, upe.Validate())
	line := upe.String()
	r := NewReader(strings.NewReader(line))
	r.line = line

	upe.Parse(r.line)
	require.NoError(t, upe.Validate())
}

// TestUPERecordType validation
func TestUPERecordType(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.recordType = "00"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestUPEUserRecordFormatType validation
func TestUPEUserRecordFormatType(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.UserRecordFormatType = "002"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserRecordFormatType", e.FieldName)
}

// TestUPEOwnerIdentifierIndicator validation
func TestUPEOwnerIdentifierIndicator(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 9
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifierIndicator", e.FieldName)
}

// TestUPEOwnerIdentifierModifier validation
func TestUPEOwnerIdentifierModifier(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierModifier = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifierModifier", e.FieldName)
}

// TestUPEUserRecordFormatTypeChar validation
func TestUPEUserRecordFormatTypeChar(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.UserRecordFormatType = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserRecordFormatType", e.FieldName)
}

// TestUPEFormatTypeVersionLevel validation
func TestUPEFormatTypeVersionLevel(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.FormatTypeVersionLevel = "W"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FormatTypeVersionLevel", e.FieldName)
}

// TestUPELengthUserData validation
func TestUPELengthUserData(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.LengthUserData = "W"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "LengthUserData", e.FieldName)
}

// TestUPEOwnerIdentifierIndicatorZero validation
func TestUPEOwnerIdentifierIndicatorZero(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 0
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifier", e.FieldName)
}

// TestUPEOwnerIdentifierIndicator123 validation
func TestUPEOwnerIdentifierIndicator123(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 1
	upe.OwnerIdentifier = "W"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifier", e.FieldName)
}

// TestUPEOwnerIdentifierIndicatorFour validation
func TestUPEOwnerIdentifierIndicatorFour(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 4
	upe.OwnerIdentifier = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifier", e.FieldName)
}

// TestUPEPayeeName validation
func TestUPEPayeeName(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.PayeeName = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayeeName", e.FieldName)
}

// TestUPEBankRoutingNumber validation
func TestUPEBankRoutingNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.BankRoutingNumber = "W"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BankRoutingNumber", e.FieldName)
}

// TestUPEBankAccountNumber validation
func TestUPEBankAccountNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.BankAccountNumber = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BankAccountNumber", e.FieldName)
}

// TestUPECustomerIdentifier validation
func TestUPECustomerIdentifier(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.CustomerIdentifier = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CustomerIdentifier", e.FieldName)
}

// TestUPECustomerContactInformation validation
func TestUPECustomerContactInformation(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.CustomerContactInformation = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CustomerContactInformation", e.FieldName)
}

// TestUPEStoreMerchantProcessingSiteNumber validation
func TestUPEStoreMerchantProcessingSiteNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.StoreMerchantProcessingSiteNumber = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "StoreMerchantProcessingSiteNumber", e.FieldName)
}

// TestUPEInternalControlSequenceNumber validation
func TestUPEInternalControlSequenceNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.InternalControlSequenceNumber = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "InternalControlSequenceNumber", e.FieldName)
}

// TestUPEOperatorName validation
func TestUPEOperatorName(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OperatorName = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OperatorName", e.FieldName)
}

// TestUPEOperatorNumber validation
func TestUPEOperatorNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OperatorNumber = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OperatorNumber", e.FieldName)
}

// TestUPEManagerName validation
func TestUPEManagerName(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.ManagerName = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ManagerName", e.FieldName)
}

// TestUPEManagerNumber validation
func TestUPEManagerNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.ManagerNumber = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ManagerNumber", e.FieldName)
}

// TestUPEEquipmentNumber validation
func TestUPEEquipmentNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.EquipmentNumber = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EquipmentNumber", e.FieldName)
}

// TestUPEEndorsementIndicator validation
func TestUPEEndorsementIndicator(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.EndorsementIndicator = 7
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsementIndicator", e.FieldName)
}

// TestUPEUserField validation
func TestUPEUserField(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.UserField = "®©"
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestUPEFIRecordType validation
func TestUPEFIRecordType(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.recordType = ""
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestUPEFIUserRecordFormatType validation
func TestUPEFIUserRecordFormatType(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.UserRecordFormatType = ""
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserRecordFormatType", e.FieldName)
}

// TestUPEFIFormatTypeVersionLevel validation
func TestUPEFIFormatTypeVersionLevel(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.FormatTypeVersionLevel = ""
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FormatTypeVersionLevel", e.FieldName)
}

// TestUPEFILengthUserData validation
func TestUPEFILengthUserData(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.LengthUserData = ""
	err := upe.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "LengthUserData", e.FieldName)
}

// TestUPERuneCountInString validates RuneCountInString
func TestUPERuneCountInString(t *testing.T) {
	upe := NewUserPayeeEndorsement()
	var line = "68"
	upe.Parse(line)

	require.Equal(t, "", upe.BankAccountNumber)
}
