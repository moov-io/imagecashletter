// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"

	"github.com/gdamore/encoding"
)

// ParseError is returned for parsing reader errors.
// The first line is 1.
type ParseError struct {
	Line   int    // Line number where the error occurred
	Record string // Name of the record type being parsed
	Err    error  // The actual error
}

func (e *ParseError) Error() string {
	if e.Record == "" {
		return fmt.Sprintf("line:%d %T %s", e.Line, e.Err, e.Err)
	}
	return fmt.Sprintf("line:%d record:%s %T %s", e.Line, e.Record, e.Err, e.Err)
}

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	scanner *bufio.Scanner
	// file is ach.file model being built as r is parsed.
	File File
	// func used to decode line to desired encoding ie. ASCII,EBCDIC
	decodeLine     DecodeLineFn
	ebcdicEncoding bool
	// line is the current line being parsed from the input r
	line string
	// currentCashLetter is the current CashLetter being parsed
	currentCashLetter CashLetter
	// line number of the file being parsed
	lineNum int
	// recordName holds the current record name being parsed.
	recordName string
}

// error creates a new ParseError based on err.
func (r *Reader) error(err error) error {
	return &ParseError{
		Line:   r.lineNum,
		Record: r.recordName,
		Err:    err,
	}
}

// addCurrentCashLetter creates the current cash letter for the file being read. A successful
// currentCashLetter will be added to r.File once parsed.
func (r *Reader) addCurrentCashLetter(cashLetter CashLetter) {
	r.currentCashLetter = cashLetter
}

// addCurrentBundle creates the CurrentBundle for the file being read. A successful
// currentBundle will be added to r.File once parsed.
func (r *Reader) addCurrentBundle(bundle *Bundle) {
	r.currentCashLetter.currentBundle = bundle
}

// addCurrentRoutingNumberSummary creates the CurrentRoutingNumberSummary for the file being read. A successful
// currentRoutingNumberSummary will be added to r.File once parsed.
func (r *Reader) addCurrentRoutingNumberSummary(rns *RoutingNumberSummary) {
	r.currentCashLetter.currentRoutingNumberSummary = rns
}

// NewReader returns a new ACH Reader that reads from r.
func NewReader(r io.Reader, opts ...ReaderOption) *Reader {
	f := NewFile()
	f.Control = FileControl{}
	reader := &Reader{
		File:       *f,
		scanner:    bufio.NewScanner(r),
		decodeLine: Passthrough,
	}
	for _, opt := range opts {
		opt(reader)
	}
	return reader
}

// DecodeLineFn is used to decode a scanned line into desired encoding.
// Depending on X9 spec, cashletter could be encoded as ASCII or EBCDIC
type DecodeLineFn func(lineIn string) (lineOut string)

// Passthrough will return line as is
func Passthrough(lineIn string) (lineOut string) {
	return lineIn
}

// DecodeEBCDIC will decode a line from EBCDIC-0037 to UTF-8
func DecodeEBCDIC(lineIn string) (lineOut string) {
	lineOut, err := encoding.EBCDIC.NewDecoder().String(lineIn)
	if err != nil {
		fmt.Printf("Error decoding '%X' as EBCDIC: %v\n", lineIn, err)
	}
	return lineOut
}

// ReaderOption can be used to change default behavior of Reader
type ReaderOption func(*Reader)

//ReadVariableLineLengthOption allows Reader to split imagecashletter files based on encoded line lengths
func ReadVariableLineLengthOption() ReaderOption {
	scanVariableLengthLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// nothing to scan
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// line length can be variable
		// use the 4 control bytes at the beginning of a line to determine its length
		ctrl := data[0:4]
		dataLen := int(binary.BigEndian.Uint32(ctrl))
		lineLen := 4 + dataLen

		// the last calculated line is expected to match the remaining bytes
		if atEOF && lineLen != len(data) {
			return len(data), data, io.ErrUnexpectedEOF
		}

		// return line while accounting for control bytes
		if lineLen <= len(data) {
			return lineLen, data[4:lineLen], nil
		}

		// request more data.
		return 0, nil, nil
	}

	return func(r *Reader) {
		r.scanner.Split(scanVariableLengthLines)
	}
}

//ReadEbcdicEncodingOption allows Reader to decode scanned lines from EBCDIC to UTF-8
func ReadEbcdicEncodingOption() ReaderOption {
	return func(r *Reader) {
		r.decodeLine = DecodeEBCDIC
	}
}

// Read reads each line of the imagecashletter file and defines which parser to use based
// on the first character of each line. It also enforces imagecashletter formatting rules and returns
// the appropriate error if issues are found.
func (r *Reader) Read() (File, error) {
	r.lineNum = 0
	// read through the entire file
	for r.scanner.Scan() {
		r.line = r.scanner.Text()
		r.lineNum++

		lineLength := len(r.line)

		if lineLength < 80 {
			msg := fmt.Sprintf(msgRecordLength, lineLength)
			err := &FileError{FieldName: "RecordLength", Value: strconv.Itoa(lineLength), Msg: msg}
			return r.File, r.error(err)
		}
		if err := r.parseLine(); err != nil {
			return r.File, err
		}
	}
	if (FileHeader{}) == r.File.Header {
		// There must be at least one File Header
		r.recordName = "FileHeader"
		return r.File, r.error(&FileError{Msg: msgFileHeader})
	}
	if (FileControl{}) == r.File.Control {
		// There must be at least one File Control
		r.recordName = "FileControl"
		return r.File, r.error(&FileError{Msg: msgFileControl})
	}
	return r.File, nil
}

func (r *Reader) parseLine() error {
	switch r.line[:2] {
	case fileHeaderPos, fileHeaderEbcPos:
		if err := r.parseFileHeader(); err != nil {
			return err
		}
	case cashLetterHeaderPos, cashLetterHeaderEbcPos:
		if err := r.parseCashLetterHeader(); err != nil {
			return err
		}
	case bundleHeaderPos, bundleHeaderEbcPos:
		if err := r.parseBundleHeader(); err != nil {
			return err
		}
	case checkDetailPos, checkDetailEbcPos:
		if err := r.parseCheckDetail(); err != nil {
			return err
		}
	case checkDetailAddendumAPos, checkDetailAddendumAEbcPos:
		if err := r.parseCheckDetailAddendumA(); err != nil {
			return err
		}
	case checkDetailAddendumBPos, checkDetailAddendumBEbcPos:
		if err := r.parseCheckDetailAddendumB(); err != nil {
			return err
		}
	case checkDetailAddendumCPos, checkDetailAddendumCEbcPos:
		if err := r.parseCheckDetailAddendumC(); err != nil {
			return err
		}
	case imageViewDetailPos, imageViewDetailEbcPos:
		if err := r.parseImageViewDetail(); err != nil {
			return err
		}
	case imageViewDataPos, imageViewDataEbcPos:
		if err := r.parseImageViewData(); err != nil {
			return err
		}
	case imageViewAnalysisPos, imageViewAnalysisEbcPos:
		if err := r.parseImageViewAnalysis(); err != nil {
			return err
		}
	case returnDetailPos, returnDetailEbcPos:
		if err := r.parseReturnDetail(); err != nil {
			return err
		}
	case returnAddendumAPos, returnAddendumAPEbcos:
		if err := r.parseReturnDetailAddendumA(); err != nil {
			return err
		}
	case returnAddendumBPos, returnAddendumBEbcPos:
		if err := r.parseReturnDetailAddendumB(); err != nil {
			return err
		}
	case returnAddendumCPos, returnAddendumCEbcPos:
		if err := r.parseReturnDetailAddendumC(); err != nil {
			return err
		}
	case returnAddendumDPos, returnAddendumDEbcPos:
		if err := r.parseReturnDetailAddendumD(); err != nil {
			return err
		}
	case creditItemPos, creditItemEbcPos:
		if err := r.parseCreditItem(); err != nil {
			return err
		}
	case bundleControlPos, bundleControlEbcPos:
		if err := r.parseBundleControl(); err != nil {
			return err
		}
		if r.currentCashLetter.currentBundle == nil {
			r.error(&FileError{Msg: msgFileBundleControl})
		}
		// Add Bundle or ReturnBundle to CashLetter
		if r.currentCashLetter.currentBundle != nil {
			if err := r.currentCashLetter.currentBundle.Validate(); err != nil {
				r.recordName = "Bundles"
				return r.error(err)
			}
			r.currentCashLetter.AddBundle(r.currentCashLetter.currentBundle)
			r.currentCashLetter.currentBundle = new(Bundle)
		}
	case routingNumberSummaryPos, routingNumberSummaryEbcPos:
		if err := r.parseRoutingNumberSummary(); err != nil {
			return err
		}
		r.currentCashLetter.AddRoutingNumberSummary(r.currentCashLetter.currentRoutingNumberSummary)
		r.currentCashLetter.currentRoutingNumberSummary = new(RoutingNumberSummary)
	case cashLetterControlPos, cashLetterControlEbcPos:
		if err := r.parseCashLetterControl(); err != nil {
			return err
		}
		if err := r.currentCashLetter.Validate(); err != nil {
			r.recordName = "CashLetters"
			return r.error(err)
		}
		r.File.AddCashLetter(r.currentCashLetter)
		r.currentCashLetter = CashLetter{}
	case fileControlPos, fileControlEbcPos:
		if err := r.parseFileControl(); err != nil {
			return err
		}
	default:
		msg := fmt.Sprintf(msgUnknownRecordType, r.line[:2])
		return r.error(&FileError{FieldName: "recordType", Value: r.line[:2], Msg: msg})
	}
	return nil
}

// parseFileHeader takes the input record string and parses the FileHeader values
func (r *Reader) parseFileHeader() error {
	r.recordName = "FileHeader"
	if (FileHeader{}) != r.File.Header {
		// There can only be one File Header per File
		r.error(&FileError{Msg: msgFileHeader})
	}
	r.File.Header.Parse(r.decodeLine(r.line))
	// Ensure valid FileHeader
	if err := r.File.Header.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseCashLetterHeader takes the input record string and parses the CashLetterHeader values
func (r *Reader) parseCashLetterHeader() error {
	r.recordName = "CashLetterHeader"
	if r.currentCashLetter.CashLetterHeader != nil {
		// CashLetterHeader inside of current cash letter
		return r.error(&FileError{Msg: msgFileCashLetterInside})
	}
	clh := NewCashLetterHeader()
	clh.Parse(r.decodeLine(r.line))
	// Ensure we have a valid CashLetterHeader
	if err := clh.Validate(); err != nil {
		return r.error(err)
	}
	// Passing CashLetterHeader into NewCashLetter creates a CashLetter
	cl := NewCashLetter(clh)
	r.addCurrentCashLetter(cl)
	return nil
}

// parseBundleHeader takes the input record string and parses the BundleHeader values
func (r *Reader) parseBundleHeader() error {
	r.recordName = "BundleHeader"
	if r.currentCashLetter.currentBundle != nil {
		// BundleHeader inside of current Bundle
		if r.currentCashLetter.currentBundle.BundleHeader != nil {
			return r.error(&FileError{Msg: msgFileBundleInside})
		}
	}
	// Ensure we have a valid bundle header before building a bundle.
	bh := NewBundleHeader()
	bh.Parse(r.decodeLine(r.line))
	if err := bh.Validate(); err != nil {
		return r.error(err)
	}
	// Passing BundleHeader into NewBundle creates a Bundle
	bundle := NewBundle(bh)
	r.addCurrentBundle(bundle)
	return nil

}

// parseCheckDetail takes the input record string and parses the CheckDetail values
func (r *Reader) parseCheckDetail() error {
	r.recordName = "CheckDetail"
	if r.currentCashLetter.currentBundle == nil {
		return r.error(&FileError{Msg: msgFileBundleOutside})
	}
	cd := new(CheckDetail)
	cd.Parse(r.decodeLine(r.line))
	// Ensure valid CheckDetail
	if err := cd.Validate(); err != nil {
		return r.error(err)
	}
	// Add CheckDetail
	if r.currentCashLetter.currentBundle.BundleHeader != nil {
		r.currentCashLetter.currentBundle.AddCheckDetail(cd)
	}
	return nil
}

// parseCheckDetailAddendumA takes the input record string and parses the CheckDetailAddendumA values
func (r *Reader) parseCheckDetailAddendumA() error {
	r.recordName = "CheckDetailAddendumA"
	if r.currentCashLetter.currentBundle.GetChecks() == nil {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "CheckDetailAddendumA", Msg: msg})
	}
	cdAddendumA := NewCheckDetailAddendumA()
	cdAddendumA.Parse(r.decodeLine(r.line))
	if err := cdAddendumA.Validate(); err != nil {
		return r.error(err)
	}
	entryIndex := len(r.currentCashLetter.currentBundle.GetChecks()) - 1
	//r.currentCashLetter.currentBundle.Checks[entryIndex].CheckDetailAddendumA = cdAddendumA
	r.currentCashLetter.currentBundle.Checks[entryIndex].AddCheckDetailAddendumA(cdAddendumA)
	return nil
}

// parseCheckDetailAddendumB takes the input record string and parses the CheckDetailAddendumB values
func (r *Reader) parseCheckDetailAddendumB() error {
	r.recordName = "CheckDetailAddendumB"
	if r.currentCashLetter.currentBundle.GetChecks() == nil {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "CheckDetailAddendumB", Msg: msg})
	}
	cdAddendumB := NewCheckDetailAddendumB()
	cdAddendumB.Parse(r.decodeLine(r.line))
	if err := cdAddendumB.Validate(); err != nil {
		return r.error(err)
	}
	entryIndex := len(r.currentCashLetter.currentBundle.GetChecks()) - 1
	r.currentCashLetter.currentBundle.Checks[entryIndex].AddCheckDetailAddendumB(cdAddendumB)
	return nil
}

// parseCheckDetailAddendumC takes the input record string and parses the CheckDetailAddendumC values
func (r *Reader) parseCheckDetailAddendumC() error {
	r.recordName = "CheckDetailAddendumC"
	if r.currentCashLetter.currentBundle.GetChecks() == nil {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "CheckDetailAddendumC", Msg: msg})
	}
	cdAddendumC := NewCheckDetailAddendumC()
	cdAddendumC.Parse(r.decodeLine(r.line))
	if err := cdAddendumC.Validate(); err != nil {
		return r.error(err)
	}
	entryIndex := len(r.currentCashLetter.currentBundle.GetChecks()) - 1
	r.currentCashLetter.currentBundle.Checks[entryIndex].AddCheckDetailAddendumC(cdAddendumC)
	return nil
}

// parseReturnDetail takes the input record string and parses the ReturnDetail values
func (r *Reader) parseReturnDetail() error {
	r.recordName = "ReturnDetail"
	if r.currentCashLetter.currentBundle == nil {
		return r.error(&FileError{Msg: msgFileBundleOutside})
	}
	rd := new(ReturnDetail)
	rd.Parse(r.decodeLine(r.line))
	if err := rd.Validate(); err != nil {
		return r.error(err)
	}
	if r.currentCashLetter.currentBundle.BundleHeader != nil {
		r.currentCashLetter.currentBundle.AddReturnDetail(rd)
	}
	return nil
}

// parseReturnDetailAddendumA takes the input record string and parses the ReturnDetailAddendumA values
func (r *Reader) parseReturnDetailAddendumA() error {
	r.recordName = "ReturnDetailAddendumA"
	if r.currentCashLetter.currentBundle.GetReturns() == nil {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "ReturnDetailAddendumA", Msg: msg})
	}
	rdAddendumA := NewReturnDetailAddendumA()
	rdAddendumA.Parse(r.decodeLine(r.line))
	if err := rdAddendumA.Validate(); err != nil {
		return r.error(err)
	}
	entryIndex := len(r.currentCashLetter.currentBundle.GetReturns()) - 1
	//r.currentCashLetter.currentBundle.Returns[entryIndex].ReturnDetailAddendumA = rdAddendumA
	r.currentCashLetter.currentBundle.Returns[entryIndex].AddReturnDetailAddendumA(rdAddendumA)
	return nil
}

// parseReturnDetailAddendumB takes the input record string and parses the ReturnDetailAddendumB values
func (r *Reader) parseReturnDetailAddendumB() error {
	r.recordName = "ReturnDetailAddendumB"
	if r.currentCashLetter.currentBundle.GetReturns() == nil {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "ReturnDetailAddendumB", Msg: msg})
	}
	rdAddendumB := NewReturnDetailAddendumB()
	rdAddendumB.Parse(r.decodeLine(r.line))
	if err := rdAddendumB.Validate(); err != nil {
		return r.error(err)
	}
	entryIndex := len(r.currentCashLetter.currentBundle.GetReturns()) - 1
	r.currentCashLetter.currentBundle.Returns[entryIndex].AddReturnDetailAddendumB(rdAddendumB)
	return nil
}

// parseReturnDetailAddendumC takes the input record string and parses the ReturnDetailAddendumC values
func (r *Reader) parseReturnDetailAddendumC() error {
	r.recordName = "ReturnDetailAddendumC"
	if r.currentCashLetter.currentBundle.GetReturns() == nil {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "ReturnDetailAddendumC", Msg: msg})
	}
	rdAddendumC := NewReturnDetailAddendumC()
	rdAddendumC.Parse(r.decodeLine(r.line))
	if err := rdAddendumC.Validate(); err != nil {
		return r.error(err)
	}
	entryIndex := len(r.currentCashLetter.currentBundle.GetReturns()) - 1
	r.currentCashLetter.currentBundle.Returns[entryIndex].AddReturnDetailAddendumC(rdAddendumC)
	return nil
}

// parseReturnDetail*AddendumD takes the input record string and parses the ReturnDetail*AddendumD values
func (r *Reader) parseReturnDetailAddendumD() error {
	r.recordName = "ReturnDetailAddendumD"

	if r.currentCashLetter.currentBundle.GetReturns() == nil {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "ReturnDetailAddendumD", Msg: msg})
	}
	rdAddendumD := NewReturnDetailAddendumD()
	rdAddendumD.Parse(r.decodeLine(r.line))
	if err := rdAddendumD.Validate(); err != nil {
		return r.error(err)
	}
	entryIndex := len(r.currentCashLetter.currentBundle.GetReturns()) - 1
	r.currentCashLetter.currentBundle.Returns[entryIndex].AddReturnDetailAddendumD(rdAddendumD)
	return nil
}

// parseImageViewDetail takes the input record string and parses the ImageViewDetail values
func (r *Reader) parseImageViewDetail() error {
	r.recordName = "ImageViewDetail"
	if err := r.ImageViewDetail(); err != nil {
		return err
	}
	return nil
}

// ImageViewDetail takes the input record string and parses ImageViewDetail for a check
func (r *Reader) ImageViewDetail() error {
	if r.currentCashLetter.currentBundle.GetChecks() != nil {
		ivDetail := NewImageViewDetail()
		ivDetail.Parse(r.decodeLine(r.line))
		if err := ivDetail.Validate(); err != nil {
			return r.error(err)
		}
		entryIndex := len(r.currentCashLetter.currentBundle.GetChecks()) - 1
		r.currentCashLetter.currentBundle.Checks[entryIndex].AddImageViewDetail(ivDetail)

	} else if r.currentCashLetter.currentBundle.GetReturns() != nil {
		ivDetail := NewImageViewDetail()
		ivDetail.Parse(r.decodeLine(r.line))
		if err := ivDetail.Validate(); err != nil {
			return r.error(err)
		}
		entryIndex := len(r.currentCashLetter.currentBundle.GetReturns()) - 1
		r.currentCashLetter.currentBundle.Returns[entryIndex].AddImageViewDetail(ivDetail)
	} else {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "ImageViewDetail", Msg: msg})
	}

	return nil
}

// parseImageViewData takes the input record string and parses the ImageViewData values
func (r *Reader) parseImageViewData() error {
	r.recordName = "ImageViewData"
	if err := r.ImageViewData(); err != nil {
		return err
	}
	return nil
}

// ImageViewData takes the input record string and parses ImageViewData for a check
func (r *Reader) ImageViewData() error {
	if r.currentCashLetter.currentBundle.GetChecks() != nil {
		ivData := NewImageViewData()
		ivData.ParseAndDecode(r.line, r.decodeLine)
		if err := ivData.Validate(); err != nil {
			return r.error(err)
		}
		entryIndex := len(r.currentCashLetter.currentBundle.GetChecks()) - 1
		r.currentCashLetter.currentBundle.Checks[entryIndex].AddImageViewData(ivData)

	} else if r.currentCashLetter.currentBundle.GetReturns() != nil {
		ivData := NewImageViewData()
		ivData.ParseAndDecode(r.line, r.decodeLine)
		if err := ivData.Validate(); err != nil {
			return r.error(err)
		}
		entryIndex := len(r.currentCashLetter.currentBundle.GetReturns()) - 1
		r.currentCashLetter.currentBundle.Returns[entryIndex].AddImageViewData(ivData)
	} else {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "ImageViewData", Msg: msg})
	}

	return nil
}

// parseImageViewAnalysis takes the input record string and parses ImageViewAnalysis values
func (r *Reader) parseImageViewAnalysis() error {
	r.recordName = "ImageViewAnalysis"
	if err := r.ImageViewAnalysis(); err != nil {
		return err
	}
	return nil
}

// ImageViewAnalysis takes the input record string and parses ImageViewAnalysis for a check
func (r *Reader) ImageViewAnalysis() error {
	if r.currentCashLetter.currentBundle.GetChecks() != nil {
		ivAnalysis := NewImageViewAnalysis()
		ivAnalysis.Parse(r.decodeLine(r.line))
		if err := ivAnalysis.Validate(); err != nil {
			return r.error(err)
		}
		entryIndex := len(r.currentCashLetter.currentBundle.GetChecks()) - 1
		r.currentCashLetter.currentBundle.Checks[entryIndex].AddImageViewAnalysis(ivAnalysis)

	} else if r.currentCashLetter.currentBundle.GetReturns() != nil {
		ivAnalysis := NewImageViewAnalysis()
		ivAnalysis.Parse(r.decodeLine(r.line))
		if err := ivAnalysis.Validate(); err != nil {
			return r.error(err)
		}
		entryIndex := len(r.currentCashLetter.currentBundle.GetReturns()) - 1
		r.currentCashLetter.currentBundle.Returns[entryIndex].AddImageViewAnalysis(ivAnalysis)
	} else {
		msg := fmt.Sprint(msgFileBundleOutside)
		return r.error(&FileError{FieldName: "ImageViewAnalysis", Msg: msg})
	}

	return nil
}

// parseCreditItem takes the input record string and parses the CreditItem values
func (r *Reader) parseCreditItem() error {
	// Current implementation has the credit letter outside the bundle but within the cash letter
	r.recordName = "CreditItem"
	if r.currentCashLetter.CashLetterHeader == nil {
		return r.error(&FileError{Msg: msgFileCreditItem})
	}
	ci := new(CreditItem)
	ci.Parse(r.decodeLine(r.line))
	if err := ci.Validate(); err != nil {
		return r.error(err)
	}
	r.currentCashLetter.AddCreditItem(ci)
	return nil
}

// parseBundleControl takes the input record string and parses the BundleControl values
func (r *Reader) parseBundleControl() error {
	r.recordName = "BundleControl"

	if r.currentCashLetter.currentBundle == nil || r.currentCashLetter.currentBundle.BundleControl == nil {
		return r.error(&FileError{Msg: msgFileBundleControl})
	}
	r.currentCashLetter.currentBundle.GetControl().Parse(r.decodeLine(r.line))
	if err := r.currentCashLetter.currentBundle.GetControl().Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseRoutingNumberSummary takes the input record string and parses the RoutingNumberSummary values
func (r *Reader) parseRoutingNumberSummary() error {
	r.recordName = "RoutingNumberSummary"
	if r.currentCashLetter.CashLetterHeader == nil {
		return r.error(&FileError{Msg: msgFileRoutingNumberSummary})
	}

	rns := NewRoutingNumberSummary()
	rns.Parse(r.decodeLine(r.line))
	if err := rns.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseCashLetterControl takes the input record string and parses the CashLetterControl values
func (r *Reader) parseCashLetterControl() error {
	r.recordName = "CashLetterControl"
	if r.currentCashLetter.CashLetterHeader == nil {
		// CashLetterControl without a current CashLetter
		return r.error(&FileError{Msg: msgFileCashLetterControl})
	}
	r.currentCashLetter.GetControl().Parse(r.decodeLine(r.line))
	// Ensure valid CashLetterControl
	if err := r.currentCashLetter.GetControl().Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseFileControl takes the input record string and parses the FileControl values
func (r *Reader) parseFileControl() error {
	r.recordName = "FileControl"
	if (FileControl{}) != r.File.Control {
		// Can be only one file control per file
		return r.error(&FileError{Msg: msgFileControl})
	}
	r.File.Control.Parse(r.decodeLine(r.line))
	// Ensure valid FileControl
	if err := r.File.Control.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}
