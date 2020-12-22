// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"bytes"
	"io/ioutil"
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

func TestICLWrite_VariableLengthOption(t *testing.T) {
	fileBytes, err := ioutil.ReadFile(filepath.Join("test", "testdata", "valid-ascii.x937"))
	if err != nil {
		t.Fatalf("Can not open local file: %s: \n", err)
	}

	fd := bytes.NewReader(fileBytes)
	r := NewReader(fd, ReadVariableLineLengthOption())
	file, err := r.Read()
	if err != nil {
		t.Errorf("Issue reading file: %+v \n", err)
	}

	b := &bytes.Buffer{}
	w := NewWriter(b, WriteVariableLineLengthOption())

	if err := w.Write(&file); err != nil {
		t.Errorf("Issue writing ICL: %+v \n", err)
	}

	if !bytes.Equal(fileBytes, b.Bytes()) {
		t.Errorf("ICLs does not match")
	}
}

func TestICLWrite_EbcdicEncodingOption(t *testing.T) {
	fileBytes, err := ioutil.ReadFile(filepath.Join("test", "testdata", "valid-ebcdic.x937"))
	if err != nil {
		t.Fatalf("Can not open local file: %s: \n", err)
	}

	fd := bytes.NewReader(fileBytes)
	r := NewReader(fd, ReadVariableLineLengthOption(), ReadEbcdicEncodingOption())
	file, err := r.Read()
	if err != nil {
		t.Errorf("Issue reading file: %+v \n", err)
	}

	b := &bytes.Buffer{}
	w := NewWriter(b, WriteVariableLineLengthOption(), WriteEbcdicEncodingOption())

	if err := w.Write(&file); err != nil {
		t.Errorf("Issue writing ICL: %+v \n", err)
	}

	if !bytes.Equal(fileBytes, b.Bytes()) {
		t.Errorf("ICLs does not match")
	}
}

func TestWriter__CollateErr(t *testing.T) {
	cd := &CheckDetail{
		// Create a CheckDetail without a corresponding ImageData or ImageViewAnalysis
		// so when we attempt to collate them it doesn't crash.
		ImageViewDetail: []ImageViewDetail{
			mockImageViewDetail(),
			mockImageViewDetail(),
		},
		// To trigger the crash this issue fixes we need two ImageViewDetails, and one ImageData.
		// Having one ImageViewAnalysis would work as well
		ImageViewData: []ImageViewData{
			mockImageViewData(),
		},
		ImageViewAnalysis: []ImageViewAnalysis{
			mockImageViewAnalysis(),
		},
	}

	var buf bytes.Buffer
	w := NewWriter(&buf)

	err := w.writeCheckImageView(cd)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "ImageViewData does not match Image View Detail count of 1") {
		t.Errorf("unexpected error: %q", err)
	}
}
