// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"bufio"
	"io"
	"strings"
)

// A Writer writes an x9.file to an ICL encoded file.
//
// As returned by NewWriter, a Writer writes x9file structs into
// x9 formatted files.
//

// Writer struct
type Writer struct {
	w       *bufio.Writer
	lineNum int //current line being written
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(w),
	}
}

// Writer writes a single x9.file record to w
func (w *Writer) Write(file *File) error {
	if err := file.Validate(); err != nil {
		return err
	}
	w.lineNum = 0
	// Iterate over all records in the file
	if _, err := w.w.WriteString(file.Header.String() + "\n"); err != nil {
		return err
	}
	w.lineNum++

	if err := w.writeCashLetter(file); err != nil {
		return err
	}
	if _, err := w.w.WriteString(file.Control.String() + "\n"); err != nil {
		return err
	}
	w.lineNum++
	// pad the final block
	for i := 0; i < (10-(w.lineNum%10)) && w.lineNum%10 != 0; i++ {
		if _, err := w.w.WriteString(strings.Repeat("9", 94) + "\n"); err != nil {
			return err
		}
	}
	return w.w.Flush()
}

// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occurred during the Flush, call Error.
func (w *Writer) Flush() {
	w.w.Flush()
}

// writeCashLetter writes a CashLetter to a file
func (w *Writer) writeCashLetter(file *File) error {
	//CashLetters
	for _, cl := range file.CashLetters {
		if _, err := w.w.WriteString(cl.GetHeader().String() + "\n"); err != nil {
			return err
		}
		w.lineNum++

		// Write Bundles
		if err := w.writeBundle(cl); err != nil {
			return err
		}

		if _, err := w.w.WriteString(cl.GetControl().String() + "\n"); err != nil {
			return err
		}
		w.lineNum++
	}
	return nil
}

// writeBundle writes a Bundle to a file
func (w *Writer) writeBundle(cl CashLetter) error {

	for _, b := range cl.GetBundles() {
		if _, err := w.w.WriteString(b.GetHeader().String() + "\n"); err != nil {
			return err
		}
		w.lineNum++

		// Write CheckDetails
		if err := w.writeCheckDetail(b); err != nil {
			return err
		}

		if _, err := w.w.WriteString(b.GetControl().String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}

// writeCheckDetail writes a CheckDetail to a Bundle
func (w *Writer) writeCheckDetail(b Bundle) error {

	for _, cd := range b.GetChecks() {
		if _, err := w.w.WriteString(cd.String() + "\n"); err != nil {
			return err
		}
		w.lineNum++

		// Write CheckDetailsAddendum (A, B, C)
		if err := w.writeCheckDetailAddendum(cd); err != nil {
			return err
		}

	}
	return nil
}

// writeCheckDetailAddendum writes a CheckDetailAddendum (A, B, C) to a CheckDetail
func (w *Writer) writeCheckDetailAddendum(cd *CheckDetail) error {
	for _, cdAddendumA := range cd.GetCheckDetailAddendumA() {
		if _, err := w.w.WriteString(cdAddendumA.String() + "\n"); err != nil {
			return err
		}
	}
	w.lineNum++

	for _, cdAddendumB := range cd.GetCheckDetailAddendumB() {
		if _, err := w.w.WriteString(cdAddendumB.String() + "\n"); err != nil {
			return err
		}
	}
	w.lineNum++

	for _, cdAddendumC := range cd.GetCheckDetailAddendumC() {
		if _, err := w.w.WriteString(cdAddendumC.String() + "\n"); err != nil {
			return err
		}
	}
	w.lineNum++

	return nil
}

//Items ImageView*
//ReturnBundles
//Items - ReturnDetail, ReturnAddendum*, and ImageView*
