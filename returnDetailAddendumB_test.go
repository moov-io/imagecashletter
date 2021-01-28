// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
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

// TestParseAddendumBJSONWith3339Date tests parsing ReturnAddendumB with a PayorBankBusinessDate in RFC3339 format
func TestParseAddendumBJSONWith3339Date(t *testing.T) {
	addB := ReturnDetailAddendumB{}
	var testBusinessDateJSON = `{
		"id": "",
		"payorBankName": "",
		"auxiliaryOnUs": "",
		"payorBankSequenceNumber": "3713365076",
		"payorAccountName": "",
		"payorBankBusinessDate": "2021-01-21T00:00:00Z"
}`
	assert.NoError(
		t,
		json.Unmarshal([]byte(testBusinessDateJSON), &addB),
		"Unable to unmarshal ReturnDetailAddendumB",
	)
	assert.Equal(
		t,
		"3713365076",
		addB.PayorBankSequenceNumber,
		"PayorBankSequenceNumber should match JSON",
	)
	parsedTime, timeParseErr := time.Parse(
		time.RFC3339,
		"2021-01-21T00:00:00Z",
	)
	assert.NoError(t, timeParseErr, "Unable to parse test time")
	assert.True(t, addB.PayorBankBusinessDate.Equal(parsedTime), "PayorBankBusinessDate should match JSON")

	marshalled, marshalErr := json.Marshal(addB)
	assert.NoError(t, marshalErr, "Unable to marshal ReturnDetailAddendumB to JSON")
	assert.Contains(
		t,
		string(marshalled),
		"2021-01-21T00:00:00Z",
		"JSON should contain PayorBankBusinessDate",
	)
}

// TestParseAddendumBJSONWithX9Date tests parsing ReturnAddendumB with a PayorBankBusinessDate in X9's YYYYMMDD format
func TestParseAddendumBJSONWithX9Date(t *testing.T) {
	addB := ReturnDetailAddendumB{}
	var testBusinessDateJSON = `{
		"id": "",
		"payorBankName": "",
		"auxiliaryOnUs": "",
		"payorBankSequenceNumber": "3713365076",
		"payorAccountName": "",
		"payorBankBusinessDate": "20210121"
}`
	assert.NoError(
		t,
		json.Unmarshal([]byte(testBusinessDateJSON), &addB),
		"Unable to unmarshal ReturnDetailAddendumB",
	)
	assert.Equal(
		t,
		"3713365076",
		addB.PayorBankSequenceNumber,
		"PayorBankSequenceNumber should match JSON",
	)
	parsedTime, timeParseErr := time.Parse(
		time.RFC3339,
		"2021-01-21T00:00:00Z",
	)
	assert.NoError(t, timeParseErr, "Unable to parse test time")
	assert.True(t, addB.PayorBankBusinessDate.Equal(parsedTime), "PayorBankBusinessDate should match JSON")

	marshalled, marshalErr := json.Marshal(addB)
	assert.NoError(t, marshalErr, "Unable to marshal ReturnDetailAddendumB to JSON")
	assert.Contains(
		t,
		string(marshalled),
		"2021-01-21T00:00:00Z",
		"JSON should contain PayorBankBusinessDate",
	)
}

// TestParseAddendumBJSONWith3339Date tests parsing ReturnAddendumB with a PayorBankBusinessDate that is empty
func TestParseAddendumBJSONWithEmptyDate(t *testing.T) {
	addB := ReturnDetailAddendumB{}
	var testNoBusinessDateJSON =`{
		"id": "",
		"payorBankName": "",
		"auxiliaryOnUs": "",
		"payorBankSequenceNumber": "3713365088",
		"payorAccountName": "",
		"payorBankBusinessDate": ""
}`

	assert.NoError(
		t,
		json.Unmarshal([]byte(testNoBusinessDateJSON), &addB),
		"Unable to unmarshal ReturnDetailAddendumB",
	)
	assert.Equal(
		t,
		"3713365088",
		addB.PayorBankSequenceNumber,
		"PayorBankSequenceNumber should match JSON",
	)
	assert.True(
		t,
		addB.PayorBankBusinessDate.IsZero(),
		"PayorBankBusinessDate should be a time.Time zero value",
	)

	marshalled, marshalErr := json.Marshal(addB)
	assert.NoError(t, marshalErr, "Unable to marshal ReturnDetailAddendumB to JSON")
	assert.NotContains(
		t,
		string(marshalled),
		"0001-01-01",
		"JSON should not contain a time.Time zero value",
	)
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
