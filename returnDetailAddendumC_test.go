// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
)

// mockReturnDetailAddendumC creates a ReturnDetailAddendumC
func mockReturnDetailAddendumC() ReturnDetailAddendumC {
	rdAddendumC := NewReturnDetailAddendumC()
	rdAddendumC.ImageReferenceKeyIndicator = 1
	rdAddendumC.MicrofilmArchiveSequenceNumber = "1A"
	rdAddendumC.LengthImageReferenceKey = 0034
	rdAddendumC.ImageReferenceKey = "0"
	rdAddendumC.Description = "RD Addendum C"
	rdAddendumC.UserField = ""
	return rdAddendumC
}

// testMockReturnDetailAddendumCcreates a ReturnDetailAddendumC
func testMockReturnDetailAddendumC(t testing.TB) {
	rdAddendumC := mockReturnDetailAddendumC()
	if err := rdAddendumC.Validate(); err != nil {
		t.Error("MockReturnDetailAddendumC does not validate and will break other tests: ", err)
	}
	if rdAddendumC.recordType != "34" {
		t.Error("recordType does not validate")
	}
	if rdAddendumC.ImageReferenceKeyIndicator != 1 {
		t.Error("ImageReferenceKeyIndicator does not validate")
	}
	if rdAddendumC.MicrofilmArchiveSequenceNumber != "1A" {
		t.Error("MicrofilmArchiveSequenceNumber does not validate")
	}
	if rdAddendumC.LengthImageReferenceKey != 0034 {
		t.Error("LengthImageReferenceKey does not validate")
	}
	if rdAddendumC.ImageReferenceKey != "0" {
		t.Error("ImageReferenceKey does not validate")
	}
	if rdAddendumC.Description != "RD Addendum C" {
		t.Error("Description does not validate")
	}
	if rdAddendumC.UserField != "" {
		t.Error("UserField does not validate")
	}
}

// TestMockReturnDetailAddendumC tests creating a ReturnDetailAddendumC
func TestMockReturnDetailAddendumC(t *testing.T) {
	testMockReturnDetailAddendumC(t)
}

// BenchmarkMockReturnDetailAddendumC benchmarks creating a ReturnDetailAddendumC
func BenchmarkMockReturnDetailAddendumC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockReturnDetailAddendumC(b)
	}
}

// parseReturnDetailAddendumC validates parsing a ReturnDetailAddendumC
func parseReturnDetailAddendumC(t testing.TB) {
	var line = "3411A             00340                                 RD Addendum C           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	cd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(cd)

	if err := r.parseReturnDetailAddendumC(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumC[0]
	if record.recordType != "34" {
		t.Errorf("RecordType Expected '34' got: %v", record.recordType)
	}
	if record.ImageReferenceKeyIndicatorField() != "1" {
		t.Errorf("ImageReferenceKeyIndicator Expected '1' got: %v",
			record.ImageReferenceKeyIndicatorField())
	}
	if record.MicrofilmArchiveSequenceNumberField() != "1A             " {
		t.Errorf("MicrofilmArchiveSequenceNumber Expected '1A             ' got: %v",
			record.MicrofilmArchiveSequenceNumberField())
	}
	if record.LengthImageReferenceKeyField() != "0034" {
		t.Errorf("ImageReferenceKeyLength Expected '0034' got: %v", record.LengthImageReferenceKeyField())
	}
	if record.ImageReferenceKeyField() != "0                                 " {
		t.Errorf("ImageReferenceKey Expected '0                                 ' got: %v",
			record.ImageReferenceKeyField())
	}
	if record.DescriptionField() != "RD Addendum C  " {
		t.Errorf("Description Expected 'RD Addendum C  ' got: %v", record.DescriptionField())
	}
	if record.UserFieldField() != "    " {
		t.Errorf("UserField Expected '    ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != "     " {
		t.Errorf("reserved Expected '     ' got: %v", record.reservedField())
	}
}

// TestParseReturnDetailAddendumC tests validating parsing a ReturnDetailAddendumC
func TestParseReturnDetailAddendumC(t *testing.T) {
	parseReturnDetailAddendumC(t)
}

// BenchmarkParseReturnDetailAddendumC benchmarks validating parsing a ReturnDetailAddendumC
func BenchmarkParseReturnDetailAddendumC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseReturnDetailAddendumC(b)
	}
}

// testRDAddendumCString validates that a known parsed ReturnDetailAddendumC can return to a string of the same value
func testRDAddendumCString(t testing.TB) {
	var line = "3411A             00340                                 RD Addendum C           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	cd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(cd)

	if err := r.parseReturnDetailAddendumC(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumC[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestRDAddendumCString tests validating that a known parsed ReturnDetailAddendumC can return to a string of the
// same value
func TestRDAddendumCString(t *testing.T) {
	testRDAddendumCString(t)
}

// BenchmarkRDAddendumCString benchmarks validating that a known parsed ReturnDetailAddendumC
// can return to a string of the same value
func BenchmarkRDAddendumCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDAddendumCString(b)
	}
}
