// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"testing"
	"time"
)

// mockCheckDetailAddendumA creates a CheckDetailAddendumA
func mockCheckDetailAddendumA() *CheckDetailAddendumA {
	cdAddendumA := NewCheckDetailAddendumA()
	cdAddendumA.RecordNumber = 1
	cdAddendumA.ReturnLocationRoutingNumber = "121042882"
	cdAddendumA.BOFDEndorsementDate = time.Now()
	cdAddendumA.BOFDItemSequenceNumber = 1
	cdAddendumA.BOFDAccountNumber = "938383"
	cdAddendumA.BOFDBranchCode = "01"
	cdAddendumA.PayeeName = "Test Payee"
	cdAddendumA.TruncationIndicator = "Y"
	cdAddendumA.BOFDConversionIndicator = "1"
	cdAddendumA.BOFDCorrectionIndicator = 0
	cdAddendumA.UserField = ""
	return cdAddendumA
}

// testMockCheckDetailAddendumA creates an ICL CheckDetailAddendumA
func testMockCheckDetailAddendumA(t testing.TB) {
	cdAddendumA := mockCheckDetailAddendumA()
	if err := cdAddendumA.Validate(); err != nil {
		t.Error("mockBundleHeader does not validate and will break other tests: ", err)
	}
	if cdAddendumA.recordType != "26" {
		t.Error("recordType does not validate and will break other tests")
	}
	if cdAddendumA.RecordNumber != 1 {
		t.Error("RecordNumber does not validate and will break other tests")
	}
	if cdAddendumA.ReturnLocationRoutingNumber != "121042882" {
		t.Error("ReturnLocationRoutingNumber does not validate and will break other tests")
	}
	if cdAddendumA.BOFDItemSequenceNumber != 1 {
		t.Error("BOFDItemSequenceNumber does not validate and will break other tests")
	}
	if cdAddendumA.BOFDAccountNumber != "938383" {
		t.Error("BOFDAccountNumber does not validate and will break other tests")
	}
	if cdAddendumA.BOFDBranchCode != "01" {
		t.Error("BOFDBranchCode does not validate and will break other tests")
	}
	if cdAddendumA.PayeeName != "Test Payee" {
		t.Error("PayeeName does not validate and will break other tests")
	}
	if cdAddendumA.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate and will break other tests")
	}
	if cdAddendumA.BOFDConversionIndicator != "1" {
		t.Error("BOFDConversionIndicator does not validate and will break other tests")
	}
	if cdAddendumA.BOFDCorrectionIndicator != 0 {
		t.Error("BOFDCorrectionIndicator does not validate and will break other tests")
	}
	if cdAddendumA.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
}

// TestMockCheckDetailAddendumA  tests creating an ICL CheckDetailAddendumA
func TestMockCheckDetailAddendumA(t *testing.T) {
	testMockCheckDetailAddendumA(t)
}

// BenchmarkMockCheckDetailAddendumA benchmarks creating an ICL CheckDetailAddendumA
func BenchmarkMockCheckDetailAddendumA(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCheckDetailAddendumA(b)
	}
}
