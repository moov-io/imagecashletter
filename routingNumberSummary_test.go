// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"log"
	"strings"
	"testing"
)

// mockRoutingNumberSummary creates a RoutingNumberSummary
func mockRoutingNumberSummary() *RoutingNumberSummary {
	rns := NewRoutingNumberSummary()
	rns.CashLetterRoutingNumber = "231380104"
	rns.RoutingNumberTotalAmount = 100000
	rns.RoutingNumberItemCount = 1
	rns.UserField = ""
	return rns
}

// TestRoutingNumberSummary creates a ReturnRoutingNumberSummary
func TestRoutingNumberSummary(t *testing.T) {
	rns := mockRoutingNumberSummary()
	if err := rns.Validate(); err != nil {
		t.Error("mockRoutingNumberSummary does not validate and will break other tests: ", err)
	}
}

// TestParseRoutingNumberSummary validates parsing a RoutingNumberSummary
func TestParseRoutingNumberSummary(t *testing.T) {
	var line = "8523138010400000000100000000001                                                 "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	rns := mockRoutingNumberSummary()
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	r.currentCashLetter.AddRoutingNumberSummary(rns)
	r.addCurrentRoutingNumberSummary(rns)

	if err := r.parseRoutingNumberSummary(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentCashLetter.currentRoutingNumberSummary

	if record.recordType != "85" {
		t.Errorf("RecordType Expected '85' got: %v", record.recordType)
	}
	if record.CashLetterRoutingNumberField() != "231380104" {
		t.Errorf("CashLetterRoutingNumber Expected '231380104' got: %v", record.CashLetterRoutingNumberField())
	}
	if record.RoutingNumberTotalAmountField() != "00000000100000" {
		t.Errorf("RoutingNumberTotalAmount Expected '00000000100000' got: %v", record.RoutingNumberTotalAmountField())
	}
	if record.RoutingNumberItemCountField() != "000001" {
		t.Errorf("RoutingNumberItemCount Expected '000001' got: %v", record.RoutingNumberItemCountField())
	}
}

// testRoutingNumberSummaryString validates that a known parsed RoutingNumberSummary can return to a string of the same value
func testRoutingNumberSummaryString(t testing.TB) {
	var line = "8523138010400000000100000000001                                                 "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	rns := mockRoutingNumberSummary()
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	r.currentCashLetter.AddRoutingNumberSummary(rns)
	r.addCurrentRoutingNumberSummary(rns)

	if err := r.parseRoutingNumberSummary(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentRoutingNumberSummary

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestRoutingNumberSummaryString tests validating that a known parsed RoutingNumberSummary an return to a string of the
// same value
func TestRoutingNumberSummaryString(t *testing.T) {
	testRoutingNumberSummaryString(t)
}

// BenchmarkRoutingNumberSummaryString benchmarks validating that a known parsed RoutingNumberSummary
// can return to a string of the same value
func BenchmarkRoutingNumberSummaryString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRoutingNumberSummaryString(b)
	}
}

// TestRoutingNumberSummaryRecordType validation
func TestRoutingNumberSummaryRecordType(t *testing.T) {
	rns := mockRoutingNumberSummary()
	rns.recordType = "00"
	if err := rns.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRoutingNumberSummaryUserField validation
func TestRoutingNumberSummaryUserField(t *testing.T) {
	rns := mockRoutingNumberSummary()
	rns.UserField = "®©"
	if err := rns.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestRoutingNumberSummaryFIRecordType validation
func TestRoutingNumberSummaryFIRecordType(t *testing.T) {
	rns := mockRoutingNumberSummary()
	rns.recordType = ""
	if err := rns.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRoutingNumberSummaryFICashLetterRoutingNumber validation
func TestRoutingNumberSummaryFICashLetterRoutingNumber(t *testing.T) {
	rns := mockRoutingNumberSummary()
	rns.CashLetterRoutingNumber = ""
	if err := rns.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CashLetterRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
