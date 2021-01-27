// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
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

func TestReturnDetailAddendumBParseErr(t *testing.T) {
	var r ReturnDetailAddendumB
	r.Parse("Asdjashfakjfa")
	if r.PayorBankName != "" {
		t.Errorf("r.PayorBankName=%s", r.PayorBankName)
	}
}

// TestMockReturnDetailAddendumB creates a ReturnDetailAddendumB
func TestMockReturnDetailAddendumB(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	if err := rdAddendumB.Validate(); err != nil {
		t.Error("MockReturnDetailAddendumB does not validate and will break other tests: ", err)
	}
	if rdAddendumB.recordType != "33" {
		t.Error("recordType does not validate")
	}
	if rdAddendumB.PayorBankName != "Payor Bank Name" {
		t.Error("PayorBankName does not validate")
	}
	if rdAddendumB.AuxiliaryOnUs != "123456789" {
		t.Error("AuxiliaryOnUs does not validate")
	}
	if rdAddendumB.PayorBankSequenceNumber != "1              " {
		t.Error("PayorBankSequenceNumber does not validate")
	}
	if rdAddendumB.PayorAccountName != "Payor Account Name" {
		t.Error("PayorAccountName does not validate")
	}
}

// TestParseReturnDetailAddendumB validates parsing a ReturnDetailAddendumB
func TestParseReturnDetailAddendumB(t *testing.T) {
	var line = "33Payor Bank Name         1234567891              20180905Payor Account Name    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(rd)

	if err := r.parseReturnDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumB[0]

	if record.recordType != "33" {
		t.Errorf("RecordType Expected '33' got: %v", record.recordType)
	}
	if record.PayorBankNameField() != "Payor Bank Name   " {
		t.Errorf("PayorBankName Expected 'Payor Bank Name   ' got: %v", record.PayorBankNameField())
	}
	if record.AuxiliaryOnUsField() != "      123456789" {
		t.Errorf("AuxiliaryOnUs Expected '      123456789' got: %v", record.AuxiliaryOnUsField())
	}
	if record.PayorBankSequenceNumberField() != "1              " {
		t.Errorf("PayorBankSequenceNumber Expected '1              ' got: %v", record.PayorBankSequenceNumberField())
	}
	if record.PayorAccountNameField() != "Payor Account Name    " {
		t.Errorf("PayorAccountName Expected 'Payor Account Name    ' got: %v", record.PayorAccountNameField())
	}
}

// testRDAddendumBString validates that a known parsed ReturnDetailAddendumB can return to a string of the same value
func testRDAddendumBString(t testing.TB) {
	var line = "33Payor Bank Name         1234567891              20180905Payor Account Name    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(rd)

	if err := r.parseReturnDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumB[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestParseAddendumBJSON tests parsing a known ReturnDetailAddendumB
func TestParseAddendumBJSON(t *testing.T) {
	addB := ReturnDetailAddendumB{}
	var testBusinessDateJSON =`{
		"id": "",
		"payorBankName": "",
		"auxiliaryOnUs": "",
		"payorBankSequenceNumber": "3713365076",
		"payorAccountName": "",
		"payorBankBusinessDate": "2021-01-21T00:00:00Z"
}`
	if err := json.Unmarshal([]byte(testBusinessDateJSON), &addB); err != nil {
		t.Errorf("Unable to unmarshal ReturnDetailAddendumB JSON: %v", err)
	}
	if addB.PayorBankSequenceNumber != "3713365076" {
		t.Errorf("PayorBankSequenceNumber Expected '3713365076' got: %v", addB.PayorBankSequenceNumber)
	}
	parsedTime, timeParseErr := time.Parse(time.RFC3339,"2021-01-21T00:00:00Z")
	if timeParseErr != nil {
		t.Errorf("Unable to parse test time: %v", timeParseErr)
	}
	if !addB.PayorBankBusinessDate.Equal(parsedTime) {
		t.Errorf("PayorBankBusinessDate Expected '2021-01-21T00:00:00Z' got: %v", addB.PayorBankBusinessDate)
	}

	var testNoBusinessDateJSON =`{
		"id": "",
		"payorBankName": "",
		"auxiliaryOnUs": "",
		"payorBankSequenceNumber": "3713365088",
		"payorAccountName": "",
		"payorBankBusinessDate": ""
}`
	if err := json.Unmarshal([]byte(testNoBusinessDateJSON), &addB); err != nil {
		t.Errorf("Unable to unmarshal ReturnDetailAddendumB JSON: %v", err)
	}
	if addB.PayorBankSequenceNumber != "3713365088" {
		t.Errorf("PayorBankSequenceNumber Expected '3713365088' got: %v", addB.PayorBankSequenceNumber)
	}
	if !addB.PayorBankBusinessDate.IsZero() {
		t.Errorf("PayorBankBusinessDate Expected zero-date got: %v", addB.PayorBankBusinessDate)
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

// TestRDAddendumBRecordType validation
func TestRDAddendumBRecordType(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.recordType = "00"
	if err := rdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumBPayorBankName validation
func TestRDAddendumBPayorBankName(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.PayorBankName = "®©"
	if err := rdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumBPayorAccountName validation
func TestRDAddendumBPayorAccountName(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.PayorAccountName = "®©"
	if err := rdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorAccountName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestRDAddendumBFIRecordType validation
func TestRDAddendumBFIRecordType(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.recordType = ""
	if err := rdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumBFIPayorBankSequenceNumber validation
func TestRDAddendumBFIPayorBankSequenceNumber(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.PayorBankSequenceNumber = "               "
	if err := rdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumBFIPayorBankBusinessDate validation
func TestRDAddendumPayorBankBusinessDate(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.PayorBankBusinessDate = time.Time{}
	if err := rdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankBusinessDate" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
