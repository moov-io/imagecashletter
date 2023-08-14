// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockCheckDetailAddendumB creates a CheckDetailAddendumB
func mockCheckDetailAddendumB() CheckDetailAddendumB {
	cdAddendumB := NewCheckDetailAddendumB()
	cdAddendumB.ImageReferenceKeyIndicator = 1
	cdAddendumB.MicrofilmArchiveSequenceNumber = "1A             "
	cdAddendumB.LengthImageReferenceKey = "0034"
	cdAddendumB.ImageReferenceKey = "0"
	cdAddendumB.Description = "CD Addendum B"
	cdAddendumB.UserField = ""
	return cdAddendumB
}

func TestCheckDetailAddendumBParseErr(t *testing.T) {
	var c CheckDetailAddendumB
	c.Parse("asdhfakjfsa")
	require.Equal(t, 0, c.ImageReferenceKeyIndicator)
	c.Parse("2711A             00340                                 CD Addendum B")
	require.Equal(t, 1, c.ImageReferenceKeyIndicator)
	require.Equal(t, "", c.ImageReferenceKey)
}

// TestMockCheckDetailAddendumB creates a CheckDetailAddendumB
func TestMockCheckDetailAddendumB(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	require.NoError(t, cdAddendumB.Validate())
	require.Equal(t, "27", cdAddendumB.recordType)
	require.Equal(t, 1, cdAddendumB.ImageReferenceKeyIndicator)
	require.Equal(t, "1A             ", cdAddendumB.MicrofilmArchiveSequenceNumber)
	require.Equal(t, "0034", cdAddendumB.LengthImageReferenceKey)
	require.Equal(t, "0", cdAddendumB.ImageReferenceKey)
	require.Equal(t, "CD Addendum B", cdAddendumB.Description)
	require.Equal(t, "", cdAddendumB.UserField)
}

// TestParseCheckDetailAddendumB validates parsing a CheckDetailAddendumB
func TestParseCheckDetailAddendumB(t *testing.T) {
	var line = "2711A             00340                                 CD Addendum B           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	require.NoError(t, r.parseCheckDetailAddendumB())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumB[0]
	require.Equal(t, "27", record.recordType)
	require.Equal(t, "1", record.ImageReferenceKeyIndicatorField())
	require.Equal(t, "1A             ", record.MicrofilmArchiveSequenceNumberField())
	require.Equal(t, "0034", record.LengthImageReferenceKeyField())
	require.Equal(t, "0                                 ", record.ImageReferenceKeyField())
	require.Equal(t, "CD Addendum B  ", record.DescriptionField())
	require.Equal(t, "    ", record.UserFieldField())
	require.Equal(t, "     ", record.reservedField())
}

// testCDAddendumBString validates that a known parsed CheckDetailAddendumB can return to a string of the same value
func testCDAddendumBString(t testing.TB) {
	var line = "2711A             00340                                 CD Addendum B           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	require.NoError(t, r.parseCheckDetailAddendumB())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumB[0]

	require.Equal(t, line, record.String())
}

// TestCDAddendumB String tests validating that a known parsed CheckDetailAddendumB can return to a string of the
// same value
func TestCDAddendumBString(t *testing.T) {
	testCDAddendumBString(t)
}

// BenchmarkCDAddendumB String benchmarks validating that a known parsed CheckDetailAddendumB
// can return to a string of the same value
func BenchmarkCDAddendumBString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumBString(b)
	}
}

// TestCDAddendumBRecordType validation
func TestCDAddendumBRecordType(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.recordType = "00"
	err := cdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCDAddendumBImageReferenceKeyIndicator validation
func TestCDAddendumBImageReferenceKeyIndicator(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.ImageReferenceKeyIndicator = 5
	err := cdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageReferenceKeyIndicator", e.FieldName)
}

// TestCDAddendumBImageReferenceKey validation
func TestCDAddendumBImageReferenceKey(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.ImageReferenceKey = "®©"
	err := cdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageReferenceKey", e.FieldName)
}

// TestCDAddendumBDescription validation
func TestCDAddendumBDescription(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.Description = "®©"
	err := cdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "Description", e.FieldName)
}

// TestCDAddendumBUserField validation
func TestCDAddendumBUserField(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.UserField = "®©"
	err := cdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestCDAddendumBFIRecordType validation
func TestCDAddendumBFIRecordType(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.recordType = ""
	err := cdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestCDAddendumBFIMicrofilmArchiveSequenceNumber validation
func TestCDAddendumBFIMicrofilmArchiveSequenceNumber(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.MicrofilmArchiveSequenceNumber = "               "
	err := cdAddendumB.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "MicrofilmArchiveSequenceNumber", e.FieldName)
}

// End FieldInclusion

// TestNBSMFieldTrim validation
func TestNBSMFieldTrim(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.AuxiliaryOnUs = "12345678901234567890"
	require.Len(t, rdAddendumB.AuxiliaryOnUsField(), 15)
}
