// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"testing"
	"time"
)

// mockReturnDetailAddendumB creates a ReturnDetailAddendumB
func mockReturnDetailAddendumB() ReturnDetailAddendumB {
	rdAddendumB := NewReturnDetailAddendumB()
	rdAddendumB.PayorBankName = "Payor Bank Name"
	rdAddendumB.AuxiliaryOnUs = "123456789"
	rdAddendumB.PayorBankSequenceNumber = "1              "
	rdAddendumB.PayorBankBusinessDate = time.Now()
	rdAddendumB.PayorAccountName = "Payor Account Name"
	return rdAddendumB
}

// testMockReturnDetailAddendumB creates a ReturnDetailAddendumB
func testMockReturnDetailAddendumB(t testing.TB) {
	rdAddendumB := mockReturnDetailAddendumB()
	if err := rdAddendumB.Validate(); err != nil {
		t.Error("MockReturnDetailAddendumB does not validate and will break other tests: ", err)
	}
	if rdAddendumB.recordType != "33" {
		t.Error("recordType does not validate and will break other tests")
	}
	if rdAddendumB.PayorBankName != "Payor Bank Name" {
		t.Error("PayorBankName does not validate and will break other tests")
	}
	if rdAddendumB.AuxiliaryOnUs != "123456789" {
		t.Error("AuxiliaryOnUs does not validate and will break other tests")
	}
	if rdAddendumB.PayorBankSequenceNumber != "1              " {
		t.Error("PayorBankSequenceNumber does not validate and will break other tests")
	}
	if rdAddendumB.PayorAccountName != "Payor Account Name" {
		t.Error("PayorAccountName does not validate and will break other tests")
	}
}

// TestMockReturnDetailAddendumB tests creating a ReturnDetailAddendumB
func TestMockReturnDetailAddendumB(t *testing.T) {
	testMockReturnDetailAddendumB(t)
}

// BenchmarkMockReturnDetailAddendumB benchmarks creating a ReturnDetailAddendumB
func BenchmarkMockReturnDetailAddendumB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockReturnDetailAddendumB(b)
	}
}
