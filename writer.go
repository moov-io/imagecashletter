// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"bufio"
	"encoding/binary"
	"io"
)

// A Writer writes an imagecashletter.file to an encoded file.
//
// As returned by NewWriter, a Writer writes imagecashletterfile structs into
// imagecashletter formatted files.

// Writer struct
type Writer struct {
	w       *bufio.Writer
	lineNum int //current line being written
	// format in which to write file as
	format Format
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(w),
	}
}

// SetFormat of imagecashletter file
func (w *Writer) SetFormat(format Format) {
	switch format {
	case Discover:
		w.format = format
	case DSTU:
		w.format = format
	}
}

func (w *Writer) writeLine(line string) error {
	if w.format == DSTU {
		// convert line length into bytes
		// As per spec, write as control bytes before line
		lineLength := len(line)
		ctrl := make([]byte, 4)
		binary.BigEndian.PutUint32(ctrl, uint32(lineLength))
		if _, err := w.w.Write(ctrl); err != nil {
			return err
		}
	}

	w.w.WriteString(line)

	if w.format == Discover {
		w.w.WriteString("\n")
	}

	w.lineNum++

	return nil
}

// Writer writes a single imagecashletter.file record to w
func (w *Writer) Write(file *File) error {
	if file == nil {
		return ErrNilFile
	}
	if err := file.Validate(); err != nil {
		return err
	}
	w.lineNum = 0
	// Iterate over all records in the file
	if err := w.writeLine(file.Header.String()); err != nil {
		return err
	}
	if err := w.writeCashLetter(file); err != nil {
		return err
	}
	if err := w.writeLine(file.Control.String()); err != nil {
		return err
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
	for _, cl := range file.CashLetters {
		if err := w.writeLine(cl.GetHeader().String()); err != nil {
			return err
		}
		for _, ci := range cl.GetCreditItems() {
			if err := w.writeLine(ci.String()); err != nil {
				return err
			}
		}
		if err := w.writeBundle(cl); err != nil {
			return err
		}
		for _, rns := range cl.GetRoutingNumberSummary() {
			if err := w.writeLine(rns.String()); err != nil {
				return err
			}
		}
		if err := w.writeLine(cl.GetControl().String()); err != nil {
			return err
		}
	}
	return nil
}

// writeBundle writes a Bundle to a CashLetter
func (w *Writer) writeBundle(cl CashLetter) error {
	for _, b := range cl.GetBundles() {
		if err := w.writeLine(b.GetHeader().String()); err != nil {
			return err
		}
		if len(b.Checks) > 0 {
			if err := w.writeCheckDetail(b); err != nil {
				return err
			}
		}
		if len(b.Returns) > 0 {
			if err := w.writeReturnDetail(b); err != nil {
				return err
			}
		}
		if err := w.writeLine(b.GetControl().String()); err != nil {
			return err
		}
	}
	return nil
}

// writeCheckDetail writes a CheckDetail to a Bundle
func (w *Writer) writeCheckDetail(b *Bundle) error {
	for _, cd := range b.GetChecks() {
		if err := w.writeLine(cd.String()); err != nil {
			return err
		}
		// Write CheckDetailsAddendum (A, B, C)
		if err := w.writeCheckDetailAddendum(cd); err != nil {
			return err
		}
		//w.lineNum++
		if err := w.writeCheckImageView(cd); err != nil {
			return err
		}
	}
	return nil
}

// writeCheckDetailAddendum writes a CheckDetailAddendum (A, B, C) to a CheckDetail
func (w *Writer) writeCheckDetailAddendum(cd *CheckDetail) error {
	for _, cdAddendumA := range cd.GetCheckDetailAddendumA() {
		if err := w.writeLine(cdAddendumA.String()); err != nil {
			return err
		}
	}
	for _, cdAddendumB := range cd.GetCheckDetailAddendumB() {
		if err := w.writeLine(cdAddendumB.String()); err != nil {
			return err
		}
	}
	for _, cdAddendumC := range cd.GetCheckDetailAddendumC() {
		if err := w.writeLine(cdAddendumC.String()); err != nil {
			return err
		}
	}
	return nil
}

// writeCheckImageView writes ImageViews (Detail, Data, Analysis) to a CheckDetail
func (w *Writer) writeCheckImageView(cd *CheckDetail) error {
	for _, ivDetail := range cd.GetImageViewDetail() {
		if err := w.writeLine(ivDetail.String()); err != nil {
			return err
		}
	}
	for _, ivData := range cd.GetImageViewData() {
		if err := w.writeLine(ivData.String()); err != nil {
			return err
		}
	}
	for _, ivAnalysis := range cd.GetImageViewAnalysis() {
		if err := w.writeLine(ivAnalysis.String()); err != nil {
			return err
		}
	}
	return nil
}

// writeReturnDetail writes a ReturnDetail to a ReturnBundle
func (w *Writer) writeReturnDetail(b *Bundle) error {
	for _, rd := range b.GetReturns() {
		if err := w.writeLine(rd.String()); err != nil {
			return err
		}
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
		if err := w.writeLine(rdAddendumA.String()); err != nil {
			return err
		}
	}
	for _, rdAddendumB := range rd.GetReturnDetailAddendumB() {
		if err := w.writeLine(rdAddendumB.String()); err != nil {
			return err
		}
	}
	for _, rdAddendumC := range rd.GetReturnDetailAddendumC() {
		if err := w.writeLine(rdAddendumC.String()); err != nil {
			return err
		}
	}
	for _, rdAddendumD := range rd.GetReturnDetailAddendumD() {
		if err := w.writeLine(rdAddendumD.String()); err != nil {
			return err
		}
	}
	return nil
}

// writeReturnImageView writes ImageViews (Detail, Data, Analysis) to a ReturnDetail
func (w *Writer) writeReturnImageView(rd *ReturnDetail) error {
	for _, ivDetail := range rd.GetImageViewDetail() {
		if err := w.writeLine(ivDetail.String()); err != nil {
			return err
		}
	}
	for _, ivData := range rd.GetImageViewData() {
		if err := w.writeLine(ivData.String()); err != nil {
			return err
		}
	}
	for _, ivAnalysis := range rd.GetImageViewAnalysis() {
		if err := w.writeLine(ivAnalysis.String()); err != nil {
			return err
		}
	}
	return nil
}
