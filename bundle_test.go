// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "testing"

// mockBundleChecks
func mockBundleChecks() *Bundle {
	mockBundleChecks := &Bundle{}
	mockBundleChecks.SetHeader(mockBundleHeader())
	mockBundleChecks.AddCheckDetail(mockCheckDetail())
	mockBundleChecks.Checks[0].AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	mockBundleChecks.Checks[0].AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	mockBundleChecks.Checks[0].AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	mockBundleChecks.Checks[0].AddImageViewDetail(mockImageViewDetail())
	mockBundleChecks.Checks[0].AddImageViewData(mockImageViewData())
	mockBundleChecks.Checks[0].AddImageViewAnalysis(mockImageViewAnalysis())
	if err := mockBundleChecks.build(); err != nil {
		panic(err)
	}
	return mockBundleChecks
}

// mockBundleReturns
func mockBundleReturns() *Bundle {
	mockBundleReturns := &Bundle{}
	mockBundleReturns.SetHeader(mockBundleHeader())
	mockBundleReturns.AddReturnDetail(mockReturnDetail())
	mockBundleReturns.Returns[0].AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	mockBundleReturns.Returns[0].AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	mockBundleReturns.Returns[0].AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	mockBundleReturns.Returns[0].AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	mockBundleReturns.Returns[0].AddImageViewDetail(mockImageViewDetail())
	mockBundleReturns.Returns[0].AddImageViewData(mockImageViewData())
	mockBundleReturns.Returns[0].AddImageViewAnalysis(mockImageViewAnalysis())
	if err := mockBundleReturns.build(); err != nil {
		panic(err)
	}
	return mockBundleReturns
}

// TestMockBundleChecks creates a Bundle of checks
func TestMockBundleChecks(t *testing.T) {
	bundle := mockBundleChecks()
	if err := bundle.Validate(); err != nil {
		t.Error("Bundle does not validate and will break other tests: ", err)
	}
}

// TestMockBundleReturns creates a Bundle of returns
func TestMockBundleReturns(t *testing.T) {
	bundle := mockBundleReturns()
	if err := bundle.Validate(); err != nil {
		t.Error("Bundle does not validate and will break other tests: ", err)
	}
}
