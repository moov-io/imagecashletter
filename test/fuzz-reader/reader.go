// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzzreader

import (
	"bytes"

	"github.com/moov-io/imagecashletter"
)

// Return codes (from go-fuzz docs)
//
// The function must return 1 if the fuzzer should increase priority
// of the given input during subsequent fuzzing (for example, the input is
// lexically correct and was parsed successfully); -1 if the input must not be
// added to corpus even if gives new coverage; and 0 otherwise; other values are
// reserved for future use.
func Fuzz(data []byte) int {
	r := imagecashletter.NewReader(bytes.NewReader(data))
	f, err := r.Read()
	if err != nil {
		// if f != nil {
		// 	panic(fmt.Sprintf("f != nil on err != nil: %v", f))
		// }
		return 0
	}

	if f.ID != "" {
		return 1
	}
	if n := checkFileHeader(f.Header); n > 0 {
		return n
	}
	if len(f.CashLetters) > 0 || len(f.Bundles) > 0 {
		return 1
	}

	return 0
}

func checkFileHeader(h imagecashletter.FileHeader) int {
	if h.ImmediateDestination != "" || h.ImmediateOrigin != "" {
		return 1
	}
	if !h.FileCreationDate.IsZero() || !h.FileCreationTime.IsZero() {
		return 1
	}
	if h.ImmediateDestinationName != "" || h.ImmediateOriginName != "" {
		return 1
	}
	if h.CountryCode != "" {
		return 1
	}
	return 0
}
