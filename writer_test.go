// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestICLWrite writes an ICL File
func TestICLWrite(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())

	// Create CheckDetail
	cd := mockCheckDetail()
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// CheckDetail 2
	cdTwo := mockCheckDetail()
	cdTwo.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cdTwo.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cdTwo.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cdTwo.AddImageViewDetail(mockImageViewDetail())
	cdTwo.AddImageViewData(mockImageViewData())
	cdTwo.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle.AddCheckDetail(cdTwo)

	// Create ReturnDetail
	rd := mockReturnDetail()
	rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	rd.AddImageViewDetail(mockImageViewDetail())
	rd.AddImageViewData(mockImageViewData())
	rd.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle := NewBundle(mockBundleHeader())
	returnBundle.BundleHeader.BundleSequenceNumber = "2"
	returnBundle.AddReturnDetail(rd)

	rdTwo := mockReturnDetail()
	rdTwo.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
	rdTwo.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
	rdTwo.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
	rdTwo.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
	rdTwo.AddImageViewDetail(mockImageViewDetail())
	rdTwo.AddImageViewData(mockImageViewData())
	rdTwo.AddImageViewAnalysis(mockImageViewAnalysis())
	returnBundle.AddReturnDetail(rdTwo)

	// Create CashLetter
	cl := NewCashLetter(mockCashLetterHeader())
	cl.AddBundle(bundle)
	cl.AddBundle(returnBundle)
	if err := cl.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	file.AddCashLetter(cl)

	clTwo := NewCashLetter(mockCashLetterHeader())
	clTwo.CashLetterHeader.CashLetterID = "A2"
	clTwo.AddBundle(bundle)
	clTwo.AddBundle(returnBundle)
	if err := clTwo.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	file.AddCashLetter(clTwo)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	/*
		// We want to write the file to an io.Writer
		w := NewWriter(os.Stdout)
		if err := w.Write(file); err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}
		w.Flush()*/

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestICLWriteCreditItem writes an ICL file with a CreditItem
func TestICLWriteCreditItem(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())

	// CreditItem
	ci := mockCreditItem()

	// Create CheckDetail
	cd := mockCheckDetail()
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// CheckDetail 2
	cdTwo := mockCheckDetail()
	cdTwo.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cdTwo.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cdTwo.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cdTwo.AddImageViewDetail(mockImageViewDetail())
	cdTwo.AddImageViewData(mockImageViewData())
	cdTwo.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle.AddCheckDetail(cdTwo)

	// Create CashLetter
	cl := NewCashLetter(mockCashLetterHeader())
	cl.AddCreditItem(ci)
	cl.AddBundle(bundle)
	if err := cl.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	file.AddCashLetter(cl)

	clTwo := NewCashLetter(mockCashLetterHeader())
	clTwo.CashLetterHeader.CashLetterID = "A2"
	clTwo.AddBundle(bundle)

	if err := clTwo.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	file.AddCashLetter(clTwo)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	/*	// We want to write the file to an io.Writer
		w := NewWriter(os.Stdout)
		if err := w.Write(file); err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}
		w.Flush()*/

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestICLWriteRoutingNumberSummary writes an ICL file with a RoutingNumberSummary
func TestICLWriteRoutingNumber(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())

	// RoutingNumberSummary
	rns := mockRoutingNumberSummary()

	// Create CheckDetail
	cd := mockCheckDetail()
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	cd.AddImageViewDetail(mockImageViewDetail())
	cd.AddImageViewData(mockImageViewData())
	cd.AddImageViewAnalysis(mockImageViewAnalysis())
	bundle := NewBundle(mockBundleHeader())
	bundle.AddCheckDetail(cd)

	// Create CashLetter
	cl := NewCashLetter(mockCashLetterHeader())
	cl.AddBundle(bundle)
	cl.AddRoutingNumberSummary(rns)
	if err := cl.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	file.AddCashLetter(cl)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// We want to write the file to an io.Writer
	w := NewWriter(os.Stdout)
	/*		if err := w.Write(file); err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}*/
	w.Flush()

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

}

func TestICLWrite_DSTU(t *testing.T) {
	fd, err := os.Open(filepath.Join("test", "testdata", "valid-dstu.x937"))
	if err != nil {
		t.Fatalf("Can not open local file: %s: \n", err)
	}
	defer fd.Close()

	fdInfo, err := fd.Stat()
	if err != nil {
		t.Errorf("Could not stat file: %s: \n", err)
	}

	r := NewReader(fd)
	r.SetFormat(DSTU)
	ICLFile, err := r.Read()
	if err != nil {
		t.Errorf("Issue reading file: %+v \n", err)
	}
	// ensure we have a validated file structure
	if ICLFile.Validate(); err != nil {
		t.Errorf("Could not validate entire read file: %v", err)
	}

	b := &bytes.Buffer{}
	w := NewWriter(b)
	w.SetFormat(DSTU)
	if err := w.Write(&ICLFile); err != nil {
		t.Errorf("Issue writing ICL: %+v \n", err)
	}

	if b.Len() != int(fdInfo.Size()) {
		t.Errorf("File size does not match")
	}
}
