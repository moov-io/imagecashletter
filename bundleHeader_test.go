// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
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
	bh.CycleNumber = "001"
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
	if bh.CycleNumber != "001" {
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
