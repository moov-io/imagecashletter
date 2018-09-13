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

// mockCheckDetailAddendumA creates a CheckDetailAddendumA
func mockCheckDetailAddendumA() CheckDetailAddendumA {
	cdAddendumA := NewCheckDetailAddendumA()
	cdAddendumA.RecordNumber = 1
	cdAddendumA.ReturnLocationRoutingNumber = "121042882"
	cdAddendumA.BOFDEndorsementDate = time.Now()
	cdAddendumA.BOFDItemSequenceNumber = 1
	cdAddendumA.BOFDAccountNumber = "938383"
	cdAddendumA.BOFDBranchCode = "01"
	cdAddendumA.PayeeName = "Test Payee"
	cdAddendumA.TruncationIndicator = "Y"
	cdAddendumA.BOFDConversionIndicator = "1"
	cdAddendumA.BOFDCorrectionIndicator = 0
	cdAddendumA.UserField = ""
	return cdAddendumA
}

// testMockCheckDetailAddendumA creates a CheckDetailAddendumA
func testMockCheckDetailAddendumA(t testing.TB) {
	cdAddendumA := mockCheckDetailAddendumA()
	if err := cdAddendumA.Validate(); err != nil {
		t.Error("mockBundleHeader does not validate and will break other tests: ", err)
	}
	if cdAddendumA.recordType != "26" {
		t.Error("recordType does not validate and will break other tests")
	}
	if cdAddendumA.RecordNumber != 1 {
		t.Error("RecordNumber does not validate and will break other tests")
	}
	if cdAddendumA.ReturnLocationRoutingNumber != "121042882" {
		t.Error("ReturnLocationRoutingNumber does not validate and will break other tests")
	}
	if cdAddendumA.BOFDItemSequenceNumber != 1 {
		t.Error("BOFDItemSequenceNumber does not validate and will break other tests")
	}
	if cdAddendumA.BOFDAccountNumber != "938383" {
		t.Error("BOFDAccountNumber does not validate and will break other tests")
	}
	if cdAddendumA.BOFDBranchCode != "01" {
		t.Error("BOFDBranchCode does not validate and will break other tests")
	}
	if cdAddendumA.PayeeName != "Test Payee" {
		t.Error("PayeeName does not validate and will break other tests")
	}
	if cdAddendumA.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate and will break other tests")
	}
	if cdAddendumA.BOFDConversionIndicator != "1" {
		t.Error("BOFDConversionIndicator does not validate and will break other tests")
	}
	if cdAddendumA.BOFDCorrectionIndicator != 0 {
		t.Error("BOFDCorrectionIndicator does not validate and will break other tests")
	}
	if cdAddendumA.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
}

// TestMockCheckDetailAddendumA  tests creating a CheckDetailAddendumA
func TestMockCheckDetailAddendumA(t *testing.T) {
	testMockCheckDetailAddendumA(t)
}

// BenchmarkMockCheckDetailAddendumA benchmarks creating a CheckDetailAddendumA
func BenchmarkMockCheckDetailAddendumA(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCheckDetailAddendumA(b)
	}
}

// parseCheckDetailAddendumA validates parsing a CheckDetailAddendumA
func parseCheckDetailAddendumA(t testing.TB) {
	var line = "26112104288220180905000000000000001938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumA(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumA[0]

	if record.recordType != "26" {
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
	if record.BOFDItemSequenceNumberField() != "000000000000001" {
		t.Errorf("BOFDItemSequenceNumber Expected '1' got: %v", record.BOFDItemSequenceNumberField())
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

// TestParseCheckDetailAddendumA test validates parsing a CheckDetailAddendumA
func TestParseCheckDetailAddendumA(t *testing.T) {
	parseCheckDetailAddendumA(t)
}

// BenchmarkParseCheckDetailAddendumA benchmark validates parsing a CheckDetailAddendumA
func BenchmarkParseCheckDetailAddendumA(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseCheckDetailAddendumA(b)
	}
}

// testCDAddendumAString validates that a known parsed CheckDetailAddendumA can return to a string of the same value
func testCDAddendumAString(t testing.TB) {
	var line = "26112104288220180905000000000000001938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumA(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumA[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}

}

// TestCDAddendumAString tests validating that a known parsed CheckDetailAddendumA can return to a string of the
// same value
func TestCDAddendumAString(t *testing.T) {
	testCDAddendumAString(t)
}

// BenchmarkCDAddendumAString benchmarks validating that a known parsed CheckDetail
// can return to a string of the same value
func BenchmarkCDAddendumAString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumAString(b)
	}
}
