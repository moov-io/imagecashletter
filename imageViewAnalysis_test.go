// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
)

// mockImageViewAnalysis crates an imageViewAnalysis
func mockImageViewAnalysis() ImageViewAnalysis {
	ivAnalysis := NewImageViewAnalysis()
	ivAnalysis.GlobalImageQuality = 2
	ivAnalysis.GlobalImageUsability = 2
	ivAnalysis.ImagingBankSpecificTest = 0
	ivAnalysis.PartialImage = 2
	ivAnalysis.ExcessiveImageSkew = 2
	ivAnalysis.PiggybackImage = 2
	ivAnalysis.TooLightOrTooDark = 2
	ivAnalysis.StreaksAndOrBands = 2
	ivAnalysis.BelowMinimumImageSize = 2
	ivAnalysis.ExceedsMaximumImageSize = 2
	ivAnalysis.ImageEnabledPOD = 1
	ivAnalysis.SourceDocumentBad = 0
	ivAnalysis.DateUsability = 2
	ivAnalysis.PayeeUsability = 2
	ivAnalysis.ConvenienceAmountUsability = 2
	ivAnalysis.AmountInWordsUsability = 2
	ivAnalysis.SignatureUsability = 2
	ivAnalysis.PayorNameAddressUsability = 2
	ivAnalysis.MICRLineUsability = 2
	ivAnalysis.MemoLineUsability = 2
	ivAnalysis.PayorBankNameAddressUsability = 2
	ivAnalysis.PayeeEndorsementUsability = 2
	ivAnalysis.BOFDEndorsementUsability = 2
	ivAnalysis.TransitEndorsementUsability = 2
	return ivAnalysis
}

// testMockImageViewAnalysis creates an ImageViewAnalysis
func TestMockImageViewAnalysis(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	if err := ivAnalysis.Validate(); err != nil {
		t.Error("mockImageViewAnalysis does not validate and will break other tests: ", err)
	}
}

// TestIVAnalysisString validates that a known parsed ImageViewAnalysis can return to a string of the same value
func TestIVAnalysisString(t *testing.T) {
	var line = "542202222222             10222222222222                                         "
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

	if err := r.parseImageViewAnalysis(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].ImageViewAnalysis[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}
