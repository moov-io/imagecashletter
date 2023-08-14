// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockReturnDetailAddendumC creates a ReturnDetailAddendumC
func mockReturnDetailAddendumC() ReturnDetailAddendumC {
	rdAddendumC := NewReturnDetailAddendumC()
	rdAddendumC.ImageReferenceKeyIndicator = 1
	rdAddendumC.MicrofilmArchiveSequenceNumber = "1A"
	rdAddendumC.LengthImageReferenceKey = "0034"
	rdAddendumC.ImageReferenceKey = "0"
	rdAddendumC.Description = "RD Addendum C"
	rdAddendumC.UserField = ""
	return rdAddendumC
}

func TestReturnDetailAddendumCParseErr(t *testing.T) {
	var r ReturnDetailAddendumC
	r.Parse("ASdsadasda")
	require.Equal(t, 0, r.ImageReferenceKeyIndicator)
	r.Parse("3411A             00340                                 RD Addendum C")
	require.Equal(t, 1, r.ImageReferenceKeyIndicator)
	require.Equal(t, "", r.ImageReferenceKey)
}

// TestMockReturnDetailAddendumCcreates a ReturnDetailAddendumC
func TestMockReturnDetailAddendumC(t *testing.T) {
	rdAddendumC := mockReturnDetailAddendumC()
	require.NoError(t, rdAddendumC.Validate())
	require.Equal(t, "34", rdAddendumC.recordType)
	require.Equal(t, 1, rdAddendumC.ImageReferenceKeyIndicator)
	require.Equal(t, "1A", rdAddendumC.MicrofilmArchiveSequenceNumber)
	require.Equal(t, "0034", rdAddendumC.LengthImageReferenceKey)
	require.Equal(t, "0", rdAddendumC.ImageReferenceKey)
	require.Equal(t, "RD Addendum C", rdAddendumC.Description)
	require.Equal(t, "", rdAddendumC.UserField)
}

// TestParseReturnDetailAddendumC validates parsing a ReturnDetailAddendumC
func TestParseReturnDetailAddendumC(t *testing.T) {
	var line = "3411A             00340                                 RD Addendum C           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	cd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(cd)

	require.NoError(t, r.parseReturnDetailAddendumC())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumC[0]
	require.Equal(t, "34", record.recordType)
	require.Equal(t, "1", record.ImageReferenceKeyIndicatorField())
	require.Equal(t, "1A             ", record.MicrofilmArchiveSequenceNumberField())
	require.Equal(t, "0034", record.LengthImageReferenceKeyField())
	require.Equal(t, "0                                 ", record.ImageReferenceKeyField())
	require.Equal(t, "RD Addendum C  ", record.DescriptionField())
	require.Equal(t, "    ", record.UserFieldField())
	require.Equal(t, "     ", record.reservedField())
}

// testRDAddendumCString validates that a known parsed ReturnDetailAddendumC can return to a string of the same value
func testRDAddendumCString(t testing.TB) {
	var line = "3411A             00340                                 RD Addendum C           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	cd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(cd)

	require.NoError(t, r.parseReturnDetailAddendumC())
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumC[0]

	require.Equal(t, line, record.String())
}

// TestRDAddendumCString tests validating that a known parsed ReturnDetailAddendumC can return to a string of the
// same value
func TestRDAddendumCString(t *testing.T) {
	testRDAddendumCString(t)
}

// BenchmarkRDAddendumCString benchmarks validating that a known parsed ReturnDetailAddendumC
// can return to a string of the same value
func BenchmarkRDAddendumCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDAddendumCString(b)
	}
}

// TestRDAddendumCRecordType validation
func TestRDAddendumCRecordType(t *testing.T) {
	rdAddendumC := mockReturnDetailAddendumC()
	rdAddendumC.recordType = "00"
	err := rdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDAddendumCImageReferenceKeyIndicator validation
func TestRDAddendumCImageReferenceKeyIndicator(t *testing.T) {
	rdAddendumC := mockReturnDetailAddendumC()
	rdAddendumC.ImageReferenceKeyIndicator = 5
	err := rdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageReferenceKeyIndicator", e.FieldName)
}

// TestRDAddendumCImageReferenceKey validation
func TestRDAddendumCImageReferenceKey(t *testing.T) {
	rdAddendumC := mockReturnDetailAddendumC()
	rdAddendumC.ImageReferenceKey = "®©"
	err := rdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageReferenceKey", e.FieldName)
}

// TestRDAddendumCDescription validation
func TestRDAddendumCDescription(t *testing.T) {
	rdAddendumC := mockReturnDetailAddendumC()
	rdAddendumC.Description = "®©"
	err := rdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "Description", e.FieldName)
}

// TestRDAddendumCUserField validation
func TestRDAddendumCUserField(t *testing.T) {
	rdAddendumC := mockReturnDetailAddendumC()
	rdAddendumC.UserField = "®©"
	err := rdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestRDAddendumCFIRecordType validation
func TestRDAddendumCFIRecordType(t *testing.T) {
	rdAddendumC := mockReturnDetailAddendumC()
	rdAddendumC.recordType = ""
	err := rdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestRDAddendumCFIMicrofilmArchiveSequenceNumber validation
func TestRDAddendumCFIMicrofilmArchiveSequenceNumber(t *testing.T) {
	rdAddendumC := mockReturnDetailAddendumC()
	rdAddendumC.MicrofilmArchiveSequenceNumber = "               "
	err := rdAddendumC.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "MicrofilmArchiveSequenceNumber", e.FieldName)
}

// TestRDAddendumCRuneCountInString validates RuneCountInString
func TestRDAddendumCRuneCountInString(t *testing.T) {
	rdAddendumC := NewReturnDetailAddendumC()
	var line = "34"
	rdAddendumC.Parse(line)

	require.Equal(t, "", rdAddendumC.Description)
}
