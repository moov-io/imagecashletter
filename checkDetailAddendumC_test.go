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

// mockCheckDetailAddendumC creates a CheckDetailAddendumC
func mockCheckDetailAddendumC() CheckDetailAddendumC {
	cdAddendumC := NewCheckDetailAddendumC()
	cdAddendumC.RecordNumber = 1
	cdAddendumC.EndorsingBankRoutingNumber = "121042882"
	cdAddendumC.BOFDEndorsementBusinessDate = time.Now()
	cdAddendumC.EndorsingItemSequenceNumber = 1
	cdAddendumC.TruncationIndicator = "Y"
	cdAddendumC.EndorsingConversionIndicator = "1"
	cdAddendumC.EndorsingCorrectionIndicator = 0
	cdAddendumC.ReturnReason = "A"
	cdAddendumC.UserField = ""
	cdAddendumC.EndorsingBankIdentifier = 0
	return cdAddendumC
}

// testMockCheckDetailAddendumC creates a CheckDetailAddendumC
func testMockCheckDetailAddendumC(t testing.TB) {
	cdAddendumC := mockCheckDetailAddendumC()
	if err := cdAddendumC.Validate(); err != nil {
		t.Error("mockBundleHeader does not validate and will break other tests: ", err)
	}
	if cdAddendumC.recordType != "28" {
		t.Error("recordType does not validate and will break other tests")
	}
	if cdAddendumC.RecordNumber != 1 {
		t.Error("RecordNumber does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingBankRoutingNumber != "121042882" {
		t.Error("EndorsingBankRoutingNumber does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingItemSequenceNumber != 1 {
		t.Error("EndorsingItemSequenceNumber does not validate and will break other tests")
	}
	if cdAddendumC.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate and will break other tests")
	}
	if cdAddendumC.ReturnReason != "A" {
		t.Error("ReturnReason does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingConversionIndicator != "1" {
		t.Error("EndorsingConversionIndicator does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingCorrectionIndicator != 0 {
		t.Error("EndorsingCorrectionIndicator does not validate and will break other tests")
	}
	if cdAddendumC.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingBankIdentifier != 0 {
		t.Error("EndorsingBankIdentifier does not validate and will break other tests")
	}
}

// TestMockCheckDetailAddendumC  tests creating an ICL CheckDetailAddendumC
func TestMockCheckDetailAddendumC(t *testing.T) {
	testMockCheckDetailAddendumC(t)
}

// BenchmarkMockCheckDetailAddendumC benchmarks creating a CheckDetailAddendumC
func BenchmarkMockCheckDetailAddendumC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCheckDetailAddendumC(b)
	}
}

// parseCheckDetailAddendumC validates parsing a CheckDetailAddendumC
func parseCheckDetailAddendumC(t testing.TB) {
	var line = "280112104288220180905000000000000001Y10A                   0                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumC(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumC[0]

	if record.recordType != "28" {
		t.Errorf("RecordType Expected '26' got: %v", record.recordType)
	}
	if record.RecordNumberField() != "01" {
		t.Errorf("RecordNumber Expected '01' got: %v", record.RecordNumberField())
	}

	if record.EndorsingBankRoutingNumberField() != "121042882" {
		t.Errorf("EndorsingBankRoutingNumbeRoutingNumber Expected '121042882' got: %v",
			record.EndorsingBankRoutingNumberField())
	}
	if record.BOFDEndorsementBusinessDateField() != "20180905" {
		t.Errorf("BOFDEndorsementBusinessDate Expected '20180905' got: %v",
			record.BOFDEndorsementBusinessDateField())
	}
	if record.EndorsingItemSequenceNumberField() != "000000000000001" {
		t.Errorf("EndorsingItemSequenceNumber Expected '1' got: %v", record.EndorsingItemSequenceNumberField())
	}
	if record.TruncationIndicatorField() != "Y" {
		t.Errorf("TruncationIndicator Expected 'Y' got: %v", record.TruncationIndicatorField())
	}
	if record.EndorsingConversionIndicatorField() != "1" {
		t.Errorf("EndorsingConversionIndicator  Expected '1' got: %v", record.EndorsingConversionIndicatorField())
	}
	if record.EndorsingCorrectionIndicatorField() != "0" {
		t.Errorf("EndorsingCorrectionIndicator Expected '0' got: %v", record.EndorsingCorrectionIndicatorField())
	}
	if record.ReturnReasonField() != "A" {
		t.Errorf("ReturnReason  Expected 'A' got: %v", record.ReturnReasonField())
	}
	if record.UserFieldField() != "                   " {
		t.Errorf("UserField Expected '                   ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != "                    " {
		t.Errorf("reserved Expected '                    ' got: %v", record.reservedField())
	}
}

// TestParseCheckDetailAddendumC test validates parsing a CheckDetailAddendumC
func TestParseCheckDetailAddendumC(t *testing.T) {
	parseCheckDetailAddendumC(t)
}

// BenchmarkParseCheckDetailAddendumC benchmark validates parsing a CheckDetailAddendumC
func BenchmarkParseCheckDetailAddendumC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseCheckDetailAddendumC(b)
	}
}

// testCDAddendumCString validates that a known parsed CheckDetailAddendumC can return to a string of the same value
func testCDAddendumCString(t testing.TB) {
	var line = "280112104288220180905000000000000001Y10A                   0                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumC(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumC[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCDAddendumCString tests validating that a known parsed CheckDetailAddendumC can return to a string of the
// same value
func TestCDAddendumCString(t *testing.T) {
	testCDAddendumCString(t)
}

// BenchmarkCDAddendumCString benchmarks validating that a known parsed CheckDetail
// can return to a string of the same value
func BenchmarkCDAddendumCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumCString(b)
	}
}
