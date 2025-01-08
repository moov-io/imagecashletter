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

// mockCheckDetailAddendumA creates a CheckDetailAddendumA
func mockCheckDetailAddendumA() CheckDetailAddendumA {
	cdAddendumA := NewCheckDetailAddendumA()
	cdAddendumA.RecordNumber = 1
	cdAddendumA.ReturnLocationRoutingNumber = "121042882"
	cdAddendumA.BOFDEndorsementDate = time.Now()
	cdAddendumA.BOFDItemSequenceNumber = "1              "
	cdAddendumA.BOFDAccountNumber = "938383"
	cdAddendumA.BOFDBranchCode = "01"
	cdAddendumA.PayeeName = "Test Payee"
	cdAddendumA.TruncationIndicator = "Y"
	cdAddendumA.BOFDConversionIndicator = "1"
	cdAddendumA.BOFDCorrectionIndicator = 0
	cdAddendumA.UserField = ""
	return cdAddendumA
}

func TestCheckDetailAddendumParseErr(t *testing.T) {
	var c CheckDetailAddendumA
	c.Parse("asdshfaksjs")
	require.Equal(t, 0, c.RecordNumber)
}

// TestMockCheckDetailAddendumA creates a CheckDetailAddendumA
func TestMockCheckDetailAddendumA(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	require.NoError(t, cdAddendumA.Validate())
	require.Equal(t, "26", cdAddendumA.recordType)
	require.Equal(t, 1, cdAddendumA.RecordNumber)
	require.Equal(t, "121042882", cdAddendumA.ReturnLocationRoutingNumber)
	require.Equal(t, "1              ", cdAddendumA.BOFDItemSequenceNumber)
	require.Equal(t, "938383", cdAddendumA.BOFDAccountNumber)
	require.Equal(t, "01", cdAddendumA.BOFDBranchCode)
	require.Equal(t, "Test Payee", cdAddendumA.PayeeName)
	require.Equal(t, "Y", cdAddendumA.TruncationIndicator)
	require.Equal(t, "1", cdAddendumA.BOFDConversionIndicator)
	require.Equal(t, 0, cdAddendumA.BOFDCorrectionIndicator)
	require.Equal(t, "", cdAddendumA.UserField)
}

// TestParseCheckDetailAddendumA validates parsing a CheckDetailAddendumA
func TestParseCheckDetailAddendumA(t *testing.T) {
	var line = "261121042882201809051              938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	require.NoError(t, r.parseCheckDetailAddendumA())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumA[0]

	require.Equal(t, "26", record.recordType)
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

// testCDAddendumAString validates that a known parsed CheckDetailAddendumA can return to a string of the same value
func testCDAddendumAString(t testing.TB) {
	var line = "261121042882201809051              938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	require.NoError(t, r.parseCheckDetailAddendumA())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumA[0]

	require.Equal(t, line, record.String())
}

// TestCDAddendumAString tests validating that a known parsed CheckDetailAddendumA can return to a string of the
// same value
func TestCDAddendumAString(t *testing.T) {
	testCDAddendumAString(t)
}

// BenchmarkCDAddendumAString benchmarks validating that a known parsed CheckDetailAddendumA
// can return to a string of the same value
func BenchmarkCDAddendumAString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumAString(b)
	}
}

// TestCDAddendumARecordType validation
func TestCDAddendumARecordType(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.recordType = "00"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCDAddendumAReturnLocationRoutingNumber validation
func TestCDAddendumAReturnLocationRoutingNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = "X"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnLocationRoutingNumber", e.FieldName)
}

// TestCDAddendumABOFDAccountNumber validation
func TestCDAddendumABOFDAccountNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDAccountNumber = "®©"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDAccountNumber", e.FieldName)
}

// TestCDAddendumABOFDBranchCode validation
func TestCDAddendumABOFDBranchCode(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDBranchCode = "®©"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDBranchCode", e.FieldName)
}

// TestCDAddendumAPayeeName validation
func TestCDAddendumAPayeeName(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.PayeeName = "®©"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayeeName", e.FieldName)
}

// TestCDAddendumATruncationIndicator validation
func TestCDAddendumATruncationIndicator(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.TruncationIndicator = "A"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
}

// TestCDAddendumATruncationIndicator validation
func TestCDAddendumATruncationIndicatorFRB(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.TruncationIndicator = ""
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
	t.Setenv(FRBCompatibilityMode, "true")
	require.NoError(t, cdAddendumA.Validate())
}

// TestCDAddendumABOFDConversionIndicator validation
func TestCDAddendumABOFDConversionIndicator(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDConversionIndicator = "99"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDConversionIndicator", e.FieldName)
}

// TestCDAddendumABOFDCorrectionIndicator validation
func TestCDAddendumABOFDCorrectionIndicator(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDCorrectionIndicator = 10
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDCorrectionIndicator", e.FieldName)
}

// TestCDAddendumAUserField validation
func TestCDAddendumAUserField(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.UserField = "®©"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestCDAddendumAFIRecordType validation
func TestCDAddendumAFIRecordType(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.recordType = ""
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCDAddendumAFIRecordNumber validation
func TestCDAddendumAFIRecordNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.RecordNumber = 0
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "RecordNumber", e.FieldName)
}

// TestCDAddendumAFIReturnLocationRoutingNumber validation
func TestCDAddendumAFIReturnLocationRoutingNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = ""
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnLocationRoutingNumber", e.FieldName)
}

// TestCDAddendumAFIReturnLocationRoutingNumberZero validation
func TestCDAddendumAFIReturnLocationRoutingNumberZero(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = "000000000"
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnLocationRoutingNumber", e.FieldName)
}

func TestCDAddendumAFIReturnLocationRoutingNumberZeroFRB(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = "000000000"

	// enable FRB mode and verify it passes
	t.Setenv(FRBCompatibilityMode, "true")
	require.NoError(t, cdAddendumA.Validate())

	t.Setenv(FRBCompatibilityMode, "")
	require.ErrorContains(t, cdAddendumA.Validate(), "ReturnLocationRoutingNumber")
}

// TestCDAddendumAFIBOFDEndorsementDate validation
func TestCDAddendumAFIBOFDEndorsementDate(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDEndorsementDate = time.Time{}
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDEndorsementDate", e.FieldName)
}

// TestCDAddendumAFIBOFDItemSequenceNumber validation
func TestCDAddendumAFIBOFDItemSequenceNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDItemSequenceNumber = "               "
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDItemSequenceNumber", e.FieldName)
}

// TestCDAddendumAFITruncationIndicator validation
func TestCDAddendumAFITruncationIndicator(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.TruncationIndicator = ""
	err := cdAddendumA.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
}

// End FieldInclusion

// TestAlphaFieldTrim validation
func TestAlphaFieldTrim(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.PayeeName = "Payee Name James Steel"
	require.Len(t, cdAddendumA.PayeeNameField(), 15)
}

// TestStringFieldTrim validation
func TestStringFieldTrim(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = "12345678912345"
	require.Len(t, cdAddendumA.ReturnLocationRoutingNumberField(), 9)
}

func TestParseCheckDetailAddendumA_BOFDAccountEBCDIC(t *testing.T) {
	t.Setenv("FRB_COMPATIBILITY_MODE", "true")
	line := "\xf2\xf6" + // Record Type 26
		strings.Repeat("\xf1", 33) + // Fill with '1's
		"@@@@@@@@@@@" + // Spaces
		"\xad\x85\x94\x97\xa3\xa8\xbd" + // [empty] in IBM1047
		strings.Repeat("@", 20) + // More spaces
		"\xe8\xf2\xf0@@@@" // End padding
	r := NewReader(strings.NewReader(line), ReadEbcdicEncodingOption())
	r.line = line

	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	err := r.parseCheckDetailAddendumA()
	require.NoError(t, err)

	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumA[0]
	require.Equal(t, "[empty]", record.BOFDAccountNumber)
	t.Setenv("FRB_COMPATIBILITY_MODE", "")
}

func TestParseCheckDetailAddendumA_BOFDAccountEBCDIC_NoFlag(t *testing.T) {
	t.Setenv("FRB_COMPATIBILITY_MODE", "false")
	line := "\xf2\xf6" +
		strings.Repeat("\xf1", 33) +
		"@@@@@@@@@@@" +
		"\xad\x85\x94\x97\xa3\xa8\xbd" +
		strings.Repeat("@", 20) +
		"\xe8\xf2\xf0@@@@"

	r := NewReader(strings.NewReader(line), ReadEbcdicEncodingOption())
	r.line = line

	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	err := r.parseCheckDetailAddendumA()
	require.Error(t, err, "Expected an error when FRB_COMPATIBILITY_MODE is false")
	t.Setenv("FRB_COMPATIBILITY_MODE", "")
}
