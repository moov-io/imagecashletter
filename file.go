// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

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

// ToDo: Errors specific to Files

// File is an ICL file
type File struct {
	// ID is a client defined string used as a reference to this record
	ID string `json:"id"`
	// FileHeader is an ICL FileHeader
	FileHeader FileHeader `json:"fileHeader"`
	// CashLetters are ICl Cash Letters
	CashLetters []CashLetter `json:"cashLetters"`
	// FileHeader is an ICL FileControl
	FileControl FileControl `json:"fileControl"`
}

// NewFile constructs a file template with a FileHeader and FileControl.
func NewFile() *File {
	return &File{
		FileHeader:  NewFileHeader(),
		FileControl: NewFileControl(),
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
