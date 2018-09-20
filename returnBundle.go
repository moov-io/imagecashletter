// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// ReturnBundle contains forward items (checks)
type ReturnBundle struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// BundleHeader is a Bundle Header Record for returns
	BundleHeader *BundleHeader `json:"bundleHeader,omitempty"`
	// Items are Items: Check Detail Records, Check Detail Addendum Records, and Image Views
	Returns []*ReturnDetail `json:"items,omitempty"`
	// BundleControl is a Bundle Control Record
	BundleControl *BundleControl `json:"bundleControl,omitempty"`
}

// NewReturnBundle takes a BundleHeader and returns a ReturnBundle
func NewReturnBundle(bh *BundleHeader) *ReturnBundle {
	rb := new(ReturnBundle)
	rb.SetControl(NewBundleControl())
	rb.SetHeader(bh)
	return rb
}

// Validate performs X9  validations and format rule checks and returns an error if not Validated
func (rb *ReturnBundle) Validate() error {
	return nil
}

// AddReturnDetail appends a ReturnDetail to the ReturnBundle
func (rb *ReturnBundle) AddReturnDetail(rd *ReturnDetail) {
	rb.Returns = append(rb.Returns, rd)
}

// SetHeader appends an BundleHeader to the ReturnBundle
func (rb *ReturnBundle) SetHeader(bundleHeader *BundleHeader) {
	rb.BundleHeader = bundleHeader
}

// GetHeader returns the current ReturnBundle header
func (rb *ReturnBundle) GetHeader() *BundleHeader {
	return rb.BundleHeader
}

// SetControl appends an BundleControl to the ReturnBundle
func (rb *ReturnBundle) SetControl(bundleControl *BundleControl) {
	rb.BundleControl = bundleControl
}

// GetControl returns the current ReturnBundle Control
func (rb *ReturnBundle) GetControl() *BundleControl {
	return rb.BundleControl
}

// GetReturns returns a slice of check details for the ReturnBundle
func (rb *ReturnBundle) GetReturns() []*ReturnDetail {
	return rb.Returns
}
