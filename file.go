// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "fmt"

// ICL File Records that are identified as Mandatory are required to support Federal Reserve processing of an image
// file.
//
// https://en.wikipedia.org/wiki/Substitute_check
//
// http://www.frbservices.org
//
// The Federal Reserve Banks uses the Accredited Standards Committee X9’s Specifications (X9.100-187–2016) for
// Electronic Exchange of Check and Image Data in providing its suite of Check 21 services.
//
// Record Types
const (
	fileHeaderPos           = "01"
	cashLetterHeaderPos     = "10"
	bundleHeaderPos         = "20"
	checkDetailPos          = "25"
	checkDetailAddendumAPos = "26"
	checkDetailAddendumBPos = "27"
	checkDetailAddendumCPos = "28"
	returnPos               = "31"
	returnAddendumAPos      = "32"
	returnAddendumBPos      = "33"
	returnAddendumCPos      = "34"
	returnAddendumDPos      = "35"
	//no longer supported by the standard - accountTotalsDetailPos  = "40"
	//no longer supported by the standard  - nonHitTotalsDetailPos   = "41"
	imageViewDetailPos   = "50"
	imageViewDataPos     = "52"
	imageViewAnalysisPos = "54"
	bundleControlPos     = "70"
	//no longer supported by the standard - boxSummaryPos           = "75"
	routingNumberSummaryPos = "85"
	cashLetterControlPos    = "90"
	fileControlPos          = "99"
)

// Errors strings specific to parsing a Batch container
var (
	//msgFileCalculatedControlEquality = "calculated %v is out-of-balance with control %v"
	// specific messages
	msgRecordLength = "must be at least 80 characters and found %d"
	//msgFileBatchOutside  = "outside of current batch"
	msgFileCashLetterInside   = "inside of current cash letter"
	msgFileControl       = "none or more than one file control exists"
	msgFileHeader        = "none or more than one file headers exists"
	msgUnknownRecordType = "%s is an unknown record type"
)

// FileError is an error describing issues validating a file
type FileError struct {
	FieldName string
	Value     string
	Msg       string
}

func (e *FileError) Error() string {
	return fmt.Sprintf("%s %s", e.FieldName, e.Msg)
}

// File is an ICL file
type File struct {
	// ID is a client defined string used as a reference to this record
	ID string `json:"id"`
	// FileHeader is an ICL FileHeader
	Header FileHeader `json:"fileHeader"`
	// CashLetters are ICl Cash Letters
	CashLetters []CashLetter `json:"cashLetters,omitempty"`
	// ToDo: Current logic to add a Bundle is in CashLetter
	// Bundles are ICL Bundles
	//Bundles []Bundle `json"bundles"`
	// FileControl is an ICL FileControl
	Control FileControl `json:"fileControl"`
}

// NewFile constructs a file template with a FileHeader and FileControl.
func NewFile() *File {
	return &File{
		Header:  NewFileHeader(),
		Control: NewFileControl(),
	}
}

// Create creates a valid ICL File
func (f *File) Create() error {
	return nil
}

// Validate validates an ICL File
func (f *File) Validate() error {
	return nil
}

// SetHeader allows for header to be built.
func (f *File) SetHeader(h FileHeader) *File {
	f.Header = h
	return f
}

// AddCashLetter appends a CashLetter to the x9.File
func (f *File) AddCashLetter(cashLetter CashLetter) []CashLetter {
	f.CashLetters = append(f.CashLetters, cashLetter)
	return f.CashLetters
}

// ToDo: Current logic to add a Bundle is in CashLetter
/*// AddBundle appends a Bundle to the x9.File
func (f *File) AddBundle(bundle Bundle) []Bundle {
	f.Bundles = append(f.Bundles, bundle)
	return f.Bundles
}*/
