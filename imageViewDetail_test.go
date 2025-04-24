// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// mockImageViewDetail crates an imageViewDetail
func mockImageViewDetail() ImageViewDetail {
	ivDetail := NewImageViewDetail()
	ivDetail.ImageIndicator = 1
	ivDetail.ImageCreatorRoutingNumber = "031300012"
	ivDetail.ImageCreatorDate = time.Now()
	ivDetail.ImageViewFormatIndicator = "00"
	ivDetail.ImageViewCompressionAlgorithm = "00"
	// use of ivDetail.ImageViewDataSize is not recommended
	ivDetail.ImageViewDataSize = "0000000"
	ivDetail.ViewSideIndicator = 0
	ivDetail.ViewDescriptor = "00"
	ivDetail.DigitalSignatureIndicator = 0
	ivDetail.DigitalSignatureMethod = "00"
	ivDetail.SecurityKeySize = 00000
	ivDetail.ProtectedDataStart = 0000000
	ivDetail.ProtectedDataLength = 0000000
	ivDetail.ImageRecreateIndicator = 0
	ivDetail.UserField = ""
	ivDetail.OverrideIndicator = "0"
	return ivDetail
}

// TestMockImageViewDetail creates an ImageViewData
func TestMockImageViewDetail(t *testing.T) {
	ivDetail := mockImageViewDetail()
	require.NoError(t, ivDetail.Validate())
	require.Equal(t, "50", ivDetail.recordType)
	require.Equal(t, 1, ivDetail.ImageIndicator)
	require.Equal(t, "031300012", ivDetail.ImageCreatorRoutingNumber)
	require.Equal(t, "00", ivDetail.ImageViewFormatIndicator)
	require.Equal(t, "00", ivDetail.ImageViewCompressionAlgorithm)
	require.Equal(t, "0000000", ivDetail.ImageViewDataSize)
	require.Equal(t, 0, ivDetail.ViewSideIndicator)
	require.Equal(t, "00", ivDetail.ViewDescriptor)
	require.Equal(t, 0, ivDetail.DigitalSignatureIndicator)
	require.Equal(t, "00", ivDetail.DigitalSignatureMethod)
	require.Equal(t, 00000, ivDetail.SecurityKeySize)
	require.Equal(t, 0000000, ivDetail.ProtectedDataStart)
	require.Equal(t, 0000000, ivDetail.ProtectedDataLength)
	require.Equal(t, 0, ivDetail.ImageRecreateIndicator)
	require.Equal(t, "", ivDetail.UserField)
	require.Equal(t, "", ivDetail.reserved)
	require.Equal(t, "0", ivDetail.OverrideIndicator)
	require.Equal(t, "", ivDetail.reservedTwo)

}

// TestParseIVDetail validates parsing an ImageViewDetail
func TestParseIVDetail(t *testing.T) {
	var line = "501031300012201809050000000000000000000000000000000000000         0             "
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
	require.NoError(t, r.parseImageViewDetail())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].ImageViewDetail[0]

	require.Equal(t, "50", record.recordType)
	require.Equal(t, "1", record.ImageIndicatorField())
	require.Equal(t, "031300012", record.ImageCreatorRoutingNumberField())
	require.Equal(t, "00", record.ImageViewFormatIndicatorField())
	require.Equal(t, "00", record.ImageViewCompressionAlgorithmField())
	require.Equal(t, "0000000", record.ImageViewDataSizeField())
	require.Equal(t, "0", record.ViewSideIndicatorField())
	require.Equal(t, "00", record.ViewDescriptorField())
	require.Equal(t, "0", record.DigitalSignatureIndicatorField())
	require.Equal(t, "00", record.DigitalSignatureMethodField())
	require.Equal(t, "     ", record.SecurityKeySizeField())
	require.Equal(t, "0000000", record.ProtectedDataStartField())
	require.Equal(t, "0000000", record.ProtectedDataLengthField())
	require.Equal(t, "0", record.ImageRecreateIndicatorField())
	require.Equal(t, "        ", record.UserFieldField())
	require.Equal(t, " ", record.reservedField())
	require.Equal(t, "0", record.OverrideIndicatorField())
	require.Equal(t, "             ", record.reservedTwoField())
}

// testIVDetailString validates that a known parsed ImageViewDetail can return to a string of the same value
func testIVDetailString(t testing.TB) {
	var line = "5010313000122018090500000000000000000     000000000000000         0             "
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

	require.NoError(t, r.parseImageViewDetail())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].ImageViewDetail[0]

	require.Equal(t, line, record.String())
}

// TestIVDetailString tests validating that a known parsed ImageViewDetail can return to a string of the
// same value
func TestIVDetailString(t *testing.T) {
	testIVDetailString(t)
}

// BenchmarkIVDetailString benchmarks validating that a known parsed ImageViewDetail
// can return to a string of the same value
func BenchmarkIVDetailString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIVDetailString(b)
	}
}

// TestIVDetailRecordType validation
func TestIVDetailRecordType(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.recordType = "00"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestIVDetailImageIndicator validation
func TestIVDetailImageIndicator(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageIndicator = 9
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageIndicator", e.FieldName)
}

// TestIVDetailImageViewFormatIndicator validation
func TestIVDetailImageViewFormatIndicator(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageViewFormatIndicator = "30"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageViewFormatIndicator", e.FieldName)
}

// TestIVDetailImageViewCompressionAlgorithm validation
func TestIVDetailImageViewCompressionAlgorithm(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageViewCompressionAlgorithm = "30"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageViewCompressionAlgorithm", e.FieldName)
}

// TestIVDetailViewSideIndicator validation
func TestIVDetailViewSideIndicator(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ViewSideIndicator = 5
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ViewSideIndicator", e.FieldName)
}

// TestIVDetailViewDescriptor validation
func TestIVDetailViewDescriptor(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ViewDescriptor = "20"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ViewDescriptor", e.FieldName)
}

// TestIVDetailDigitalSignatureIndicator validation
func TestIVDetailDigitalSignatureIndicator(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.DigitalSignatureIndicator = 5
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DigitalSignatureIndicator", e.FieldName)
}

// TestIVDetailDigitalSignatureMethod validation
func TestIVDetailDigitalSignatureMethod(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.DigitalSignatureMethod = "10"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DigitalSignatureMethod", e.FieldName)
}

// TestIVDetailDigitalSignatureMethodFRB validation
func TestIVDetailDigitalSignatureMethodFRB(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.DigitalSignatureMethod = "0"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DigitalSignatureMethod", e.FieldName)
	// "0" should be accepted in FRB compatibility mode
	t.Setenv(FRBCompatibilityMode, "true")
	require.NoError(t, ivDetail.Validate())
}

// TestIVDetailImageRecreateIndicator validation
func TestIVDetailImageRecreateIndicator(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageRecreateIndicator = 5
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageRecreateIndicator", e.FieldName)
}

// TestIVDetailOverrideIndicator validation
func TestIVDetailOverrideIndicator(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.OverrideIndicator = "W"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "OverrideIndicator", e.FieldName)
}

// TestIVDetailUserField validation
func TestIVDetailUserField(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.UserField = "®©"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestIVDetailFIRecordType validation
func TestIVDetailFIRecordType(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.recordType = ""
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestIVDetailFIImageCreatorRoutingNumber validation
func TestIVDetailFIImageCreatorRoutingNumber(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageCreatorRoutingNumber = ""
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageCreatorRoutingNumber", e.FieldName)
}

// TestIVDetailFIImageCreatorRoutingNumberFRB validation
func TestIVDetailFIImageCreatorRoutingNumberFRB(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageCreatorRoutingNumber = "00000000"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageCreatorRoutingNumber", e.FieldName)
	t.Setenv(FRBCompatibilityMode, "true")
	require.NoError(t, ivDetail.Validate())
}

// TestIVDetailFIImageCreatorRoutingNumberZero validation
func TestIVDetailFIImageCreatorRoutingNumberZero(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageCreatorRoutingNumber = "000000000"
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageCreatorRoutingNumber", e.FieldName)
}

// TestIVDetailFIImageCreatorDate validation
func TestIVDetailFIImageCreatorDate(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageCreatorDate = time.Time{}
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageCreatorDate", e.FieldName)
}

func TestIVDetailFIImageCreatorDateFRB(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ImageCreatorDate = time.Time{}
	t.Setenv(FRBCompatibilityMode, "true")
	require.NoError(t, ivDetail.Validate())
}

// TestIVDetailFIViewDescriptor validation
func TestIVDetailFIViewDescriptor(t *testing.T) {
	ivDetail := mockImageViewDetail()
	ivDetail.ViewDescriptor = ""
	err := ivDetail.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ViewDescriptor", e.FieldName)
}

// TestIVDetailRuneCountInString validates RuneCountInString
func TestIVDetailRuneCountInString(t *testing.T) {
	ivDetail := NewImageViewDetail()
	var line = "50"
	ivDetail.Parse(line)

	require.Equal(t, "", ivDetail.ImageCreatorRoutingNumber)
}
