// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
)

// mockBundleControl creates a BundleControl
func mockBundleControl() *BundleControl {
	bc := NewBundleControl()
	// BundleItemsCount - CheckDetail
	// ToDo: CheckDetailAddendum* and ImageView*
	bc.BundleItemsCount = 1
	bc.BundleTotalAmount = 100000 // 1000.00
	// ToDo: CheckDetail
	bc.MICRValidTotalAmount = 0
	// ToDo: ImageView*
	bc.BundleImagesCount = 0
	bc.UserField = ""
	bc.CreditTotalIndicator = 0
	return bc
}

// testMockBundleControl creates an ICL BundleControl
func testMockBundleControl(t testing.TB) {
	bc := mockBundleControl()
	if err := bc.Validate(); err != nil {
		t.Error("mockBundleControl does not validate and will break other tests: ", err)
	}
	if bc.recordType != "70" {
		t.Error("recordType does not validate and will break other tests")
	}
	if bc.BundleItemsCount != 1 {
		t.Error("BundleItemsCount does not validate and will break other tests")
	}
	if bc.BundleTotalAmount != 100000 {
		t.Error("BundleTotalAmount does not validate and will break other tests")
	}
	if bc.MICRValidTotalAmount != 0 {
		t.Error("MICRValidTotalAmount does not validate and will break other tests")
	}
	if bc.BundleImagesCount != 0 {
		t.Error("BundleImagesCount does not validate and will break other tests")
	}
	if bc.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
	if bc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate and will break other tests")
	}
}

// TestMockBundleControl tests creating an ICL BundleControl
func TestMockBundleControl(t *testing.T) {
	testMockBundleControl(t)
}

// BenchmarkMockBundleControl benchmarks creating an ICL BundleControl
func BenchmarkMockBundleControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockBundleControl(b)
	}
}

// testParseBundleControl parses a known BundleControl record string
func testParseBundleControl(t testing.TB) {
	var line = "70000100000010000000000000000000000                    0                        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	err := r.parseBundleControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.BundleControl

	if record.recordType != "70" {
		t.Errorf("RecordType Expected '70' got: %v", record.recordType)
	}
	if record.BundleItemsCountField() != "0001" {
		t.Errorf("BundleItemsCountCount Expected '0001' got: %v", record.BundleItemsCountField())
	}
	if record.BundleTotalAmountField() != "000000100000" {
		t.Errorf("BundleTotalAmount Expected '000000100000' got: %v", record.BundleTotalAmountField())
	}
	if record.MICRValidTotalAmountField() != "000000000000" {
		t.Errorf("MICRValidTotalAmount Expected '000000000000' got: %v", record.MICRValidTotalAmountField())
	}
	if record.BundleImagesCountField() != "00000" {
		t.Errorf("BundleImagesCount Expected '00000' got: %v", record.BundleImagesCountField())
	}
	if record.UserFieldField() != "                    " {
		t.Errorf("UserField Expected '                    ' got: %v", record.UserFieldField())
	}
	if record.CreditTotalIndicatorField() != "0" {
		t.Errorf("CreditTotalIndicator Expected '0' got: %v", record.CreditTotalIndicatorField())
	}
	if record.reservedField() != "                        " {
		t.Errorf("Reserved Expected '                        ' got: %v", record.reservedField())
	}
}

// TestParseBundleControl tests parsing a known BundleControl record string
func TestParseBundleControl(t *testing.T) {
	testParseBundleControl(t)
}

// BenchmarkParseBundleControl benchmarks parsing a known BundleControl record string
func BenchmarkParseBundleControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseBundleControl(b)
	}
}

// testBCString validates that a known parsed BundleControl can be return to a string of the same value
func testBCString(t testing.TB) {
	var line = "70000100000010000000000000000000000                    0                        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	err := r.parseBundleControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.BundleControl
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}

// TestBCString tests validating that a known parsed BundleControl can be return to a string of the same value
func TestBCString(t *testing.T) {
	testBCString(t)
}

// BenchmarkBCString benchmarks validating that a known parsed BundleControl can be return to a string of the same value
func BenchmarkBCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCString(b)
	}
}
