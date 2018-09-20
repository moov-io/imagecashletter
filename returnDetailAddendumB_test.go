// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"log"
	"strings"
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

// testRDAddendumBString validates that a known parsed ReturnDetailAddendumB can return to a string of the same value
func testRDAddendumBString(t testing.TB) {
	//var line = "3301121042882201809051              Y10A                   0                    "
	var line = "33Payor Bank Name         1234567891              20180905Payor Account Name    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewReturnBundle(bh)
	r.currentCashLetter.AddReturnBundle(rb)
	r.addCurrentReturnBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentReturnBundle.AddReturnDetail(rd)

	if err := r.parseReturnDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentReturnBundle.GetReturns()[0].ReturnDetailAddendumB[0]

	fmt.Printf("Lineee: %v \n", line)
	fmt.Printf("String: %v \n", record.String())

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestRDAddendumBString tests validating that a known parsed ReturnDetailAddendumB can return to a string of the
// same value
func TestRDAddendumBString(t *testing.T) {
	testRDAddendumBString(t)
}

// BenchmarkRDAddendumBString benchmarks validating that a known parsed ReturnDetailAddendumB
// can return to a string of the same value
func BenchmarkRDAddendumBString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDAddendumBString(b)
	}
}
