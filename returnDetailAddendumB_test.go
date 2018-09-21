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

// mockReturnDetailAddendumB creates a ReturnDetailAddendumB
func mockReturnDetailAddendumB() ReturnDetailAddendumB {
	rdAddendumB := NewReturnDetailAddendumB()
	rdAddendumB.PayorBankName = "Payor Bank Name"
	rdAddendumB.AuxiliaryOnUs = "123456789"
	rdAddendumB.PayorBankSequenceNumber = "1              "
	rdAddendumB.PayorBankBusinessDate = time.Now()
	rdAddendumB.PayorAccountName = "Payor Account Name"
	return rdAddendumB
}

// testMockReturnDetailAddendumB creates a ReturnDetailAddendumB
func testMockReturnDetailAddendumB(t testing.TB) {
	rdAddendumB := mockReturnDetailAddendumB()
	if err := rdAddendumB.Validate(); err != nil {
		t.Error("MockReturnDetailAddendumB does not validate and will break other tests: ", err)
	}
	if rdAddendumB.recordType != "33" {
		t.Error("recordType does not validate")
	}
	if rdAddendumB.PayorBankName != "Payor Bank Name" {
		t.Error("PayorBankName does not validate")
	}
	if rdAddendumB.AuxiliaryOnUs != "123456789" {
		t.Error("AuxiliaryOnUs does not validate")
	}
	if rdAddendumB.PayorBankSequenceNumber != "1              " {
		t.Error("PayorBankSequenceNumber does not validate")
	}
	if rdAddendumB.PayorAccountName != "Payor Account Name" {
		t.Error("PayorAccountName does not validate")
	}
}

// TestMockReturnDetailAddendumB tests creating a ReturnDetailAddendumB
func TestMockReturnDetailAddendumB(t *testing.T) {
	testMockReturnDetailAddendumB(t)
}

// BenchmarkMockReturnDetailAddendumB benchmarks creating a ReturnDetailAddendumB
func BenchmarkMockReturnDetailAddendumB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockReturnDetailAddendumB(b)
	}
}

// parseReturnDetailAddendumB validates parsing a ReturnDetailAddendumB
func parseReturnDetailAddendumB(t testing.TB) {
	var line = "33Payor Bank Name         1234567891              20180905Payor Account Name    "
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

	if err := r.parseReturnDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentReturnBundle.GetReturns()[0].ReturnDetailAddendumB[0]

	if record.recordType != "33" {
		t.Errorf("RecordType Expected '33' got: %v", record.recordType)
	}
	if record.PayorBankNameField() != "Payor Bank Name   " {
		t.Errorf("PayorBankName Expected 'Payor Bank Name   ' got: %v", record.PayorBankNameField())
	}
	if record.AuxiliaryOnUsField() != "      123456789" {
		t.Errorf("AuxiliaryOnUs Expected '      123456789' got: %v", record.AuxiliaryOnUsField())
	}
	if record.PayorBankSequenceNumberField() != "1              " {
		t.Errorf("PayorBankSequenceNumber Expected '1              ' got: %v", record.PayorBankSequenceNumberField())
	}
	if record.PayorAccountNameField() != "Payor Account Name    " {
		t.Errorf("PayorAccountName Expected 'Payor Account Name    ' got: %v", record.PayorAccountNameField())
	}
}

// TestParseReturnDetailAddendumB tests validating parsing a ReturnDetailAddendumB
func TestParseReturnDetailAddendumB(t *testing.T) {
	parseReturnDetailAddendumB(t)
}

// BenchmarkParseReturnDetailAddendumB benchmarks validatingparsing a ReturnDetailAddendumB
func BenchmarkParseReturnDetailAddendumB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseReturnDetailAddendumB(b)
	}
}

// testRDAddendumBString validates that a known parsed ReturnDetailAddendumB can return to a string of the same value
func testRDAddendumBString(t testing.TB) {
	var line = "33Payor Bank Name         1234567891              20180905Payor Account Name    "
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

	if err := r.parseReturnDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentReturnBundle.GetReturns()[0].ReturnDetailAddendumB[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestRDAddendumBString tests validating that a known parsed ReturnDetailAddendumB can return to a string of the
// same value
func TestRDAddendumBString(t *testing.T) {
	testRDAddendumBString(t)
}

// BenchmarkRDAddendumBString benchmarks validating that a known parsed ReturnDetailAddendumB
// can return to a string of the same value
func BenchmarkRDAddendumBString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDAddendumBString(b)
	}
}
