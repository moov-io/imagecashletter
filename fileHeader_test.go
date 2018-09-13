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

// mockFileHeader creates a FileHeader
func mockFileHeader() FileHeader {
	fh := NewFileHeader()
	fh.StandardLevel = "35"
	fh.TestFileIndicator = "T"
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now()
	fh.FileCreationTime = time.Now()
	fh.ResendIndicator = "N"
	fh.ImmediateDestinationName = "Citadel"
	fh.ImmediateOriginName = "Wells Fargo"
	fh.FileIDModifier = ""
	fh.CountryCode = "US"
	fh.UserField = ""
	fh.CompanionDocumentIndicator = ""
	return fh
}

// testMockFileHeader creates an ICL FileHeader
func testMockFileHeader(t testing.TB) {
	fh := mockFileHeader()
	if err := fh.Validate(); err != nil {
		t.Error("mockFileHeader does not validate and will break other tests: ", err)
	}
	if fh.recordType != "01" {
		t.Error("recordType does not validate and will break other tests")
	}
	if fh.StandardLevel != "35" {
		t.Error("StandardLevel does not validate and will break other tests")
	}
	if fh.TestFileIndicator != "T" {
		t.Error("TestFileIndicator does not validate and will break other tests")
	}
	if fh.ResendIndicator != "N" {
		t.Error("ResendIndicator does not validate and will break other tests")
	}
	if fh.ImmediateDestination != "231380104" {
		t.Error("DestinationRoutingNumber does not validate and will break other tests")
	}
	if fh.ImmediateOrigin != "121042882" {
		t.Error("ECEInstitutionRoutingNumber does not validate and will break other tests")
	}
	if fh.ImmediateDestinationName != "Citadel" {
		t.Error("ImmediateDestinationName does not validate and will break other tests")
	}
	if fh.ImmediateOriginName != "Wells Fargo" {
		t.Error("ImmediateOriginName does not validate and will break other tests")
	}
	if fh.FileIDModifier != "" {
		t.Error("FileIDModifier does not validate and will break other tests")
	}
	if fh.CountryCode != "US" {
		t.Error("CountryCode does not validate and will break other tests")
	}
	if fh.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
	if fh.CompanionDocumentIndicator != "" {
		t.Error("CompanionDocumentIndicator does not validate and will break other tests")
	}
}

// TestMockFileHeader tests creating an ICL FileHeader
func TestMockFileHeader(t *testing.T) {
	testMockFileHeader(t)
}

// BenchmarkMockFileHeader benchmarks creating an ICL FileHeader
func BenchmarkMockFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockFileHeader(b)
	}
}

// parseFileHeader validates parsing a FileHeader
func parseFileHeader(t testing.TB) {
	var line = "0135T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseFileHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.File.Header

	if record.recordType != "01" {
		t.Errorf("RecordType Expected '01' got: %v", record.recordType)
	}
	if record.StandardLevelField() != "35" {
		t.Errorf("StandardLevel Expected '35' got: %v", record.StandardLevelField())
	}
	if record.TestFileIndicatorField() != "T" {
		t.Errorf("TestFileIndicator 'T' got: %v", record.TestFileIndicatorField())
	}
	if record.ImmediateDestinationField() != "231380104" {
		t.Errorf("ImmediateDestination Expected '231380104' got: %v", record.ImmediateDestinationField())
	}
	if record.ImmediateOriginField() != "121042882" {
		t.Errorf("ImmediateOrigin Expected '121042882' got: %v", record.ImmediateOriginField())
	}
	if record.FileCreationDateField() != "20180905" {
		t.Errorf("FileCreationDate Expected '20180905' got:'%v'", record.FileCreationDateField())
	}
	if record.FileCreationTimeField() != "1523" {
		t.Errorf("FileCreationTime Expected '1523' got:'%v'", record.FileCreationTimeField())
	}
	if record.ResendIndicatorField() != "N" {
		t.Errorf("ResendIndicator Expected 'N' got: %v", record.ResendIndicatorField())
	}
	if record.ImmediateDestinationNameField() != "Citadel           " {
		t.Errorf("ImmediateDestinationName Expected 'Citadel           ' got:'%v'", record.ImmediateDestinationNameField())
	}
	if record.ImmediateOriginNameField() != "Wells Fargo       " {
		t.Errorf("ImmediateOriginName Expected 'Wells Fargo       ' got: '%v'", record.ImmediateOriginNameField())
	}
	if record.FileIDModifierField() != " " {
		t.Errorf("FileIDModifier Expected ' ' got:'%v'", record.FileIDModifierField())
	}
	if record.CountryCodeField() != "US" {
		t.Errorf("CountryCode Expected 'US' got:'%v'", record.CountryCodeField())
	}
	if record.UserFieldField() != "    " {
		t.Errorf("UserField Expected '    ' got:'%v'", record.UserFieldField())
	}
	if record.CompanionDocumentIndicatorField() != " " {
		t.Errorf("CompanionDocumentIndicator Expected ' ' got:'%v'", record.CompanionDocumentIndicatorField())
	}
}

// TestParseFileHeader test validates parsing a FileHeader
func TestParseFileHeader(t *testing.T) {
	parseFileHeader(t)
}

// BenchmarkParseFileHeader benchmark validates parsing a FileHeader
func BenchmarkParseFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseFileHeader(b)
	}
}

// testFHString validates that a known parsed FileHeader can return to a string of the same value
func testFHString(t testing.TB) {
	var line = "0135T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseFileHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.File.Header

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestFHString tests validating that a known parsed FileHeader can return to a string of the same value
func TestFHString(t *testing.T) {
	testFHString(t)
}

// BenchmarkFHString benchmarks validating that a known parsed FileHeader
// can return to a string of the same value
func BenchmarkFHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHString(b)
	}
}
