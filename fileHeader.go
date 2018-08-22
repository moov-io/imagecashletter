// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "time"

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format
// ToDo: ASCII vs EBCDIC

// Errors specific to a FileHeader Record

// FileHeader Record is mandatory
type FileHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// recordType defines the type of record.
	recordType string
	// standardLevel identifies the standard level of the file.  Current support is for DSTU X9.37 - 2003
	// Values: 03, 30, 35
	// 03: DSTU X9.37 - 2003
	// 30: X9.100-187-2008
	// 35: X9.100-187-2013 and 2016
	StandardLevel string `json:"standardLevel"`
	// TestFileIndicator identifies whether the file is a test or production file.
	// Values:
	// T: Test File
	// P: Production File
	TestFileIndicator string `json:"testIndicator"`
	// ImmediateDestination contains the routing and transit number of the Federal Reserve Bank
	// (FRB) or receiver to which the file is being sent.  Format: TTTTAAAAC, where:
	//  TTTT Federal Reserve Prefix
	//  AAAA ABA Institution Identifier
	//  C Check Digit
	//  For a number that identifies a non-financial institution: NNNNNNNNN
	ImmediateDestination string `json:"immediateDestination"`
	// ImmediateOrigin contains the routing and transit number of the Federal Reserve Bank
	// (FRB) or receiver from which the file is being sent.  Format: TTTTAAAAC, where:
	// TTTT Federal Reserve Prefix
	// AAAA ABA Institution Identifier
	// C Check Digit
	// For a number that identifies a non-financial institution: NNNNNNNNN
	ImmediateOrigin string `json:"immediateOrigin"`
	// FileCreationDate is the date that the immediate origin institution creates the file.  Default time shall be in
	// Eastern Time zone format. Other time zones may be used under clearing arrangements.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	FileCreationDate time.Time `json:"fileCreationDate"`
	// FileCreationTime is the time the immediate origin institution creates the file. Default time shall be in
	// Eastern Time zone format. Other time zones may be used under clearing arrangements.
	// Format: hhmm, where: hh hour, mm minute
	// Values:
	// hh '00' through '23'
	// mm '00' through '59'
	FileCreationTime time.Time `json:"fileCreationTime"`
	// ResendIndicator indicates whether the file has been previously transmitted.
	// Values:
	// Y: Yes
	// N: No
	ResendIndicator string `json:"ResendIndicator"`
	// ImmediateDestinationName identifies the short name of the institution that receives the file.
	ImmediateDestinationName string `json:"immediateDestinationName"`
	// ImmediateOriginName identifies the short name of the institution that sends the file.
	ImmediateOriginName string `json:"ImmediateOriginName"`
	// FileIDModifier is a code that permits multiple files, created on the same date, same time and between the
	// same institutions, to be distinguished one from another. If all of the following fields in a previous file are
	// equal to the same fields in this file: FileHeader ImmediateDestination, ImmediateOrigin, FileCreationDate, and
	// FileCreationTime, it must be defined.
	FileIDModifier string `json:"fileIDModifier"`
	// CountryCode is a 2-character code as approved by the International Organization for Standardization (ISO) used
	// to identify the country in which the payer bank is located.
	// Example: US = United States
	// Values for other countries can be found on the International Organization for Standardization
	// website: www.iso.org.
	CountryCode string `json:"countryCode"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// CompanionDocumentIndicator identifies a field used to indicate the Companion Document being used.
	//
	// Shall be present only under clearing arrangements. Companion Document usage and values
	// defined by clearing arrangements.
	// Values:
	// 0–9 Reserved for United States use
	// A–J Reserved for Canadian use
	// Other - as defined by clearing arrangements.
	CompanionDocumentIndicator string `json:"companionDocumentIndicator"`
	// validator is composed for x9 data validation
	validator
	// converters is composed for x9 to golang Converters
	converters
}

// NewFileHeader returns a new FileHeader with default values for non exported fields
func NewFileHeader() FileHeader {
	fh := FileHeader{
		recordType:    "01",
		StandardLevel: "03",
	}
	return fh
}

// Parse takes the input record string and parses the FileHeader values
func (fh *FileHeader) Parse(record string) {
	// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format)
	// (character position 1-2) Always "01"
	fh.recordType = "01"
	// (03-04)
	fh.StandardLevel = fh.parseStringField(record[2:4])
	//

	//

	//

	//

	//

	//

	//

	//

	//

}

// String writes the FileHeader struct to a variable length string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
