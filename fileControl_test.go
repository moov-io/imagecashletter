// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"log"
	"strings"
	"testing"
)

// mockFileControl creates a FileControl
func mockFileControl() FileControl {
	fc := NewFileControl()
	fc.CashLetterCount = 1
	fc.TotalRecordCount = 7
	fc.TotalItemCount = 1
	fc.FileTotalAmount = 100000 //1000.00
	fc.ImmediateOriginContactName = "Contact Name"
	fc.ImmediateOriginContactPhoneNumber = "5558675552"
	fc.CreditTotalIndicator = 0
	return fc
}

// TestMockFileControl creates a FileControl
func TestMockFileControl(t *testing.T) {
	fc := mockFileControl()
	if err := fc.Validate(); err != nil {
		t.Error("mockFileControl does not validate and will break other tests: ", err)
	}
	if fc.recordType != "99" {
		t.Error("recordType does not validate")
	}
	if fc.CashLetterCount != 1 {
		t.Error("CashLetterCount does not validate")
	}
	if fc.TotalRecordCount != 7 {
		t.Error("TotalRecordCount does not validate")
	}
	if fc.TotalItemCount != 1 {
		t.Error("TotalItemCount does not validate")
	}
	if fc.FileTotalAmount != 100000 {
		t.Error("FileTotalAmount does not validate")
	}
	if fc.ImmediateOriginContactName != "Contact Name" {
		t.Error("ImmediateOriginContactName does not validate")
	}
	if fc.ImmediateOriginContactPhoneNumber != "5558675552" {
		t.Error("ImmediateOriginContactPhoneNumber does not validate")
	}
	if fc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate")
	}
}

// testParseFileControl parses a known FileControl record string
func testParseFileControl(t testing.TB) {
	var line = "9900000100000007000000010000000000100000Contact Name  55586755520               "
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseFileControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.File.Control

	if record.recordType != "99" {
		t.Errorf("RecordType Expected '99' got: %v", record.recordType)
	}
	if record.CashLetterCountField() != "000001" {
		t.Errorf("CashLetterCount Expected '000001' got: %v", record.CashLetterCountField())
	}
	if record.TotalRecordCountField() != "00000007" {
		t.Errorf("TotalRecordCount Expected '00000007' got: %v", record.TotalRecordCountField())
	}
	if record.TotalItemCountField() != "00000001" {
		t.Errorf("TotalItemCount Expected '00000001' got: %v", record.TotalItemCountField())
	}
	if record.FileTotalAmountField() != "0000000000100000" {
		t.Errorf("FileTotalAmount Expected '0000000000100000' got: %v", record.FileTotalAmountField())
	}
	if record.ImmediateOriginContactNameField() != "Contact Name  " {
		t.Errorf("ImmediateOriginContactName Expected 'Contact Name  ' got: %v", record.ImmediateOriginContactNameField())
	}
	if record.ImmediateOriginContactPhoneNumberField() != "5558675552" {
		t.Errorf("ImmediateOriginContactPhoneNumber Expected '5558675552' got: %v", record.ImmediateOriginContactPhoneNumberField())
	}
	if record.CreditTotalIndicatorField() != "0" {
		t.Errorf("CreditTotalIndicator Expected '0' got: %v", record.CreditTotalIndicatorField())
	}
	if record.reservedField() != "               " {
		t.Errorf("Reserved Expected '               ' got: %v", record.reservedField())
	}
}

// TestParseFileControl tests parsing a known FileControl record string
func TestParseFileControl(t *testing.T) {
	testParseFileControl(t)
}

// BenchmarkParseFileControl benchmarks parsing a known FileControl record string
func BenchmarkParseFileControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseFileControl(b)
	}
}

// testFCString validates that a known parsed FileControl can be return to a string of the same value
func testFCString(t testing.TB) {
	var line = "9900000100000007000000010000000000100000Contact Name  55586755520               "
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseFileControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.File.Control
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}

// TestFCString tests validating that a known parsed FileControl can be return to a string of the same value
func TestFCString(t *testing.T) {
	testFCString(t)
}

// BenchmarkFCString benchmarks validating that a known parsed FileControl can be return to a string of the same value
func BenchmarkFCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCString(b)
	}
}

// TestFCRecordType validation
func TestFCRecordType(t *testing.T) {
	fc := mockFileControl()
	fc.recordType = "00"
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestImmediateOriginContactName validation
func TestImmediateOriginContactName(t *testing.T) {
	fc := mockFileControl()
	fc.ImmediateOriginContactName = "®©"
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateOriginContactName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestImmediateOriginContactPhoneNumber validation
func TestImmediateOriginContactPhoneNumber(t *testing.T) {
	fc := mockFileControl()
	fc.ImmediateOriginContactPhoneNumber = "--"
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateOriginContactPhoneNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCreditTotalIndicator validation
func TestCreditTotalIndicator(t *testing.T) {
	fc := mockFileControl()
	fc.CreditTotalIndicator = 9
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CreditTotalIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFCFieldInclusionRecordType validates FieldInclusion
func TestFCFieldInclusionRecordType(t *testing.T) {
	fc := mockFileControl()
	fc.recordType = ""
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCashLetterCount validates FieldInclusion
func TestFieldInclusionCashLetterCount(t *testing.T) {
	fc := mockFileControl()
	fc.CashLetterCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CashLetterCount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionTotalRecordCount validates FieldInclusion
func TestFieldInclusionTotalRecordCount(t *testing.T) {
	fc := mockFileControl()
	fc.TotalRecordCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TotalRecordCount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionTotalItemCount validates FieldInclusion
func TestFieldInclusionTotalItemCount(t *testing.T) {
	fc := mockFileControl()
	fc.TotalItemCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TotalItemCount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionFileTotalAmount validates FieldInclusion
func TestFieldInclusionFileTotalAmount(t *testing.T) {
	fc := mockFileControl()
	fc.FileTotalAmount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FileTotalAmount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFileControlRuneCountInString validates RuneCountInString
func TestFileControlRuneCountInString(t *testing.T) {
	fc := NewFileControl()
	var line = "99"
	fc.Parse(line)

	if fc.CashLetterCount != 0 {
		t.Error("Parsed with an invalid RuneCountInString")
	}
}
