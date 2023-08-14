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

// mockReturnDetail creates a ReturnDetail
func mockReturnDetail() *ReturnDetail {
	rd := NewReturnDetail()
	rd.PayorBankRoutingNumber = "03130001"
	rd.PayorBankCheckDigit = "2"
	rd.OnUs = "5558881"
	rd.ItemAmount = 100000
	rd.ReturnReason = "A"
	rd.AddendumCount = 4
	rd.DocumentationTypeIndicator = "G"
	rd.ForwardBundleDate = time.Now()
	rd.EceInstitutionItemSequenceNumber = "1              "
	rd.ExternalProcessingCode = ""
	rd.ReturnNotificationIndicator = "2"
	rd.ArchiveTypeIndicator = "B"
	rd.TimesReturned = 0
	return rd
}

func TestReturnDetailParse(t *testing.T) {
	var r ReturnDetail
	r.Parse("asshafaksjfas")
	require.Equal(t, "", r.PayorBankRoutingNumber)
}

// TestMockReturnDetail creates a ReturnDetail
func TestMockReturnDetail(t *testing.T) {
	rd := mockReturnDetail()
	require.NoError(t, rd.Validate())
	require.Equal(t, "31", rd.recordType)
	require.Equal(t, "03130001", rd.PayorBankRoutingNumber)
	require.Equal(t, "2", rd.PayorBankCheckDigit)
	require.Equal(t, "5558881", rd.OnUs)
	require.Equal(t, 100000, rd.ItemAmount)
	require.Equal(t, "A", rd.ReturnReason)
	require.Equal(t, 4, rd.AddendumCount)
	require.Equal(t, "G", rd.DocumentationTypeIndicator)
	require.Equal(t, "1              ", rd.EceInstitutionItemSequenceNumber)
	require.Equal(t, "", rd.ExternalProcessingCode)
	require.Equal(t, "2", rd.ReturnNotificationIndicator)
	require.Equal(t, "B", rd.ArchiveTypeIndicator)
	require.Equal(t, 0, rd.TimesReturned)
}

// TestParseReturnDetail validates parsing a ReturnDetail
func TestParseReturnDetail(t *testing.T) {
	var line = "31031300012             55588810000100000A04G201809051               2B0        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)

	require.NoError(t, r.parseReturnDetail())
	record := r.currentCashLetter.currentBundle.GetReturns()[0]

	require.Equal(t, "31", record.recordType)
	require.Equal(t, "03130001", record.PayorBankRoutingNumberField())
	require.Equal(t, "2", record.PayorBankCheckDigitField())
	require.Equal(t, "             5558881", record.OnUsField())
	require.Equal(t, "0000100000", record.ItemAmountField())
	require.Equal(t, "A", record.ReturnReasonField())
	require.Equal(t, "04", record.AddendumCountField())
	require.Equal(t, "G", record.DocumentationTypeIndicatorField())
	require.Equal(t, "1              ", record.EceInstitutionItemSequenceNumberField())
	require.Equal(t, " ", record.ExternalProcessingCodeField())
	require.Equal(t, "2", record.ReturnNotificationIndicatorField())
	require.Equal(t, "B", record.ArchiveTypeIndicatorField())
	require.Equal(t, "0", record.TimesReturnedField())

}

// testReturnDetailString validates that a known parsed ReturnDetail can return to a string of the same value
func testReturnDetailString(t testing.TB) {
	var line = "31031300012             55588810000100000A04G201809051               2B0        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)

	require.NoError(t, r.parseReturnDetail())
	record := r.currentCashLetter.currentBundle.GetReturns()[0]

	require.Equal(t, line, record.String())
}

// TestReturnDetailString tests validating that a known parsed ReturnDetail can return to a string of the
// same value
func TestReturnDetailString(t *testing.T) {
	testReturnDetailString(t)
}

// BenchmarkReturnDetailString benchmarks validating that a known parsed ReturnDetailAddendumB
// can return to a string of the same value
func BenchmarkReturnDetailString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReturnDetailString(b)
	}
}

// TestRDRecordType validation
func TestRDRecordType(t *testing.T) {
	rd := mockReturnDetail()
	rd.recordType = "00"
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDDocumentationTypeIndicator validation
func TestRDDocumentationTypeIndicator(t *testing.T) {
	rd := mockReturnDetail()
	rd.DocumentationTypeIndicator = "P"
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DocumentationTypeIndicator", e.FieldName)
}

// TestRDDocumentationTypeIndicatorZ validation
func TestRDDocumentationTypeIndicatorZ(t *testing.T) {
	rd := mockReturnDetail()
	rd.DocumentationTypeIndicator = "Z"
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DocumentationTypeIndicator", e.FieldName)
}

// TestRDReturnNotificationIndicator validation
func TestRDReturnNotificationIndicator(t *testing.T) {
	rd := mockReturnDetail()
	rd.ReturnNotificationIndicator = "0"
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnNotificationIndicator", e.FieldName)
}

// TestRDArchiveTypeIndicator validation
func TestRDArchiveTypeIndicator(t *testing.T) {
	rd := mockReturnDetail()
	rd.ArchiveTypeIndicator = "W"
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ArchiveTypeIndicator", e.FieldName)
}

// TestRDTimesReturned validation
func TestRDTimesReturned(t *testing.T) {
	rd := mockReturnDetail()
	rd.TimesReturned = 5
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TimesReturned", e.FieldName)
}

// TestReturnReasonInvalid validation
func TestReturnReasonInvalid(t *testing.T) {
	rd := mockReturnDetail()
	rd.ReturnReason = "88"
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnReason", e.FieldName)
}

// Field Inclusion

// TestRDFIRecordType validation
func TestRDFIRecordType(t *testing.T) {
	rd := mockReturnDetail()
	rd.recordType = ""
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDFIPayorBankRoutingNumber validation
func TestRDFIPayorBankRoutingNumber(t *testing.T) {
	rd := mockReturnDetail()
	rd.PayorBankRoutingNumber = ""
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankRoutingNumber", e.FieldName)
}

// TestRDFIPayorBankRoutingNumberZero validation
func TestRDFIPayorBankRoutingNumberZero(t *testing.T) {
	rd := mockReturnDetail()
	rd.PayorBankRoutingNumber = "00000000"
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankRoutingNumber", e.FieldName)
}

// TestRDFIPayorBankCheckDigit validation
func TestRDFIPayorBankCheckDigit(t *testing.T) {
	rd := mockReturnDetail()
	rd.PayorBankCheckDigit = ""
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankCheckDigit", e.FieldName)
}

// TestRDFIReturnReason validation
func TestRDFIReturnReason(t *testing.T) {
	rd := mockReturnDetail()
	rd.ReturnReason = ""
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnReason", e.FieldName)
}

// TestRDFIEceInstitutionItemSequenceNumber validation
func TestRDFIEceInstitutionItemSequenceNumber(t *testing.T) {
	rd := mockReturnDetail()
	rd.EceInstitutionItemSequenceNumber = "               "
	err := rd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EceInstitutionItemSequenceNumber", e.FieldName)
}
