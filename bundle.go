// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

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
	return fmt.Sprintf("BatchNumber %s %s %s", e.BundleSequenceNumber, e.FieldName, e.Msg)
}

// Errors specific to parsing a Batch container
var (
	msgBundleEntries = "must have Check Detail or Return Detail to be built"
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
// ToDo:  Follow up on returning a pointer when implementing tests and examples
func NewBundle(bh *BundleHeader) *Bundle {
	b := new(Bundle)
	b.SetControl(NewBundleControl())
	b.SetHeader(bh)
	return b
}

// Validate performs X9  validations and format rule checks and returns an error if not Validated
func (b *Bundle) Validate() error {
	if (len(b.Checks) <= 0) && (len(b.Returns) <= 0) {
		return &BundleError{BundleSequenceNumber: b.BundleHeader.BundleSequenceNumber, FieldName: "entries", Msg: msgBundleEntries}
	}

	return nil
}

// ToDo: Add verify

// build creates valid bundle by building sequence numbers and BundleControl. An error is returned if
// the bundle being built has invalid records.
func (b *Bundle) build() error {
	// Requires a valid BundleHeader
	if err := b.BundleHeader.Validate(); err != nil {
		return err
	}
	if (len(b.Checks) <= 0) && (len(b.Returns) <= 0) {
		return &BundleError{BundleSequenceNumber: b.BundleHeader.BundleSequenceNumber, FieldName: "entries", Msg: msgBundleEntries}
	}

	//
	itemCount := 0
	bundleTotalAmount := 0
	micrValidTotalAmount := 0
	bundleImagesCount := 0

	// ToDo: Sequences

	// Forward Items
	for _, cd := range b.Checks {
		itemCount = itemCount + 1
		itemCount = itemCount + len(cd.CheckDetailAddendumA) + len(cd.CheckDetailAddendumB) + len(cd.CheckDetailAddendumC)
		itemCount = itemCount + len(cd.ImageViewDetail) + len(cd.ImageViewData) + len(cd.ImageViewAnalysis)
		bundleTotalAmount = bundleTotalAmount + cd.ItemAmount
		if cd.MICRValidIndicator == 1 {
			micrValidTotalAmount = micrValidTotalAmount + cd.ItemAmount
		}
		// Validate CheckDetailAddendum* and ImageView*
		if err := b.ValidateForwardItems(cd); err != nil {
			return err
		}
	}

	// Return Items
	for _, rd := range b.Returns {
		itemCount = itemCount + 1
		itemCount = itemCount + len(rd.ReturnDetailAddendumA) + len(rd.ReturnDetailAddendumB) + len(rd.ReturnDetailAddendumC)
		itemCount = itemCount + len(rd.ImageViewDetail) + len(rd.ImageViewData) + len(rd.ImageViewAnalysis)
		bundleTotalAmount = bundleTotalAmount + rd.ItemAmount
		// ToDo: micrValidTotalAmount for returns?
		// Validate ReturnDetailAddendum* and ImageView*
		if err := b.ValidateReturnItems(rd); err != nil {
			return err
		}

	}

	// build a BundleControl record
	bc := NewBundleControl()
	bc.BundleItemsCount = itemCount
	bc.BundleTotalAmount = bundleTotalAmount
	bc.MICRValidTotalAmount = micrValidTotalAmount
	bc.BundleImagesCount = bundleImagesCount
	// ToDo:  Add Credit Functionality?
	bc.CreditTotalIndicator = 0
	b.BundleControl = bc
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (b *Bundle) Create() error {
	// generates sequence numbers and batch control
	if err := b.build(); err != nil {
		return err
	}

	return b.Validate()
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
	return b.Checks
}

// AddReturnDetail appends a ReturnDetail to the Bundle
func (b *Bundle) AddReturnDetail(rd *ReturnDetail) {
	b.Returns = append(b.Returns, rd)
}

// GetReturns returns a slice of return details for the Bundle
func (b *Bundle) GetReturns() []*ReturnDetail {
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
	for _, ivData := range cd.ImageViewDetail {
		if err := ivData.Validate(); err != nil {
			return err
		}
	}
	for _, ivAnalysis := range cd.ImageViewDetail {
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
