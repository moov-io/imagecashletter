// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

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
	clc.CashLetterItemsCount = 7
	clc.CashLetterTotalAmount = 100000 // 1000.00
	clc.CashLetterImagesCount = 1
	clc.ECEInstitutionName = "Wells Fargo"
	clc.SettlementDate = time.Now()
	clc.CreditTotalIndicator = 0
	return clc
}

// TestMockCashLetterControl creates a CashLetterControl
func TestMockCashLetterControl(t *testing.T) {
	clc := mockCashLetterControl()
	if err := clc.Validate(); err != nil {
		t.Error("mockCashLetterControl does not validate and will break other tests: ", err)
	}
	if clc.recordType != "90" {
		t.Error("recordType does not validate")
	}
	if clc.CashLetterBundleCount != 1 {
		t.Error("CashLetterBundleCount does not validate")
	}
	if clc.CashLetterItemsCount != 7 {
		t.Error("CashLetterItemsCount does not validate")
	}
	if clc.CashLetterTotalAmount != 100000 {
		t.Error("CashLetterTotalAmount does not validate")
	}
	if clc.CashLetterImagesCount != 1 {
		t.Error("CashLetterImagesCount does not validate")
	}
	if clc.ECEInstitutionName != "Wells Fargo" {
		t.Error("ImmediateOriginContactName does not validate")
	}
	if clc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate")
	}
}

// TestParseCashLetterControl parses a known CashLetterControl record string
func TestParseCashLetterControl(t *testing.T) {
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

// TestCLCRecordType validation
func TestCLCRecordType(t *testing.T) {
	clc := mockCashLetterControl()
	clc.recordType = "00"
	if err := clc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestECEInstitutionName validation
func TestECEInstitutionName(t *testing.T) {
	clc := mockCashLetterControl()
	clc.ECEInstitutionName = "®©"
	if err := clc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ECEInstitutionName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCLCCreditTotalIndicator validation
func TestCLCCreditTotalIndicator(t *testing.T) {
	clc := mockCashLetterControl()
	clc.CreditTotalIndicator = 9
	if err := clc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CreditTotalIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCLCFieldInclusionRecordType validates FieldInclusion
func TestCLCFieldInclusionRecordType(t *testing.T) {
	clc := mockCashLetterControl()
	clc.recordType = ""
	if err := clc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCashLetterItemsCount validates FieldInclusion
func TestFieldInclusionCashLetterItemsCount(t *testing.T) {
	clc := mockCashLetterControl()
	clc.CashLetterItemsCount = 0
	if err := clc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CashLetterItemsCount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCashLetterTotalAmount validates FieldInclusion
func TestFieldInclusionCashLetterTotalAmount(t *testing.T) {
	clc := mockCashLetterControl()
	clc.CashLetterTotalAmount = 0
	if err := clc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CashLetterTotalAmount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionSettlementDate validates FieldInclusion
func TestFieldInclusionRecordTypeSettlementDate(t *testing.T) {
	clc := mockCashLetterControl()
	clc.SettlementDate = time.Time{}
	if err := clc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "SettlementDate" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCashLetterControlRuneCountInString validates RuneCountInString
func TestCashLetterControlRuneCountInString(t *testing.T) {
	clc := NewCashLetterControl()
	var line = "90"
	clc.Parse(line)

	if clc.CashLetterBundleCount != 0 {
		t.Error("Parsed with an invalid RuneCountInString")
	}
}
