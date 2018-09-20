// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
)

// mockFileControl creates a FileControl
func mockFileControl() FileControl {
	fc := NewFileControl()
	fc.CashLetterCount = 1
	// TotalRecordCount - FileHeader, CashLetterHeader, BundleHeader, CheckDetail, CashLetterControl, BundleControl,
	// FileControl
	// ToDo: CheckDetailAddendum* and ImageView*
	fc.TotalRecordCount = 7
	// TotalItemCount - CheckDetail
	// ToDo: CheckDetailAddendum* and ImageView*
	fc.TotalItemCount = 1
	fc.FileTotalAmount = 100000 //1000.00
	fc.ImmediateOriginContactName = "Contact Name"
	fc.ImmediateOriginContactPhoneNumber = "5558675552"
	fc.CreditTotalIndicator = 0
	return fc
}

// testMockFileControl creates a FileControl
func testMockFileControl(t testing.TB) {
	fc := mockFileControl()
	if err := fc.Validate(); err != nil {
		t.Error("mockFileControl does not validate and will break other tests: ", err)
	}
	if fc.recordType != "99" {
		t.Error("recordType does not validate and will break other tests")
	}
	if fc.CashLetterCount != 1 {
		t.Error("CashLetterCount does not validate and will break other tests")
	}
	if fc.TotalRecordCount != 7 {
		t.Error("TotalRecordCount does not validate and will break other tests")
	}
	if fc.TotalItemCount != 1 {
		t.Error("TotalItemCount does not validate and will break other tests")
	}
	if fc.FileTotalAmount != 100000 {
		t.Error("FileTotalAmount does not validate and will break other tests")
	}
	if fc.ImmediateOriginContactName != "Contact Name" {
		t.Error("ImmediateOriginContactName does not validate and will break other tests")
	}
	if fc.ImmediateOriginContactPhoneNumber != "5558675552" {
		t.Error("ImmediateOriginContactPhoneNumber does not validate and will break other tests")
	}
	if fc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate and will break other tests")
	}
}

// TestMockFileControl tests creating a FileControl
func TestMockFileControl(t *testing.T) {
	testMockFileControl(t)
}

// BenchmarkMockFileControl benchmarks creating an ICL FileControl
func BenchmarkMockFileControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockFileControl(b)
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
