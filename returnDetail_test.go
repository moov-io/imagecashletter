// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
	"time"
)

// mockReturnDetail creates a ReturnDetail
func mockReturnDetail() *ReturnDetail {
	rd := NewReturnDetail()
	rd.PayorBankRoutingNumber = "03130001"
	rd.PayorBankCheckDigit = "2"
	rd.OnUs = "5558881"
	rd.ItemAmount = 100000
	rd.ReturnReason = "A"
	rd.AddendumCount = 4
	rd.DocumentationTypeIndicator = "G"
	rd.ForwardBundleDate = time.Now()
	rd.EceInstitutionItemSequenceNumber = "1              "
	rd.ExternalProcessingCode = ""
	rd.ReturnNotificationIndicator = 2
	rd.ArchiveTypeIndicator = "B"
	rd.TimesReturned = 0
	return rd
}

// testMockReturnDetail creates a ReturnDetail
func testMockReturnDetail(t testing.TB) {
	rd := mockReturnDetail()
	if err := rd.Validate(); err != nil {
		t.Error("mockReturnDetail does not validate and will break other tests: ", err)
	}
	if rd.recordType != "31" {
		t.Error("recordType does not validate")
	}
	if rd.PayorBankRoutingNumber != "03130001" {
		t.Error("PayorBankRoutingNumber does not validate")
	}
	if rd.PayorBankCheckDigit != "2" {
		t.Error("PayorBankCheckDigit does not validate")
	}
	if rd.OnUs != "5558881" {
		t.Error("OnUs does not validate")
	}
	if rd.ItemAmount != 100000 {
		t.Error("ItemAmount does not validate")
	}
	if rd.ReturnReason != "A" {
		t.Error("ReturnReason does not validate")
	}
	if rd.AddendumCount != 4 {
		t.Error("AddendumCount does not validate")
	}
	if rd.DocumentationTypeIndicator != "G" {
		t.Error("DocumentationTypeIndicator does not validate")
	}
	if rd.EceInstitutionItemSequenceNumber != "1              " {
		t.Error("EceInstitutionItemSequenceNumber does not validate")
	}
	if rd.ExternalProcessingCode != "" {
		t.Error("ExternalProcessingCode does not validate")
	}
	if rd.ReturnNotificationIndicator != 2 {
		t.Error("ReturnNotificationIndicator does not validate")
	}
	if rd.ArchiveTypeIndicator != "B" {
		t.Error("ArchiveTypeIndicator does not validate")
	}
	if rd.TimesReturned != 0 {
		t.Error("TimesReturned does not validate")
	}
}

// TestMockReturnDetail tests creating a ReturnDetail
func TestMockReturnDetail(t *testing.T) {
	testMockReturnDetail(t)
}

// BenchmarkMockReturnDetail benchmarks creating a ReturnDetail
func BenchmarkMockReturnDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockReturnDetail(b)
	}
}

// TestParseReturnDetail validates parsing a ReturnDetail
func TestParseReturnDetail(t *testing.T) {
	var line = "31031300012             55588810000100000A04G201809051               2B0        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)

	if err := r.parseReturnDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0]

	if record.recordType != "31" {
		t.Errorf("RecordType Expected '31' got: %v", record.recordType)
	}
	if record.PayorBankRoutingNumberField() != "03130001" {
		t.Errorf("PayorBankRoutingNumber Expected '03130001' got: %v", record.PayorBankRoutingNumberField())
	}
	if record.PayorBankCheckDigitField() != "2" {
		t.Errorf("PayorBank Expected '2' got: %v", record.PayorBankCheckDigitField())
	}
	if record.OnUsField() != "             5558881" {
		t.Errorf("OnUs Expected '             5558881' got: %v", record.OnUsField())
	}
	if record.ItemAmountField() != "0000100000" {
		t.Errorf("ItemAmount Expected '0000100000' got: %v", record.ItemAmountField())
	}
	if record.ReturnReasonField() != "A" {
		t.Errorf("ReturnReason Expected 'A' got: %v", record.ReturnReasonField())
	}
	if record.AddendumCountField() != "04" {
		t.Errorf("AddendumCount Expected '04' got: %v", record.AddendumCountField())
	}
	if record.DocumentationTypeIndicatorField() != "G" {
		t.Errorf("DocumentationTypeIndicator Expected 'G' got: %v", record.DocumentationTypeIndicatorField())
	}
	if record.EceInstitutionItemSequenceNumberField() != "1              " {
		t.Errorf("EceInstitutionItemSequenceNumber Expected '1              ' got: %v", record.EceInstitutionItemSequenceNumberField())
	}
	if record.ExternalProcessingCodeField() != " " {
		t.Errorf("ExternalProcessingCode Expected ' ' got: %v", record.ExternalProcessingCodeField())
	}
	if record.ReturnNotificationIndicatorField() != "2" {
		t.Errorf("ReturnNotificationIndicator Expected '2' got: %v", record.ReturnNotificationIndicatorField())
	}
	if record.ArchiveTypeIndicatorField() != "B" {
		t.Errorf("ArchiveTypeIndicator Expected 'R' got: %v", record.ArchiveTypeIndicatorField())
	}
	if record.TimesReturnedField() != "0" {
		t.Errorf("TimesReturned Expected '0' got: %v", record.TimesReturnedField())
	}

}

// testReturnDetailString validates that a known parsed ReturnDetail can return to a string of the same value
func testReturnDetailString(t testing.TB) {
	var line = "31031300012             55588810000100000A04G201809051               2B0        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)

	if err := r.parseReturnDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestReturnDetailString tests validating that a known parsed ReturnDetail can return to a string of the
// same value
func TestReturnDetailString(t *testing.T) {
	testReturnDetailString(t)
}

// BenchmarkReturnDetailString benchmarks validating that a known parsed ReturnDetailAddendumB
// can return to a string of the same value
func BenchmarkReturnDetailString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReturnDetailString(b)
	}
}
