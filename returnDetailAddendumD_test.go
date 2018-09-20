// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"testing"
	"time"
)

// mockReturnDetailAddendumD creates a ReturnDetailAddendumD
func mockReturnDetailAddendumD() ReturnDetailAddendumD {
	rdAddendumD := NewReturnDetailAddendumD()
	rdAddendumD.RecordNumber = 1
	rdAddendumD.EndorsingBankRoutingNumber = "121042882"
	rdAddendumD.BOFDEndorsementBusinessDate = time.Now()
	rdAddendumD.EndorsingBankItemSequenceNumber = "1              "
	rdAddendumD.TruncationIndicator = "Y"
	rdAddendumD.EndorsingBankConversionIndicator = "1"
	rdAddendumD.EndorsingBankCorrectionIndicator = 0
	rdAddendumD.ReturnReason = "A"
	rdAddendumD.UserField = ""
	rdAddendumD.EndorsingBankIdentifier = 0
	rdAddendumD.UserField = ""
	return rdAddendumD
}

// testMockReturnDetailAddendumD creates a ReturnDetailAddendumD
func testMockReturnDetailAddendumD(t testing.TB) {
	rdAddendumD := mockReturnDetailAddendumD()
	if err := rdAddendumD.Validate(); err != nil {
		t.Error("MockReturnDetailAddendumD does not validate and will break other tests: ", err)
	}
	if rdAddendumD.recordType != "35" {
		t.Error("recordType does not validate and will break other tests")
	}
	if rdAddendumD.RecordNumber != 1 {
		t.Error("RecordNumber does not validate and will break other tests")
	}
	if rdAddendumD.EndorsingBankRoutingNumber != "121042882" {
		t.Error("EndorsingBankRoutingNumber does not validate and will break other tests")
	}
	if rdAddendumD.EndorsingBankItemSequenceNumber != "1              " {
		t.Error("EndorsingBankItemSequenceNumber does not validate and will break other tests")
	}
	if rdAddendumD.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate and will break other tests")
	}
	if rdAddendumD.ReturnReason != "A" {
		t.Error("ReturnReason does not validate and will break other tests")
	}
	if rdAddendumD.EndorsingBankConversionIndicator != "1" {
		t.Error("EndorsingBankConversionIndicator does not validate and will break other tests")
	}
	if rdAddendumD.EndorsingBankCorrectionIndicator != 0 {
		t.Error("EndorsingBankCorrectionIndicator does not validate and will break other tests")
	}
	if rdAddendumD.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
	if rdAddendumD.EndorsingBankIdentifier != 0 {
		t.Error("EndorsingBankIdentifier does not validate and will break other tests")
	}
}

// TestMockReturnDetailAddendumD tests creating a ReturnDetailAddendumD
func TestMockReturnDetailAddendumD(t *testing.T) {
	testMockReturnDetailAddendumD(t)
}

// BenchmarkMockReturnDetailAddendumD benchmarks creating a ReturnDetailAddendumD
func BenchmarkMockReturnDetailAddendumD(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockReturnDetailAddendumD(b)
	}
}
