// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/base64"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// mockImageViewData crates an imageViewData
func mockImageViewData() ImageViewData {
	ivData := NewImageViewData()
	ivData.EceInstitutionRoutingNumber = "121042882"
	ivData.BundleBusinessDate = time.Now()
	ivData.CycleNumber = "1"
	ivData.EceInstitutionItemSequenceNumber = "1             "
	ivData.SecurityOriginatorName = "Sec Orig Name"
	ivData.SecurityAuthenticatorName = "Sec Auth Name"
	ivData.SecurityKeyName = "SECURE"
	ivData.ClippingOrigin = 0
	ivData.ClippingCoordinateH1 = ""
	ivData.ClippingCoordinateH2 = ""
	ivData.ClippingCoordinateV1 = ""
	ivData.ClippingCoordinateV2 = ""
	ivData.LengthImageReferenceKey = "0000"
	ivData.ImageReferenceKey = ""
	ivData.LengthDigitalSignature = "0    "
	ivData.DigitalSignature = []byte("")
	ivData.LengthImageData = "0000001"
	ivData.ImageData = []byte("")
	return ivData
}

// testMockImageViewData creates an ImageViewData
func TestMockImageViewData(t *testing.T) {
	ivData := mockImageViewData()
	require.NoError(t, ivData.Validate())
	require.Equal(t, "121042882", ivData.EceInstitutionRoutingNumber)
	require.Equal(t, "1", ivData.CycleNumber)
	require.Equal(t, "1             ", ivData.EceInstitutionItemSequenceNumber)
	require.Equal(t, "Sec Orig Name", ivData.SecurityOriginatorName)
	require.Equal(t, "Sec Auth Name", ivData.SecurityAuthenticatorName)
	require.Equal(t, "SECURE", ivData.SecurityKeyName)
	require.Equal(t, 0, ivData.ClippingOrigin)
	require.Empty(t, ivData.ClippingCoordinateH1)
	require.Empty(t, ivData.ClippingCoordinateH2)
	require.Empty(t, ivData.ClippingCoordinateV1)
	require.Empty(t, ivData.ClippingCoordinateV2)
	require.Equal(t, "0000", ivData.LengthImageReferenceKey)
	require.Equal(t, "", ivData.ImageReferenceKey)
	require.Equal(t, "0    ", ivData.LengthDigitalSignature)
	require.Equal(t, []byte(""), ivData.DigitalSignature)
	require.Equal(t, "0000001", ivData.LengthImageData)
	require.Equal(t, []byte(""), ivData.ImageData)
}

// testIVDataString validates that a known parsed ImageViewData can return to a string of the same value
func testIVDataString(t testing.TB) {
	var line = "5212104288220180915  1                                                              0                00000    0000001 "
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

	require.NoError(t, r.parseImageViewData())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].ImageViewData[0]
	require.Equal(t, line, record.String())
}

func TestIVDParseCrash(t *testing.T) {
	iv := &ImageViewData{}
	iv.Parse(`20000000000040000000020000100300001003000000000000000000000000000000000000000000`)
	require.Equal(t, "", iv.ImageReferenceKey)

	prefix := "20000000000040000000020000100300001003"

	iv.Parse(prefix + strings.Repeat("0", 110-len(prefix)-1))
	require.Equal(t, "", iv.ImageReferenceKey)

	iv.Parse(prefix + strings.Repeat("0", 117-len(prefix)-1))
	require.NotEmpty(t, iv.LengthDigitalSignature)

	d := mockImageViewData()
	iv.Parse(d.String()[:117])
	require.Empty(t, iv.ImageData)
}

// TestIVDataString tests validating that a known parsed ImageViewData an return to a string of the
// same value
func TestIVDataString(t *testing.T) {
	testIVDataString(t)
}

// BenchmarkIVDataString benchmarks validating that a known parsed ImageViewData
// can return to a string of the same value
func BenchmarkIVDataString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIVDataString(b)
	}
}

// TestIVDataRecordType validation
func TestIVDataRecordType(t *testing.T) {
	ivData := mockImageViewData()
	ivData.recordType = "00"
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestIVDataCycleNumber validation
func TestIVDataCycleNumber(t *testing.T) {
	ivData := mockImageViewData()
	ivData.CycleNumber = "--"
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "CycleNumber", e.FieldName)
}

// TestIVDataSecurityOriginatorName validation
func TestIVSecurityOriginatorName(t *testing.T) {
	ivData := mockImageViewData()
	ivData.SecurityOriginatorName = "®©"
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "SecurityOriginatorName", e.FieldName)
}

// TestIVDataSecurityAuthenticatorName validation
func TestIVSecurityAuthenticatorName(t *testing.T) {
	ivData := mockImageViewData()
	ivData.SecurityAuthenticatorName = "®©"
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "SecurityAuthenticatorName", e.FieldName)
}

// TestIVDataSecurityKeyName validation
func TestIVSecurityKeyName(t *testing.T) {
	ivData := mockImageViewData()
	ivData.SecurityKeyName = "®©"
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "SecurityKeyName", e.FieldName)
}

// TestIVDataImageReferenceKey validation
func TestIVImageReferenceKey(t *testing.T) {
	ivData := mockImageViewData()
	ivData.ImageReferenceKey = "®©"
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageReferenceKey", e.FieldName)
}

// Field Inclusion

// TestIVDataFIRecordType validation
func TestIVDataFIRecordType(t *testing.T) {
	ivData := mockImageViewData()
	ivData.recordType = ""
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestIVDataFIEceInstitutionRoutingNumber validation
func TestIVDataFIEceInstitutionRoutingNumber(t *testing.T) {
	ivData := mockImageViewData()
	ivData.EceInstitutionRoutingNumber = ""
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "EceInstitutionRoutingNumber", e.FieldName)
}

// TestIVDataFIBundleBusinessDate validation
func TestIVDataFIBundleBusinessDate(t *testing.T) {
	ivData := mockImageViewData()
	ivData.BundleBusinessDate = time.Time{}
	err := ivData.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BundleBusinessDate", e.FieldName)
}

func TestImageViewData_Base64String(t *testing.T) {
	data := base64Encode("hello, world")

	ivData := mockImageViewData()
	ivData.ImageData = data

	output := ivData.String()
	require.Contains(t, output, "hello, world")
}

func TestDecodeImageData(t *testing.T) {
	data := base64Encode("hello, world")

	ivData := &ImageViewData{
		ImageData: data,
	}

	decoded, err := ivData.DecodeImageData()
	require.NoError(t, err)
	require.NotEmpty(t, decoded)
	require.Equal(t, "hello, world", string(decoded))
}

func base64Encode(in string) []byte {
	input := []byte(in)
	out := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(out, input)
	return out
}
