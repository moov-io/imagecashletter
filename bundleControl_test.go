// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBundleControl creates a BundleControl
func mockBundleControl() *BundleControl {
	bc := NewBundleControl()
	bc.BundleItemsCount = 7
	bc.BundleTotalAmount = 100000    // 1000.00
	bc.MICRValidTotalAmount = 100000 // 1000.00
	bc.BundleImagesCount = 1
	bc.UserField = ""
	bc.CreditTotalIndicator = 0
	return bc
}

// TestMockBundleControl creates an BundleControl
func TestMockBundleControl(t *testing.T) {
	bc := mockBundleControl()
	require.NoError(t, bc.Validate())
	require.Equal(t, "70", bc.recordType)
	require.Equal(t, 7, bc.BundleItemsCount)
	require.Equal(t, 100000, bc.BundleTotalAmount)
	require.Equal(t, 100000, bc.MICRValidTotalAmount)
	require.Equal(t, 1, bc.BundleImagesCount)
	require.Empty(t, bc.UserField)
	require.Equal(t, 0, bc.CreditTotalIndicator)
}

// TestParseBundleControl parses a known BundleControl record string
func TestParseBundleControl(t *testing.T) {
	var line = "70000100000010000000000000000000000                    0                        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	require.NoError(t, r.parseBundleControl())
	record := r.currentCashLetter.currentBundle.BundleControl

	require.Equal(t, "70", record.recordType)
	require.Equal(t, "0001", record.BundleItemsCountField())
	require.Equal(t, "000000100000", record.BundleTotalAmountField())
	require.Equal(t, "000000000000", record.MICRValidTotalAmountField())
	require.Equal(t, "00000", record.BundleImagesCountField())
	require.Equal(t, "                    ", record.UserFieldField())
	require.Equal(t, "0", record.CreditTotalIndicatorField())
	require.Equal(t, "                        ", record.reservedField())
}

func TestParseBundleControlError(t *testing.T) {
	bc := NewBundleControl()
	bc.Parse(" invalid line ")

	require.Equal(t, 0, bc.BundleItemsCount)
	require.Equal(t, 0, bc.BundleTotalAmount)
	require.Equal(t, 0, bc.MICRValidTotalAmount)
	require.Equal(t, 0, bc.BundleImagesCount)
	require.Empty(t, bc.UserField)
	require.Equal(t, 0, bc.CreditTotalIndicator)
}

// testBCString validates that a known parsed BundleControl can be return to a string of the same value
func testBCString(t testing.TB) {
	var line = "70000100000010000000000000000000000                    0                        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	require.NoError(t, r.parseBundleControl())
	record := r.currentCashLetter.currentBundle.BundleControl
	require.Equal(t, line, record.String())
}

// TestBCString tests validating that a known parsed BundleControl can be return to a string of the same value
func TestBCString(t *testing.T) {
	testBCString(t)
}

// BenchmarkBCString benchmarks validating that a known parsed BundleControl can be return to a string of the same value
func BenchmarkBCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCString(b)
	}
}

// TestBCRecordType validation
func TestBCRecordType(t *testing.T) {
	bc := mockBundleControl()
	bc.recordType = "00"
	err := bc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "recordType", fieldErr.FieldName)
}

// TestBCUserField validation
func TestBCUserFieldI(t *testing.T) {
	bc := mockBundleControl()
	bc.UserField = "®©"
	err := bc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "UserField", fieldErr.FieldName)
}

// TestBCCreditTotalIndicator validation
func TestBCCreditTotalIndicator(t *testing.T) {
	bc := mockBundleControl()
	bc.CreditTotalIndicator = 9
	err := bc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "CreditTotalIndicator", fieldErr.FieldName)
}

// TestBCFieldInclusionRecordType validates FieldInclusion
func TestBCFieldInclusionRecordType(t *testing.T) {
	bc := mockBundleControl()
	bc.recordType = ""
	err := bc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "recordType", fieldErr.FieldName)
}

// TestFieldInclusionBundleItemsCount validates FieldInclusion
func TestFieldInclusionBundleItemsCount(t *testing.T) {
	bc := mockBundleControl()
	bc.BundleItemsCount = 0
	err := bc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "BundleItemsCount", fieldErr.FieldName)
}

// TestFieldInclusionBundleTotalAmount validates FieldInclusion
func TestFieldInclusionBundleTotalAmount(t *testing.T) {
	bc := mockBundleControl()
	bc.BundleTotalAmount = 0
	err := bc.Validate()
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)
	require.Equal(t, "BundleTotalAmount", fieldErr.FieldName)
}

// TestBundleControlRuneCountInString validates RuneCountInString
func TestBundleControlRuneCountInString(t *testing.T) {
	bc := NewBundleControl()
	var line = "70"
	bc.Parse(line)

	require.Equal(t, 0, bc.BundleItemsCount)
}
