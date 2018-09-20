// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
	"time"
)

// mockCashLetterControl creates a CashLetterControl
func mockCashLetterControl() *CashLetterControl {
	clc := NewCashLetterControl()
	clc.CashLetterBundleCount = 1
	// CashLetterItemsCount - CheckDetail
	// ToDo: CheckDetailAddendum* and ImageView*
	clc.CashLetterItemsCount = 1
	clc.CashLetterTotalAmount = 100000 // 1000.00
	// ToDo: ImageView*
	clc.CashLetterImagesCount = 0
	clc.ECEInstitutionName = "Wells Fargo"
	clc.SettlementDate = time.Now()
	clc.CreditTotalIndicator = 0
	return clc
}

// testMockCashLetterControl creates a CashLetterControl
func testMockCashLetterControl(t testing.TB) {
	clc := mockCashLetterControl()
	if err := clc.Validate(); err != nil {
		t.Error("mockCashLetterControl does not validate and will break other tests: ", err)
	}
	if clc.recordType != "90" {
		t.Error("recordType does not validate and will break other tests")
	}
	if clc.CashLetterBundleCount != 1 {
		t.Error("CashLetterBundleCount does not validate and will break other tests")
	}
	if clc.CashLetterItemsCount != 1 {
		t.Error("CashLetterItemsCount does not validate and will break other tests")
	}
	if clc.CashLetterTotalAmount != 100000 {
		t.Error("CashLetterTotalAmount does not validate and will break other tests")
	}
	if clc.CashLetterImagesCount != 0 {
		t.Error("CashLetterImagesCount does not validate and will break other tests")
	}
	if clc.ECEInstitutionName != "Wells Fargo" {
		t.Error("ImmediateOriginContactName does not validate and will break other tests")
	}
	if clc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate and will break other tests")
	}
}

// TestMockCashLetterControl tests creating a CashLetterControl
func TestMockCashLetterControl(t *testing.T) {
	testMockCashLetterControl(t)
}

// BenchmarkMockCashLetterControl benchmarks creating a CashLetterControl
func BenchmarkMockCashLetterControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCashLetterControl(b)
	}
}

// testParseCashLetterControl parses a known CashLetterControl record string
func testParseCashLetterControl(t testing.TB) {
	var line = "900000010000000100000000100000000000000Wells Fargo       201809050              "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	err := r.parseCashLetterControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.CashLetterControl

	if record.recordType != "90" {
		t.Errorf("RecordType Expected '90' got: %v", record.recordType)
	}
	if record.CashLetterBundleCountField() != "000001" {
		t.Errorf("CashLetterBundleCount Expected '000001' got: %v", record.CashLetterBundleCountField())
	}
	if record.CashLetterItemsCountField() != "00000001" {
		t.Errorf("CashLetterItemsCount Expected '00000001' got: %v", record.CashLetterItemsCountField())
	}
	if record.CashLetterTotalAmountField() != "00000000100000" {
		t.Errorf("CashLetterTotalAmount Expected '00000000100000' got: %v", record.CashLetterTotalAmountField())
	}
	if record.CashLetterImagesCountField() != "000000000" {
		t.Errorf("CashLetterImagesCount Expected '000000000' got: %v", record.CashLetterImagesCountField())
	}
	if record.ECEInstitutionNameField() != "Wells Fargo       " {
		t.Errorf("ECEInstitutionName Expected 'Wells Fargo       ' got: %v", record.ECEInstitutionNameField())
	}
	if record.SettlementDateField() != "20180905" {
		t.Errorf("SettlementDate Expected '20180905' got: %v", record.SettlementDateField())
	}
	if record.CreditTotalIndicatorField() != "0" {
		t.Errorf("CreditTotalIndicator Expected '0' got: %v", record.CreditTotalIndicatorField())
	}
	if record.reservedField() != "              " {
		t.Errorf("Reserved Expected '              ' got: %v", record.reservedField())
	}
}

// TestParseCashLetterControl tests parsing a known CashLetterControl record string
func TestParseCashLetterControl(t *testing.T) {
	testParseCashLetterControl(t)
}

// BenchmarkParseCashLetterControl benchmarks parsing a known CashLetterControl record string
func BenchmarkParseCashLetterControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseCashLetterControl(b)
	}
}

// testCLCString validates that a known parsed CashLetterControl can be return to a string of the same value
func testCLCString(t testing.TB) {
	var line = "900000010000000100000000100000000000000Wells Fargo       201809050              "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	err := r.parseCashLetterControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.CashLetterControl
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}

// TestCLCString tests validating that a known parsed CashLetterControl can be return to a string of the same value
func TestCLCString(t *testing.T) {
	testCLCString(t)
}

// BenchmarkCLCString benchmarks validating that a known parsed CashLetterControl can be return to a string of the same value
func BenchmarkCLCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCLCString(b)
	}
}
