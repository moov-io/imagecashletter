// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// CashLetter contains CashLetterHeader, CashLetterControl and Bundle records.
type CashLetter struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// CashLetterHeader is a Cash Letter Header Record
	CashLetterHeader *CashLetterHeader `json:"cashLetterHeader,omitempty"`
	// Bundle is an array of ICL Bundle
	Bundles []Bundle `json:"bundles,omitempty"`
	// ReturnBundles is an array of Return Bundle
	// ReturnBundles      []ReturnBundle      `json:"returnBundle,omitempty"`
	// currentBundle is the currentBundle being parsed
	currentBundle Bundle
	// currentReturnBundle is the current ReturnBundle being parsed
	// currentReturnBundle ReturnBundle
	// CashLetterControl is a Cash Letter Control Record
	CashLetterControl *CashLetterControl `json:"cashLetterControl,omitempty"`
	// Converters is composed for x9 to GoLang Converters
	converters
}

// NewCashLetter takes a CashLetterHeader and returns a CashLetter
// ToDo:  Follow up on returning a pointer when implementing tests and examples
func NewCashLetter(clh *CashLetterHeader) CashLetter {
	cl := CashLetter{}
	cl.SetControl(NewCashLetterControl())
	cl.SetHeader(clh)
	return cl
}

// Validate performs X9 validations and format rule checks and returns an error if not Validated
func (cl *CashLetter) Validate() error {
	// ToDo:  If CashLetterRecordTypeIndicator is "N", There should be no bundle, it is an empty cash letter
	return nil
}

// SetHeader appends a CashLetterHeader to the CashLetter
func (cl *CashLetter) SetHeader(cashLetterHeader *CashLetterHeader) {
	cl.CashLetterHeader = cashLetterHeader
}

// GetHeader returns the current CashLetter header
func (cl *CashLetter) GetHeader() *CashLetterHeader {
	return cl.CashLetterHeader
}

// SetControl appends a CashLetterControl to the CashLetter
func (cl *CashLetter) SetControl(cashLetterControl *CashLetterControl) {
	cl.CashLetterControl = cashLetterControl
}

// GetControl returns the current CashLetter Control
func (cl *CashLetter) GetControl() *CashLetterControl {
	return cl.CashLetterControl
}

// AddBundle appends a Bundle to the x9.File.CashLetter
func (cl *CashLetter) AddBundle(bundle Bundle) []Bundle {
	cl.Bundles = append(cl.Bundles, bundle)
	return cl.Bundles
}

// GetBundles returns a slice of Bundles for the CashLetter
func (cl *CashLetter) GetBundles() []Bundle {
	return cl.Bundles
}
