// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/gdamore/encoding"
)

// Writer writes an ImageCashLetter/X9 File to an encoded format.
//
// Callers should use NewWriter to create a new instance and apply WriterOptions
// as needed to properly encode files for their usecase.
type Writer struct {
	w                  *bufio.Writer
	lineNum            int // current line being written
	VariableLineLength bool
	EbcdicEncoding     bool
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer, opts ...WriterOption) *Writer {
	writer := &Writer{
		w: bufio.NewWriter(w),
	}
	for _, opt := range opts {
		opt(writer)
	}
	return writer
}

// WriterOption allows Writer to be configured to write in different formats
type WriterOption func(w *Writer)

// WriteVariableLineLengthOption allows Writer to write control bytes ahead of record to describe how long the line is
// Follows DSTU microformat as defined https://www.frbservices.org/assets/financial-services/check/setup/frb-x937-standards-reference.pdf
func WriteVariableLineLengthOption() WriterOption {
	return func(w *Writer) {
		w.VariableLineLength = true
	}
}

// WriteEbcdicEncodingOption allows Writer to write file in EBCDIC
// Follows DSTU microformat as defined https://www.frbservices.org/assets/financial-services/check/setup/frb-x937-standards-reference.pdf
func WriteEbcdicEncodingOption() WriterOption {
	return func(w *Writer) {
		w.EbcdicEncoding = true
	}
}

func (w *Writer) writeLine(record FileRecord) error {
	line := record.String()
	lineLength := len(line)

	if w.VariableLineLength {
		ctrl := make([]byte, 4)
		binary.BigEndian.PutUint32(ctrl, uint32(lineLength))
		if _, err := w.w.Write(ctrl); err != nil {
			return err
		}
	}

	if w.EbcdicEncoding {
		if ivData, ok := record.(*ImageViewData); ok {
			// need to encode everything other than binary image into EBCDIC
			encoded, err := encoding.EBCDIC.NewEncoder().String(ivData.toString(false))
			if err != nil {
				return err
			}
			if _, err := w.w.WriteString(encoded); err != nil {
				return err
			}
			if _, err := w.w.Write(ivData.ImageData); err != nil {
				return err
			}
		} else {
			// no binary data in record, encode entire line
			encoded, err := encoding.EBCDIC.NewEncoder().String(line)
			if err != nil {
				return err
			}
			if _, err := w.w.WriteString(encoded); err != nil {
				return err
			}
		}
	} else {
		// ASCII encoding by default
		if _, err := w.w.WriteString(line); err != nil {
			return err
		}
	}

	if !w.VariableLineLength {
		if _, err := w.w.WriteString("\n"); err != nil {
			return err
		}
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
	if err := w.writeLine(&file.Header); err != nil {
		return err
	}
	if err := w.writeCashLetter(file); err != nil {
		return err
	}
	if err := w.writeLine(&file.Control); err != nil {
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
		if err := w.writeLine(cl.GetHeader()); err != nil {
			return err
		}
		for _, ci := range cl.GetCreditItems() {
			if err := w.writeLine(ci); err != nil {
				return err
			}
		}
		for _, credit := range cl.GetCredits() {
			if err := w.writeLine(credit); err != nil {
				return err
			}
		}
		if err := w.writeBundle(cl); err != nil {
			return err
		}
		for _, rns := range cl.GetRoutingNumberSummary() {
			if err := w.writeLine(rns); err != nil {
				return err
			}
		}
		if err := w.writeLine(cl.GetControl()); err != nil {
			return err
		}
	}
	return nil
}

// writeBundle writes a Bundle to a CashLetter
func (w *Writer) writeBundle(cl CashLetter) error {
	for _, b := range cl.GetBundles() {
		if err := w.writeLine(b.GetHeader()); err != nil {
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
		if err := w.writeLine(b.GetControl()); err != nil {
			return err
		}
	}
	return nil
}

// writeCheckDetail writes a CheckDetail to a Bundle
func (w *Writer) writeCheckDetail(b *Bundle) error {
	for _, cd := range b.GetChecks() {
		if err := w.writeLine(cd); err != nil {
			return err
		}
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
	addendumA := cd.GetCheckDetailAddendumA()
	for i := range addendumA {
		if err := w.writeLine(&addendumA[i]); err != nil {
			return err
		}
	}

	addendumB := cd.GetCheckDetailAddendumB()
	for i := range addendumB {
		if err := w.writeLine(&addendumB[i]); err != nil {
			return err
		}
	}

	addendumC := cd.GetCheckDetailAddendumC()
	for i := range addendumC {
		if err := w.writeLine(&addendumC[i]); err != nil {
			return err
		}
	}
	return nil
}

// writeReturnDetail writes a ReturnDetail to a ReturnBundle
func (w *Writer) writeReturnDetail(b *Bundle) error {
	for _, rd := range b.GetReturns() {
		if err := w.writeLine(rd); err != nil {
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
	addendumA := rd.GetReturnDetailAddendumA()
	for i := range addendumA {
		if err := w.writeLine(&addendumA[i]); err != nil {
			return err
		}
	}

	addendumB := rd.GetReturnDetailAddendumB()
	for i := range addendumB {
		if err := w.writeLine(&addendumB[i]); err != nil {
			return err
		}
	}

	addendumC := rd.GetReturnDetailAddendumC()
	for i := range addendumC {
		if err := w.writeLine(&addendumC[i]); err != nil {
			return err
		}
	}

	addendumD := rd.GetReturnDetailAddendumD()
	for i := range addendumD {
		if err := w.writeLine(&addendumD[i]); err != nil {
			return err
		}
	}
	return nil
}

// writeCheckImageView writes ImageViews (Detail, Data, Analysis) to a CheckDetail
func (w *Writer) writeCheckImageView(cd *CheckDetail) error {
	return w.writeImageView(cd.GetImageViewDetail(), cd.GetImageViewData(), cd.GetImageViewAnalysis())
}

// writeReturnImageView writes ImageViews (Detail, Data, Analysis) to a ReturnDetail
func (w *Writer) writeReturnImageView(rd *ReturnDetail) error {
	return w.writeImageView(rd.GetImageViewDetail(), rd.GetImageViewData(), rd.GetImageViewAnalysis())
}

// writeImageView writes ImageViews (Detail, Data, Analysis) in the order expected by the FRB
func (w *Writer) writeImageView(ivDetail []ImageViewDetail, ivData []ImageViewData, ivAnalysis []ImageViewAnalysis) error {

	// TODO: Add validator to ensure that each imageViewDetail has a corresponding imageViewData and imageViewAnalysis
	// for now enforce that all images have data and analysis or no images have data and analysis

	if len(ivData) > 0 && len(ivData) != len(ivDetail) {
		// should be same number of imageViewData as imageViewDetail
		msg := fmt.Sprintf(msgBundleImageDetailCount, len(ivData))
		return &BundleError{FieldName: "ImageViewData", Msg: msg}
	}

	if len(ivAnalysis) > 0 && len(ivAnalysis) != len(ivDetail) {
		// should same number of imageViewAnalysis and imageViewDetail
		msg := fmt.Sprintf(msgBundleImageDetailCount, len(ivAnalysis))
		return &BundleError{FieldName: "ImageViewAnalysis", Msg: msg}
	}

	// FRB asks that imageViewDetail should immediately be followed by its corresponding data and analysis
	for i := range ivDetail {
		if err := w.writeLine(&ivDetail[i]); err != nil {
			return err
		}
		if len(ivData) > 0 && len(ivData) >= i-1 {
			ivData := ivData[i]
			if err := w.writeLine(&ivData); err != nil {
				return err
			}
		}
		if len(ivAnalysis) > 0 && len(ivAnalysis) >= i-1 {
			ivAnalysis := ivAnalysis[i]
			if err := w.writeLine(&ivAnalysis); err != nil {
				return err
			}
		}
	}

	return nil
}
