// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockUserGeneral creates a UserGeneral
func mockUserGeneral() *UserGeneral {
	ug := NewUserGeneral()
	ug.OwnerIdentifierIndicator = 3
	ug.OwnerIdentifier = "230918276"
	ug.OwnerIdentifierModifier = "ZZ1"
	ug.UserRecordFormatType = "000"
	ug.FormatTypeVersionLevel = "1"
	ug.LengthUserData = "0000038"
	ug.UserData = "This is a payment for your information"
	return ug
}

func TestUserGeneralParseErr(t *testing.T) {
	var ug UserGeneral
	ug.Parse("askdhfaskjas")
	require.Equal(t, 0, ug.OwnerIdentifierIndicator)
}

// TestMockUserGeneral creates a UserGeneral
func TestMockUserGeneral(t *testing.T) {
	ug := mockUserGeneral()
	require.NoError(t, ug.Validate())
	require.Equal(t, "68", ug.recordType)
	require.Equal(t, 3, ug.OwnerIdentifierIndicator)
	require.Equal(t, "230918276", ug.OwnerIdentifier)
	require.Equal(t, "ZZ1", ug.OwnerIdentifierModifier)
	require.Equal(t, "000", ug.UserRecordFormatType)
	require.Equal(t, "1", ug.FormatTypeVersionLevel)
	require.Equal(t, "0000038", ug.LengthUserData)
	require.Equal(t, "This is a payment for your information", ug.UserData)
}

// TestUGString validation
func TestUGString(t *testing.T) {
	line := "683230918276ZZ1                 0001  0000038This is a payment for your information"
	ug := NewUserGeneral()
	r := NewReader(strings.NewReader(line))
	r.line = line
	ug.Parse(r.line)

	require.Equal(t, line, ug.String())
}

// TestUGParse validation
func TestUGParse(t *testing.T) {
	ug := mockUserGeneral()
	require.NoError(t, ug.Validate())
	line := ug.String()
	r := NewReader(strings.NewReader(line))
	r.line = line
	ug.Parse(r.line)

	require.NoError(t, ug.Validate())
}

// TestUGRecordType validation
func TestUGRecordType(t *testing.T) {
	ug := mockUserGeneral()
	ug.recordType = "00"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestUGUserRecordFormatType validation
func TestUGUserRecordFormatType(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserRecordFormatType = "001"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserRecordFormatType", e.FieldName)
}

// TestUGOwnerIdentifierIndicator validation
func TestUGOwnerIdentifierIndicator(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierIndicator = 9
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifierIndicator", e.FieldName)
}

// TestUGOwnerIdentifierModifier validation
func TestUGOwnerIdentifierModifier(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierModifier = "®©"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifierModifier", e.FieldName)
}

// TestUGUserRecordFormatTypeChar validation
func TestUGUserRecordFormatTypeChar(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserRecordFormatType = "®©"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserRecordFormatType", e.FieldName)
}

// TestUGFormatTypeVersionLevel validation
func TestUGFormatTypeVersionLevel(t *testing.T) {
	ug := mockUserGeneral()
	ug.FormatTypeVersionLevel = "W"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FormatTypeVersionLevel", e.FieldName)
}

// TestUGLengthUserData validation
func TestUGLengthUserData(t *testing.T) {
	ug := mockUserGeneral()
	ug.LengthUserData = "W"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "LengthUserData", e.FieldName)
}

// TestUGUserData validation
func TestUGUserData(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserData = "A®©X"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserData", e.FieldName)
}

// TestUGOwnerIdentifierIndicatorZero validation
func TestUGOwnerIdentifierIndicatorZero(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierIndicator = 0
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifier", e.FieldName)
}

// TestUGOwnerIdentifierIndicator123 validation
func TestUGOwnerIdentifierIndicator123(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierIndicator = 1
	ug.OwnerIdentifier = "W"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifier", e.FieldName)
}

// TestUGOwnerIdentifierIndicatorFour validation
func TestUGOwnerIdentifierIndicatorFour(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierIndicator = 4
	ug.OwnerIdentifier = "®©"
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OwnerIdentifier", e.FieldName)
}

// Field Inclusion

// TestUGFIRecordType validation
func TestUGFIRecordType(t *testing.T) {
	ug := mockUserGeneral()
	ug.recordType = ""
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestUGFIUserRecordFormatType validation
func TestUGFIUserRecordFormatType(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserRecordFormatType = ""
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserRecordFormatType", e.FieldName)
}

// TestUGFIFormatTypeVersionLevel validation
func TestUGFIFormatTypeVersionLevel(t *testing.T) {
	ug := mockUserGeneral()
	ug.FormatTypeVersionLevel = ""
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FormatTypeVersionLevel", e.FieldName)
}

// TestUGFILengthUserData validation
func TestUGFILengthUserData(t *testing.T) {
	ug := mockUserGeneral()
	ug.LengthUserData = ""
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "LengthUserData", e.FieldName)
}

// TestUGFIUserData validation
func TestUGFIUserData(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserData = ""
	err := ug.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserData", e.FieldName)
}
