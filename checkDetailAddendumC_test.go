// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"testing"
	"time"
)

// mockCheckDetailAddendumC creates a CheckDetailAddendumC
func mockCheckDetailAddendumC() CheckDetailAddendumC {
	cdAddendumC := NewCheckDetailAddendumC()
	cdAddendumC.RecordNumber = 1
	cdAddendumC.EndorsingBankRoutingNumber = "121042882"
	cdAddendumC.BOFDEndorsementBusinessDate = time.Now()
	cdAddendumC.EndorsingItemSequenceNumber = 1
	cdAddendumC.TruncationIndicator = "Y"
	cdAddendumC.EndorsingConversionIndicator = "1"
	cdAddendumC.EndorsingCorrectionIndicator = 0
	cdAddendumC.ReturnReason = "A"
	cdAddendumC.UserField = ""
	cdAddendumC.EndorsingBankIdentifier = 0
	return cdAddendumC
}

// testMockCheckDetailAddendumC creates an ICL CheckDetailAddendumC
func testMockCheckDetailAddendumC(t testing.TB) {
	cdAddendumC := mockCheckDetailAddendumC()
	if err := cdAddendumC.Validate(); err != nil {
		t.Error("mockBundleHeader does not validate and will break other tests: ", err)
	}
	if cdAddendumC.recordType != "28" {
		t.Error("recordType does not validate and will break other tests")
	}
	if cdAddendumC.RecordNumber != 1 {
		t.Error("RecordNumber does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingBankRoutingNumber != "121042882" {
		t.Error("EndorsingBankRoutingNumber does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingItemSequenceNumber != 1 {
		t.Error("EndorsingItemSequenceNumber does not validate and will break other tests")
	}
	if cdAddendumC.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate and will break other tests")
	}
	if cdAddendumC.ReturnReason != "A" {
		t.Error("ReturnReason does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingConversionIndicator != "1" {
		t.Error("EndorsingConversionIndicator does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingCorrectionIndicator != 0 {
		t.Error("EndorsingCorrectionIndicator does not validate and will break other tests")
	}
	if cdAddendumC.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
	if cdAddendumC.EndorsingBankIdentifier != 0 {
		t.Error("EndorsingBankIdentifier does not validate and will break other tests")
	}
}

// TestMockCheckDetailAddendumC  tests creating an ICL CheckDetailAddendumC
func TestMockCheckDetailAddendumC(t *testing.T) {
	testMockCheckDetailAddendumC(t)
}

// BenchmarkMockCheckDetailAddendumC benchmarks creating an ICL CheckDetailAddendumC
func BenchmarkMockCheckDetailAddendumC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCheckDetailAddendumC(b)
	}
}
