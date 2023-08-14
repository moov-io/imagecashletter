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

// mockBundleHeader creates a BundleHeader
func mockBundleHeader() *BundleHeader {
	bh := NewBundleHeader()
	bh.CollectionTypeIndicator = "01"
	bh.DestinationRoutingNumber = "231380104"
	bh.ECEInstitutionRoutingNumber = "121042882"
	bh.BundleBusinessDate = time.Now()
	bh.BundleCreationDate = time.Now()
	bh.BundleID = "9999"
	bh.BundleSequenceNumber = "1"
	bh.CycleNumber = "01"
	bh.UserField = ""
	return bh
}

// testMockBundleHeader creates a BundleHeader
func testMockBundleHeader(t testing.TB) {
	bh := mockBundleHeader()
	require.NoError(t, bh.Validate())
	require.Equal(t, "20", bh.recordType)
	require.Equal(t, "01", bh.CollectionTypeIndicator)
	require.Equal(t, "231380104", bh.DestinationRoutingNumber)
	require.Equal(t, "121042882", bh.ECEInstitutionRoutingNumber)
	require.Equal(t, "9999", bh.BundleID)
	require.Equal(t, "1", bh.BundleSequenceNumber)
	require.Equal(t, "01", bh.CycleNumber)
	require.Equal(t, "", bh.UserField)
}

// TestMockBundleHeader tests creating a BundleHeader
func TestMockBundleHeader(t *testing.T) {
	testMockBundleHeader(t)
}

// BenchmarkMockBundleHeader benchmarks creating a BundleHeader
func BenchmarkMockBundleHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockBundleHeader(b)
	}
}

// parseBundleHeader validates parsing a BundleHeader
func parseBundleHeader(t testing.TB) {
	var line = "200123138010412104288220180905201809059999      1   01                          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))

	require.NoError(t, r.parseBundleHeader())

	record := r.currentCashLetter.currentBundle.BundleHeader
	require.Equal(t, "20", record.recordType)
	require.Equal(t, "01", record.CollectionTypeIndicatorField())
	require.Equal(t, "231380104", record.DestinationRoutingNumberField())
	require.Equal(t, "121042882", record.ECEInstitutionRoutingNumberField())
	require.Equal(t, "20180905", record.BundleBusinessDateField())
	require.Equal(t, "20180905", record.BundleCreationDateField())
	require.Equal(t, "9999      ", record.BundleIDField())
	require.Equal(t, "1   ", record.BundleSequenceNumberField())
	require.Equal(t, "01", record.CycleNumberField())
	require.Equal(t, "         ", record.ReturnLocationRoutingNumberField())
	require.Equal(t, "     ", record.UserFieldField())
	require.Equal(t, "            ", record.reservedField())
}

// TestParseBundleHeader tests validating parsing a BundleHeader
func TestParseBundleHeader(t *testing.T) {
	parseBundleHeader(t)
}

// BenchmarkParseBundleHeader benchmarks validating parsing a BundleHeader
func BenchmarkParseBundleHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseBundleHeader(b)
	}
}

// testBHString validates that a known parsed BundleHeader can return to a string of the same value
func testBHString(t testing.TB) {
	var line = "200123138010412104288220180905201809059999      1   01                          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	require.NoError(t, r.parseBundleHeader())
	record := r.currentCashLetter.currentBundle.BundleHeader
	require.Equal(t, line, record.String())
}

// TestBHString tests validating that a known parsed BundleHeader can return to a string of the same value
func TestBHString(t *testing.T) {
	testBHString(t)
}

// BenchmarkBHString benchmarks validating that a known parsed BundleHeader
// can return to a string of the same value
func BenchmarkBHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHString(b)
	}
}

// TestBHRecordType validation
func TestBHRecordType(t *testing.T) {
	bh := mockBundleHeader()
	bh.recordType = "00"
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestBHCollectionTypeIndicator validation
func TestBHCollectionTypeIndicator(t *testing.T) {
	bh := mockBundleHeader()
	bh.CollectionTypeIndicator = "87"
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CollectionTypeIndicator", e.FieldName)
}

// TestBundleID validation
func TestBundleID(t *testing.T) {
	bh := mockBundleHeader()
	bh.BundleID = "--"
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BundleID", e.FieldName)
}

// TestCycleNumber validation
func TestCycleNumber(t *testing.T) {
	bh := mockBundleHeader()
	bh.CycleNumber = "--"
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CycleNumber", e.FieldName)
}

// TestBHUserField validation
func TestBHUserFieldI(t *testing.T) {
	bh := mockBundleHeader()
	bh.UserField = "®©"
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// TestBHFieldInclusionRecordType validates FieldInclusion
func TestBHFieldInclusionRecordType(t *testing.T) {
	bh := mockBundleHeader()
	bh.recordType = ""
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestBHFieldInclusionCollectionTypeIndicator validates FieldInclusion
func TestBHFieldInclusionCollectionTypeIndicator(t *testing.T) {
	bh := mockBundleHeader()
	bh.CollectionTypeIndicator = ""
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CollectionTypeIndicator", e.FieldName)
}

// TestBHFieldInclusionDestinationRoutingNumber validates FieldInclusion
func TestBHFieldInclusionDestinationRoutingNumber(t *testing.T) {
	bh := mockBundleHeader()
	bh.DestinationRoutingNumber = ""
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DestinationRoutingNumber", e.FieldName)
}

// TestBHFieldInclusionDestinationRoutingNumberZero validates FieldInclusion
func TestBHFieldInclusionDestinationRoutingNumberZero(t *testing.T) {
	bh := mockBundleHeader()
	bh.DestinationRoutingNumber = "000000000"
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DestinationRoutingNumber", e.FieldName)
}

// TestBHFieldInclusionECEInstitutionRoutingNumber validates FieldInclusion
func TestBHFieldInclusionECEInstitutionRoutingNumber(t *testing.T) {
	bh := mockBundleHeader()
	bh.ECEInstitutionRoutingNumber = ""
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ECEInstitutionRoutingNumber", e.FieldName)
}

// TestBHFieldInclusionECEInstitutionRoutingNumberZero validates FieldInclusion
func TestBHFieldInclusionECEInstitutionRoutingNumberZero(t *testing.T) {
	bh := mockBundleHeader()
	bh.ECEInstitutionRoutingNumber = "000000000"
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ECEInstitutionRoutingNumber", e.FieldName)
}

// TestBHFieldInclusionBundleBusinessDate validates FieldInclusion
func TestBHFieldInclusionBundleBusinessDate(t *testing.T) {
	bh := mockBundleHeader()
	bh.BundleBusinessDate = time.Time{}
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BundleBusinessDate", e.FieldName)
}

// TestBHFieldInclusionBundleCreationDate validates FieldInclusion
func TestBHFieldInclusionBundleCreationDate(t *testing.T) {
	bh := mockBundleHeader()
	bh.BundleCreationDate = time.Time{}
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BundleCreationDate", e.FieldName)
}

// TestBHFieldInclusionBundleSequenceNumber validates FieldInclusion
func TestBHFieldInclusionBundleSequenceNumber(t *testing.T) {
	bh := mockBundleHeader()
	bh.BundleSequenceNumber = "    "
	err := bh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BundleSequenceNumber", e.FieldName)
}

// TestBundleHeaderRuneCountInString validates RuneCountInString
func TestBundleHeaderRuneCountInString(t *testing.T) {
	bh := NewBundleHeader()
	var line = "20"
	bh.Parse(line)

	require.Equal(t, "", bh.CycleNumber)
}
