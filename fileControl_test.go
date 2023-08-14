// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFileControl creates a FileControl
func mockFileControl() FileControl {
	fc := NewFileControl()
	fc.CashLetterCount = 1
	fc.TotalRecordCount = 7
	fc.TotalItemCount = 1
	fc.FileTotalAmount = 100000 // 1000.00
	fc.ImmediateOriginContactName = "Contact Name"
	fc.ImmediateOriginContactPhoneNumber = "5558675552"
	fc.CreditTotalIndicator = 0
	return fc
}

// TestMockFileControl creates a FileControl
func TestMockFileControl(t *testing.T) {
	fc := mockFileControl()
	require.NoError(t, fc.Validate())
	require.Equal(t, "99", fc.recordType)
	require.Equal(t, 1, fc.CashLetterCount)
	require.Equal(t, 7, fc.TotalRecordCount)
	require.Equal(t, 1, fc.TotalItemCount)
	require.Equal(t, 100000, fc.FileTotalAmount)
	require.Equal(t, "Contact Name", fc.ImmediateOriginContactName)
	require.Equal(t, "5558675552", fc.ImmediateOriginContactPhoneNumber)
	require.Equal(t, 0, fc.CreditTotalIndicator)
}

// testParseFileControl parses a known FileControl record string
func testParseFileControl(t testing.TB) {
	var line = "9900000100000007000000010000000000100000Contact Name  55586755520               "
	r := NewReader(strings.NewReader(line))
	r.line = line
	require.NoError(t, r.parseFileControl())
	record := r.File.Control

	require.Equal(t, "99", record.recordType)
	require.Equal(t, "000001", record.CashLetterCountField())
	require.Equal(t, "00000007", record.TotalRecordCountField())
	require.Equal(t, "00000001", record.TotalItemCountField())
	require.Equal(t, "0000000000100000", record.FileTotalAmountField())
	require.Equal(t, "Contact Name  ", record.ImmediateOriginContactNameField())
	require.Equal(t, "5558675552", record.ImmediateOriginContactPhoneNumberField())
	require.Equal(t, "0", record.CreditTotalIndicatorField())
	require.Equal(t, "               ", record.reservedField())
}

// TestParseFileControl tests parsing a known FileControl record string
func TestParseFileControl(t *testing.T) {
	testParseFileControl(t)
}

// BenchmarkParseFileControl benchmarks parsing a known FileControl record string
func BenchmarkParseFileControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseFileControl(b)
	}
}

// testFCString validates that a known parsed FileControl can be return to a string of the same value
func testFCString(t testing.TB) {
	var line = "9900000100000007000000010000000000100000Contact Name  55586755520               "
	r := NewReader(strings.NewReader(line))
	r.line = line
	require.NoError(t, r.parseFileControl())
	record := r.File.Control
	require.Equal(t, line, record.String())
}

// TestFCString tests validating that a known parsed FileControl can be return to a string of the same value
func TestFCString(t *testing.T) {
	testFCString(t)
}

// BenchmarkFCString benchmarks validating that a known parsed FileControl can be return to a string of the same value
func BenchmarkFCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCString(b)
	}
}

// TestFCRecordType validation
func TestFCRecordType(t *testing.T) {
	fc := mockFileControl()
	fc.recordType = "00"
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestImmediateOriginContactName validation
func TestImmediateOriginContactName(t *testing.T) {
	fc := mockFileControl()
	fc.ImmediateOriginContactName = "®©"
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImmediateOriginContactName", e.FieldName)
}

// TestImmediateOriginContactPhoneNumber validation
func TestImmediateOriginContactPhoneNumber(t *testing.T) {
	fc := mockFileControl()
	fc.ImmediateOriginContactPhoneNumber = "--"
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImmediateOriginContactPhoneNumber", e.FieldName)
}

// TestCreditTotalIndicator validation
func TestCreditTotalIndicator(t *testing.T) {
	fc := mockFileControl()
	fc.CreditTotalIndicator = 9
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CreditTotalIndicator", e.FieldName)
}

// TestFCFieldInclusionRecordType validates FieldInclusion
func TestFCFieldInclusionRecordType(t *testing.T) {
	fc := mockFileControl()
	fc.recordType = ""
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestFieldInclusionCashLetterCount validates FieldInclusion
func TestFieldInclusionCashLetterCount(t *testing.T) {
	fc := mockFileControl()
	fc.CashLetterCount = 0
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CashLetterCount", e.FieldName)
}

// TestFieldInclusionTotalRecordCount validates FieldInclusion
func TestFieldInclusionTotalRecordCount(t *testing.T) {
	fc := mockFileControl()
	fc.TotalRecordCount = 0
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TotalRecordCount", e.FieldName)
}

// TestFieldInclusionTotalItemCount validates FieldInclusion
func TestFieldInclusionTotalItemCount(t *testing.T) {
	fc := mockFileControl()
	fc.TotalItemCount = 0
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TotalItemCount", e.FieldName)
}

// TestFieldInclusionFileTotalAmount validates FieldInclusion
func TestFieldInclusionFileTotalAmount(t *testing.T) {
	fc := mockFileControl()
	fc.FileTotalAmount = 0
	err := fc.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FileTotalAmount", e.FieldName)
}

// TestFileControlRuneCountInString validates RuneCountInString
func TestFileControlRuneCountInString(t *testing.T) {
	fc := NewFileControl()
	var line = "99"
	fc.Parse(line)

	require.Equal(t, 0, fc.CashLetterCount)
}
