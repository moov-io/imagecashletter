// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.Equal(t, "", r.PayorBankName)
}

// TestMockReturnDetailAddendumB creates a ReturnDetailAddendumB
func TestMockReturnDetailAddendumB(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	require.NoError(t, rdAddendumB.Validate())
	require.Equal(t, "33", rdAddendumB.recordType)
	require.Equal(t, "Payor Bank Name", rdAddendumB.PayorBankName)
	require.Equal(t, "123456789", rdAddendumB.AuxiliaryOnUs)
	require.Equal(t, "1              ", rdAddendumB.PayorBankSequenceNumber)
	require.Equal(t, "Payor Account Name", rdAddendumB.PayorAccountName)
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

	require.NoError(t, r.parseReturnDetailAddendumB())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumB[0]

	require.Equal(t, "33", record.recordType)
	require.Equal(t, "Payor Bank Name   ", record.PayorBankNameField())
	require.Equal(t, "      123456789", record.AuxiliaryOnUsField())
	require.Equal(t, "1              ", record.PayorBankSequenceNumberField())
	require.Equal(t, "Payor Account Name    ", record.PayorAccountNameField())
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

	require.NoError(t, r.parseReturnDetailAddendumB())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumB[0]

	require.Equal(t, line, record.String())
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
	require.NoError(
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
	require.NoError(t, timeParseErr, "Unable to parse test time")
	assert.True(t, addB.PayorBankBusinessDate.Equal(parsedTime), "PayorBankBusinessDate should match JSON")

	marshalled, marshalErr := json.Marshal(addB)
	require.NoError(t, marshalErr, "Unable to marshal ReturnDetailAddendumB to JSON")
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
	require.NoError(
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
	require.NoError(t, timeParseErr, "Unable to parse test time")
	assert.True(t, addB.PayorBankBusinessDate.Equal(parsedTime), "PayorBankBusinessDate should match JSON")

	marshalled, marshalErr := json.Marshal(addB)
	require.NoError(t, marshalErr, "Unable to marshal ReturnDetailAddendumB to JSON")
	assert.Contains(
		t,
		string(marshalled),
		"2021-01-21T00:00:00Z",
		"JSON should contain PayorBankBusinessDate",
	)
}

// TestParseAddendumBJSONWith3339Date tests parsing ReturnAddendumB with a PayorBankBusinessDate that is empty
func TestParseAddendumBJSONWithEmptyDate(t *testing.T) {
	addB := &ReturnDetailAddendumB{}
	var testNoBusinessDateJSON = `{
		"id": "",
		"payorBankName": "",
		"auxiliaryOnUs": "",
		"payorBankSequenceNumber": "3713365088",
		"payorAccountName": "",
		"payorBankBusinessDate": ""
	}`

	require.NoError(
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
	require.NoError(t, marshalErr, "Unable to marshal ReturnDetailAddendumB to JSON")
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
	err := rdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDAddendumBPayorBankName validation
func TestRDAddendumBPayorBankName(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.PayorBankName = "®©"
	err := rdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankName", e.FieldName)
}

// TestRDAddendumBPayorAccountName validation
func TestRDAddendumBPayorAccountName(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.PayorAccountName = "®©"
	err := rdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorAccountName", e.FieldName)
}

// Field Inclusion

// TestRDAddendumBFIRecordType validation
func TestRDAddendumBFIRecordType(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.recordType = ""
	err := rdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDAddendumBFIPayorBankSequenceNumber validation
func TestRDAddendumBFIPayorBankSequenceNumber(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.PayorBankSequenceNumber = "               "
	err := rdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankSequenceNumber", e.FieldName)
}

// TestRDAddendumBFIPayorBankBusinessDate validation
func TestRDAddendumPayorBankBusinessDate(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	date := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	rdAddendumB.PayorBankBusinessDate = date
	err := rdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankBusinessDate", e.FieldName)
	require.Contains(t, e.Msg, msgInvalidDate)
}
