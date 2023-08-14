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

// mockReturnDetailAddendumD creates a ReturnDetailAddendumD
func mockReturnDetailAddendumD() ReturnDetailAddendumD {
	rdAddendumD := NewReturnDetailAddendumD()
	rdAddendumD.RecordNumber = 1
	rdAddendumD.EndorsingBankRoutingNumber = "121042882"
	rdAddendumD.BOFDEndorsementBusinessDate = time.Now()
	rdAddendumD.EndorsingBankItemSequenceNumber = "1              "
	rdAddendumD.TruncationIndicator = "Y"
	rdAddendumD.EndorsingBankConversionIndicator = "1"
	rdAddendumD.EndorsingBankCorrectionIndicator = 0
	rdAddendumD.ReturnReason = "A"
	rdAddendumD.UserField = ""
	rdAddendumD.EndorsingBankIdentifier = 0
	return rdAddendumD
}

// mockReturnDetailAddendumDWithoutEndorsingBankItemSequenceNumber creates a ReturnDetailAddendumD
func mockReturnDetailAddendumDWithoutEndorsingBankItemSequenceNumber() ReturnDetailAddendumD {
	rdAddendumD := NewReturnDetailAddendumD()
	rdAddendumD.RecordNumber = 1
	rdAddendumD.EndorsingBankRoutingNumber = "121042882"
	rdAddendumD.BOFDEndorsementBusinessDate = time.Now()
	rdAddendumD.TruncationIndicator = "Y"
	rdAddendumD.EndorsingBankConversionIndicator = "1"
	rdAddendumD.EndorsingBankCorrectionIndicator = 0
	rdAddendumD.ReturnReason = "A"
	rdAddendumD.UserField = ""
	rdAddendumD.EndorsingBankIdentifier = 0
	return rdAddendumD
}

func TestReturnDetailAddendumDParseErr(t *testing.T) {
	var r ReturnDetailAddendumD
	r.Parse("ASdasdas")
	require.Equal(t, 0, r.RecordNumber)
}

// TestMockReturnDetailAddendumD creates a ReturnDetailAddendumD
func TestMockReturnDetailAddendumD(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	require.NoError(t, rdAddendumD.Validate())
	require.Equal(t, "35", rdAddendumD.recordType)
	require.Equal(t, 1, rdAddendumD.RecordNumber)
	require.Equal(t, "121042882", rdAddendumD.EndorsingBankRoutingNumber)
	require.Equal(t, "1              ", rdAddendumD.EndorsingBankItemSequenceNumber)
	require.Equal(t, "Y", rdAddendumD.TruncationIndicator)
	require.Equal(t, "A", rdAddendumD.ReturnReason)
	require.Equal(t, "1", rdAddendumD.EndorsingBankConversionIndicator)
	require.Equal(t, 0, rdAddendumD.EndorsingBankCorrectionIndicator)
	require.Equal(t, "", rdAddendumD.UserField)
	require.Equal(t, 0, rdAddendumD.EndorsingBankIdentifier)
}

// TestMockReturnDetailAddendumD creates a ReturnDetailAddendumD
func TestMockReturnDetailAddendumDWithoutEndorsingBankItemSequenceNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumDWithoutEndorsingBankItemSequenceNumber()
	require.NoError(t, rdAddendumD.Validate())
	require.Equal(t, "35", rdAddendumD.recordType)
	require.Equal(t, 1, rdAddendumD.RecordNumber)
	require.Equal(t, "121042882", rdAddendumD.EndorsingBankRoutingNumber)
	require.Equal(t, "Y", rdAddendumD.TruncationIndicator)
	require.Equal(t, "A", rdAddendumD.ReturnReason)
	require.Equal(t, "1", rdAddendumD.EndorsingBankConversionIndicator)
	require.Equal(t, 0, rdAddendumD.EndorsingBankCorrectionIndicator)
	require.Equal(t, "", rdAddendumD.UserField)
	require.Equal(t, 0, rdAddendumD.EndorsingBankIdentifier)
}

// TestParseReturnDetailAddendumD validates parsing a ReturnDetailAddendumD
func TestParseReturnDetailAddendumD(t *testing.T) {
	var line = "3501121042882201809051              Y10A                   0                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	cd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(cd)

	require.NoError(t, r.parseReturnDetailAddendumD())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumD[0]

	require.Equal(t, "35", record.recordType)
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

// TestParseReturnDetailAddendumDWithoutEndorsingBankItemSequenceNumber validates parsing a ReturnDetailAddendumD
func TestParseReturnDetailAddendumDWithoutEndorsingBankItemSequenceNumber(t *testing.T) {
	var line = "350112104288220180905               Y10A                   0                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	cd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(cd)

	require.NoError(t, r.parseReturnDetailAddendumD())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumD[0]

	require.Equal(t, "35", record.recordType)
	require.Equal(t, "01", record.RecordNumberField())

	require.Equal(t, "121042882", record.EndorsingBankRoutingNumberField())
	require.Equal(t, "20180905", record.BOFDEndorsementBusinessDateField())
	require.Equal(t, "               ", record.EndorsingBankItemSequenceNumberField())
	require.Equal(t, "Y", record.TruncationIndicatorField())
	require.Equal(t, "1", record.EndorsingBankConversionIndicatorField())
	require.Equal(t, "0", record.EndorsingBankCorrectionIndicatorField())
	require.Equal(t, "A", record.ReturnReasonField())
	require.Equal(t, "                   ", record.UserFieldField())
	require.Equal(t, "                    ", record.reservedField())
}

// testRDAddendumDString validates that a known parsed ReturnDetailAddendumD can return to a string of the same value
func testRDAddendumDString(t testing.TB) {
	var line = "3501121042882201809051              Y10A                   0                    "
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

	require.NoError(t, r.parseReturnDetailAddendumD())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumD[0]

	require.Equal(t, line, record.String())
}

// TestRDAddendumDString tests validating that a known parsed ReturnDetailAddendumD can return to a string of the
// same value
func TestRDAddendumDString(t *testing.T) {
	testRDAddendumDString(t)
}

// BenchmarkRDAddendumDString benchmarks validating that a known parsed ReturnDetailAddendumD
// can return to a string of the same value
func BenchmarkRDAddendumDString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDAddendumDString(b)
	}
}

// TestRDAddendumDRecordType validation
func TestRDAddendumDRecordType(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.recordType = "00"
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDAddendumDReturnLocationRoutingNumber validation
func TestRDAddendumDReturnLocationRoutingNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankRoutingNumber = "X"
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankRoutingNumber", e.FieldName)
}

// TestRDAddendumDTruncationIndicator validation
func TestRDAddendumDTruncationIndicator(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.TruncationIndicator = "A"
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
}

// TestRDAddendumDBOFDConversionIndicator validation
func TestRDAddendumDBOFDConversionIndicator(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankConversionIndicator = "99"
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankConversionIndicator", e.FieldName)
}

// TestRDAddendumDBOFDCorrectionIndicator validation
func TestRDAddendumDBOFDCorrectionIndicator(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankCorrectionIndicator = 10
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankCorrectionIndicator", e.FieldName)
}

// TestRDAddendumDReturnReason validation
func TestRDAddendumDReturnReason(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.ReturnReason = "--"
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ReturnReason", e.FieldName)
}

// TestRDAddendumDUserField validation
func TestRDAddendumDUserField(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.UserField = "®©"
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// TestRDAddendumDEndorsingBankIdentifier validation
func TestRDAddendumDEndorsingBankIdentifier(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankIdentifier = 9
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankIdentifier", e.FieldName)
}

// Field Inclusion

// TestRDAddendumDFIRecordType validation
func TestRDAddendumDFIRecordType(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.recordType = ""
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDAddendumDFIRecordNumber validation
func TestRDAddendumDFIRecordNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.RecordNumber = 0
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "RecordNumber", e.FieldName)
}

// TestRDAddendumDFIReturnLocationRoutingNumber validation
func TestRDAddendumDFIReturnLocationRoutingNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankRoutingNumber = ""
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankRoutingNumber", e.FieldName)
}

// TestRDAddendumDFIReturnLocationRoutingNumberZero validation
func TestRDAddendumDFIReturnLocationRoutingNumberZero(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankRoutingNumber = "000000000"
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankRoutingNumber", e.FieldName)
}

// TestRDAddendumDFIBOFDEndorsementDate validation
func TestRDAddendumDFIBOFDEndorsementDate(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.BOFDEndorsementBusinessDate = time.Time{}
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDEndorsementBusinessDate", e.FieldName)
}

// TestRDAddendumDFIBOFDItemSequenceNumber validation
func TestRDAddendumDFIBOFDItemSequenceNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankItemSequenceNumber = "         s     "
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EndorsingBankItemSequenceNumber", e.FieldName)
}

func TestRDAddendumDFIBOFDItemSequenceNumber_emptyAllowed(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankItemSequenceNumber = "               "
	require.NoError(t, rdAddendumD.Validate())
}

// TestRDAddendumDFITruncationIndicator validation
func TestRDAddendumDFITruncationIndicator(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.TruncationIndicator = ""
	err := rdAddendumD.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TruncationIndicator", e.FieldName)
}

// TestRDAddendumDRuneCountInString validates RuneCountInString
func TestRDAddendumDRuneCountInString(t *testing.T) {
	rdAddendumD := NewReturnDetailAddendumD()
	var line = "35"
	rdAddendumD.Parse(line)

	require.Equal(t, "", rdAddendumD.EndorsingBankRoutingNumber)
}
