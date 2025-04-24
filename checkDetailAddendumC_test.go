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

// mockCheckDetailAddendumC creates a CheckDetailAddendumC
func mockCheckDetailAddendumC() CheckDetailAddendumC {
	cdAddendumC := NewCheckDetailAddendumC()
	cdAddendumC.RecordNumber = 1
	cdAddendumC.EndorsingBankRoutingNumber = "121042882"
	cdAddendumC.BOFDEndorsementBusinessDate = time.Now()
	cdAddendumC.EndorsingBankItemSequenceNumber = "1              "
	cdAddendumC.TruncationIndicator = "Y"
	cdAddendumC.EndorsingBankConversionIndicator = "1"
	cdAddendumC.EndorsingBankCorrectionIndicator = 0
	cdAddendumC.ReturnReason = "A"
	cdAddendumC.UserField = ""
	cdAddendumC.EndorsingBankIdentifier = 0
	return cdAddendumC
}

// TestMockCheckDetailAddendumC creates a CheckDetailAddendumC
func TestMockCheckDetailAddendumC(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	require.NoError(t, cdAddendumC.Validate())
	require.Equal(t, "28", cdAddendumC.recordType)
	require.Equal(t, 1, cdAddendumC.RecordNumber)
	require.Equal(t, "121042882", cdAddendumC.EndorsingBankRoutingNumber)
	require.Equal(t, "1              ", cdAddendumC.EndorsingBankItemSequenceNumber)
	require.Equal(t, "Y", cdAddendumC.TruncationIndicator)
	require.Equal(t, "A", cdAddendumC.ReturnReason)
	require.Equal(t, "1", cdAddendumC.EndorsingBankConversionIndicator)
	require.Equal(t, 0, cdAddendumC.EndorsingBankCorrectionIndicator)
	require.Equal(t, "", cdAddendumC.UserField)
	require.Equal(t, 0, cdAddendumC.EndorsingBankIdentifier)
}

func TestParseCheckDetailAddendumCErr(t *testing.T) {
	var c CheckDetailAddendumC
	c.Parse("asdsakjahsfa")
	require.Equal(t, 0, c.RecordNumber)
}

// TestParseCheckDetailAddendumC validates parsing a CheckDetailAddendumC
func TestParseCheckDetailAddendumC(t *testing.T) {
	var line = "2801121042882201809051              Y10A                   0                    "
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

	require.NoError(t, r.parseCheckDetailAddendumC())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumC[0]

	require.Equal(t, "28", record.recordType)
	require.Equal(t, "01", record.RecordNumberField())

	require.Equal(t, "121042882", record.EndorsingBankRoutingNumberField())
	require.Equal(t, "20180905", record.BOFDEndorsementBusinessDateField())
	require.Equal(t, "1              ", record.EndorsingBankItemSequenceNumberField())
	require.Equal(t, "Y", record.TruncationIndicatorField())
	require.Equal(t, "1", record.EndorsingBankConversionIndicatorField())
	require.Equal(t, "0", record.EndorsingBankCorrectionIndicatorField())
	require.Equal(t, "A", record.ReturnReasonField())
	require.Equal(t, "                   ", record.UserFieldField())
	require.Equal(t, "                    ", record.reservedField())
}

// testCDAddendumCString validates that a known parsed CheckDetailAddendumC can return to a string of the same value
func testCDAddendumCString(t testing.TB) {
	var line = "2801121042882201809051              Y10A                   0                    "
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

	require.NoError(t, r.parseCheckDetailAddendumC())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumC[0]

	require.Equal(t, line, record.String())
}

// TestCDAddendumCString tests validating that a known parsed CheckDetailAddendumC can return to a string of the
// same value
func TestCDAddendumCString(t *testing.T) {
	testCDAddendumCString(t)
}

// BenchmarkCDAddendumCString benchmarks validating that a known parsed CheckDetailAddendumC
// can return to a string of the same value
func BenchmarkCDAddendumCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumCString(b)
	}
}

// TestCDAddendumCRecordType validation
func TestCDAddendumCRecordType(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.recordType = "00"
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCDAddendumCEndorsingBankRoutingNumber validation
func TestCDAddendumCEndorsingBankRoutingNumber(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankRoutingNumber = "Z"
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankRoutingNumber", e.FieldName)
}

// TestCDAddendumCTruncationIndicator validation
func TestCDAddendumCTruncationIndicator(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.TruncationIndicator = "A"
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
}

// TestCDAddendumCEndorsingBankConversionIndicator validation
func TestCDAddendumCPayeeName(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankConversionIndicator = "10"
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankConversionIndicator", e.FieldName)
}

// TestCDAddendumCEndorsingBankCorrectionIndicator validation
func TestCDAddendumCEndorsingBankCorrectionIndicator(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankCorrectionIndicator = 6
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankCorrectionIndicator", e.FieldName)
}

// TestCDAddendumCBOFDReturnReason validation
func TestCDAddendumCBOFDReturnReason(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.ReturnReason = "®©"
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnReason", e.FieldName)
}

// TestCDAddendumCUserField validation
func TestCDAddendumCUserField(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.UserField = "®©"
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// TestCDAddendumCEndorsingBankIdentifier validation
func TestCDAddendumCEndorsingBankIdentifier(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankIdentifier = 10
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankIdentifier", e.FieldName)
}

// FieldInclusion

// TestCDAddendumCFIRecordType validation
func TestCDAddendumCFIRecordType(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.recordType = ""
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCDAddendumCFIRecordNumber validation
func TestCDAddendumCFIRecordNumber(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.RecordNumber = 0
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "RecordNumber", e.FieldName)
}

// TestCDAddendumCFIEndorsingBankRoutingNumber validation
func TestCDAddendumCFIEndorsingBankRoutingNumber(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankRoutingNumber = ""
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankRoutingNumber", e.FieldName)
}

// TestCDAddendumCFIEndorsingBankRoutingNumberZero validation
func TestCDAddendumCFIEndorsingBankRoutingNumberZero(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankRoutingNumber = "000000000"
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankRoutingNumber", e.FieldName)
}

// TestCDAddendumCFIEndorsingBankRoutingNumberZero validation
func TestCDAddendumCFIEndorsingBankRoutingNumberZeroFRBEnabled(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankRoutingNumber = "000000000"
	t.Setenv(FRBCompatibilityMode, "true")
	err := cdAddendumC.Validate()
	require.Nil(t, err)
}

// TestCDAddendumCFIBOFDEndorsementBusinessDate validation
func TestCDAddendumCFIBOFDEndorsementBusinessDate(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.BOFDEndorsementBusinessDate = time.Time{}
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDEndorsementBusinessDate", e.FieldName)
}

// TestCDAddendumCFIEndorsingBankItemSequenceNumber validation
func TestCDAddendumCFIEndorsingBankItemSequenceNumber(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankItemSequenceNumber = "               "
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankItemSequenceNumber", e.FieldName)
}

// TestCDAddendumCFITruncationIndicator validation
func TestCDAddendumCFITruncationIndicator(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.TruncationIndicator = ""
	err := cdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
}
