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

	"github.com/stretchr/testify/require"
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
	require.NoError(t, cl.Create())
	file.AddCashLetter(cl)

	clTwo := NewCashLetter(mockCashLetterHeader())
	clTwo.CashLetterHeader.CashLetterID = "A2"
	clTwo.AddBundle(bundle)
	clTwo.AddBundle(returnBundle)
	require.NoError(t, clTwo.Create())
	file.AddCashLetter(clTwo)

	// Create file
	require.NoError(t, file.Create())
	require.NoError(t, file.Validate())

	b := &bytes.Buffer{}
	f := NewWriter(b)
	require.NoError(t, f.Write(file))

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	require.NoError(t, err)
	require.NoError(t, r.File.Validate())
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
	require.NoError(t, cl.Create())
	file.AddCashLetter(cl)

	clTwo := NewCashLetter(mockCashLetterHeader())
	clTwo.CashLetterHeader.CashLetterID = "A2"
	clTwo.AddBundle(bundle)

	require.NoError(t, clTwo.Create())
	file.AddCashLetter(clTwo)

	// Create file
	require.NoError(t, file.Create())
	require.NoError(t, file.Validate())

	b := &bytes.Buffer{}
	f := NewWriter(b)
	require.NoError(t, f.Write(file))

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	require.NoError(t, err)
	require.NoError(t, r.File.Validate())
}

// TestICLWriteCreditRecord writes an ICL file with a Credit record
func TestICLWriteCreditRecord(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())

	// CreditItem
	ci := mockCredit()

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
	cl.AddCredit(ci)
	cl.AddBundle(bundle)
	require.NoError(t, cl.Create())
	file.AddCashLetter(cl)

	clTwo := NewCashLetter(mockCashLetterHeader())
	clTwo.CashLetterHeader.CashLetterID = "A2"
	clTwo.AddBundle(bundle)

	require.NoError(t, clTwo.Create())
	file.AddCashLetter(clTwo)

	// Create file
	require.NoError(t, file.Create())
	require.NoError(t, file.Validate())

	b := &bytes.Buffer{}
	f := NewWriter(b)
	require.NoError(t, f.Write(file))

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	require.NoError(t, err)
	require.NoError(t, r.File.Validate())
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
	require.NoError(t, cl.Create())
	file.AddCashLetter(cl)

	// Create file
	require.NoError(t, file.Create())
	require.NoError(t, file.Validate())

	b := &bytes.Buffer{}
	f := NewWriter(b)
	require.NoError(t, f.Write(file))

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	require.NoError(t, err)
	require.NoError(t, r.File.Validate())
}

func TestICLWrite_VariableLengthOption(t *testing.T) {
	fileBytes, err := os.ReadFile(filepath.Join("test", "testdata", "valid-ascii.x937"))
	require.NoError(t, err)

	fd := bytes.NewReader(fileBytes)
	r := NewReader(fd, ReadVariableLineLengthOption())
	file, err := r.Read()
	require.NoError(t, err)

	b := &bytes.Buffer{}
	w := NewWriter(b, WriteVariableLineLengthOption())
	require.NoError(t, w.Write(&file))
	require.Equal(t, fileBytes, b.Bytes())
}

func TestICLWrite_EbcdicEncodingOption(t *testing.T) {
	fileBytes, err := os.ReadFile(filepath.Join("test", "testdata", "valid-ebcdic.x937"))
	require.NoError(t, err)

	fd := bytes.NewReader(fileBytes)
	r := NewReader(fd, ReadVariableLineLengthOption(), ReadEbcdicEncodingOption())
	file, err := r.Read()
	require.NoError(t, err)

	b := &bytes.Buffer{}
	w := NewWriter(b, WriteVariableLineLengthOption(), WriteEbcdicEncodingOption())
	require.NoError(t, w.Write(&file))
	require.Equal(t, fileBytes, b.Bytes())
}

func TestWriter_CollateErr(t *testing.T) {
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
	require.ErrorContains(t, w.writeCheckImageView(cd), "ImageViewData does not match Image View Detail count of 1")
}
