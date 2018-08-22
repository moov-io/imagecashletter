// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.
//
// Package x9 reads and writes x9 files.
//
// https://en.wikipedia.org/wiki/Substitute_check
//
// http://www.frbservices.org.
//
// The Federal Reserve Banks uses the Accredited Standards Committee X9’s Specifications for Electronic Exchange of
// Check and Image Data in providing its suite of Check 21 services.

package x9

// ICL File Records that are identified as Mandatory are required to support Federal Reserve processing of an image
// file.

// First two position of all Record Types. These codes are uniquely assigned to
// the first 2 bytes of each row in a file.
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
	accountTotalsDetailPos  = "40"
	nonHitTotalsDetailPos   = "41"
	imageViewDetailPos      = "50"
	imageViewDataPos        = "52"
	imageViewAnalysisPos    = "54"
	bundleControlPos        = "70"
	boxSummaryPos           = "75"
	routingNumberSummaryPos = "85"
	cashLetterControlPos    = "90"
	fileControlPos          = "99"
)

// File is an ICL file
type File struct {
	// ID is a client defined string used as a reference to this record
	ID string `json:"id"`
	// FileHeader is an ICL FileHeader
	FileHeader FileHeader `json:"fileHeader"`
	// FileHeader is an ICL FileControl
	FileControl FileControl `json:"fileControl"`
}
