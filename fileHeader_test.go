// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
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
