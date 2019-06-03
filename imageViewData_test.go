// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"
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
	if err := ivData.Validate(); err != nil {
		t.Error("mockImageViewData does not validate and will break other tests: ", err)
	}
	if ivData.EceInstitutionRoutingNumber != "121042882" {
		t.Error("EceInstitutionRoutingNumber does not validate")
	}
	if ivData.CycleNumber != "1" {
		t.Error("CycleNumber does not validate")
	}
	if ivData.EceInstitutionItemSequenceNumber != "1             " {
		t.Error(" does not validate")
	}
	if ivData.SecurityOriginatorName != "Sec Orig Name" {
		t.Error("EceInstitutionItemSequenceNumber does not validate")
	}
	if ivData.SecurityAuthenticatorName != "Sec Auth Name" {
		t.Error("SecurityAuthenticatorName does not validate")
	}
	if ivData.SecurityKeyName != "SECURE" {
		t.Error("SecurityKeyName does not validate")
	}
	if ivData.ClippingOrigin != 0 {
		t.Error("ClippingOrigin does not validate")
	}
	if ivData.ClippingCoordinateH1 != "" {
		t.Error(" does not validate")
	}
	if ivData.ClippingCoordinateH2 != "" {
		t.Error("ClippingCoordinateH2 does not validate")
	}
	if ivData.ClippingCoordinateV1 != "" {
		t.Error("ClippingCoordinateV1 does not validate")
	}
	if ivData.ClippingCoordinateV2 != "" {
		t.Error("ClippingCoordinateV2 does not validate")
	}
	if ivData.LengthImageReferenceKey != "0000" {
		t.Error("LengthImageReferenceKey does not validate")
	}
	if ivData.ImageReferenceKey != "" {
		t.Error("ImageReferenceKey does not validate")
	}
	if ivData.LengthDigitalSignature != "0    " {
		t.Error("LengthDigitalSignature does not validate")
	}
	if bytes.Compare(ivData.DigitalSignature, []byte("")) < 0 {
		t.Error("DigitalSignature does not validate")
	}
	if ivData.LengthImageData != "0000001" {
		t.Error("LengthImageData does not validate")
	}
	if bytes.Compare(ivData.ImageData, []byte("")) < 0 {
		t.Error("ImageData does not validate")
	}
}

// testIVDataString validates that a known parsed ImageViewData can return to a string of the same value
func testIVDataString(t testing.TB) {
	//var line = "5212345678020140410  44000000                                                       0                00000    0005591"
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

	if err := r.parseImageViewData(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].ImageViewData[0]

	/*	fmt.Printf("Lineee: %v \n", line)
		fmt.Printf("String: %v \n", record.String())*/

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
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
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVDataCycleNumber validation
func TestIVDataCycleNumber(t *testing.T) {
	ivData := mockImageViewData()
	ivData.CycleNumber = "--"
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CycleNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVDataSecurityOriginatorName validation
func TestIVSecurityOriginatorName(t *testing.T) {
	ivData := mockImageViewData()
	ivData.SecurityOriginatorName = "®©"
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "SecurityOriginatorName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVDataSecurityAuthenticatorName validation
func TestIVSecurityAuthenticatorName(t *testing.T) {
	ivData := mockImageViewData()
	ivData.SecurityAuthenticatorName = "®©"
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "SecurityAuthenticatorName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVDataSecurityKeyName validation
func TestIVSecurityKeyName(t *testing.T) {
	ivData := mockImageViewData()
	ivData.SecurityKeyName = "®©"
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "SecurityKeyName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVDataImageReferenceKey validation
func TestIVImageReferenceKey(t *testing.T) {
	ivData := mockImageViewData()
	ivData.ImageReferenceKey = "®©"
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImageReferenceKey" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestIVDataFIRecordType validation
func TestIVDataFIRecordType(t *testing.T) {
	ivData := mockImageViewData()
	ivData.recordType = ""
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVDataFIEceInstitutionRoutingNumber validation
func TestIVDataFIEceInstitutionRoutingNumber(t *testing.T) {
	ivData := mockImageViewData()
	ivData.EceInstitutionRoutingNumber = ""
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EceInstitutionRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVDataFIBundleBusinessDate validation
func TestIVDataFIBundleBusinessDate(t *testing.T) {
	ivData := mockImageViewData()
	ivData.BundleBusinessDate = time.Time{}
	if err := ivData.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BundleBusinessDate" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
