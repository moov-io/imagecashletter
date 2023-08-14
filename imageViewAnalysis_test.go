// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
	require.NoError(t, ivAnalysis.Validate())
	require.Equal(t, 2, ivAnalysis.GlobalImageQuality)
	require.Equal(t, 2, ivAnalysis.GlobalImageUsability)
	require.Equal(t, 0, ivAnalysis.ImagingBankSpecificTest)
	require.Equal(t, 2, ivAnalysis.PartialImage)
	require.Equal(t, 2, ivAnalysis.ExcessiveImageSkew)
	require.Equal(t, 2, ivAnalysis.PiggybackImage)
	require.Equal(t, 2, ivAnalysis.TooLightOrTooDark)
	require.Equal(t, 2, ivAnalysis.StreaksAndOrBands)
	require.Equal(t, 2, ivAnalysis.BelowMinimumImageSize)
	require.Equal(t, 2, ivAnalysis.ExceedsMaximumImageSize)

	_ = additionalIVAnalysisFields(ivAnalysis, t)
}

func additionalIVAnalysisFields(ivAnalysis ImageViewAnalysis, t *testing.T) string {
	require.Equal(t, 1, ivAnalysis.ImageEnabledPOD)
	require.Equal(t, 0, ivAnalysis.SourceDocumentBad)
	require.Equal(t, 2, ivAnalysis.DateUsability)
	require.Equal(t, 2, ivAnalysis.PayeeUsability)
	require.Equal(t, 2, ivAnalysis.ConvenienceAmountUsability)
	require.Equal(t, 2, ivAnalysis.AmountInWordsUsability)
	require.Equal(t, 2, ivAnalysis.SignatureUsability)
	require.Equal(t, 2, ivAnalysis.PayorNameAddressUsability)
	require.Equal(t, 2, ivAnalysis.MICRLineUsability)
	require.Equal(t, 2, ivAnalysis.MemoLineUsability)
	require.Equal(t, 2, ivAnalysis.PayorBankNameAddressUsability)
	require.Equal(t, 2, ivAnalysis.PayeeEndorsementUsability)
	require.Equal(t, 2, ivAnalysis.BOFDEndorsementUsability)
	require.Equal(t, 2, ivAnalysis.TransitEndorsementUsability)
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

	require.NoError(t, r.parseImageViewAnalysis())
	record := r.currentCashLetter.currentBundle.GetChecks()[0].ImageViewAnalysis[0]

	require.Equal(t, line, record.String())
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
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestIVAnalysisGlobalImageQuality validation
func TestIVAnalysisGlobalImageQuality(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.GlobalImageQuality = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "GlobalImageQuality", e.FieldName)
}

// TestIVAnalysisGlobalImageUsability validation
func TestIVAnalysisGlobalImageUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.GlobalImageUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "GlobalImageUsability", e.FieldName)
}

// TestIVAnalysisImagingBankSpecificTest validation
func TestIVAnalysisImagingBankSpecificTest(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ImagingBankSpecificTest = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImagingBankSpecificTest", e.FieldName)
}

// TestIVAnalysisPartialImage validation
func TestIVAnalysisPartialImage(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PartialImage = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PartialImage", e.FieldName)
}

// TestIVAnalysisExcessiveImageSkew validation
func TestIVAnalysisExcessiveImageSkew(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ExcessiveImageSkew = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ExcessiveImageSkew", e.FieldName)
}

// TestIVAnalysisPiggybackImage validation
func TestIVAnalysisPiggybackImage(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PiggybackImage = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PiggybackImage", e.FieldName)
}

// TestIVAnalysisTooLightOrTooDark validation
func TestIVAnalysisTooLightOrTooDarke(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.TooLightOrTooDark = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TooLightOrTooDark", e.FieldName)
}

// TestIVAnalysisStreaksAndOrBands validation
func TestIVAnalysisStreaksAndOrBands(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.StreaksAndOrBands = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "StreaksAndOrBands", e.FieldName)
}

// TestIVAnalysisBelowMinimumImageSize validation
func TestIVAnalysisBelowMinimumImageSize(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.BelowMinimumImageSize = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BelowMinimumImageSize", e.FieldName)
}

// TestIVAnalysisExceedsMaximumImageSize validation
func TestIVAnalysisExceedsMaximumImageSize(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ExceedsMaximumImageSize = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ExceedsMaximumImageSize", e.FieldName)
}

// TestIVAnalysisImageEnabledPOD validation
func TestIVAnalysisImageEnabledPOD(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ImageEnabledPOD = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ImageEnabledPOD", e.FieldName)
}

// TestIVAnalysisSourceDocumentBad validation
func TestIVAnalysisSourceDocumentBad(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.SourceDocumentBad = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "SourceDocumentBad", e.FieldName)
}

// TestIVAnalysisDateUsability validation
func TestIVAnalysisDateUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.DateUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "DateUsability", e.FieldName)
}

// TestIVAnalysisPayeeUsability validation
func TestIVAnalysisPayeeUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PayeeUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayeeUsability", e.FieldName)
}

// TestIVAnalysisConvenienceAmountUsability validation
func TestIVAnalysisConvenienceAmountUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.ConvenienceAmountUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "ConvenienceAmountUsability", e.FieldName)
}

// TestIVAnalysisAmountInWordsUsability validation
func TestIVAnalysisAmountInWordsUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.AmountInWordsUsability = 57
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "AmountInWordsUsability", e.FieldName)
}

// TestIVAnalysisSignatureUsability validation
func TestIVAnalysisSignatureUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.SignatureUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "SignatureUsability", e.FieldName)
}

// TestIVAnalysisPayorNameAddressUsability validation
func TestIVAnalysisPayorNameAddressUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PayorNameAddressUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorNameAddressUsability", e.FieldName)
}

// TestIVAnalysisMICRLineUsability validation
func TestIVAnalysisMICRLineUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.MICRLineUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "MICRLineUsability", e.FieldName)
}

// TestIVAnalysisMemoLineUsability validation
func TestIVAnalysisMemoLineUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.MemoLineUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "MemoLineUsability", e.FieldName)
}

// TestIVAnalysisPayorBankNameAddressUsability validation
func TestIVAnalysisPayorBankNameAddressUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PayorBankNameAddressUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayorBankNameAddressUsability", e.FieldName)
}

// TestIVAnalysisPayeeEndorsementUsability validation
func TestIVAnalysisPayeeEndorsementUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.PayeeEndorsementUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "PayeeEndorsementUsability", e.FieldName)
}

// TestIVAnalysisBOFDEndorsementUsability validation
func TestIVAnalysisBOFDEndorsementUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.BOFDEndorsementUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "BOFDEndorsementUsability", e.FieldName)
}

// TestIVAnalysisTransitEndorsementUsability validation
func TestIVAnalysisTransitEndorsementUsability(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.TransitEndorsementUsability = 5
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "TransitEndorsementUsability", e.FieldName)
}

// TestIVAnalysisUserField validation
func TestIVAnalysisUserField(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.UserField = "®©"
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "UserField", e.FieldName)
}

// Field Inclusion

// TestIVAnalysisFIRecordType validation
func TestIVAnalysisFIRecordType(t *testing.T) {
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.recordType = ""
	err := ivAnalysis.Validate()
	var e *FieldError
	require.ErrorAs(t, err, &e)
	require.Equal(t, "recordType", e.FieldName)
}

// TestIVAnalysisRuneCountInString validates RuneCountInString
func TestIVAnalysisRuneCountInString(t *testing.T) {
	ivAnalysis := NewImageViewAnalysis()
	var line = "54"
	ivAnalysis.Parse(line)

	require.Equal(t, 0, ivAnalysis.AmountInWordsUsability)
}
