// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"fmt"
)

// BundleError is an Error that describes bundle validation issues
type BundleError struct {
	BundleSequenceNumber string
	FieldName            string
	Msg                  string
}

func (e *BundleError) Error() string {
	return fmt.Sprintf("BundleNumber %s %s %s", e.BundleSequenceNumber, e.FieldName, e.Msg)
}

// Addendum Counts
const (
	CheckDetailAddendumACount  = 9
	CheckDetailAddendumBCount  = 1
	CheckDetailAddendumCCount  = 99
	ReturnDetailAddendumACount = 9
	ReturnDetailAddendumBCount = 1
	ReturnDetailAddendumCCount = 1
	ReturnDetailAddendumDCount = 99
)

// Errors specific to parsing a Bundle
var (
	msgBundleEntries          = "must have Check Detail or Return Detail to be built"
	msgBundleAddendum         = "%v found is greater than maximum of %v"
	msgBundleAddendumCount    = "%v does not match Addenda Records"
	msgBundleImageDetailCount = "does not match Image View Detail count of %v"
)

// Bundle contains forward items (checks)
type Bundle struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// BundleHeader is a Bundle Header Record
	BundleHeader *BundleHeader `json:"bundleHeader,omitempty"`
	// Checks are Check Items: Check Detail Records, Check Detail Addendum Records, and Image Views
	Checks []*CheckDetail `json:"checks,omitempty"`
	// Returns are Return Items: Return Detail Records, Return Detail Addendum Records, and Image Views
	Returns []*ReturnDetail `json:"returns,omitempty"`
	// BundleControl is a Bundle Control Record
	BundleControl *BundleControl `json:"bundleControl,omitempty"`
}

// NewBundle takes a BundleHeader and returns a Bundle
func NewBundle(bh *BundleHeader) *Bundle {
	b := new(Bundle)
	b.SetControl(NewBundleControl())
	b.SetHeader(bh)
	return b
}

func (b *Bundle) setRecordType() {
	if b == nil {
		return
	}
	b.BundleHeader.setRecordType()
	for i := range b.Checks {
		b.Checks[i].setRecordType()
	}
	for i := range b.Returns {
		b.Returns[i].setRecordType()
	}
	b.BundleControl.setRecordType()
}

// Validate performs imagecashletter validations and format rule checks and returns an error if not Validated
func (b *Bundle) Validate() error {
	if (len(b.Checks) <= 0) && (len(b.Returns) <= 0) {
		seqNumber := ""
		if b.BundleHeader != nil {
			seqNumber = b.BundleHeader.BundleSequenceNumber
		}
		return &BundleError{BundleSequenceNumber: seqNumber, FieldName: "entries", Msg: msgBundleEntries}
	}

	if len(b.Checks) > 0 {
		if err := b.checkDetailAddendumCount(); err != nil {
			return err
		}
	} else {
		if err := b.returnDetailAddendumCount(); err != nil {
			return err
		}
	}
	return nil
}

// build creates a valid Bundle by building  BundleControl. An error is returned if
// the bundle being built has invalid records.
func (b *Bundle) build() error {
	if b == nil {
		return nil
	}

	// Requires a valid BundleHeader
	if err := b.BundleHeader.Validate(); err != nil {
		return err
	}
	if (len(b.Checks) <= 0) && (len(b.Returns) <= 0) {
		seqNumber := ""
		if b.BundleHeader != nil {
			seqNumber = b.BundleHeader.BundleSequenceNumber
		}
		return &BundleError{BundleSequenceNumber: seqNumber, FieldName: "entries", Msg: msgBundleEntries}
	}

	itemCount := 0
	bundleTotalAmount := 0
	micrValidTotalAmount := 0
	bundleImagesCount := 0
	// The current Implementation doe snot support CreditItems as part of a bundle so BundleControl.CreditIndicator = 0
	creditIndicator := 0

	// Forward Items
	for _, cd := range b.Checks {

		// Validate CheckDetailAddendum* and ImageView*
		if err := b.ValidateForwardItems(cd); err != nil {
			return err
		}

		itemCount = itemCount + 1
		bundleTotalAmount = bundleTotalAmount + cd.ItemAmount
		if cd.MICRValidIndicator == 1 {
			micrValidTotalAmount = micrValidTotalAmount + cd.ItemAmount
		}

		bundleImagesCount = bundleImagesCount + len(cd.ImageViewDetail)
	}

	// Return Items
	for _, rd := range b.Returns {

		// Validate ReturnDetailAddendum* and ImageView*
		if err := b.ValidateReturnItems(rd); err != nil {
			return err
		}
		itemCount = itemCount + 1
		bundleTotalAmount = bundleTotalAmount + rd.ItemAmount
		bundleImagesCount = bundleImagesCount + len(rd.ImageViewDetail)
	}

	// build a BundleControl record
	bc := NewBundleControl()
	bc.BundleItemsCount = itemCount
	bc.BundleTotalAmount = bundleTotalAmount
	bc.MICRValidTotalAmount = micrValidTotalAmount
	bc.BundleImagesCount = bundleImagesCount
	bc.CreditTotalIndicator = creditIndicator
	b.BundleControl = bc
	return nil
}

// SetHeader appends an BundleHeader to the Bundle
func (b *Bundle) SetHeader(bundleHeader *BundleHeader) {
	b.BundleHeader = bundleHeader
}

// GetHeader returns the current Bundle header
func (b *Bundle) GetHeader() *BundleHeader {
	return b.BundleHeader
}

// SetControl appends an BundleControl to the Bundle
func (b *Bundle) SetControl(bundleControl *BundleControl) {
	b.BundleControl = bundleControl
}

// GetControl returns the current Bundle Control
func (b *Bundle) GetControl() *BundleControl {
	return b.BundleControl
}

// AddCheckDetail appends a CheckDetail to the Bundle
func (b *Bundle) AddCheckDetail(cd *CheckDetail) {
	b.Checks = append(b.Checks, cd)
}

// GetChecks returns a slice of check details for the Bundle
func (b *Bundle) GetChecks() []*CheckDetail {
	if b == nil {
		return nil
	}
	return b.Checks
}

// AddReturnDetail appends a ReturnDetail to the Bundle
func (b *Bundle) AddReturnDetail(rd *ReturnDetail) {
	b.Returns = append(b.Returns, rd)
}

// GetReturns returns a slice of return details for the Bundle
func (b *Bundle) GetReturns() []*ReturnDetail {
	if b == nil {
		return nil
	}
	return b.Returns
}

// ValidateForwardItems calls Validate function for check items
func (b *Bundle) ValidateForwardItems(cd *CheckDetail) error {
	// Validate items
	for _, addendumA := range cd.CheckDetailAddendumA {
		if err := addendumA.Validate(); err != nil {
			return err
		}
	}
	for _, addendumB := range cd.CheckDetailAddendumB {
		if err := addendumB.Validate(); err != nil {
			return err
		}
	}
	for _, addendumC := range cd.CheckDetailAddendumC {
		if err := addendumC.Validate(); err != nil {
			return err
		}
	}
	for _, ivDetail := range cd.ImageViewDetail {
		if err := ivDetail.Validate(); err != nil {
			return err
		}
	}
	for _, ivData := range cd.ImageViewData {
		if err := ivData.Validate(); err != nil {
			return err
		}
	}
	for _, ivAnalysis := range cd.ImageViewAnalysis {
		if err := ivAnalysis.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// ValidateReturnItems calls Validate function for return items
func (b *Bundle) ValidateReturnItems(rd *ReturnDetail) error {
	// Validate items
	for _, addendumA := range rd.ReturnDetailAddendumA {
		if err := addendumA.Validate(); err != nil {
			return err
		}
	}
	for _, addendumB := range rd.ReturnDetailAddendumB {
		if err := addendumB.Validate(); err != nil {
			return err
		}
	}
	for _, addendumC := range rd.ReturnDetailAddendumC {
		if err := addendumC.Validate(); err != nil {
			return err
		}
	}
	for _, addendumD := range rd.ReturnDetailAddendumD {
		if err := addendumD.Validate(); err != nil {
			return err
		}
	}
	for _, ivDetail := range rd.ImageViewDetail {
		if err := ivDetail.Validate(); err != nil {
			return err
		}
	}
	for _, ivData := range rd.ImageViewDetail {
		if err := ivData.Validate(); err != nil {
			return err
		}
	}
	for _, ivAnalysis := range rd.ImageViewDetail {
		if err := ivAnalysis.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// checkDetailAddendumCount validates CheckDetail AddendumCount
func (b *Bundle) checkDetailAddendumCount() error {
	bundleSequenceNumber := "-"
	if b.BundleHeader != nil {
		bundleSequenceNumber = b.BundleHeader.BundleSequenceNumber
	}

	// Check Items
	for _, cd := range b.Checks {
		if cd.AddendumCount != len(cd.CheckDetailAddendumA)+len(cd.CheckDetailAddendumB)+len(cd.CheckDetailAddendumC) {
			msg := fmt.Sprintf(msgBundleAddendumCount, cd.AddendumCount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "AddendumCount", Msg: msg}
		}
		if len(cd.CheckDetailAddendumA) > CheckDetailAddendumACount {
			msg := fmt.Sprintf(msgBundleAddendum, len(cd.CheckDetailAddendumA), CheckDetailAddendumACount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "CheckDetailAddendumA", Msg: msg}
		}
		if len(cd.CheckDetailAddendumB) > CheckDetailAddendumBCount {
			msg := fmt.Sprintf(msgBundleAddendum, len(cd.CheckDetailAddendumB), CheckDetailAddendumBCount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "CheckDetailAddendumB", Msg: msg}
		}
		if len(cd.CheckDetailAddendumC) > CheckDetailAddendumCCount {
			msg := fmt.Sprintf(msgBundleAddendum, len(cd.CheckDetailAddendumC), CheckDetailAddendumCCount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "CheckDetailAddendumC", Msg: msg}
		}

	}
	return nil
}

// returnDetailAddendumCount validates ReturnDetail AddendumCount
func (b *Bundle) returnDetailAddendumCount() error {
	bundleSequenceNumber := "-"
	if b.BundleHeader != nil {
		bundleSequenceNumber = b.BundleHeader.BundleSequenceNumber
	}

	for _, rd := range b.Returns {
		if rd.AddendumCount != len(rd.ReturnDetailAddendumA)+len(rd.ReturnDetailAddendumB)+len(rd.ReturnDetailAddendumC)+len(rd.ReturnDetailAddendumD) {
			msg := fmt.Sprintf(msgBundleAddendumCount, rd.AddendumCount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "AddendumCount", Msg: msg}
		}
		if len(rd.ReturnDetailAddendumA) > ReturnDetailAddendumACount {
			msg := fmt.Sprintf(msgBundleAddendum, len(rd.ReturnDetailAddendumA), ReturnDetailAddendumACount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "ReturnDetailAddendumA", Msg: msg}
		}
		if len(rd.ReturnDetailAddendumB) > ReturnDetailAddendumBCount {
			msg := fmt.Sprintf(msgBundleAddendum, len(rd.ReturnDetailAddendumB), ReturnDetailAddendumBCount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "ReturnDetailAddendumB", Msg: msg}
		}
		if len(rd.ReturnDetailAddendumC) > ReturnDetailAddendumCCount {
			msg := fmt.Sprintf(msgBundleAddendum, len(rd.ReturnDetailAddendumC), ReturnDetailAddendumCCount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "ReturnDetailAddendumC", Msg: msg}
		}
		if len(rd.ReturnDetailAddendumD) > ReturnDetailAddendumDCount {
			msg := fmt.Sprintf(msgBundleAddendum, len(rd.ReturnDetailAddendumD), ReturnDetailAddendumDCount)
			return &BundleError{BundleSequenceNumber: bundleSequenceNumber, FieldName: "ReturnDetailAddendumD", Msg: msg}
		}
	}
	return nil
}
