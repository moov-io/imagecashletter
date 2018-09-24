// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
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
}

// TestIVDataString validates that a known parsed ImageViewData can return to a string of the same value
func TestIVDataString(t *testing.T) {
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
