// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "testing"

// mockReturnDetailAddendumC creates a ReturnDetailAddendumC
func mockReturnDetailAddendumC() ReturnDetailAddendumC {
	rdAddendumC := NewReturnDetailAddendumC()
	rdAddendumC.ImageReferenceKeyIndicator = 1
	rdAddendumC.MicrofilmArchiveSequenceNumber = "1A"
	rdAddendumC.LengthImageReferenceKey = 0034
	rdAddendumC.ImageReferenceKey = "0"
	rdAddendumC.Description = "Return Addendum C"
	rdAddendumC.UserField = ""
	return rdAddendumC
}

// testMockReturnDetailAddendumCcreates a ReturnDetailAddendumC
func testMockReturnDetailAddendumC(t testing.TB) {
	rdAddendumC := mockReturnDetailAddendumC()
	if err := rdAddendumC.Validate(); err != nil {
		t.Error("MockReturnDetailAddendumB does not validate and will break other tests: ", err)
	}
	if rdAddendumC.recordType != "34" {
		t.Error("recordType does not validate and will break other tests")
	}
	if rdAddendumC.ImageReferenceKeyIndicator != 1 {
		t.Error("ImageReferenceKeyIndicator does not validate and will break other tests")
	}
	if rdAddendumC.MicrofilmArchiveSequenceNumber != "1A" {
		t.Error("MicrofilmArchiveSequenceNumber does not validate and will break other tests")
	}
	if rdAddendumC.LengthImageReferenceKey != 0034 {
		t.Error("LengthImageReferenceKey does not validate and will break other tests")
	}
	if rdAddendumC.ImageReferenceKey != "0" {
		t.Error("ImageReferenceKey does not validate and will break other tests")
	}
	if rdAddendumC.Description != "Return Addendum C" {
		t.Error("Description does not validate and will break other tests")
	}
	if rdAddendumC.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
}

// TestMockReturnDetailAddendumC tests creating a ReturnDetailAddendumB
func TestMockReturnDetailAddendumC(t *testing.T) {
	testMockReturnDetailAddendumB(t)
}

// BenchmarkMockReturnDetailAddendumC benchmarks creating a ReturnDetailAddendumB
func BenchmarkMockReturnDetailAddendumC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockReturnDetailAddendumC(b)
	}
}
