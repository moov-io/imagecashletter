// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
	"time"
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
	if err := ivDetail.Validate(); err != nil {
		t.Error("mockImageViewDetail does not validate and will break other tests: ", err)
	}
	if ivDetail.recordType != "50" {
		t.Error("recordType does not validate")
	}
	if ivDetail.ImageIndicator != 1 {
		t.Error("ImageIndicator does not validate")
	}
	if ivDetail.ImageCreatorRoutingNumber != "031300012" {
		t.Error("ImageCreatorRoutingNumber does not validate")
	}
	if ivDetail.ImageViewFormatIndicator != "00" {
		t.Error("ImageViewFormatIndicator does not validate")
	}
	if ivDetail.ImageViewCompressionAlgorithm != "00" {
		t.Error("ImageViewCompressionAlgorithm does not validate")
	}
	if ivDetail.ImageViewDataSize != "0000000" {
		t.Error("ImageViewDataSize does not validate")
	}
	if ivDetail.ViewSideIndicator != 0 {
		t.Error("ViewSideIndicator does not validate")
	}
	if ivDetail.ViewDescriptor != "00" {
		t.Error("ViewDescriptor does not validate")
	}
	if ivDetail.DigitalSignatureIndicator != 0 {
		t.Error("DigitalSignatureIndicator does not validate")
	}
	if ivDetail.DigitalSignatureMethod != "00" {
		t.Error("DigitalSignatureMethod does not validate")
	}
	if ivDetail.SecurityKeySize != 00000 {
		t.Error(" does not validate")
	}
	if ivDetail.ProtectedDataStart != 0000000 {
		t.Error("ProtectedDataStart does not validate")
	}
	if ivDetail.ProtectedDataLength != 0000000 {
		t.Error("ProtectedDataLength does not validate")
	}
	if ivDetail.ImageRecreateIndicator != 0 {
		t.Error("ImageRecreateIndicator does not validate")
	}
	if ivDetail.UserField != "" {
		t.Error("UserField does not validate")
	}
	if ivDetail.reserved != "" {
		t.Error("reserved does not validate")
	}
	if ivDetail.OverrideIndicator != "0" {
		t.Error("OverrideIndicator does not validate")
	}
	if ivDetail.reservedTwo != "" {
		t.Error("reservedTwo does not validate")
	}

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
	if err := r.parseImageViewDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].ImageViewDetail[0]

	if record.recordType != "50" {
		t.Errorf("RecordType Expected '50' got: %v", record.recordType)
	}
	if record.ImageIndicatorField() != "1" {
		t.Errorf("ImageIndicator Expected '1' got: %v", record.ImageIndicatorField())
	}
	if record.ImageCreatorRoutingNumberField() != "031300012" {
		t.Errorf("ImageCreatorRoutingNumber Expected '031300012' got: %v", record.ImageCreatorRoutingNumberField())
	}
	if record.ImageViewFormatIndicatorField() != "00" {
		t.Errorf("ImageViewFormatIndicator Expected '00' got: %v", record.ImageViewFormatIndicatorField())
	}
	if record.ImageViewCompressionAlgorithmField() != "00" {
		t.Errorf("ImageViewCompressionAlgorithm Expected '00' got: %v", record.ImageViewDataSizeField())
	}
	if record.ImageViewDataSizeField() != "0000000" {
		t.Errorf("ImageViewDataSize Expected '0000000' got: %v", record.ImageViewDataSizeField())
	}
	if record.ViewSideIndicatorField() != "0" {
		t.Errorf("ViewSideIndicator Expected '0' got: %v", record.ViewSideIndicatorField())
	}
	if record.ViewDescriptorField() != "00" {
		t.Errorf("ViewDescriptor Expected '00' got: %v", record.ViewDescriptorField())
	}
	if record.DigitalSignatureIndicatorField() != "0" {
		t.Errorf("DigitalSignatureIndicator Expected '0' got: %v", record.DigitalSignatureIndicatorField())
	}
	if record.DigitalSignatureMethodField() != "00" {
		t.Errorf("DigitalSignatureMethod Expected '00' got: %v", record.DigitalSignatureMethodField())
	}
	if record.SecurityKeySizeField() != "00000" {
		t.Errorf("SecurityKeySize Expected '0' got: %v", record.SecurityKeySizeField())
	}
	if record.ProtectedDataStartField() != "0000000" {
		t.Errorf("ProtectedDataStart Expected '0' got: %v", record.ProtectedDataStartField())
	}
	if record.ProtectedDataLengthField() != "0000000" {
		t.Errorf("ProtectedDataLength Expected '0' got: %v", record.ProtectedDataLengthField())
	}
	if record.ImageRecreateIndicatorField() != "0" {
		t.Errorf("ImageRecreateIndicator Expected '0' got: %v", record.ImageRecreateIndicatorField())
	}
	if record.UserFieldField() != "        " {
		t.Errorf("UserField Expected ' ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != " " {
		t.Errorf("reserved Expected ' ' got: %v", record.reservedField())
	}
	if record.OverrideIndicatorField() != "0" {
		t.Errorf("OverrideIndicator Expected '0' got: %v", record.OverrideIndicatorField())
	}
	if record.reservedTwoField() != "             " {
		t.Errorf("reservedTwo Expected '             ' got: %v", record.reservedTwoField())
	}
}

// TestIVDetailString validates that a known parsed ImageViewDetail can return to a string of the same value
func TestIVDetailString(t *testing.T) {
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

	if err := r.parseImageViewDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].ImageViewDetail[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}
