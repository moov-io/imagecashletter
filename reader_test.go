// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func TestICLFiles(t *testing.T) {
	tests := []struct {
		filename string
	}{
		{"BNK20180905121042882-A.icl"},
		{"without-micrValidIndicator.icl"},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			fd, err := os.Open(filepath.Join("test", "testdata", tt.filename))
			require.NoError(t, err)
			defer fd.Close()

			r := NewReader(fd, ReadVariableLineLengthOption())
			iclFile, err := r.Read()
			require.NoError(t, err)

			if testing.Verbose() {
				t.Logf("r.File.Header=%#v", r.File.Header)
				t.Logf("r.File.Control=%#v", r.File.Control)
			}

			// ensure we have a validated file structure
			require.NoError(t, iclFile.Validate())
		})
	}
}

func TestICL_ReadVariableLineLengthOption(t *testing.T) {
	fd, err := os.Open(filepath.Join("test", "testdata", "valid-ascii.x937"))
	require.NoError(t, err)
	defer fd.Close()

	r := NewReader(fd, ReadVariableLineLengthOption())
	iclFile, err := r.Read()
	require.NoError(t, err)

	if testing.Verbose() {
		t.Logf("r.File.Header=%#v", r.File.Header)
		t.Logf("r.File.Control=%#v", r.File.Control)
	}

	// ensure we have a validated file structure
	require.NoError(t, iclFile.Validate())

	actual, err := json.MarshalIndent(iclFile, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("test", "testdata", "valid-x937.json"))
	require.NoError(t, err)

	require.Equal(t, string(expected), string(actual))
}

func TestICL_EBCDICEncodingOption(t *testing.T) {
	fd, err := os.Open(filepath.Join("test", "testdata", "valid-ebcdic.x937"))
	require.NoError(t, err)
	defer fd.Close()

	r := NewReader(fd, ReadVariableLineLengthOption(), ReadEbcdicEncodingOption())
	iclFile, err := r.Read()
	require.NoError(t, err)

	if testing.Verbose() {
		t.Logf("r.File.Header=%#v", r.File.Header)
		t.Logf("r.File.Control=%#v", r.File.Control)
	}

	// ensure we have a validated file structure
	require.NoError(t, iclFile.Validate())
	actual, err := json.MarshalIndent(iclFile, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("test", "testdata", "valid-x937.json"))
	require.NoError(t, err)

	require.Equal(t, string(expected), string(actual))
}

func getFileError(t *testing.T, err error) *FileError {
	var fileErr *FileError
	require.ErrorAs(t, err, &fileErr)

	return fileErr
}

func getFieldError(t *testing.T, err error) *FieldError {
	var fieldErr *FieldError
	require.ErrorAs(t, err, &fieldErr)

	return fieldErr
}

// TestRecordTypeUnknown validates record type unknown
func TestRecordTypeUnknown(t *testing.T) {
	line := "1735T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()

	fileErr := getFileError(t, err)
	require.Equal(t, "recordType", fileErr.FieldName)
}

// TestFileLineShort validates file line is short
func TestFileLineShort(t *testing.T) {
	line := "1 line is only 70 characters ........................................!"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()

	fileErr := getFileError(t, err)
	require.Equal(t, "RecordLength", fileErr.FieldName)
}

func TestReaderCrash_parseBundleControl(t *testing.T) {
	require.Error(t, new(Reader).parseBundleControl())
}

// TestFileFileHeaderErr validates error flows back from the parser
func TestFileFileHeaderErr(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = ""
	r := NewReader(strings.NewReader(fh.String()))
	// necessary to have a file control not nil
	r.File.Control = mockFileControl()
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestTwoFileHeaders validates one file header
func TestTwoFileHeaders(t *testing.T) {
	line := "0135T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	twoHeaders := line + "\n" + line
	r := NewReader(strings.NewReader(twoHeaders))
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileControl, fileErr.Msg)
}

// TestCashLetterHeaderErr validates error flows back from the parser
func TestCashLetterHeaderErr(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.DestinationRoutingNumber = ""
	r := NewReader(strings.NewReader(clh.String()))
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestCashLetterHeaderDuplicate validates when two CashLetterHeader exists in a current CashLetter
func TestCashLetterHeaderDuplicate(t *testing.T) {
	// create a new CashLetter header string
	clh := mockCashLetterHeader()
	r := NewReader(strings.NewReader(clh.String()))
	// instantiate a CashLetter in the reader
	r.addCurrentCashLetter(NewCashLetter(clh))
	// read should fail because it is parsing a second CashLetter Header and there can only be one.
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileCashLetterInside, fileErr.Msg)
}

// TestBundleHeaderErr validates error flows back from the parser
func TestBundleHeaderErr(t *testing.T) {
	bh := mockBundleHeader()
	bh.DestinationRoutingNumber = ""
	r := NewReader(strings.NewReader(bh.String()))
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestBundleHeaderDuplicate validates when two BundleHeader exists in a current Bundle
func TestBundleHeaderDuplicate(t *testing.T) {
	// create a new CashLetter header string
	bh := mockBundleHeader()
	r := NewReader(strings.NewReader(bh.String()))
	// instantiate a CashLetter in the reader
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bhTwo := mockBundleHeader()
	r.addCurrentBundle(NewBundle(bhTwo))
	// read should fail because it is parsing a second CashLetter Header and there can only be one.
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleInside, fileErr.Msg)
}

// TestCheckDetailError validates error flows back from the parser
func TestCheckDetailError(t *testing.T) {
	cd := mockCheckDetail()
	cd.PayorBankRoutingNumber = ""
	r := NewReader(strings.NewReader(cd.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestCheckDetailAddendumABundleError validates error flows back from the parser
func TestCheckDetailAddendumABundleError(t *testing.T) {
	cd := mockCheckDetail()
	cdaddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdaddendumA)
	r := NewReader(strings.NewReader(cdaddendumA.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()

	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestCheckDetailAddendumBBundleError validates error flows back from the parser
func TestCheckDetailAddendumBBundleError(t *testing.T) {
	cd := mockCheckDetail()
	cdaddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdaddendumA)
	cdaddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdaddendumB)
	r := NewReader(strings.NewReader(cdaddendumB.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestCheckDetailAddendumCBundleError validates error flows back from the parser
func TestCheckDetailAddendumCBundleError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdAddendumB)
	cdAddendumC := mockCheckDetailAddendumC()
	cd.AddCheckDetailAddendumC(cdAddendumC)
	r := NewReader(strings.NewReader(cdAddendumC.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestCheckDetailAddendumAError validates error flows back from the parser
func TestCheckDetailAddendumAError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = ""
	cd.AddCheckDetailAddendumA(cdAddendumA)
	r := NewReader(strings.NewReader(cdAddendumA.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddCheckDetail(cd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestCheckDetailAddendumBError validates error flows back from the parser
func TestCheckDetailAddendumBError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.MicrofilmArchiveSequenceNumber = "               "
	cd.AddCheckDetailAddendumB(cdAddendumB)
	r := NewReader(strings.NewReader(cdAddendumB.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddCheckDetail(cd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestCheckDetailAddendumCError validates error flows back from the parser
func TestCheckDetailAddendumCError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdAddendumB)
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankRoutingNumber = ""
	cd.AddCheckDetailAddendumC(cdAddendumC)
	r := NewReader(strings.NewReader(cdAddendumC.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddCheckDetail(cd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestReturnDetailError validates error flows back from the parser
func TestReturnDetailError(t *testing.T) {
	rd := mockReturnDetail()
	rd.PayorBankRoutingNumber = ""
	r := NewReader(strings.NewReader(rd.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestReturnDetailAddendumABundleError validates error flows back from the parser
func TestReturnDetailAddendumABundleError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	r := NewReader(strings.NewReader(rdAddendumA.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestReturnDetailAddendumBBundleError validates error flows back from the parser
func TestReturnDetailAddendumBBundleError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumB := mockReturnDetailAddendumB()
	rd.AddReturnDetailAddendumB(rdAddendumB)
	r := NewReader(strings.NewReader(rdAddendumB.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestReturnDetailAddendumCBundleError validates error flows back from the parser
func TestReturnDetailAddendumCBundleError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumB := mockReturnDetailAddendumB()
	rd.AddReturnDetailAddendumB(rdAddendumB)
	rdAddendumC := mockReturnDetailAddendumC()
	rd.AddReturnDetailAddendumC(rdAddendumC)
	r := NewReader(strings.NewReader(rdAddendumC.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestReturnDetailAddendumDBundleError validates error flows back from the parser
func TestReturnDetailAddendumDBundleError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumB := mockReturnDetailAddendumB()
	rd.AddReturnDetailAddendumB(rdAddendumB)
	rdAddendumC := mockReturnDetailAddendumC()
	rd.AddReturnDetailAddendumC(rdAddendumC)
	rdAddendumD := mockReturnDetailAddendumD()
	rd.AddReturnDetailAddendumD(rdAddendumD)
	r := NewReader(strings.NewReader(rdAddendumD.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestReturnDetailAddendumAError validates error flows back from the parser
func TestReturnDetailAddendumAError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rdAddendumA.ReturnLocationRoutingNumber = ""
	rd.AddReturnDetailAddendumA(rdAddendumA)
	r := NewReader(strings.NewReader(rdAddendumA.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddReturnDetail(rd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestReturnDetailAddendumBError validates error flows back from the parser
func TestReturnDetailAddendumBError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.PayorBankSequenceNumber = "               "
	rd.AddReturnDetailAddendumB(rdAddendumB)
	r := NewReader(strings.NewReader(rdAddendumB.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddReturnDetail(rd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestReturnDetailAddendumCError validates error flows back from the parser
func TestReturnDetailAddendumCError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumB := mockReturnDetailAddendumB()
	rd.AddReturnDetailAddendumB(rdAddendumB)
	rdAddendumC := mockReturnDetailAddendumC()
	rdAddendumC.MicrofilmArchiveSequenceNumber = "               "
	rd.AddReturnDetailAddendumC(rdAddendumC)
	r := NewReader(strings.NewReader(rdAddendumC.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddReturnDetail(rd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestReturnDetailAddendumDError validates error flows back from the parser
func TestReturnDetailAddendumDError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumB := mockReturnDetailAddendumB()
	rd.AddReturnDetailAddendumB(rdAddendumB)
	rdAddendumC := mockReturnDetailAddendumC()
	rd.AddReturnDetailAddendumC(rdAddendumC)
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankRoutingNumber = "000000000"
	rd.AddReturnDetailAddendumD(rdAddendumD)
	r := NewReader(strings.NewReader(rdAddendumD.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddReturnDetail(rd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestCheckDetailBundleError validates error flows back from the parser
func TestCheckDetailBundleError(t *testing.T) {
	cd := mockCheckDetail()
	r := NewReader(strings.NewReader(cd.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestReturnDetailBundleError validates error flows back from the parser
func TestReturnDetailBundleError(t *testing.T) {
	rd := mockReturnDetail()
	r := NewReader(strings.NewReader(rd.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestCheckDetailIVDetailError validates error flows back from the parser
func TestCheckDetailIVDetailError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdAddendumB)
	ivDetail := mockImageViewDetail()
	ivDetail.ViewDescriptor = ""
	cd.AddImageViewDetail(ivDetail)
	r := NewReader(strings.NewReader(ivDetail.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddCheckDetail(cd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestCheckDetailIVDataError validates error flows back from the parser
func TestCheckDetailIVDataError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdAddendumB)
	ivd := mockImageViewDetail()
	cd.AddImageViewDetail(ivd)
	ivData := mockImageViewData()
	ivData.EceInstitutionRoutingNumber = "000000000"
	cd.AddImageViewData(ivData)
	r := NewReader(strings.NewReader(ivData.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddCheckDetail(cd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestCheckDetailIVAnalysisError validates error flows back from the parser
func TestCheckDetailIVAnalysisError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdAddendumB)
	ivd := mockImageViewDetail()
	cd.AddImageViewDetail(ivd)
	ivData := mockImageViewData()
	cd.AddImageViewData(ivData)
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.GlobalImageQuality = 9
	cd.AddImageViewAnalysis(ivAnalysis)
	r := NewReader(strings.NewReader(ivAnalysis.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddCheckDetail(cd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Equal(t, "GlobalImageQuality", fieldErr.FieldName)
}

// TestReturnDetailIVDetailError validates error flows back from the parser
func TestReturnDetailIVDetailError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumC := mockReturnDetailAddendumC()
	rd.AddReturnDetailAddendumC(rdAddendumC)
	ivDetail := mockImageViewDetail()
	ivDetail.ViewDescriptor = ""
	rd.AddImageViewDetail(ivDetail)
	r := NewReader(strings.NewReader(ivDetail.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddReturnDetail(rd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestReturnDetailIVDataError validates error flows back from the parser
func TestReturnDetailIVDataError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumC := mockReturnDetailAddendumC()
	rd.AddReturnDetailAddendumC(rdAddendumC)
	ivDetail := mockImageViewDetail()
	rd.AddImageViewDetail(ivDetail)
	ivData := mockImageViewData()
	ivData.EceInstitutionRoutingNumber = "000000000"
	r := NewReader(strings.NewReader(ivData.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddReturnDetail(rd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Contains(t, fieldErr.Msg, msgFieldInclusion)
}

// TestReturnDetailIVAnalysisError validates error flows back from the parser
func TestReturnDetailIVAnalysisError(t *testing.T) {
	rd := mockReturnDetail()
	rdAddendumA := mockReturnDetailAddendumA()
	rd.AddReturnDetailAddendumA(rdAddendumA)
	rdAddendumC := mockReturnDetailAddendumC()
	rd.AddReturnDetailAddendumC(rdAddendumC)
	ivDetail := mockImageViewDetail()
	rd.AddImageViewDetail(ivDetail)
	ivData := mockImageViewData()
	rd.AddImageViewData(ivData)
	ivAnalysis := mockImageViewAnalysis()
	ivAnalysis.GlobalImageQuality = 9
	rd.AddImageViewAnalysis(ivAnalysis)
	r := NewReader(strings.NewReader(ivAnalysis.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	b.AddReturnDetail(rd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fieldErr := getFieldError(t, err)
	require.Equal(t, "GlobalImageQuality", fieldErr.FieldName)
}

// TestIVDetailBundleError validates error flows back from the parser
func TestIVDetailBundleError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdAddendumB)
	ivDetail := mockImageViewDetail()
	cd.AddImageViewDetail(ivDetail)
	r := NewReader(strings.NewReader(ivDetail.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	// b.AddCheckDetail(cd)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestIVDataBundleError validates error flows back from the parser
func TestIVDataBundleError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdAddendumB)
	ivDetail := mockImageViewDetail()
	cd.AddImageViewDetail(ivDetail)
	ivData := mockImageViewData()
	cd.AddImageViewData(ivData)
	r := NewReader(strings.NewReader(ivData.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)

	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestIVAnalysisBundleError validates error flows back from the parser
func TestIVAnalysisBundleError(t *testing.T) {
	cd := mockCheckDetail()
	cdAddendumA := mockCheckDetailAddendumA()
	cd.AddCheckDetailAddendumA(cdAddendumA)
	cdAddendumB := mockCheckDetailAddendumB()
	cd.AddCheckDetailAddendumB(cdAddendumB)
	ivDetail := mockImageViewDetail()
	cd.AddImageViewDetail(ivDetail)
	ivData := mockImageViewData()
	cd.AddImageViewData(ivData)
	ivAnalysis := mockImageViewAnalysis()
	cd.AddImageViewAnalysis(ivAnalysis)
	r := NewReader(strings.NewReader(ivAnalysis.String()))
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)

	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	_, err := r.Read()
	fileErr := getFileError(t, err)
	require.Equal(t, msgFileBundleOutside, fileErr.Msg)
}

// TestICLCreditItemFile validates reading an ICL file with a CreditItem
func TestICLCreditItemFile(t *testing.T) {
	fd, err := os.Open(filepath.Join("test", "testdata", "BNK20181010121042882-A.icl"))
	require.NoError(t, err)
	defer fd.Close()

	iclFile, err := NewReader(fd, ReadVariableLineLengthOption()).Read()
	require.NoError(t, err)
	// ensure we have a validated file structure
	require.NoError(t, iclFile.Validate())
}

// TestICLCreditRecord61File validates reading an ICL file with a Credit record (type 61)
func TestICLCreditRecord61File(t *testing.T) {
	fd, err := os.Open(filepath.Join("test", "testdata", "creditRecord61.icl"))
	require.NoError(t, err)
	defer fd.Close()

	iclFile, err := NewReader(fd, ReadVariableLineLengthOption()).Read()
	require.NoError(t, err)

	// ensure we have a validated file structure
	require.NoError(t, iclFile.Validate())
	require.Len(t, iclFile.CashLetters, 2)
	require.Len(t, iclFile.CashLetters[0].Credits, 1)
}

func TestICLBase64ImageData(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "base64-encoded-images.json"))
	require.NoError(t, err)

	file, err := FileFromJSON(bs)
	require.NoError(t, err)

	var buf bytes.Buffer
	require.NoError(t, NewWriter(&buf).Write(file))
	require.Contains(t, buf.String(), "hello, world")
}

// TestICLFile_LargeCheckImage validates that reading a file with a large
// check detail record fails by default with bufio.ErrTooLong, and succeeds
// if a sufficiently-large buffer is created via BufferSizeOption.
//
// It creates this file on the fly to avoid bloating the repository.
func TestICLFile_LargeCheckImage(t *testing.T) {
	fd, err := os.Open(filepath.Join("test", "testdata", "BNK20180905121042882-A.icl"))
	require.NoError(t, err)
	defer fd.Close()

	r := NewReader(fd, ReadVariableLineLengthOption())
	iclFile, err := r.Read()
	require.NoError(t, err)

	if testing.Verbose() {
		t.Logf("r.File.Header=%#v", r.File.Header)
		t.Logf("r.File.Control=%#v", r.File.Control)
	}

	require.NoError(t, iclFile.Validate())

	data := make([]byte, 128*1024)
	_, err = rand.Read(data)
	require.NoError(t, err)

	iclFile.CashLetters[0].Bundles[0].Checks[0].ImageViewData[0].LengthImageData = strconv.Itoa(len(data))
	iclFile.CashLetters[0].Bundles[0].Checks[0].ImageViewData[0].ImageData = data

	var buf bytes.Buffer
	w := NewWriter(&buf, WriteVariableLineLengthOption())
	require.NoError(t, w.Write(&iclFile))

	fileReader := bytes.NewReader(buf.Bytes())
	r = NewReader(fileReader, ReadVariableLineLengthOption())
	_, err = r.Read()
	require.Error(t, err)

	fileErr := getFileError(t, err)
	require.Equal(t, bufio.ErrTooLong.Error(), fileErr.Msg)

	fileReader.Reset(buf.Bytes())
	r = NewReader(fileReader, ReadVariableLineLengthOption(), BufferSizeOption(256*1024))
	_, err = r.Read()
	require.NoError(t, err)
}

func Test_DecodeEBCDIC(t *testing.T) {
	// test with valid sample
	decoded, err := DecodeEBCDIC(string([]byte{0xF0, 0xF1, 0xF2}))
	require.NoError(t, err)
	require.Equal(t, "012", decoded)

	// replaced invalid code with utf8.RuneError
	decoded, err = DecodeEBCDIC(string([]byte{0x04, 0x06, 0xff}))
	require.NoError(t, err)
	r, n := utf8.DecodeRuneInString(decoded)
	require.Equal(t, 3, n)
	require.Equal(t, utf8.RuneError, r)
}
