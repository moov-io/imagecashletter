// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"testing"
	"time"
)

// mockCashLetterControl creates a CashLetterControl
func mockCashLetterControl() *CashLetterControl {
	clc := NewCashLetterControl()
	clc.CashLetterBundleCount = 1
	// CashLetterItemsCount - CheckDetail
	// ToDo: CheckDetailAddendum* and ImageView*
	clc.CashLetterItemsCount = 1
	clc.CashLetterTotalAmount = 100000 // 1000.00
	// ToDo: ImageView*
	clc.CashLetterImagesCount = 0
	clc.ECEInstitutionName = "Wells Fargo"
	clc.SettlementDate = time.Now()
	clc.CreditTotalIndicator = 0
	return clc
}

// testMockCashLetterControl creates an ICL CashLetterControl
func testMockCashLetterControl(t testing.TB) {
	clc := mockCashLetterControl()
	if err := clc.Validate(); err != nil {
		t.Error("mockCashLetterControl does not validate and will break other tests: ", err)
	}
	if clc.recordType != "90" {
		t.Error("recordType does not validate and will break other tests")
	}
	if clc.CashLetterBundleCount != 1 {
		t.Error("CashLetterBundleCount does not validate and will break other tests")
	}
	if clc.CashLetterItemsCount != 1 {
		t.Error("CashLetterItemsCount does not validate and will break other tests")
	}
	if clc.CashLetterTotalAmount != 100000 {
		t.Error("CashLetterTotalAmount does not validate and will break other tests")
	}
	if clc.CashLetterImagesCount != 0 {
		t.Error("CashLetterImagesCount does not validate and will break other tests")
	}
	if clc.ECEInstitutionName != "Wells Fargo" {
		t.Error("ImmediateOriginContactName does not validate and will break other tests")
	}
	if clc.CreditTotalIndicator != 0 {
		t.Error("CreditTotalIndicator does not validate and will break other tests")
	}
}

// TestMockCashLetterControl tests creating an ICL CashLetterControl
func TestMockCashLetterControl(t *testing.T) {
	testMockCashLetterControl(t)
}

// BenchmarkMockCashLetterControl benchmarks creating an ICL CashLetterControl
func BenchmarkMockCashLetterControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCashLetterControl(b)
	}
}
