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

// mockCashLetterHeader creates a CashLetterHeader
func mockCashLetterHeader() *CashLetterHeader {
	clh := NewCashLetterHeader()
	clh.CollectionTypeIndicator = "01"
	clh.DestinationRoutingNumber = "231380104"
	clh.ECEInstitutionRoutingNumber = "121042882"
	clh.CashLetterBusinessDate = time.Now()
	clh.CashLetterCreationDate = time.Now()
	clh.CashLetterCreationTime = time.Now()
	clh.CashLetterRecordTypeIndicator = "I"
	clh.CashLetterDocumentationTypeIndicator = "G"
	clh.CashLetterID = "A1"
	clh.OriginatorContactName = "Contact Name"
	clh.OriginatorContactPhoneNumber = "5558675552"
	clh.FedWorkType = ""
	clh.ReturnsIndicator = ""
	clh.UserField = ""
	return clh
}

// testMockCashLetterHeader creates a CashLetterHeader
func testMockCashLetterHeader(t testing.TB) {
	clh := mockCashLetterHeader()
	if err := clh.Validate(); err != nil {
		t.Error("mockCashLetterHeader does not validate and will break other tests: ", err)
	}
	if clh.recordType != "10" {
		t.Error("recordType does not validate and will break other tests")
	}
	if clh.CollectionTypeIndicator != "01" {
		t.Error("CollectionTypeIndicator does not validate and will break other tests")
	}
	if clh.DestinationRoutingNumber != "231380104" {
		t.Error("DestinationRoutingNumber does not validate and will break other tests")
	}
	if clh.ECEInstitutionRoutingNumber != "121042882" {
		t.Error("ECEInstitutionRoutingNumber does not validate and will break other tests")
	}
	if clh.CashLetterRecordTypeIndicator != "I" {
		t.Error("RecordTypeIndicator does not validate and will break other tests")
	}
	if clh.CashLetterDocumentationTypeIndicator != "G" {
		t.Error("DocumentationTypeIndicator does not validate and will break other tests")
	}
	if clh.CashLetterID != "A1" {
		t.Error("CashLetterID does not validate and will break other tests")
	}
	if clh.OriginatorContactName != "Contact Name" {
		t.Error("OriginatorContactName does not validate and will break other tests")
	}
	if clh.OriginatorContactPhoneNumber != "5558675552" {
		t.Error("OriginatorContactPhoneNumber does not validate and will break other tests")
	}
	if clh.FedWorkType != "" {
		t.Error("FedWorkType does not validate and will break other tests")
	}
	if clh.ReturnsIndicator != "" {
		t.Error("ReturnsIndicator does not validate and will break other tests")
	}
	if clh.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
	if clh.reserved != "" {
		t.Error("Reserved does not validate and will break other tests")
	}
}

// TestMockCashLetterHeader tests creating an ICL CashLetterHeader
func TestMockCashLetterHeader(t *testing.T) {
	testMockCashLetterHeader(t)
}

// BenchmarkMockCashLetterHeader benchmarks creating a CashLetterHeader
func BenchmarkMockCashLetterHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCashLetterHeader(b)
	}
}

// parseCashLetterHeader validates parsing a CashLetterHeader
func parseCashLetterHeader(t testing.TB) {
	var line = "100123138010412104288220180905201809051523IGA1      Contact Name  5558675552    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseCashLetterHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.CashLetterHeader

	if record.recordType != "10" {
		t.Errorf("RecordType Expected '10' got: %v", record.recordType)
	}
	if record.CollectionTypeIndicatorField() != "01" {
		t.Errorf("CollectionTypeIndicator Expected '01' got: %v", record.CollectionTypeIndicatorField())
	}
	if record.DestinationRoutingNumberField() != "231380104" {
		t.Errorf("DestinationRoutingNumber '231380104' got: %v", record.DestinationRoutingNumberField())
	}
	if record.ECEInstitutionRoutingNumberField() != "121042882" {
		t.Errorf("ECEInstitutionRoutingNumber Expected '121042882' got: %v", record.ECEInstitutionRoutingNumberField())
	}
	if record.CashLetterBusinessDateField() != "20180905" {
		t.Errorf("CashLetterBusinessDate Expected '20180905' got:'%v'", record.CashLetterBusinessDateField())
	}
	if record.CashLetterCreationDateField() != "20180905" {
		t.Errorf("CashLetterCreationDate Expected '20180905' got:'%v'", record.CashLetterCreationDateField())
	}
	if record.CashLetterCreationTimeField() != "1523" {
		t.Errorf("CashLetterCreationTime Expected '1523' got:'%v'", record.CashLetterCreationTimeField())
	}
	if record.CashLetterRecordTypeIndicatorField() != "I" {
		t.Errorf("CashLetterRecordTypeIndicator Expected 'I' got: %v", record.CashLetterRecordTypeIndicatorField())
	}
	if record.CashLetterDocumentationTypeIndicatorField() != "G" {
		t.Errorf("CashLetterDocumentationTypeIndicator Expected 'G' got:'%v'", record.CashLetterDocumentationTypeIndicatorField())
	}
	if record.CashLetterIDField() != "A1      " {
		t.Errorf("CashLetterID Expected 'A1      ' got:'%v'", record.CashLetterIDField())
	}
	if record.OriginatorContactNameField() != "Contact Name  " {
		t.Errorf("OriginatorContactName Expected 'Contact Name  ' got: '%v'", record.OriginatorContactNameField())
	}
	if record.OriginatorContactPhoneNumberField() != "5558675552" {
		t.Errorf("OriginatorContactPhoneNumber Expected '5558675552' got: '%v'", record.OriginatorContactPhoneNumberField())
	}
	if record.FedWorkTypeField() != " " {
		t.Errorf("FedWorkType Expected ' ' got:'%v'", record.FedWorkTypeField())
	}
	if record.ReturnsIndicatorField() != " " {
		t.Errorf("ReturnsIndicator ' ' got:'%v'", record.ReturnsIndicatorField())
	}
	if record.UserFieldField() != " " {
		t.Errorf("UserField Expected ' ' got:'%v'", record.UserFieldField())
	}
	if record.reservedField() != " " {
		t.Errorf("reserved Expected ' ' got:'%v'", record.reservedField())
	}
}

// TestParseCashLetterHeader test validates parsing a CashLetterHeader
func TestParseCashLetterHeader(t *testing.T) {
	parseCashLetterHeader(t)
}

// BenchmarkParseCashLetterHeader benchmark validates parsing a CashLetterHeader
func BenchmarkParseCashLetterHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseCashLetterHeader(b)
	}
}

// testCLHString validates that a known parsed CashLetterHeader can return to a string of the same value
func testCLHString(t testing.TB) {
	var line = "100123138010412104288220180905201809051523IGA1      Contact Name  5558675552    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseCashLetterHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.CashLetterHeader
	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCLHString tests validating that a known parsed CashLetterHeader can return to a string of the same value
func TestCLHString(t *testing.T) {
	testCLHString(t)
}

// BenchmarkCLHString benchmarks validating that a known parsed CashLetterHeader
// can return to a string of the same value
func BenchmarkCLHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCLHString(b)
	}
}
