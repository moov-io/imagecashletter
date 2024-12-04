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

// mockReturnDetailAddendumA creates a ReturnDetailAddendumA
func mockReturnDetailAddendumA() ReturnDetailAddendumA {
	rdAddendumA := NewReturnDetailAddendumA()
	rdAddendumA.RecordNumber = 1
	rdAddendumA.ReturnLocationRoutingNumber = "121042882"
	rdAddendumA.BOFDEndorsementDate = time.Now()
	rdAddendumA.BOFDItemSequenceNumber = "1              "
	rdAddendumA.BOFDAccountNumber = "938383"
	rdAddendumA.BOFDBranchCode = "01"
	rdAddendumA.PayeeName = "Test Payee"
	rdAddendumA.TruncationIndicator = "Y"
	rdAddendumA.BOFDConversionIndicator = "1"
	rdAddendumA.BOFDCorrectionIndicator = 0
	rdAddendumA.UserField = ""
	return rdAddendumA
}

// mockReturnDetailAddendumAWithoutBOFDItemSequenceNumber creates a ReturnDetailAddendumA
func mockReturnDetailAddendumAWithoutBOFDItemSequenceNumber() ReturnDetailAddendumA {
	rdAddendumA := NewReturnDetailAddendumA()
	rdAddendumA.RecordNumber = 1
	rdAddendumA.ReturnLocationRoutingNumber = "121042882"
	rdAddendumA.BOFDEndorsementDate = time.Now()
	rdAddendumA.BOFDAccountNumber = "938383"
	rdAddendumA.BOFDBranchCode = "01"
	rdAddendumA.PayeeName = "Test Payee"
	rdAddendumA.TruncationIndicator = "Y"
	rdAddendumA.BOFDConversionIndicator = "1"
	rdAddendumA.BOFDCorrectionIndicator = 0
	rdAddendumA.UserField = ""
	return rdAddendumA
}

func TestReturnDetailAddendumAParseErr(t *testing.T) {
	var r ReturnDetailAddendumA
	r.Parse("asdlsajhfakjfa")
	require.Equal(t, 0, r.RecordNumber)
}

// TestMockReturnDetailAddendumA creates a ReturnDetailAddendumA
func TestMockReturnDetailAddendumA(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	require.NoError(t, rdAddendumA.Validate())
	require.Equal(t, "32", rdAddendumA.recordType)
	require.Equal(t, 1, rdAddendumA.RecordNumber)
	require.Equal(t, "121042882", rdAddendumA.ReturnLocationRoutingNumber)
	require.Equal(t, "1              ", rdAddendumA.BOFDItemSequenceNumber)
	require.Equal(t, "938383", rdAddendumA.BOFDAccountNumber)
	require.Equal(t, "01", rdAddendumA.BOFDBranchCode)
	require.Equal(t, "Test Payee", rdAddendumA.PayeeName)
	require.Equal(t, "Y", rdAddendumA.TruncationIndicator)
	require.Equal(t, "1", rdAddendumA.BOFDConversionIndicator)
	require.Equal(t, 0, rdAddendumA.BOFDCorrectionIndicator)
	require.Equal(t, "", rdAddendumA.UserField)
}

// TestMockReturnDetailAddendumA creates a ReturnDetailAddendumA
func TestMockReturnDetailAddendumAWithoutBOFDItemSequenceNumber(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumAWithoutBOFDItemSequenceNumber()
	require.NoError(t, rdAddendumA.Validate())
	require.Equal(t, "32", rdAddendumA.recordType)
	require.Equal(t, 1, rdAddendumA.RecordNumber)
	require.Equal(t, "121042882", rdAddendumA.ReturnLocationRoutingNumber)
	require.Equal(t, "938383", rdAddendumA.BOFDAccountNumber)
	require.Equal(t, "01", rdAddendumA.BOFDBranchCode)
	require.Equal(t, "Test Payee", rdAddendumA.PayeeName)
	require.Equal(t, "Y", rdAddendumA.TruncationIndicator)
	require.Equal(t, "1", rdAddendumA.BOFDConversionIndicator)
	require.Equal(t, 0, rdAddendumA.BOFDCorrectionIndicator)
	require.Equal(t, "", rdAddendumA.UserField)
}

// TestParseReturnDetailAddendumAWithoutBOFDItemSequenceNumber validates parsing a ReturnDetailAddendumA
func TestParseReturnDetailAddendumAWithoutBOFDItemSequenceNumber(t *testing.T) {
	var line = "32112104288220180905               938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(rd)

	require.NoError(t, r.parseReturnDetailAddendumA())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumA[0]

	require.Equal(t, "32", record.recordType)
	require.Equal(t, "1", record.RecordNumberField())
	require.Equal(t, "121042882", record.ReturnLocationRoutingNumberField())
	require.Equal(t, "20180905", record.BOFDEndorsementDateField())
	require.Equal(t, "               ", record.BOFDItemSequenceNumberField())
	require.Equal(t, "938383            ", record.BOFDAccountNumberField())
	require.Equal(t, "01   ", record.BOFDBranchCodeField())
	require.Equal(t, "Test Payee     ", record.PayeeNameField())
	require.Equal(t, "Y", record.TruncationIndicatorField())
	require.Equal(t, "1", record.BOFDConversionIndicatorField())
	require.Equal(t, "0", record.BOFDCorrectionIndicatorField())
	require.Equal(t, " ", record.UserFieldField())
	require.Equal(t, "   ", record.reservedField())
}

// TestParseReturnDetailAddendumA validates parsing a ReturnDetailAddendumA
func TestParseReturnDetailAddendumA(t *testing.T) {
	var line = "321121042882201809051              938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(rd)

	require.NoError(t, r.parseReturnDetailAddendumA())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumA[0]

	require.Equal(t, "32", record.recordType)
	require.Equal(t, "1", record.RecordNumberField())
	require.Equal(t, "121042882", record.ReturnLocationRoutingNumberField())
	require.Equal(t, "20180905", record.BOFDEndorsementDateField())
	require.Equal(t, "1              ", record.BOFDItemSequenceNumberField())
	require.Equal(t, "938383            ", record.BOFDAccountNumberField())
	require.Equal(t, "01   ", record.BOFDBranchCodeField())
	require.Equal(t, "Test Payee     ", record.PayeeNameField())
	require.Equal(t, "Y", record.TruncationIndicatorField())
	require.Equal(t, "1", record.BOFDConversionIndicatorField())
	require.Equal(t, "0", record.BOFDCorrectionIndicatorField())
	require.Equal(t, " ", record.UserFieldField())
	require.Equal(t, "   ", record.reservedField())
}

// testRDAddendumAString validates that a known parsed ReturnDetailAddendumA can return to a string of the same value
func testRDAddendumAString(t testing.TB) {
	var line = "321121042882201809051              938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(rd)

	require.NoError(t, r.parseReturnDetailAddendumA())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumA[0]

	require.Equal(t, line, record.String())

}

// TestRDAddendumAString tests validating that a known parsed ReturnDetailAddendumA can return to a string of the
// same value
func TestRDAddendumAString(t *testing.T) {
	testRDAddendumAString(t)
}

// BenchmarkRDAddendumAString benchmarks validating that a known parsed ReturnDetailAddendumA
// can return to a string of the same value
func BenchmarkRDAddendumAString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDAddendumAString(b)
	}
}

// TestRDAddendumARecordType validation
func TestRDAddendumARecordType(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.recordType = "00"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDAddendumAReturnLocationRoutingNumber validation
func TestRDAddendumAReturnLocationRoutingNumber(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.ReturnLocationRoutingNumber = "X"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnLocationRoutingNumber", e.FieldName)
}

// TestRDAddendumABOFDAccountNumber validation
func TestRDAddendumABOFDAccountNumber(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.BOFDAccountNumber = "®©"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDAccountNumber", e.FieldName)
}

// TestRDAddendumABOFDBranchCode validation
func TestRDAddendumABOFDBranchCode(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.BOFDBranchCode = "®©"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDBranchCode", e.FieldName)
}

// TestRDAddendumAPayeeName validation
func TestRDAddendumAPayeeName(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.PayeeName = "®©"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayeeName", e.FieldName)
}

// TestRDAddendumATruncationIndicator validation
func TestRDAddendumATruncationIndicator(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.TruncationIndicator = "A"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
}

// TestRDAddendumABOFDConversionIndicator validation
func TestRDAddendumABOFDConversionIndicator(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.BOFDConversionIndicator = "99"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDConversionIndicator", e.FieldName)
}

// TestRDAddendumABOFDCorrectionIndicator validation
func TestRDAddendumABOFDCorrectionIndicator(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.BOFDCorrectionIndicator = 10
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDCorrectionIndicator", e.FieldName)
}

// TestRDAddendumAUserField validation
func TestRDAddendumAUserField(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.UserField = "®©"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestRDAddendumAFIRecordType validation
func TestRDAddendumAFIRecordType(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.recordType = ""
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDAddendumAFIRecordNumber validation
func TestRDAddendumAFIRecordNumber(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.RecordNumber = 0
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "RecordNumber", e.FieldName)
}

// TestRDAddendumAFIReturnLocationRoutingNumber validation
func TestRDAddendumAFIReturnLocationRoutingNumber(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.ReturnLocationRoutingNumber = ""
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnLocationRoutingNumber", e.FieldName)
}

// TestRDAddendumAFIReturnLocationRoutingNumberZero validation
func TestRDAddendumAFIReturnLocationRoutingNumberZero(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.ReturnLocationRoutingNumber = "000000000"
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnLocationRoutingNumber", e.FieldName)
}

// TestRDAddendumAFIBOFDEndorsementDate validation
func TestRDAddendumAFIBOFDEndorsementDate(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.BOFDEndorsementDate = time.Time{}
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDEndorsementDate", e.FieldName)
}

// TestRDAddendumAFIBOFDEndorsementDateFRB validation
func TestRDAddendumAFIBOFDEndorsementDateFRB(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.BOFDEndorsementDate = time.Time{}
	t.Setenv(FRBCompatibilityMode, "true")
	require.NoError(t, rdAddendumA.Validate())
}

// TestRDAddendumAFITruncationIndicator validation
func TestRDAddendumAFITruncationIndicator(t *testing.T) {
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.TruncationIndicator = ""
	err := rdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
}
