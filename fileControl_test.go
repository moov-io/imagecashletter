// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "testing"

// mockFileControl creates a FileControl
func mockFileControl() FileControl {
	fc := NewFileControl()
	fc.CashLetterCount = 1
	// TotalRecordCount - FileHeader, CashLetterHeader, BundleHeader, CheckDetail, CashLetterControl, BundleControl,
	// FileControl
	// ToDo: CheckDetailAddendum* and ImageView*
	fc.TotalRecordCount = 7
	// TotalItemCount - CheckDetail
	// ToDo: CheckDetailAddendum* and ImageView*
	fc.TotalItemCount = 1
	fc.FileTotalAmount = 100000 //1000.00
	fc.ImmediateOriginContactName = "Contact Name"
	fc.ImmediateOriginContactPhoneNumber = "5558675552"
	fc.CreditTotalIndicator = 0
	return fc
}

// testMockFileControl creates an ICL FileControl
func testMockFileControl(t testing.TB) {
	fc := mockFileControl()
	if err := fc.Validate(); err != nil {
		t.Error("mockFileControl does not validate and will break other tests: ", err)
	}
	if fc.recordType != "99" {
		t.Error("recordType does not validate and will break other tests")
	}
	if fc.CashLetterCount != 1 {
		t.Error("CashLetterCount does not validate and will break other tests")
	}
	if fc.TotalRecordCount != 7 {
		t.Error("TotalRecordCount does not validate and will break other tests")
	}
	if fc.TotalItemCount != 1 {
		t.Error("TotalItemCount does not validate and will break other tests")
	}
	if fc.FileTotalAmount != 100000 {
		t.Error("FileTotalAmount does not validate and will break other tests")
	}
	if fc.ImmediateOriginContactName != "Contact Name" {
		t.Error("ImmediateOriginContactName does not validate and will break other tests")
	}
	if fc.ImmediateOriginContactPhoneNumber != "5558675552" {
		t.Error("ImmediateOriginContactPhoneNumber does not validate and will break other tests")
	}
	if fc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate and will break other tests")
	}
}

// TestMockFileControl tests creating an ICL FileControl
func TestMockFileControl(t *testing.T) {
	testMockFileControl(t)
}

// BenchmarkMockFileControl benchmarks creating an ICL FileControl
func BenchmarkMockFileControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockFileControl(b)
	}
}
