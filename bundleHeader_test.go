// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"strings"
	"testing"
	"time"
)

// mockBundleHeader creates a BundleHeader
func mockBundleHeader() *BundleHeader {
	bh := NewBundleHeader()
	bh.CollectionTypeIndicator = "01"
	bh.DestinationRoutingNumber = "231380104"
	bh.ECEInstitutionRoutingNumber = "121042882"
	bh.BundleBusinessDate = time.Now()
	bh.BundleCreationDate = time.Now()
	bh.BundleID = "9999"
	bh.BundleSequenceNumber = 1
	bh.CycleNumber = "01"
	bh.UserField = ""
	return bh
}

// testMockBundleHeader creates an ICL BundleHeader
func testMockBundleHeader(t testing.TB) {
	bh := mockBundleHeader()
	if err := bh.Validate(); err != nil {
		t.Error("mockBundleHeader does not validate and will break other tests: ", err)
	}
	if bh.recordType != "20" {
		t.Error("recordType does not validate and will break other tests")
	}
	if bh.CollectionTypeIndicator != "01" {
		t.Error("CollectionTypeIndicator does not validate and will break other tests")
	}
	if bh.DestinationRoutingNumber != "231380104" {
		t.Error("DestinationRoutingNumber does not validate and will break other tests")
	}
	if bh.ECEInstitutionRoutingNumber != "121042882" {
		t.Error("ECEInstitutionRoutingNumber does not validate and will break other tests")
	}
	if bh.BundleID != "9999" {
		t.Error("BundleID does not validate and will break other tests")
	}
	if bh.BundleSequenceNumber != 1 {
		t.Error("SequenceNumber does not validate and will break other tests")
	}
	if bh.CycleNumber != "01" {
		t.Error("CycleNumber does not validate and will break other tests")
	}
	if bh.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
}

// TestMockBundleHeader tests creating an ICL BundleHeader
func TestMockBundleHeader(t *testing.T) {
	testMockBundleHeader(t)
}

// BenchmarkMockBundleHeader benchmarks creating an ICL BundleHeader
func BenchmarkMockBundleHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockBundleHeader(b)
	}
}

// parseBundleHeader validates parsing a BundleHeader
func parseBundleHeader(t testing.TB) {
	var line = "200123138010412104288220180905201809059999      000101                          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))

	if err := r.parseBundleHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	record := r.currentCashLetter.currentBundle.BundleHeader
	if record.recordType != "20" {
		t.Errorf("RecordType Expected '20' got: %v", record.recordType)
	}
	if record.CollectionTypeIndicatorField() != "01" {
		t.Errorf("CollectionTypeIndicator Expected '01' got: %v", record.CollectionTypeIndicator)
	}
	if record.DestinationRoutingNumberField() != "231380104" {
		t.Errorf("DestinationRoutingNumber '231380104' got: %v", record.DestinationRoutingNumberField())
	}
	if record.ECEInstitutionRoutingNumberField() != "121042882" {
		t.Errorf("ECEInstitutionRoutingNumber Expected '121042882' got: %v", record.ECEInstitutionRoutingNumberField())
	}
	if record.BundleBusinessDateField() != "20180905" {
		t.Errorf("BundleBusinessDate Expected '20180905' got:'%v'", record.BundleBusinessDateField())
	}
	if record.BundleCreationDateField() != "20180905" {
		t.Errorf("BundleCreationDate Expected '20180905' got:'%v'", record.BundleCreationDateField())
	}
	if record.BundleIDField() != "9999      " {
		t.Errorf("BundleID Expected '9999      ' got:'%v'", record.BundleIDField())
	}
	if record.BundleSequenceNumberField() != "0001" {
		t.Errorf("BundleSequenceNumber Expected '0001' got: '%v'", record.BundleSequenceNumberField())
	}
	if record.CycleNumberField() != "01" {
		t.Errorf("CycleNumber Expected '01' got:'%v'", record.CycleNumberField())
	}
	if record.reservedField() != "         " {
		t.Errorf("reserved Expected '         ' got:'%v'", record.reservedField())
	}
	if record.UserFieldField() != "     " {
		t.Errorf("UserField Expected '     ' got:'%v'", record.UserFieldField())
	}
	if record.reservedTwoField() != "            " {
		t.Errorf("reservedTwo Expected '            ' got:'%v'", record.reservedTwoField())
	}
}

// TestParseBundleHeader test validates parsing a BundleHeader
func TestParseBundleHeader(t *testing.T) {
	parseBundleHeader(t)
}

// BenchmarkParseBundleHeader benchmark validates parsing a BundleHeader
func BenchmarkParseBundleHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseBundleHeader(b)
	}
}

// testBHString validates that a known parsed BundleHeader can return to a string of the same value
func testBHString(t testing.TB) {
	var line = "200123138010412104288220180905201809059999      000101                          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	if err := r.parseBundleHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentCashLetter.currentBundle.BundleHeader
	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestBHString tests validating that a known parsed BundleHeader can return to a string of the same value
func TestBHString(t *testing.T) {
	testBHString(t)
}

// BenchmarkBHString benchmarks validating that a known parsed BundleHeader
// can return to a string of the same value
func BenchmarkBHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHString(b)
	}
}
