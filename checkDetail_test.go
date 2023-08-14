// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockCheckDetail creates a CheckDetail
func mockCheckDetail() *CheckDetail {
	cd := NewCheckDetail()
	cd.AuxiliaryOnUs = "123456789"
	cd.ExternalProcessingCode = ""
	cd.PayorBankRoutingNumber = "03130001"
	cd.PayorBankCheckDigit = "2"
	cd.OnUs = "5558881"
	cd.ItemAmount = 100000 // 1000.00
	cd.EceInstitutionItemSequenceNumber = "1              "
	cd.DocumentationTypeIndicator = "G"
	cd.ReturnAcceptanceIndicator = "D"
	cd.MICRValidIndicator = 1
	cd.BOFDIndicator = "Y"
	cd.AddendumCount = 3
	cd.CorrectionIndicator = 0
	cd.ArchiveTypeIndicator = "B"
	return cd
}

func TestCheckDetailParseErr(t *testing.T) {
	var c CheckDetail
	c.Parse("jakjsakjfas")
	require.Equal(t, "", c.AuxiliaryOnUs)
}

// TestMockCheckDetail creates a CheckDetail
func TestMockCheckDetail(t *testing.T) {
	cd := mockCheckDetail()
	require.NoError(t, cd.Validate())
	require.Equal(t, "25", cd.recordType)
	require.Equal(t, "123456789", cd.AuxiliaryOnUs)
	require.Equal(t, "", cd.ExternalProcessingCode)
	require.Equal(t, "03130001", cd.PayorBankRoutingNumber)
	require.Equal(t, "2", cd.PayorBankCheckDigit)
	require.Equal(t, "5558881", cd.OnUs)
	require.Equal(t, 100000, cd.ItemAmount)
	require.Equal(t, "1              ", cd.EceInstitutionItemSequenceNumber)
	require.Equal(t, "G", cd.DocumentationTypeIndicator)
	require.Equal(t, "D", cd.ReturnAcceptanceIndicator)
	require.Equal(t, 1, cd.MICRValidIndicator)
	require.Equal(t, "Y", cd.BOFDIndicator)
	require.Equal(t, 3, cd.AddendumCount)
	require.Equal(t, 0, cd.CorrectionIndicator)
	require.Equal(t, "B", cd.ArchiveTypeIndicator)
}

// TestParseCheckDetail validates parsing a CheckDetail
func TestParseCheckDetail(t *testing.T) {
	var line = "25      123456789 031300012             555888100001000001              GD1Y030B"
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)

	require.NoError(t, r.parseCheckDetail())
	record := r.currentCashLetter.currentBundle.GetChecks()[0]

	require.Equal(t, "25", record.recordType)
	require.Equal(t, "      123456789", record.AuxiliaryOnUsField())
	require.Equal(t, " ", record.ExternalProcessingCodeField())
	require.Equal(t, "03130001", record.PayorBankRoutingNumberField())
	require.Equal(t, "2", record.PayorBankCheckDigitField())
	require.Equal(t, "             5558881", record.OnUsField())
	require.Equal(t, "0000100000", record.ItemAmountField())
	require.Equal(t, "1              ", record.EceInstitutionItemSequenceNumberField())
	require.Equal(t, "G", record.DocumentationTypeIndicatorField())
	require.Equal(t, "D", record.ReturnAcceptanceIndicatorField())
	require.Equal(t, "1", record.MICRValidIndicatorField())
	require.Equal(t, "Y", record.BOFDIndicatorField())
	require.Equal(t, "03", record.AddendumCountField())
	require.Equal(t, "0", record.CorrectionIndicatorField())
	require.Equal(t, "B", record.ArchiveTypeIndicatorField())
}

// testCDString validates that a known parsed CheckDetail can return to a string of the same value
func testCDString(t testing.TB) {
	var line = "25      123456789 031300012             555888100001000001              GD1Y030B"
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	require.NoError(t, r.parseCheckDetail())
	record := r.currentCashLetter.currentBundle.GetChecks()[0]

	require.Equal(t, line, record.String())
}

// TestCDString tests validating that a known parsed CheckDetail can return to a string of the same value
func TestCDString(t *testing.T) {
	testCDString(t)
}

// BenchmarkCDString benchmarks validating that a known parsed CheckDetail
// can return to a string of the same value
func BenchmarkCDString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDString(b)
	}
}

// TestCDRecordType validation
func TestCDRecordType(t *testing.T) {
	cd := mockCheckDetail()
	cd.recordType = "00"
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCDDocumentationTypeIndicator validation
func TestCDDocumentationTypeIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.DocumentationTypeIndicator = "P"
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DocumentationTypeIndicator", e.FieldName)
}

// TestCDDocumentationTypeIndicatorZ validation
func TestCDDocumentationTypeIndicatorZ(t *testing.T) {
	cd := mockCheckDetail()
	cd.DocumentationTypeIndicator = "Z"
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DocumentationTypeIndicator", e.FieldName)
}

// TestCDReturnAcceptanceIndicator validation
func TestCDReturnAcceptanceIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.ReturnAcceptanceIndicator = "M"
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnAcceptanceIndicator", e.FieldName)
}

// TestCDMICRValidIndicator validation
func TestCDMICRValidIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.MICRValidIndicator = 7
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "MICRValidIndicator", e.FieldName)
}

// TestCDBOFDIndicator validation
func TestCDBOFDIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.BOFDIndicator = "B"
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDIndicator", e.FieldName)
}

// TestCDCorrectionIndicator validation
func TestCDCorrectionIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.CorrectionIndicator = 10
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CorrectionIndicator", e.FieldName)
}

// TestCDArchiveTypeIndicator validation
func TestCDArchiveTypeIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.ArchiveTypeIndicator = "W"
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ArchiveTypeIndicator", e.FieldName)
}

// Field Inclusion

// TestCDFIRecordType validation
func TestCDFIRecordType(t *testing.T) {
	cd := mockCheckDetail()
	cd.recordType = ""
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCDFIPayorBankRoutingNumber validation
func TestCDFIPayorBankRoutingNumber(t *testing.T) {
	cd := mockCheckDetail()
	cd.PayorBankRoutingNumber = ""
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankRoutingNumber", e.FieldName)
}

// TestCDFIPayorBankRoutingNumberZero validation
func TestCDFIPayorBankRoutingNumberZero(t *testing.T) {
	cd := mockCheckDetail()
	cd.PayorBankRoutingNumber = "00000000"
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankRoutingNumber", e.FieldName)
}

// TestCDFIPayorBankCheckDigit validation
func TestCDFIPayorBankCheckDigit(t *testing.T) {
	cd := mockCheckDetail()
	cd.PayorBankCheckDigit = ""
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankCheckDigit", e.FieldName)
}

// TestCDFIEceInstitutionItemSequenceNumber validation
func TestCDFIEceInstitutionItemSequenceNumber(t *testing.T) {
	cd := mockCheckDetail()
	cd.EceInstitutionItemSequenceNumber = "               "
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EceInstitutionItemSequenceNumber", e.FieldName)
}

// TestCDFIBOFDIndicator validation
func TestCDFIBOFDIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.BOFDIndicator = ""
	err := cd.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDIndicator", e.FieldName)
}
