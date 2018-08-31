// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "testing"

// mockBundleControl creates a BundleControl
func mockBundleControl() BundleControl {
	bc := NewBundleControl()
	// BundleItemsCount - CheckDetail
	// ToDo: CheckDetailAddendum* and ImageView*
	bc.BundleItemsCount = 1
	bc.BundleTotalAmount = 100000 // 1000.00
	// ToDo: CheckDetail
	bc.MICRValidTotalAmount = 0
	// ToDo: ImageView*
	bc.BundleImagesCount = 0
	bc.UserField = ""
	bc.CreditTotalIndicator = 0
	return bc
}

// testMockBundleControl creates an ICL BundleControl
func testMockBundleControl(t testing.TB) {
	bc := mockBundleControl()
	if err := bc.Validate(); err != nil {
		t.Error("mockBundleControl does not validate and will break other tests: ", err)
	}
	if bc.recordType != "70" {
		t.Error("recordType does not validate and will break other tests")
	}
	if bc.BundleItemsCount != 1 {
		t.Error("BundleItemsCount does not validate and will break other tests")
	}
	if bc.BundleTotalAmount != 100000 {
		t.Error("BundleTotalAmount does not validate and will break other tests")
	}
	if bc.MICRValidTotalAmount != 0 {
		t.Error("MICRValidTotalAmount does not validate and will break other tests")
	}
	if bc.BundleImagesCount != 0 {
		t.Error("BundleImagesCount does not validate and will break other tests")
	}
	if bc.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
	if bc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate and will break other tests")
	}
}

// TestMockBundleControl tests creating an ICL BundleControl
func TestMockBundleControl(t *testing.T) {
	testMockBundleControl(t)
}

// BenchmarkMockBundleControl benchmarks creating an ICL BundleControl
func BenchmarkMockBundleControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockBundleControl(b)
	}
}
