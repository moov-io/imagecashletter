// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"testing"
)

// ToDo: Review Image key Code/Functionality
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

// testMockCheckDetailAddendumB creates an ICL CheckDetailAddendumB
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

// BenchmarkMockCheckDetailAddendumB benchmarks creating an ICL CheckDetailAddendumB
func BenchmarkMockCheckDetailAddendumB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCheckDetailAddendumB(b)
	}
}
