// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// CashLetter contains CashLetterHeader, CashLetterControl and Bundle records.
type CashLetter struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// CashLetterHeader is an ICl CashLetterHeader
	CashLetterHeader *CashLetterHeader `json:"cashLetterHeader,omitempty"`
	// Bundle is an ICL Bundle
	Bundle *Bundle `json:"bundle,omitempty"`
	// ReturnBundle is an ICL ReturnBundle
	//ReturnBundle      *ReturnBundle      `json:"returnBundle,omitempty"`
	// CashLetterControl is an ICL CashLetterControl
	CashLetterControl *CashLetterControl `json:"cashLetterControl,omitempty"`
	// Converters is composed for x9 to GoLang Converters
	converters
}

// NewCashLetter takes a CashLetterHeader and returns a CashLetter
func NewCashLetter(clh *CashLetterHeader) CashLetter {
	cashLetter := CashLetter{}
	//bundle.SetControl(NewBundleControl())
	//bundle.SetHeader(bh)
	return cashLetter
}

// SetHeader appends an CashLetterHeader to the CashLetter
func (cashLetter *CashLetter) SetHeader(cashLetterHeader *CashLetterHeader) {
	cashLetter.CashLetterHeader = cashLetterHeader
}

// SetControl appends an CashLetterControl to the CashLetter
func (cashLetter *CashLetter) SetControl(cashLetterControl *CashLetterControl) {
	cashLetter.CashLetterControl = cashLetterControl
}
