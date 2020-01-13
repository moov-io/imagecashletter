// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

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

// TestMockImageViewAnalysis creates an ImageViewAnalysis
func TestMockImageViewAnalysis(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	if err := ivAnalysis.Validate(); err != nil {
		t.Error("mockImageViewAnalysis does not validate and will break other tests: ", err)
	}
	if ivAnalysis.GlobalImageQuality != 2 {
		t.Error("GlobalImageQuality does not validate")
	}
	if ivAnalysis.GlobalImageUsability != 2 {
		t.Error("GlobalImageUsability does not validate")
	}
	if ivAnalysis.ImagingBankSpecificTest != 0 {
		t.Error("ImagingBankSpecificTest does not validate")
	}
	if ivAnalysis.PartialImage != 2 {
		t.Error("PartialImage does not validate")
	}
	if ivAnalysis.ExcessiveImageSkew != 2 {
		t.Error("ExcessiveImageSkew does not validate")
	}
	if ivAnalysis.PiggybackImage != 2 {
		t.Error("PiggybackImage does not validate")
	}
	if ivAnalysis.TooLightOrTooDark != 2 {
		t.Error("TooLightOrTooDark does not validate")
	}
	if ivAnalysis.StreaksAndOrBands != 2 {
		t.Error("StreaksAndOrBands does not validate")
	}
	if ivAnalysis.BelowMinimumImageSize != 2 {
		t.Error("BelowMinimumImageSize does not validate")
	}
	if ivAnalysis.ExceedsMaximumImageSize != 2 {
		t.Error("ExceedsMaximumImageSize does not validate")
	}

	_ = additionalIVAnalysisFields(ivAnalysis, t)
}

func additionalIVAnalysisFields(ivAnalysis ImageViewAnalysis, t *testing.T) string {
	if ivAnalysis.ImageEnabledPOD != 1 {
		t.Error("ImageEnabledPOD does not validate")
	}
	if ivAnalysis.SourceDocumentBad != 0 {
		t.Error("SourceDocumentBad does not validate")
	}
	if ivAnalysis.DateUsability != 2 {
		t.Error("DateUsability does not validate")
	}
	if ivAnalysis.PayeeUsability != 2 {
		t.Error("PayeeUsability does not validate")
	}
	if ivAnalysis.ConvenienceAmountUsability != 2 {
		t.Error("ConvenienceAmountUsability does not validate")
	}
	if ivAnalysis.AmountInWordsUsability != 2 {
		t.Error("AmountInWordsUsability does not validate")
	}
	if ivAnalysis.SignatureUsability != 2 {
		t.Error("SignatureUsability does not validate")
	}
	if ivAnalysis.PayorNameAddressUsability != 2 {
		t.Error("PayorNameAddressUsability does not validate")
	}
	if ivAnalysis.MICRLineUsability != 2 {
		t.Error("MICRLineUsability does not validate")
	}
	if ivAnalysis.MemoLineUsability != 2 {
		t.Error("MemoLineUsability does not validate")
	}
	if ivAnalysis.PayorBankNameAddressUsability != 2 {
		t.Error("PayorBankNameAddressUsability does not validate")
	}
	if ivAnalysis.PayeeEndorsementUsability != 2 {
		t.Error("PayeeEndorsementUsability does not validate")
	}
	if ivAnalysis.BOFDEndorsementUsability != 2 {
		t.Error("BOFDEndorsementUsability does not validate")
	}
	if ivAnalysis.TransitEndorsementUsability != 2 {
		t.Error("TransitEndorsementUsability does not validate")
	}
	return ""
}

// testIVAnalysisString validates that a known parsed ImageViewAnalysis can return to a string of the same value
func testIVAnalysisString(t testing.TB) {
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

// TestIVAnalysisString tests validating that a known parsed ImageViewAnalysis an return to a string of the
// same value
func TestIVAnalysisString(t *testing.T) {
	testIVAnalysisString(t)
}

// BenchmarkIVAnalysisString benchmarks validating that a known parsed ImageViewAnalysis
// can return to a string of the same value
func BenchmarkIVAnalysisString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIVAnalysisString(b)
	}
}

// TestIVAnalysisRecordType validation
func TestIVAnalysisRecordType(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.recordType = "00"
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisGlobalImageQuality validation
func TestIVAnalysisGlobalImageQuality(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.GlobalImageQuality = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "GlobalImageQuality" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisGlobalImageUsability validation
func TestIVAnalysisGlobalImageUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.GlobalImageUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "GlobalImageUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisImagingBankSpecificTest validation
func TestIVAnalysisImagingBankSpecificTest(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ImagingBankSpecificTest = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImagingBankSpecificTest" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisPartialImage validation
func TestIVAnalysisPartialImage(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PartialImage = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PartialImage" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisExcessiveImageSkew validation
func TestIVAnalysisExcessiveImageSkew(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ExcessiveImageSkew = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ExcessiveImageSkew" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisPiggybackImage validation
func TestIVAnalysisPiggybackImage(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PiggybackImage = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PiggybackImage" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisTooLightOrTooDark validation
func TestIVAnalysisTooLightOrTooDarke(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.TooLightOrTooDark = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TooLightOrTooDark" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisStreaksAndOrBands validation
func TestIVAnalysisStreaksAndOrBands(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.StreaksAndOrBands = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "StreaksAndOrBands" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisBelowMinimumImageSize validation
func TestIVAnalysisBelowMinimumImageSize(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.BelowMinimumImageSize = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BelowMinimumImageSize" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisExceedsMaximumImageSize validation
func TestIVAnalysisExceedsMaximumImageSize(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ExceedsMaximumImageSize = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ExceedsMaximumImageSize" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisImageEnabledPOD validation
func TestIVAnalysisImageEnabledPOD(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ImageEnabledPOD = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImageEnabledPOD" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisSourceDocumentBad validation
func TestIVAnalysisSourceDocumentBad(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.SourceDocumentBad = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "SourceDocumentBad" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisDateUsability validation
func TestIVAnalysisDateUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.DateUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DateUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisPayeeUsability validation
func TestIVAnalysisPayeeUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PayeeUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayeeUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisConvenienceAmountUsability validation
func TestIVAnalysisConvenienceAmountUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ConvenienceAmountUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ConvenienceAmountUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisAmountInWordsUsability validation
func TestIVAnalysisAmountInWordsUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.AmountInWordsUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "AmountInWordsUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisSignatureUsability validation
func TestIVAnalysisSignatureUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.SignatureUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "SignatureUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisPayorNameAddressUsability validation
func TestIVAnalysisPayorNameAddressUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PayorNameAddressUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorNameAddressUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisMICRLineUsability validation
func TestIVAnalysisMICRLineUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.MICRLineUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "MICRLineUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisMemoLineUsability validation
func TestIVAnalysisMemoLineUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.MemoLineUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "MemoLineUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisPayorBankNameAddressUsability validation
func TestIVAnalysisPayorBankNameAddressUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PayorBankNameAddressUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankNameAddressUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisPayeeEndorsementUsability validation
func TestIVAnalysisPayeeEndorsementUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PayeeEndorsementUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayeeEndorsementUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisBOFDEndorsementUsability validation
func TestIVAnalysisBOFDEndorsementUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.BOFDEndorsementUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDEndorsementUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisTransitEndorsementUsability validation
func TestIVAnalysisTransitEndorsementUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.TransitEndorsementUsability = 5
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TransitEndorsementUsability" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisUserField validation
func TestIVAnalysisUserField(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.UserField = "®©"
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestIVAnalysisFIRecordType validation
func TestIVAnalysisFIRecordType(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.recordType = ""
	if err := ivAnalysis.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestIVAnalysisRuneCountInString validates RuneCountInString
func TestIVAnalysisRuneCountInString(t *testing.T) {
	ivAnalysis := NewImageViewAnalysis()
	var line = "54"
	ivAnalysis.Parse(line)

	if ivAnalysis.AmountInWordsUsability != 0 {
		t.Error("Parsed with an invalid RuneCountInString")
	}
}
