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

// mockFileHeader creates a FileHeader
func mockFileHeader() FileHeader {
	fh := NewFileHeader()
	fh.StandardLevel = "35"
	fh.TestFileIndicator = "T"
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now()
	fh.FileCreationTime = time.Now()
	fh.ResendIndicator = "N"
	fh.ImmediateDestinationName = "Citadel"
	fh.ImmediateOriginName = "Wells Fargo"
	fh.FileIDModifier = ""
	fh.CountryCode = "US"
	fh.UserField = ""
	fh.CompanionDocumentIndicator = ""
	return fh
}

// testMockFileHeader creates a FileHeader
func testMockFileHeader(t testing.TB) {
	fh := mockFileHeader()
	require.NoError(t, fh.Validate())
	require.Equal(t, "01", fh.recordType)
	require.Equal(t, "35", fh.StandardLevel)
	require.Equal(t, "T", fh.TestFileIndicator)
	require.Equal(t, "N", fh.ResendIndicator)
	require.Equal(t, "231380104", fh.ImmediateDestination)
	require.Equal(t, "121042882", fh.ImmediateOrigin)
	require.Equal(t, "Citadel", fh.ImmediateDestinationName)
	require.Equal(t, "Wells Fargo", fh.ImmediateOriginName)
	require.Equal(t, "", fh.FileIDModifier)
	require.Equal(t, "US", fh.CountryCode)
	require.Equal(t, "", fh.UserField)
	require.Equal(t, "", fh.CompanionDocumentIndicator)
}

// TestMockFileHeader tests creating a FileHeader
func TestMockFileHeader(t *testing.T) {
	testMockFileHeader(t)
}

// BenchmarkMockFileHeader benchmarks creating a FileHeader
func BenchmarkMockFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockFileHeader(b)
	}
}

// parseFileHeader validates parsing a FileHeader
func parseFileHeader(t testing.TB) {
	var line = "0135T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	require.NoError(t, r.parseFileHeader())
	record := r.File.Header

	require.Equal(t, "01", record.recordType)
	require.Equal(t, "35", record.StandardLevelField())
	require.Equal(t, "T", record.TestFileIndicatorField())
	require.Equal(t, "231380104", record.ImmediateDestinationField())
	require.Equal(t, "121042882", record.ImmediateOriginField())
	require.Equal(t, "20180905", record.FileCreationDateField())
	require.Equal(t, "1523", record.FileCreationTimeField())
	require.Equal(t, "N", record.ResendIndicatorField())
	require.Equal(t, "Citadel           ", record.ImmediateDestinationNameField())
	require.Equal(t, "Wells Fargo       ", record.ImmediateOriginNameField())
	require.Equal(t, " ", record.FileIDModifierField())
	require.Equal(t, "US", record.CountryCodeField())
	require.Equal(t, "    ", record.UserFieldField())
	require.Equal(t, " ", record.CompanionDocumentIndicatorField())
}

// TestParseFileHeader tests validating parsing a FileHeader
func TestParseFileHeader(t *testing.T) {
	parseFileHeader(t)
}

// BenchmarkParseFileHeader benchmarks validating parsing a FileHeader
func BenchmarkParseFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseFileHeader(b)
	}
}

// testFHString validates that a known parsed FileHeader can return to a string of the same value
func testFHString(t testing.TB) {
	var line = "0135T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	require.NoError(t, r.parseFileHeader())
	record := r.File.Header

	require.Equal(t, line, record.String())
}

// TestFHString tests validating that a known parsed FileHeader can return to a string of the same value
func TestFHString(t *testing.T) {
	testFHString(t)
}

// BenchmarkFHString benchmarks validating that a known parsed FileHeader
// can return to a string of the same value
func BenchmarkFHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHString(b)
	}
}

// TestFHRecordType validation
func TestFHRecordType(t *testing.T) {
	fh := mockFileHeader()
	fh.recordType = "00"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestStandardLevel validation
func TestStandardLevel(t *testing.T) {
	fh := mockFileHeader()
	fh.StandardLevel = "01"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "StandardLevel", e.FieldName)
}

// TestTestFileIndicator validation
func TestTestFileIndicator(t *testing.T) {
	fh := mockFileHeader()
	fh.TestFileIndicator = "S"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TestFileIndicator", e.FieldName)
}

// TestResendIndicator validation
func TestResendIndicator(t *testing.T) {
	fh := mockFileHeader()
	fh.ResendIndicator = "R"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ResendIndicator", e.FieldName)
}

// TestImmediateDestinationName validation
func TestImmediateDestinationName(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestinationName = "®©"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImmediateDestinationName", e.FieldName)
}

// TestImmediateOriginName validation
func TestImmediateOriginName(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOriginName = "®©"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImmediateOriginName", e.FieldName)
}

// TestFileIDModifier validation
func TestFileIDModifier(t *testing.T) {
	fh := mockFileHeader()
	fh.FileIDModifier = "--"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FileIDModifier", e.FieldName)
}

// TestCountryCode validation
func TestCountryCode(t *testing.T) {
	fh := mockFileHeader()
	fh.CompanionDocumentIndicator = "D"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CompanionDocumentIndicator", e.FieldName)
}

// TestCACountryCode validation
func TestCACountryCode(t *testing.T) {
	fh := mockFileHeader()
	fh.CountryCode = "CA"
	fh.CompanionDocumentIndicator = "1"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CompanionDocumentIndicator", e.FieldName)
}

// TestUserField validation
func TestUserFieldI(t *testing.T) {
	fh := mockFileHeader()
	fh.UserField = "®©"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// TestFHFieldInclusionRecordType validates FieldInclusion
func TestFHFieldInclusionRecordType(t *testing.T) {
	fh := mockFileHeader()
	fh.recordType = ""
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestFHFieldInclusionStandardLevel validates FieldInclusion
func TestFHFieldInclusionStandardLevel(t *testing.T) {
	fh := mockFileHeader()
	fh.StandardLevel = ""
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "StandardLevel", e.FieldName)
}

// TestFHFieldInclusionTestFileIndicator validates FieldInclusion
func TestFHFieldInclusionTestFileIndicator(t *testing.T) {
	fh := mockFileHeader()
	fh.TestFileIndicator = ""
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TestFileIndicator", e.FieldName)
}

// TestFHFieldInclusionResendIndicator validates FieldInclusion
func TestFHFieldInclusionResendIndicator(t *testing.T) {
	fh := mockFileHeader()
	fh.ResendIndicator = ""
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ResendIndicator", e.FieldName)
}

// TestFHFieldInclusionImmediateDestination validates FieldInclusion
func TestFHFieldInclusionImmediateDestination(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestination = ""
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImmediateDestination", e.FieldName)
}

// TestFHFieldInclusionImmediateDestinationZero validates FieldInclusion
func TestFHFieldInclusionImmediateDestinationZero(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestination = "000000000"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImmediateDestination", e.FieldName)
}

// TestFHFieldInclusionImmediateOrigin validates FieldInclusion
func TestFHFieldInclusionImmediateOrigin(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = ""
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImmediateOrigin", e.FieldName)
}

// TestFHFieldInclusionImmediateOriginZero validates FieldInclusion
func TestFHFieldInclusionImmediateOriginZero(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = "000000000"
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImmediateOrigin", e.FieldName)
}

// TestFHFieldInclusionCreationDate validates FieldInclusion
func TestFHFieldInclusionCreationDate(t *testing.T) {
	fh := mockFileHeader()
	fh.FileCreationDate = time.Time{}
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FileCreationDate", e.FieldName)
}

// TestFHFieldInclusionCreationTime validates FieldInclusion
func TestFHFieldInclusionCreationTime(t *testing.T) {
	fh := mockFileHeader()
	fh.FileCreationTime = time.Time{}
	err := fh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FileCreationTime", e.FieldName)
}

// TestFileHeaderRuneCountInString validates RuneCountInString
func TestFileHeaderRuneCountInString(t *testing.T) {
	fh := NewFileHeader()
	var line = "01"
	fh.Parse(line)

	require.Equal(t, "", fh.ImmediateOrigin)
}
