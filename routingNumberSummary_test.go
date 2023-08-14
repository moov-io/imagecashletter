// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestRoutingNumberSummaryParseErr(t *testing.T) {
	var r RoutingNumberSummary
	r.Parse("asdlahsakjajf")
	require.Equal(t, "", r.CashLetterRoutingNumber)
}

// TestRoutingNumberSummary creates a ReturnRoutingNumberSummary
func TestRoutingNumberSummary(t *testing.T) {
	rns := mockRoutingNumberSummary()
	require.NoError(t, rns.Validate())
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

	require.NoError(t, r.parseRoutingNumberSummary())
	record := r.currentCashLetter.currentRoutingNumberSummary

	require.Equal(t, "85", record.recordType)
	require.Equal(t, "231380104", record.CashLetterRoutingNumberField())
	require.Equal(t, "00000000100000", record.RoutingNumberTotalAmountField())
	require.Equal(t, "000001", record.RoutingNumberItemCountField())
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

	require.NoError(t, r.parseRoutingNumberSummary())
	record := r.currentCashLetter.currentRoutingNumberSummary

	require.Equal(t, line, record.String())
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
	err := rns.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRoutingNumberSummaryUserField validation
func TestRoutingNumberSummaryUserField(t *testing.T) {
	rns := mockRoutingNumberSummary()
	rns.UserField = "®©"
	err := rns.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestRoutingNumberSummaryFIRecordType validation
func TestRoutingNumberSummaryFIRecordType(t *testing.T) {
	rns := mockRoutingNumberSummary()
	rns.recordType = ""
	err := rns.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRoutingNumberSummaryFICashLetterRoutingNumber validation
func TestRoutingNumberSummaryFICashLetterRoutingNumber(t *testing.T) {
	rns := mockRoutingNumberSummary()
	rns.CashLetterRoutingNumber = ""
	err := rns.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CashLetterRoutingNumber", e.FieldName)
}
