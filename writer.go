// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"bufio"
	"io"
	"strings"
)

// A Writer writes an x9.file to an encoded file.
//
// As returned by NewWriter, a Writer writes x9file structs into
// x9 formatted files.

// ToDo;  Review/Test the looping/writing of Bundles, ReturnBundles, CheckDetail, ReturnDetail, Addendum, and ImageView

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
		if err := w.writeReturnBundle(cl); err != nil {
			return err
		}

		if _, err := w.w.WriteString(cl.GetControl().String() + "\n"); err != nil {
			return err
		}
		w.lineNum++
	}
	return nil
}

// writeBundle writes a Bundle to a CashLetter
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

// writeReturnBundle writes a ReturnBundle to a CashLetter
func (w *Writer) writeReturnBundle(cl CashLetter) error {

	for _, rb := range cl.GetReturnBundles() {
		if _, err := w.w.WriteString(rb.GetHeader().String() + "\n"); err != nil {
			return err
		}
		w.lineNum++

		// Write ReturnDetails
		if err := w.writeReturnDetail(rb); err != nil {
			return err
		}

		if _, err := w.w.WriteString(rb.GetControl().String() + "\n"); err != nil {
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
		if err := w.writeCheckImageView(cd); err != nil {
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

// writeCheckImageView writes ImageViews (Detail, Data, Analysis) to a CheckDetail
func (w *Writer) writeCheckImageView(cd *CheckDetail) error {
	for _, ivDetail := range cd.GetCheckDetailImageViewDetail() {
		if _, err := w.w.WriteString(ivDetail.String() + "\n"); err != nil {
			return err
		}
	}
	for _, ivData := range cd.GetCheckDetailImageViewData() {
		if _, err := w.w.WriteString(ivData.String() + "\n"); err != nil {
			return err
		}
	}
	for _, ivAnalysis := range cd.GetCheckDetailImageViewAnalysis() {
		if _, err := w.w.WriteString(ivAnalysis.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}

// writeReturnDetail writes a ReturnDetail to a ReturnBundle
func (w *Writer) writeReturnDetail(b ReturnBundle) error {
	for _, rd := range b.GetReturns() {
		if _, err := w.w.WriteString(rd.String() + "\n"); err != nil {
			return err
		}
		w.lineNum++
		// Write ReturnDetailAddendum (A, B, C, D)
		if err := w.writeReturnDetailAddendum(rd); err != nil {
			return err
		}
		if err := w.writeReturnImageView(rd); err != nil {
			return err
		}
	}
	return nil
}

// writeReturnDetailAddendum writes a ReturnDetailAddendum (A, B, C, D) to a ReturnDetail
func (w *Writer) writeReturnDetailAddendum(rd *ReturnDetail) error {
	for _, rdAddendumA := range rd.GetReturnDetailAddendumA() {
		if _, err := w.w.WriteString(rdAddendumA.String() + "\n"); err != nil {
			return err
		}
	}
	w.lineNum++
	for _, rdAddendumB := range rd.GetReturnDetailAddendumB() {
		if _, err := w.w.WriteString(rdAddendumB.String() + "\n"); err != nil {
			return err
		}
	}
	w.lineNum++
	for _, rdAddendumC := range rd.GetReturnDetailAddendumC() {
		if _, err := w.w.WriteString(rdAddendumC.String() + "\n"); err != nil {
			return err
		}
	}
	w.lineNum++
	for _, rdAddendumD := range rd.GetReturnDetailAddendumD() {
		if _, err := w.w.WriteString(rdAddendumD.String() + "\n"); err != nil {
			return err
		}
	}
	w.lineNum++

	return nil
}

// writeReturnImageView writes ImageViews (Detail, Data, Analysis) to a ReturnDetail
func (w *Writer) writeReturnImageView(rd *ReturnDetail) error {
	for _, ivDetail := range rd.GetReturnDetailImageViewDetail() {
		if _, err := w.w.WriteString(ivDetail.String() + "\n"); err != nil {
			return err
		}
	}
	for _, ivData := range rd.GetReturnDetailImageViewData() {
		if _, err := w.w.WriteString(ivData.String() + "\n"); err != nil {
			return err
		}
	}
	for _, ivAnalysis := range rd.GetReturnDetailImageViewAnalysis() {
		if _, err := w.w.WriteString(ivAnalysis.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}