// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// Bundle contains forward items (checks)
type Bundle struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// BundleHeader is an ICl BundleHeader
	BundleHeader *BundleHeader `json:"bundleHeader,omitempty"`
	// Items are ICL CheckDetail items
	Items []*CheckDetail `json:"items,omitempty"`
	// BundleControl is an ICl BundleControl
	BundleControl *BundleControl `json:"bundleControl,omitempty"`
	// Converters is composed for x9 to GoLang Converters
	converters
}

// NewBundle takes a BundleHeader and returns a Bundle
func NewBundle(bh *BundleHeader) Bundle {
	bundle := Bundle{}
	bundle.SetControl(NewBundleControl())
	bundle.SetHeader(bh)
	return bundle
}

// SetHeader appends an BundleHeader to the Bundle
func (bundle *Bundle) SetHeader(bundleHeader *BundleHeader) {
	bundle.BundleHeader = bundleHeader
}

// SetControl appends an BundleControl to the Bundle
func (bundle *Bundle) SetControl(bundleControl *BundleControl) {
	bundle.BundleControl = bundleControl
}
