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
	bc.BundleItemsCount = 7
	bc.BundleTotalAmount = 100000    // 1000.00
	bc.MICRValidTotalAmount = 100000 // 1000.00
	bc.BundleImagesCount = 1
	bc.UserField = ""
	bc.CreditTotalIndicator = 0
	return bc
}

// TestMockBundleControl creates an BundleControl
func TestMockBundleControl(t *testing.T) {
	bc := mockBundleControl()
	if err := bc.Validate(); err != nil {
		t.Error("mockBundleControl does not validate and will break other tests: ", err)
	}
	if bc.recordType != "70" {
		t.Error("recordType does not validate")
	}
	if bc.BundleItemsCount != 7 {
		t.Error("BundleItemsCount does not validate")
	}
	if bc.BundleTotalAmount != 100000 {
		t.Error("BundleTotalAmount does not validate")
	}
	if bc.MICRValidTotalAmount != 100000 {
		t.Error("MICRValidTotalAmount does not validate")
	}
	if bc.BundleImagesCount != 1 {
		t.Error("BundleImagesCount does not validate")
	}
	if bc.UserField != "" {
		t.Error("UserField does not validate")
	}
	if bc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate")
	}
}

// TestParseBundleControl parses a known BundleControl record string
func TestParseBundleControl(t *testing.T) {
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

// TestBCRecordType validation
func TestBCRecordType(t *testing.T) {
	bc := mockBundleControl()
	bc.recordType = "00"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCUserField validation
func TestBCUserFieldI(t *testing.T) {
	bc := mockBundleControl()
	bc.UserField = "®©"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCCreditTotalIndicator validation
func TestBCCreditTotalIndicator(t *testing.T) {
	bc := mockBundleControl()
	bc.CreditTotalIndicator = 9
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CreditTotalIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCFieldInclusionRecordType validates FieldInclusion
func TestBCFieldInclusionRecordType(t *testing.T) {
	bc := mockBundleControl()
	bc.recordType = ""
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionBundleItemsCount validates FieldInclusion
func TestFieldInclusionBundleItemsCount(t *testing.T) {
	bc := mockBundleControl()
	bc.BundleItemsCount = 0
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionBundleTotalAmount validates FieldInclusion
func TestFieldInclusionBundleTotalAmount(t *testing.T) {
	bc := mockBundleControl()
	bc.BundleTotalAmount = 0
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
