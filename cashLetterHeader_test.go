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

// mockCashLetterHeader creates a CashLetterHeader
func mockCashLetterHeader() *CashLetterHeader {
	clh := NewCashLetterHeader()
	clh.CollectionTypeIndicator = "01"
	clh.DestinationRoutingNumber = "231380104"
	clh.ECEInstitutionRoutingNumber = "121042882"
	clh.CashLetterBusinessDate = time.Now()
	clh.CashLetterCreationDate = time.Now()
	clh.CashLetterCreationTime = time.Now()
	clh.RecordTypeIndicator = "I"
	clh.DocumentationTypeIndicator = "G"
	clh.CashLetterID = "A1"
	clh.OriginatorContactName = "Contact Name"
	clh.OriginatorContactPhoneNumber = "5558675552"
	clh.FedWorkType = ""
	clh.ReturnsIndicator = ""
	clh.UserField = ""
	return clh
}

// TestMockCashLetterHeader creates a CashLetterHeader
func TestMockCashLetterHeader(t *testing.T) {
	clh := mockCashLetterHeader()
	require.NoError(t, clh.Validate())
	require.Equal(t, "10", clh.recordType)
	require.Equal(t, "01", clh.CollectionTypeIndicator)
	require.Equal(t, "231380104", clh.DestinationRoutingNumber)
	require.Equal(t, "121042882", clh.ECEInstitutionRoutingNumber)
	require.Equal(t, "I", clh.RecordTypeIndicator)
	require.Equal(t, "G", clh.DocumentationTypeIndicator)
	require.Equal(t, "A1", clh.CashLetterID)
	require.Equal(t, "Contact Name", clh.OriginatorContactName)
	require.Equal(t, "5558675552", clh.OriginatorContactPhoneNumber)
	require.Equal(t, "", clh.FedWorkType)
	require.Equal(t, "", clh.ReturnsIndicator)
	require.Equal(t, "", clh.UserField)
	require.Equal(t, "", clh.reserved)
}

// TestParseCashLetterHeader validates parsing a CashLetterHeader
func TestParseCashLetterHeader(t *testing.T) {
	var line = "100123138010412104288220180905201809051523IGA1      Contact Name  5558675552    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	require.NoError(t, r.parseCashLetterHeader())
	record := r.currentCashLetter.CashLetterHeader

	require.Equal(t, "10", record.recordType)
	require.Equal(t, "01", record.CollectionTypeIndicatorField())
	require.Equal(t, "231380104", record.DestinationRoutingNumberField())
	require.Equal(t, "121042882", record.ECEInstitutionRoutingNumberField())
	require.Equal(t, "20180905", record.CashLetterBusinessDateField())
	require.Equal(t, "20180905", record.CashLetterCreationDateField())
	require.Equal(t, "1523", record.CashLetterCreationTimeField())
	require.Equal(t, "I", record.RecordTypeIndicatorField())
	require.Equal(t, "G", record.DocumentationTypeIndicatorField())
	require.Equal(t, "A1      ", record.CashLetterIDField())
	require.Equal(t, "Contact Name  ", record.OriginatorContactNameField())
	require.Equal(t, "5558675552", record.OriginatorContactPhoneNumberField())
	require.Equal(t, " ", record.FedWorkTypeField())
	require.Equal(t, " ", record.ReturnsIndicatorField())
	require.Equal(t, " ", record.UserFieldField())
	require.Equal(t, " ", record.reservedField())
}

// testCLHString validates that a known parsed CashLetterHeader can return to a string of the same value
func testCLHString(t testing.TB) {
	var line = "100123138010412104288220180905201809051523IGA1      Contact Name  5558675552    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	require.NoError(t, r.parseCashLetterHeader())
	record := r.currentCashLetter.CashLetterHeader
	require.Equal(t, line, record.String())
}

// TestCLHString tests validating that a known parsed CashLetterHeader can return to a string of the same value
func TestCLHString(t *testing.T) {
	testCLHString(t)
}

// BenchmarkCLHString benchmarks validating that a known parsed CashLetterHeader
// can return to a string of the same value
func BenchmarkCLHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCLHString(b)
	}
}

// TestCLHRecordType validation
func TestCLHRecordType(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.recordType = "00"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCHCollectionTypeIndicator validation
func TestCHCollectionTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CollectionTypeIndicator = "87"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CollectionTypeIndicator", e.FieldName)
}

// TestRecordTypeIndicator validation
func TestRecordTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.RecordTypeIndicator = "W"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "RecordTypeIndicator", e.FieldName)
}

// TestDocumentationTypeIndicator validation
func TestDocumentationTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.DocumentationTypeIndicator = "WAZ"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DocumentationTypeIndicator", e.FieldName)
}

// TestCashLetterID validation
func TestCashLetterID(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterID = "--"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CashLetterID", e.FieldName)
}

// TestOriginatorContactName validation
func TestOriginatorContactName(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.OriginatorContactName = "®©"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OriginatorContactName", e.FieldName)
}

// TestOriginatorContactPhoneNumber validation
func TestOriginatorContactPhoneNumber(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.OriginatorContactPhoneNumber = "--"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OriginatorContactPhoneNumber", e.FieldName)
}

// TestFedWorkType validation
func TestFedWorkType(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.FedWorkType = "--"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "FedWorkType", e.FieldName)
}

// TestReturnsIndicator validation
func TestReturnsIndicator(t *testing.T) {
	testCases := []struct {
		name             string
		returnsIndicator string
		err              bool
	}{
		{
			name:             "Invalid ReturnsIndicator",
			returnsIndicator: "A",
			err:              true,
		},
		{
			name:             "Valid ReturnsIndicator empty",
			returnsIndicator: "",
			err:              false,
		},
		{
			name:             "Valid ReturnsIndicator E",
			returnsIndicator: "E",
			err:              false,
		},
		{
			name:             "Valid ReturnsIndicator R",
			returnsIndicator: "R",
			err:              false,
		},
		{
			name:             "Valid ReturnsIndicator J",
			returnsIndicator: "J",
			err:              false,
		},
		{
			name:             "Valid ReturnsIndicator N",
			returnsIndicator: "N",
			err:              false,
		},
	}
	clh := mockCashLetterHeader()
	var e *FieldError
	for _, tc := range testCases {
		clh.ReturnsIndicator = tc.returnsIndicator
		err := clh.Validate()
		if tc.err {
			require.ErrorAs(t, err, &e)
			require.Equal(t, "ReturnsIndicator", e.FieldName)
		} else {
			require.NoError(t, err)
		}
	}
}

// TestCLHUserField validation
func TestCLHUserFieldI(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.UserField = "®©"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// TestCLHFieldInclusionRecordType validates FieldInclusion
func TestCLHFieldInclusionRecordType(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.recordType = ""
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestFieldInclusionCollectionTypeIndicator validates FieldInclusion
func TestFieldInclusionCollectionTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CollectionTypeIndicator = ""
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CollectionTypeIndicator", e.FieldName)
}

// TestFieldInclusionRecordTypeIndicator validates FieldInclusion
func TestFieldInclusionRecordTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.RecordTypeIndicator = ""
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "RecordTypeIndicator", e.FieldName)
}

// TestFieldInclusionDestinationRoutingNumber validates FieldInclusion
func TestFieldInclusionDestinationRoutingNumber(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.DestinationRoutingNumber = ""
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DestinationRoutingNumber", e.FieldName)
}

// TestFieldInclusionDestinationRoutingZerovalidates FieldInclusion
func TestFieldInclusionDestinationRoutingNumberZero(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.DestinationRoutingNumber = "000000000"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DestinationRoutingNumber", e.FieldName)
}

// TestFieldInclusionECEInstitutionRoutingNumber validates FieldInclusion
func TestFieldInclusionECEInstitutionRoutingNumber(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.ECEInstitutionRoutingNumber = ""
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ECEInstitutionRoutingNumber", e.FieldName)
}

// TestFieldInclusionECEInstitutionRoutingNumberZero validates FieldInclusion
func TestFieldInclusionECEInstitutionRoutingNumberZero(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.ECEInstitutionRoutingNumber = "000000000"
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ECEInstitutionRoutingNumber", e.FieldName)
}

// TestFieldInclusionCashLetterBusinessDate validates FieldInclusion
func TestFieldInclusionCashLetterBusinessDate(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterBusinessDate = time.Time{}
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CashLetterBusinessDate", e.FieldName)
}

// TestFieldInclusionCashLetterCreationDate validates FieldInclusion
func TestFieldInclusionCashLetterCreationDate(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterCreationDate = time.Time{}
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CashLetterCreationDate", e.FieldName)
}

// TestFieldInclusionCashLetterCreationTime validates FieldInclusion
func TestFieldInclusionCashLetterCreationTime(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterCreationTime = time.Time{}
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CashLetterCreationTime", e.FieldName)
}

// TestFieldInclusionCashLetterID validates FieldInclusion
func TestFieldInclusionCashLetterID(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterID = ""
	err := clh.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CashLetterID", e.FieldName)
}

// TestCashLetterHeaderRuneCountInString validates RuneCountInString
func TestCashLetterHeaderRuneCountInString(t *testing.T) {
	clh := NewCashLetterHeader()
	var line = "10"
	clh.Parse(line)

	require.Equal(t, "", clh.CollectionTypeIndicator)
}
