// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

// Errors specific to a ImageViewAnalysis Record

// ImageViewAnalysis Record
type ImageViewAnalysis struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
}

// NewImageViewAnalysis returns a new ImageViewAnalysis with default values for non exported fields
func NewImageViewAnalysis() *ImageViewAnalysis {
	imageAnalysis := &ImageViewAnalysis{
		recordType: "54",
	}
	return imageAnalysis
}

// Parse takes the input record string and parses the ImageViewAnalysis values

// String writes the ImageViewAnalysis struct to a variable length string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
