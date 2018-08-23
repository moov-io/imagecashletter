// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// CashLetter contains CashLetterHeader, CashLetterControl and Bundle records.
type CashLetter struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// CashLetterHeader is an ICL Cash Letter Header Record
	CashLetterHeader *CashLetterHeader `json:"cashLetterHeader,omitempty"`
	// Bundle is an ICL Bundle
	Bundle *Bundle `json:"bundle,omitempty"`
	// ReturnBundle is an ICL Return Bundle
	//ReturnBundle      *ReturnBundle      `json:"returnBundle,omitempty"`
	// CashLetterControl is an ICL Cash Letter Control Record
	CashLetterControl *CashLetterControl `json:"cashLetterControl,omitempty"`
	// Converters is composed for x9 to GoLang Converters
	converters
}

// NewCashLetter takes a CashLetterHeader and returns a CashLetter
// ToDo:  Follow up on returning a pointer when implementing tests and examples
func NewCashLetter(clh *CashLetterHeader) *CashLetter {
	cashLetter := &CashLetter{}
	cashLetter.SetControl(NewCashLetterControl())
	cashLetter.SetHeader(clh)
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
