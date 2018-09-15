// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
)

// ToDo: Review Image key Code/Functionality - Variable length
// mockCheckDetailAddendumB creates a CheckDetailAddendumB
func mockCheckDetailAddendumB() CheckDetailAddendumB {
	cdAddendumB := NewCheckDetailAddendumB()
	cdAddendumB.ImageReferenceKeyIndicator = 1
	cdAddendumB.MicrofilmArchiveSequenceNumber = "1A"
	cdAddendumB.ImageReferenceKeyLength = "0034"
	cdAddendumB.ImageReferenceKey = "0"
	cdAddendumB.Description = "CD Addendum B"
	cdAddendumB.UserField = ""
	return cdAddendumB
}

// testMockCheckDetailAddendumB creates a CheckDetailAddendumB
func testMockCheckDetailAddendumB(t testing.TB) {
	cdAddendumB := mockCheckDetailAddendumB()
	if err := cdAddendumB.Validate(); err != nil {
		t.Error("mockBundleHeader does not validate and will break other tests: ", err)
	}
	if cdAddendumB.recordType != "27" {
		t.Error("recordType does not validate and will break other tests")
	}
	if cdAddendumB.ImageReferenceKeyIndicator != 1 {
		t.Error("ImageReferenceKeyIndicator does not validate and will break other tests")
	}
	if cdAddendumB.MicrofilmArchiveSequenceNumber != "1A" {
		t.Error("MicrofilmArchiveSequenceNumber does not validate and will break other tests")
	}
	if cdAddendumB.ImageReferenceKeyLength != "0034" {
		t.Error("ImageReferenceKeyLength does not validate and will break other tests")
	}
	if cdAddendumB.ImageReferenceKey != "0" {
		t.Error("ImageReferenceKey does not validate and will break other tests")
	}
	if cdAddendumB.Description != "CD Addendum B" {
		t.Error("Description does not validate and will break other tests")
	}
	if cdAddendumB.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
}

// TestMockCheckDetailAddendumB  tests creating an ICL CheckDetailAddendumB
func TestMockCheckDetailAddendumB(t *testing.T) {
	testMockCheckDetailAddendumB(t)
}

// BenchmarkMockCheckDetailAddendumB benchmarks creating a CheckDetailAddendumB
func BenchmarkMockCheckDetailAddendumB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCheckDetailAddendumB(b)
	}
}

// parseCheckDetailAddendumB validates parsing a CheckDetailAddendumB
func parseCheckDetailAddendumB(t testing.TB) {
	var line = "2711A             00340                                 CD Addendum B           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumB[0]
	if record.recordType != "27" {
		t.Errorf("RecordType Expected '26' got: %v", record.recordType)
	}
	if record.ImageReferenceKeyIndicatorField() != "1" {
		t.Errorf("ImageReferenceKeyIndicator Expected '26' got: %v",
			record.ImageReferenceKeyIndicatorField())
	}
	if record.MicrofilmArchiveSequenceNumberField() != "1A             " {
		t.Errorf("MicrofilmArchiveSequenceNumber Expected '1A             ' got: %v",
			record.MicrofilmArchiveSequenceNumberField())
	}
	if record.ImageReferenceKeyLengthField() != "0034" {
		t.Errorf("ImageReferenceKeyLength Expected '0034' got: %v", record.ImageReferenceKeyLengthField())
	}
	if record.ImageReferenceKeyField() != "0                                 " {
		t.Errorf("ImageReferenceKey Expected '0                                 ' got: %v",
			record.ImageReferenceKeyLengthField())
	}
	if record.DescriptionField() != "CD Addendum B  " {
		t.Errorf("Description Expected 'CD Addendum B  ' got: %v", record.DescriptionField())
	}
	if record.UserFieldField() != "    " {
		t.Errorf("UserField Expected '    ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != "     " {
		t.Errorf("reserved Expected '     ' got: %v", record.reservedField())
	}
}

// TestParseCheckDetailAddendumB test validates parsing a CheckDetailAddendumB
func TestParseCheckDetailAddendumB(t *testing.T) {
	parseCheckDetailAddendumB(t)
}

// BenchmarkParseCheckDetailAddendumB benchmark validates parsing a CheckDetailAddendumB
func BenchmarkParseCheckDetailAddendumB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseCheckDetailAddendumB(b)
	}
}

// testCDAddendumBString validates that a known parsed CheckDetailAddendumB can return to a string of the same value
func testCDAddendumBString(t testing.TB) {
	var line = "2711A             00340                                 CD Addendum B           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumB[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCDAddendumBString tests validating that a known parsed CheckDetailAddendumB can return to a string of the
// same value
func TestCDAddendumBString(t *testing.T) {
	testCDAddendumBString(t)
}

// BenchmarkCDAddendumBString benchmarks validating that a known parsed CheckDetail
// can return to a string of the same value
func BenchmarkCDAddendumBString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumBString(b)
	}
}
