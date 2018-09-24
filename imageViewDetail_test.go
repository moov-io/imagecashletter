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

// mockImageViewDetail crates an imageViewDetail
func mockImageViewDetail() ImageViewDetail {
	ivDetail := NewImageViewDetail()
	ivDetail.ImageIndicator = 1
	ivDetail.ImageCreatorRoutingNumber = "031300012"
	ivDetail.ImageCreatorDate = time.Now()
	ivDetail.ImageViewFormatIndicator = "00"
	ivDetail.ImageViewCompressionAlgorithm = "00"
	// use of ivDetail.ImageViewDataSize is not recommended
	ivDetail.ImageViewDataSize = "0"
	ivDetail.ViewSideIndicator = 0
	ivDetail.ViewDescriptor = "00"
	ivDetail.DigitalSignatureIndicator = 1
	ivDetail.DigitalSignatureMethod = "00"
	ivDetail.SecurityKeySize = 00000
	ivDetail.ProtectedDataStart = 0000000
	ivDetail.ProtectedDataLength = 0000000
	ivDetail.ImageRecreateIndicator = 0
	ivDetail.UserField = ""
	ivDetail.OverrideIndicator = "0"
	return ivDetail
}

// testMockImageViewDetail creates an ImageViewData
func TestMockImageViewDetail(t *testing.T) {
	ivDetail := mockImageViewDetail()
	if err := ivDetail.Validate(); err != nil {
		t.Error("mockImageViewDetail does not validate and will break other tests: ", err)
	}
}

// TestIVDetailString validates that a known parsed ImageViewDetail can return to a string of the same value
func TestIVDetailString(t *testing.T) {
	var line = "501031300012201809050000000000100000000000000000000000000         0             "
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
