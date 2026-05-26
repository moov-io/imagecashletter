// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBundleChecks
func mockBundleChecks(t *testing.T) *Bundle {
	t.Helper()

	bundle := &Bundle{}
	bundle.SetHeader(mockBundleHeader())
	bundle.AddCheckDetail(mockCheckDetail())
	bundle.Checks[0].AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	bundle.Checks[0].AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	bundle.Checks[0].AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	bundle.Checks[0].AddImageViewDetail(mockImageViewDetail())
	bundle.Checks[0].AddImageViewData(mockImageViewData())
	bundle.Checks[0].AddImageViewAnalysis(mockImageViewAnalysis())

	require.NoError(t, bundle.build())

	return bundle
}

// mockBundleReturns
func mockBundleReturns(t *testing.T) *Bundle {
	t.Helper()

	bundle := &Bundle{}
	bundle.SetHeader(mockBundleHeader())
	bundle.AddReturnDetail(mockReturnDetail())
	bundle.Returns[0].AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	bundle.Returns[0].AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	bundle.Returns[0].AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	bundle.Returns[0].AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	bundle.Returns[0].AddImageViewDetail(mockImageViewDetail())
	bundle.Returns[0].AddImageViewData(mockImageViewData())
	bundle.Returns[0].AddImageViewAnalysis(mockImageViewAnalysis())

	require.NoError(t, bundle.build())

	return bundle
}

// TestMockBundleChecks creates a Bundle of checks
func TestMockBundleChecks(t *testing.T) {
	bundle := mockBundleChecks(t)
	require.NoError(t, bundle.Validate())

	bundle = nil // ensure we don't panic
	require.NotPanics(t, func() {
		checks := bundle.GetChecks()
		require.Nil(t, checks)
	})
}

// TestMockBundleReturns creates a Bundle of returns
func TestMockBundleReturns(t *testing.T) {
	bundle := mockBundleReturns(t)
	require.NoError(t, bundle.Validate())

	bundle = nil // ensure we don't panic
	require.NotPanics(t, func() {
		returns := bundle.GetReturns()
		require.Nil(t, returns)
	})
}

func TestBundleValidate(t *testing.T) {
	header := mockBundleHeader()
	bundle := NewBundle(header)
	require.Error(t, bundle.Validate())
}

// TestCheckDetailAddendumCount validates CheckDetail AddendumCount
func TestCheckDetailAddendumCount(t *testing.T) {
	cd := mockCheckDetail()
	cd.AddendumCount = 2 // incorrect count should cause error
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	err := bundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "AddendumCount", bundleErr.FieldName)
}

// TestCheckDetailAddendumACount validates CheckDetailAddendumA AddendaCount
func TestCheckDetailAddendumACount(t *testing.T) {
	cd := mockCheckDetail()
	cd.AddendumCount = 12 // incorrect count should cause error
	for i := 0; i < 10; i++ {
		cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	}
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	err := bundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "CheckDetailAddendumA", bundleErr.FieldName)
}

// TestCheckDetailAddendumBCount validates CheckDetailAddendumB AddendaCount
func TestCheckDetailAddendumBCount(t *testing.T) {
	cd := mockCheckDetail()
	cd.AddendumCount = 4
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	err := bundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "CheckDetailAddendumB", bundleErr.FieldName)
}

// TestCheckDetailAddendumCCount validates CheckDetailAddendumC AddendaCount
func TestCheckDetailAddendumCCount(t *testing.T) {
	cd := mockCheckDetail()
	cd.AddendumCount = 102
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	for i := 0; i < 100; i++ {
		cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	}
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	err := bundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "CheckDetailAddendumC", bundleErr.FieldName)
}

// TestReturnDetailAddendumCount validates ReturnDetail AddendumCount
func TestReturnDetailAddendumCount(t *testing.T) {
	// Create ReturnDetail
	rd := mockReturnDetail()
	rd.AddendumCount = 3
	rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	rd.AddImageViewDetail(mockImageViewDetail())
	rd.AddImageViewData(mockImageViewData())
	rd.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle := NewBundle(mockBundleHeader())
	returnBundle.AddReturnDetail(rd)

	err := returnBundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "AddendumCount", bundleErr.FieldName)
}

// TestReturnDetailAddendumACount validates ReturnDetailAddendumA Count
func TestReturnDetailAddendumACount(t *testing.T) {
	// Create ReturnDetail
	rd := mockReturnDetail()
	rd.AddendumCount = 13
	for i := 0; i < 10; i++ {
		rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	}
	rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	rd.AddImageViewDetail(mockImageViewDetail())
	rd.AddImageViewData(mockImageViewData())
	rd.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle := NewBundle(mockBundleHeader())
	returnBundle.AddReturnDetail(rd)

	err := returnBundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "ReturnDetailAddendumA", bundleErr.FieldName)
}

// TestReturnDetailAddendumBCount validates ReturnDetailAddendumB Count
func TestReturnDetailAddendumBCount(t *testing.T) {
	// Create ReturnDetail
	rd := mockReturnDetail()
	rd.AddendumCount = 5
	rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	rd.AddImageViewDetail(mockImageViewDetail())
	rd.AddImageViewData(mockImageViewData())
	rd.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle := NewBundle(mockBundleHeader())
	returnBundle.AddReturnDetail(rd)

	err := returnBundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "ReturnDetailAddendumB", bundleErr.FieldName)
}

// TestReturnDetailAddendumCCount validates ReturnDetailAddendumC Count
func TestReturnDetailAddendumCCount(t *testing.T) {
	// Create ReturnDetail
	rd := mockReturnDetail()
	rd.AddendumCount = 5
	rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	rd.AddImageViewDetail(mockImageViewDetail())
	rd.AddImageViewData(mockImageViewData())
	rd.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle := NewBundle(mockBundleHeader())
	returnBundle.AddReturnDetail(rd)

	err := returnBundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "ReturnDetailAddendumC", bundleErr.FieldName)
}

// TestReturnDetailAddendumDCount validates ReturnDetailAddendumD Count
func TestReturnDetailAddendumDCount(t *testing.T) {
	// Create ReturnDetail
	rd := mockReturnDetail()
	rd.AddendumCount = 103
	rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	for i := 0; i < 100; i++ {
		rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	}
	rd.AddImageViewDetail(mockImageViewDetail())
	rd.AddImageViewData(mockImageViewData())
	rd.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle := NewBundle(mockBundleHeader())
	returnBundle.AddReturnDetail(rd)

	err := returnBundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "ReturnDetailAddendumD", bundleErr.FieldName)
}

// TestBundleValidateSkipCountValidation verifies that SkipCountValidation (and SkipAll) relax count checks inside bundles.
func TestBundleValidateSkipCountValidation(t *testing.T) {
	cd := mockCheckDetail()
	cd.AddendumCount = 0 // declared 0 but we'll attach one addendum -> mismatch
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())

	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// Without skip: should fail with the exact error users report
	err := bundle.Validate()
	var bundleErr *BundleError
	require.ErrorAs(t, err, &bundleErr)
	require.Equal(t, "AddendumCount", bundleErr.FieldName)

	// With the count validation skip: succeeds
	bundle.SetValidation(&ValidateOpts{SkipCountValidation: true})
	err = bundle.Validate()
	require.NoError(t, err)

	// Also still works with the broad SkipAll
	bundle.SetValidation(&ValidateOpts{SkipAll: true})
	err = bundle.Validate()
	require.NoError(t, err)
}

// TestBundleBuild_NilHeader verifies that a nil BundleHeader is always an error
// from build(), even under SkipAll (structural integrity, prevents malformed bundles).
func TestBundleBuild_NilHeader(t *testing.T) {
	bundle := &Bundle{} // no header
	err := bundle.build()
	require.Error(t, err)
	require.Equal(t, "nil BundleHeader", err.Error())
}

func TestBundleBuild_NilHeaderSkipAll(t *testing.T) {
	bundle := &Bundle{}
	bundle.SetValidation(&ValidateOpts{SkipAll: true})
	// Structural nil check happens regardless of SkipAll
	err := bundle.build()
	require.Error(t, err)
	require.Equal(t, "nil BundleHeader", err.Error())
}
