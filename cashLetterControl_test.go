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

// mockCashLetterControl creates a CashLetterControl
func mockCashLetterControl() *CashLetterControl {
	clc := NewCashLetterControl()
	clc.CashLetterBundleCount = 1
	clc.CashLetterItemsCount = 7
	clc.CashLetterTotalAmount = 100000 // 1000.00
	clc.CashLetterImagesCount = 1
	clc.ECEInstitutionName = "Wells Fargo"
	clc.SettlementDate = time.Now()
	clc.CreditTotalIndicator = 0
	return clc
}

// TestMockCashLetterControl creates a CashLetterControl
func TestMockCashLetterControl(t *testing.T) {
	clc := mockCashLetterControl()
	require.NoError(t, clc.Validate())
	require.Equal(t, "90", clc.recordType)
	require.Equal(t, 1, clc.CashLetterBundleCount)
	require.Equal(t, 7, clc.CashLetterItemsCount)
	require.Equal(t, 100000, clc.CashLetterTotalAmount)
	require.Equal(t, 1, clc.CashLetterImagesCount)
	require.Equal(t, "Wells Fargo", clc.ECEInstitutionName)
	require.Equal(t, 0, clc.CreditTotalIndicator)
}

// TestParseCashLetterControl parses a known CashLetterControl record string
func TestParseCashLetterControl(t *testing.T) {
	var line = "900000010000000100000000100000000000000Wells Fargo       201809050              "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	require.NoError(t, r.parseCashLetterControl())
	record := r.currentCashLetter.CashLetterControl

	require.Equal(t, "90", record.recordType)
	require.Equal(t, "000001", record.CashLetterBundleCountField())
	require.Equal(t, "00000001", record.CashLetterItemsCountField())
	require.Equal(t, "00000000100000", record.CashLetterTotalAmountField())
	require.Equal(t, "000000000", record.CashLetterImagesCountField())
	require.Equal(t, "Wells Fargo       ", record.ECEInstitutionNameField())
	require.Equal(t, "20180905", record.SettlementDateField())
	require.Equal(t, "0", record.CreditTotalIndicatorField())
	require.Equal(t, "              ", record.reservedField())
}

// testCLCString validates that a known parsed CashLetterControl can be return to a string of the same value
func testCLCString(t testing.TB) {
	var line = "900000010000000100000000100000000000000Wells Fargo       201809050              "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	require.NoError(t, r.parseCashLetterControl())
	record := r.currentCashLetter.CashLetterControl
	require.Equal(t, line, record.String())
}

// TestCLCString tests validating that a known parsed CashLetterControl can be return to a string of the same value
func TestCLCString(t *testing.T) {
	testCLCString(t)
}

// BenchmarkCLCString benchmarks validating that a known parsed CashLetterControl can be return to a string of the same value
func BenchmarkCLCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCLCString(b)
	}
}

// TestCLCRecordType validation
func TestCLCRecordType(t *testing.T) {
	clc := mockCashLetterControl()
	clc.recordType = "00"
	err := clc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "recordType", fieldErr.FieldName)
}

// TestECEInstitutionName validation
func TestECEInstitutionName(t *testing.T) {
	clc := mockCashLetterControl()
	clc.ECEInstitutionName = "®©"
	err := clc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "ECEInstitutionName", fieldErr.FieldName)
}

// TestCLCCreditTotalIndicator validation
func TestCLCCreditTotalIndicator(t *testing.T) {
	clc := mockCashLetterControl()
	clc.CreditTotalIndicator = 9
	err := clc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "CreditTotalIndicator", fieldErr.FieldName)
}

// TestCLCFieldInclusionRecordType validates FieldInclusion
func TestCLCFieldInclusionRecordType(t *testing.T) {
	clc := mockCashLetterControl()
	clc.recordType = ""
	err := clc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "recordType", fieldErr.FieldName)
}

// TestFieldInclusionCashLetterItemsCount validates FieldInclusion
func TestFieldInclusionCashLetterItemsCount(t *testing.T) {
	clc := mockCashLetterControl()
	clc.CashLetterItemsCount = 0
	err := clc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "CashLetterItemsCount", fieldErr.FieldName)
}

// TestFieldInclusionCashLetterTotalAmount validates FieldInclusion
func TestFieldInclusionCashLetterTotalAmount(t *testing.T) {
	clc := mockCashLetterControl()
	clc.CashLetterTotalAmount = 0
	err := clc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "CashLetterTotalAmount", fieldErr.FieldName)
}

// TestFieldInclusionSettlementDate validates FieldInclusion
func TestFieldInclusionRecordTypeSettlementDate(t *testing.T) {
	clc := mockCashLetterControl()
	// if present (non-zero), SettlementDate.Year() must be between 1993 and 9999
	clc.SettlementDate = time.Date(40010, time.November, 9, 0, 0, 0, 0, time.UTC)
	err := clc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "SettlementDate", fieldErr.FieldName)
}

// TestCashLetterControlRuneCountInString validates RuneCountInString
func TestCashLetterControlRuneCountInString(t *testing.T) {
	clc := NewCashLetterControl()
	var line = "90"
	clc.Parse(line)

	require.Equal(t, 0, clc.CashLetterBundleCount)
}

func TestCashLetterControl_isReturnCollectionType(t *testing.T) {
	tests := []struct {
		collectionType string
		expected       bool
	}{
		{"03", true},
		{"04", true},
		{"05", true},
		{"06", true},
		{"07", false},
		{"01", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.collectionType, func(t *testing.T) {
			require.Equal(t, tt.expected, isReturnCollectionType(tt.collectionType))
		})
	}
}
