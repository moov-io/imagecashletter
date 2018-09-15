// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// Errors specific to a Bundle

// Bundle contains forward items (checks)
type Bundle struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// BundleHeader is a Bundle Header Record
	BundleHeader *BundleHeader `json:"bundleHeader,omitempty"`
	// Items are Items: Check Detail Records, Check Detail Addendum Records, and Image Views
	Checks []*CheckDetail `json:"items,omitempty"`
	// BundleControl is a Bundle Control Record
	BundleControl *BundleControl `json:"bundleControl,omitempty"`
	// Converters is composed for x9 to GoLang Converters
	converters
}

// NewBundle takes a BundleHeader and returns a Bundle
// ToDo:  Follow up on returning a pointer when implementing tests and examples
func NewBundle(bh *BundleHeader) Bundle {
	b := Bundle{}
	b.SetControl(NewBundleControl())
	b.SetHeader(bh)
	return b
}

// Validate performs X9  validations and format rule checks and returns an error if not Validated
func (b *Bundle) Validate() error {
	return nil
}

// AddCheckDetail appends a CheckDetail to the Bundle
func (b *Bundle) AddCheckDetail(check *CheckDetail) {
	b.Checks = append(b.Checks, check)
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

// GetChecks returns a slice of check details for the bundle
func (b *Bundle) GetChecks() []*CheckDetail {
	return b.Checks
}
