// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockCreditItem creates a CreditItem
func mockCreditItem() *CreditItem {
	ci := NewCreditItem()
	ci.AuxiliaryOnUs = "123456789"
	ci.ExternalProcessingCode = ""
	ci.PostingBankRoutingNumber = "031300012"
	ci.OnUs = "5558881"
	ci.ItemAmount = 100000 // 1000.00
	ci.CreditItemSequenceNumber = "1              "
	ci.DocumentationTypeIndicator = "G"
	ci.AccountTypeCode = "1"
	ci.SourceWorkCode = "01"
	ci.UserField = "                "
	return ci
}

// TestMockCreditItem creates a CreditItem
func TestMockCreditItem(t *testing.T) {
	ci := mockCreditItem()
	require.NoError(t, ci.Validate())
	require.Equal(t, "62", ci.recordType)
	require.Equal(t, "123456789", ci.AuxiliaryOnUs)
	require.Equal(t, "", ci.ExternalProcessingCode)
	require.Equal(t, "031300012", ci.PostingBankRoutingNumber)
	require.Equal(t, "5558881", ci.OnUs)
	require.Equal(t, 100000, ci.ItemAmount)
	require.Equal(t, "1              ", ci.CreditItemSequenceNumber)
	require.Equal(t, "G", ci.DocumentationTypeIndicator)
	require.Equal(t, "1", ci.AccountTypeCode)
	require.Equal(t, "01", ci.SourceWorkCode)
	require.Equal(t, "                ", ci.UserField)
}

func TestCreditItemCrash(t *testing.T) {
	ci := &CreditItem{}
	ci.Parse(`100000000000400000000200001010000010100000IG000000000000000000000000000000000 00`)
	require.Equal(t, "", ci.UserField)
}

// TestParseCreditItem validates parsing a CreditItem
func TestParseCreditItem(t *testing.T) {
	var line = "62      123456789 031300012             5558881000000001000001              G101                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	ci := mockCreditItem()
	r.currentCashLetter.AddCreditItem(ci)
	require.NoError(t, r.parseCreditItem())
	record := r.currentCashLetter.GetCreditItems()[0]

	require.Equal(t, "62", record.recordType)
	require.Equal(t, "123456789", record.AuxiliaryOnUs)
	require.Equal(t, "", record.ExternalProcessingCode)
	require.Equal(t, "031300012", record.PostingBankRoutingNumber)
	require.Equal(t, "5558881", record.OnUs)
	require.Equal(t, 100000, record.ItemAmount)
	require.Equal(t, "1              ", record.CreditItemSequenceNumber)
	require.Equal(t, "G", record.DocumentationTypeIndicator)
	require.Equal(t, "1", record.AccountTypeCode)
	require.Equal(t, "01", record.SourceWorkCode)
	require.Equal(t, "                ", record.UserField)
}

// testCIString validates parsing a CreditItem
func testCIString(t testing.TB) {
	var line = "62      123456789 031300012             5558881000000001000001              G101                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	ci := mockCreditItem()
	r.currentCashLetter.AddCreditItem(ci)
	require.NoError(t, r.parseCreditItem())
	record := r.currentCashLetter.GetCreditItems()[0]

	require.Equal(t, line, record.String())
}

// TestCIString tests validating that a known parsed CheckDetail can return to a string of the same value
func TestCIString(t *testing.T) {
	testCIString(t)
}

// BenchmarkCIString benchmarks validating that a known parsed CreditItem
// can return to a string of the same value
func BenchmarkCIString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCIString(b)
	}
}

// TestCIRecordType validation
func TestCIRecordType(t *testing.T) {
	ci := mockCreditItem()
	ci.recordType = "00"
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCIDocumentationTypeIndicator validation
func TestCIDocumentationTypeIndicator(t *testing.T) {
	ci := mockCreditItem()
	ci.DocumentationTypeIndicator = "P"
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DocumentationTypeIndicator", e.FieldName)
}

// TestCIDocumentationTypeIndicatorZ validation
func TestCIDocumentationTypeIndicatorZ(t *testing.T) {
	ci := mockCreditItem()
	ci.DocumentationTypeIndicator = "Z"
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DocumentationTypeIndicator", e.FieldName)
}

// TestCIDocumentationTypeIndicatorM validation
func TestCIDocumentationTypeIndicatorM(t *testing.T) {
	ci := mockCreditItem()
	ci.DocumentationTypeIndicator = "M"
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DocumentationTypeIndicator", e.FieldName)
}

// TestCIAccountTypeCode validation
func TestCIAccountTypeCode(t *testing.T) {
	ci := mockCreditItem()
	ci.AccountTypeCode = "Z"
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "AccountTypeCode", e.FieldName)
}

// TestCISourceWorkCode validation
func TestCISourceWorkCode(t *testing.T) {
	ci := mockCreditItem()
	ci.SourceWorkCode = "99"
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "SourceWorkCode", e.FieldName)
}

// TestCIUserField validation
func TestCIUserField(t *testing.T) {
	ci := mockCreditItem()
	ci.UserField = "®©"
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestCIFIRecordType validation
func TestCIFIRecordType(t *testing.T) {
	ci := mockCreditItem()
	ci.recordType = ""
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCIFIPostingBankRoutingNumber validation
func TestCIFIPostingBankRoutingNumber(t *testing.T) {
	ci := mockCreditItem()
	ci.PostingBankRoutingNumber = ""
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PostingBankRoutingNumber", e.FieldName)
}

// TestCIFIPostingBankRoutingNumberZero validation
func TestCIFIPostingBankRoutingNumberZero(t *testing.T) {
	ci := mockCreditItem()
	ci.PostingBankRoutingNumber = "000000000"
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PostingBankRoutingNumber", e.FieldName)
}

// TestCIFICreditItemSequenceNumber validation
func TestCIFICreditItemSequenceNumber(t *testing.T) {
	ci := mockCreditItem()
	ci.CreditItemSequenceNumber = ""
	err := ci.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CreditItemSequenceNumber", e.FieldName)
}
