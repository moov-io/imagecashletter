// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
	"time"
)

// mockReturnDetailAddendumA creates a ReturnDetailAddendumA
func mockReturnDetailAddendumA() ReturnDetailAddendumA {
	RDAddendumA := NewReturnDetailAddendumA()
	RDAddendumA.RecordNumber = 1
	RDAddendumA.ReturnLocationRoutingNumber = "121042882"
	RDAddendumA.BOFDEndorsementDate = time.Now()
	RDAddendumA.BOFDItemSequenceNumber = "1              "
	RDAddendumA.BOFDAccountNumber = "938383"
	RDAddendumA.BOFDBranchCode = "01"
	RDAddendumA.PayeeName = "Test Payee"
	RDAddendumA.TruncationIndicator = "Y"
	RDAddendumA.BOFDConversionIndicator = "1"
	RDAddendumA.BOFDCorrectionIndicator = 0
	RDAddendumA.UserField = ""
	return RDAddendumA
}

// testMockReturnDetailAddendumA creates a ReturnDetailAddendumA
func testMockReturnDetailAddendumA(t testing.TB) {
	RDAddendumA := mockReturnDetailAddendumA()
	if err := RDAddendumA.Validate(); err != nil {
		t.Error("mockBundleHeader does not validate and will break other tests: ", err)
	}
	if RDAddendumA.recordType != "32" {
		t.Error("recordType does not validate and will break other tests")
	}
	if RDAddendumA.RecordNumber != 1 {
		t.Error("RecordNumber does not validate and will break other tests")
	}
	if RDAddendumA.ReturnLocationRoutingNumber != "121042882" {
		t.Error("ReturnLocationRoutingNumber does not validate and will break other tests")
	}
	if RDAddendumA.BOFDItemSequenceNumber != "1              " {
		t.Error("BOFDItemSequenceNumber does not validate and will break other tests")
	}
	if RDAddendumA.BOFDAccountNumber != "938383" {
		t.Error("BOFDAccountNumber does not validate and will break other tests")
	}
	if RDAddendumA.BOFDBranchCode != "01" {
		t.Error("BOFDBranchCode does not validate and will break other tests")
	}
	if RDAddendumA.PayeeName != "Test Payee" {
		t.Error("PayeeName does not validate and will break other tests")
	}
	if RDAddendumA.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate and will break other tests")
	}
	if RDAddendumA.BOFDConversionIndicator != "1" {
		t.Error("BOFDConversionIndicator does not validate and will break other tests")
	}
	if RDAddendumA.BOFDCorrectionIndicator != 0 {
		t.Error("BOFDCorrectionIndicator does not validate and will break other tests")
	}
	if RDAddendumA.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
}

// TestMockReturnDetailAddendumA  tests creating a ReturnDetailAddendumA
func TestMockReturnDetailAddendumA(t *testing.T) {
	testMockReturnDetailAddendumA(t)
}

// BenchmarkMockReturnDetailAddendumA benchmarks creating a ReturnDetailAddendumA
func BenchmarkMockReturnDetailAddendumA(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockReturnDetailAddendumA(b)
	}
}

// parseReturnDetailAddendumA validates parsing a ReturnDetailAddendumA
func parseReturnDetailAddendumA(t testing.TB) {
	var line = "321121042882201809051              938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewReturnBundle(bh)
	r.currentCashLetter.AddReturnBundle(rb)
	r.addCurrentReturnBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentReturnBundle.AddReturnDetail(rd)

	if err := r.parseReturnDetailAddendumA(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentReturnBundle.GetReturns()[0].ReturnDetailAddendumA[0]

	if record.recordType != "32" {
		t.Errorf("RecordType Expected '26' got: %v", record.recordType)
	}
	if record.RecordNumberField() != "1" {
		t.Errorf("RecordNumber Expected '1' got: %v", record.RecordNumberField())
	}
	if record.ReturnLocationRoutingNumberField() != "121042882" {
		t.Errorf("ReturnLocationRoutingNumber Expected '121042882' got: %v", record.ReturnLocationRoutingNumberField())
	}
	if record.BOFDEndorsementDateField() != "20180905" {
		t.Errorf("BOFDEndorsementDate Expected '20180905' got: %v", record.BOFDEndorsementDateField())
	}
	if record.BOFDItemSequenceNumberField() != "1              " {
		t.Errorf("BOFDItemSequenceNumber Expected '1               ' got: %v", record.BOFDItemSequenceNumberField())
	}
	if record.BOFDAccountNumberField() != "938383            " {
		t.Errorf("BOFDAccountNumber Expected '938383            ' got: %v", record.BOFDAccountNumberField())
	}
	if record.BOFDBranchCodeField() != "01   " {
		t.Errorf("BOFDBranchCode Expected '01   ' got: %v", record.BOFDBranchCodeField())
	}
	if record.PayeeNameField() != "Test Payee     " {
		t.Errorf("PayeeName Expected 'Test Payee     ' got: %v", record.PayeeNameField())
	}
	if record.TruncationIndicatorField() != "Y" {
		t.Errorf("TruncationIndicator Expected 'Y' got: %v", record.TruncationIndicatorField())
	}
	if record.BOFDConversionIndicatorField() != "1" {
		t.Errorf("BOFDConversionIndicator Expected '1' got: %v", record.BOFDConversionIndicatorField())
	}
	if record.BOFDCorrectionIndicatorField() != "0" {
		t.Errorf("BOFDCorrectionIndicator Expected '0' got: %v", record.BOFDCorrectionIndicatorField())
	}
	if record.UserFieldField() != " " {
		t.Errorf("UserField Expected ' ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != "   " {
		t.Errorf("reserved Expected '   ' got: %v", record.reservedField())
	}
}

// TestParseReturnDetailAddendumA test validates parsing a ReturnDetailAddendumA
func TestParseReturnDetailAddendumA(t *testing.T) {
	parseReturnDetailAddendumA(t)
}

// BenchmarkParseReturnDetailAddendumA benchmark validates parsing a ReturnDetailAddendumA
func BenchmarkParseReturnDetailAddendumA(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseReturnDetailAddendumA(b)
	}
}

// testRDAddendumAString validates that a known parsed ReturnDetailAddendumA can return to a string of the same value
func testRDAddendumAString(t testing.TB) {
	var line = "321121042882201809051              938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewReturnBundle(bh)
	r.currentCashLetter.AddReturnBundle(rb)
	r.addCurrentReturnBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentReturnBundle.AddReturnDetail(rd)

	if err := r.parseReturnDetailAddendumA(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentReturnBundle.GetReturns()[0].ReturnDetailAddendumA[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}

}

// TestRDAddendumAString tests validating that a known parsed ReturnDetailAddendumA can return to a string of the
// same value
func TestRDAddendumAString(t *testing.T) {
	testRDAddendumAString(t)
}

// BenchmarkRDAddendumAString benchmarks validating that a known parsed ReturnDetail
// can return to a string of the same value
func BenchmarkRDAddendumAString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDAddendumAString(b)
	}
}
