// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

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

func TestBundleValidate(t *testing.T) {
	header := mockBundleHeader()
	bundle := NewBundle(header)
	if err := bundle.Validate(); err == nil {
		t.Error("expected error, but got nothing")
	}
}

// TestCheckDetailAddendumCount validates CheckDetail AddendumCount
func TestCheckDetailAddendumCount(t *testing.T) {
	cd := mockCheckDetail()
	cd.AddendumCount = 2
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	if err := bundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "AddendumCount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCheckDetailAddendumACount validates CheckDetailAddendumA AddendaCount
func TestCheckDetailAddendumACount(t *testing.T) {
	cd := mockCheckDetail()
	cd.AddendumCount = 12
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

	if err := bundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "CheckDetailAddendumA" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
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

	if err := bundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "CheckDetailAddendumB" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
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

	if err := bundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "CheckDetailAddendumC" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
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
	if err := returnBundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "AddendumCount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
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
	if err := returnBundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "ReturnDetailAddendumA" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
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
	if err := returnBundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "ReturnDetailAddendumB" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
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
	if err := returnBundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "ReturnDetailAddendumC" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
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
	if err := returnBundle.Validate(); err != nil {
		if e, ok := err.(*BundleError); ok {
			if e.FieldName != "ReturnDetailAddendumD" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
