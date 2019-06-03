// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import "fmt"

// CashLetterError is an Error that describes CashLetter validation issues
type CashLetterError struct {
	CashLetterID string
	FieldName    string
	Msg          string
}

func (e *CashLetterError) Error() string {
	return fmt.Sprintf("CashLetterNumber %s %s %s", e.CashLetterID, e.FieldName, e.Msg)
}

// Errors specific to parsing a CashLetter
var (
	msgCashLetterBundleEntries = "%v cannot have bundle entries"
	msgCashLetterRoutingNumber = "%v cannot have a Routing Number Summary"
)

// CashLetter contains CashLetterHeader, CashLetterControl and Bundle records.
type CashLetter struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// CashLetterHeader is a Cash Letter Header Record
	CashLetterHeader *CashLetterHeader `json:"cashLetterHeader,omitempty"`
	// Bundles is an array of Bundle
	Bundles []*Bundle `json:"bundles,omitempty"`
	// CreditItems is an array of CreditItem
	CreditItems []*CreditItem `json:"creditItem,omitempty"`
	// RoutingNumberSummary is an array of RoutingNumberSummary
	RoutingNumberSummary []*RoutingNumberSummary `json:"routingNumberSummary,omitempty"`
	// currentBundle is the currentBundle being parsed
	currentBundle *Bundle
	// RoutingNumberSummary is an X9 RoutingNumberSummary
	currentRoutingNumberSummary *RoutingNumberSummary
	// CashLetterControl is a Cash Letter Control Record
	CashLetterControl *CashLetterControl `json:"cashLetterControl,omitempty"`
}

// NewCashLetter takes a CashLetterHeader and returns a CashLetter
func NewCashLetter(clh *CashLetterHeader) CashLetter {
	cl := CashLetter{}
	cl.SetControl(NewCashLetterControl())
	cl.SetHeader(clh)
	return cl
}

// Validate performs X9 validations and format rule checks and returns an error if not Validated
func (cl *CashLetter) Validate() error {
	if cl.CashLetterHeader.RecordTypeIndicator == "N" {
		if cl.GetBundles() != nil {
			msg := fmt.Sprintf(msgCashLetterBundleEntries, cl.CashLetterHeader.RecordTypeIndicator)
			return &CashLetterError{CashLetterID: cl.CashLetterHeader.CashLetterID,
				FieldName: "RecordTypeIndicator", Msg: msg}
		}
	}
	switch cl.CashLetterHeader.CollectionTypeIndicator {
	case
		"00", "01", "02":
	default:
		if cl.GetRoutingNumberSummary() != nil {
			msg := fmt.Sprintf(msgCashLetterRoutingNumber, cl.CashLetterHeader.CollectionTypeIndicator)
			return &CashLetterError{CashLetterID: cl.CashLetterHeader.CashLetterID,
				FieldName: "CollectionTypeIndicator", Msg: msg}
		}
	}

	return nil
}

// build a valid CashLetter by building a CashLetterControl. An error is returned if
// the CashLetter being built has invalid records.
func (cl *CashLetter) build() error {

	// Requires a valid CashLetterHeader
	if err := cl.CashLetterHeader.Validate(); err != nil {
		return err
	}

	// CashLetterControl Counts
	cashLetterBundleCount := len(cl.Bundles)
	cashLetterItemsCount := 0
	cashLetterTotalAmount := 0
	cashLetterImagesCount := 0

	// Sequence Numbers
	bundleSequenceNumber := 1
	//creditIndicator
	creditIndicator := 0

	if len(cl.GetCreditItems()) > 0 {
		cashLetterItemsCount = cashLetterItemsCount + len(cl.GetCreditItems())
		creditIndicator = 1
	}
	// Bundles
	for _, b := range cl.Bundles {

		// Set Bundle Sequence Numbers
		b.BundleHeader.SetBundleSequenceNumber(bundleSequenceNumber)

		// Check Items
		for _, cd := range b.Checks {

			// Sequence  Number
			cdSequenceNumber := 1
			// Record Numbers
			cdAddendumARecordNumber := 1
			cdAddendumCRecordNumber := 1

			// Set CheckDetail Sequence Numbers
			cd.SetEceInstitutionItemSequenceNumber(cdSequenceNumber)

			// Set Addenda SequenceNumber and RecordNumber
			for i := range cd.CheckDetailAddendumA {
				cd.CheckDetailAddendumA[i].SetBOFDItemSequenceNumber(cdSequenceNumber)
				cd.CheckDetailAddendumA[i].RecordNumber = cdAddendumARecordNumber
				cdAddendumARecordNumber++
				if cdAddendumARecordNumber > 9 {
					cdAddendumARecordNumber = 1
				}
			}
			for x := range cd.CheckDetailAddendumC {
				cd.CheckDetailAddendumC[x].SetEndorsingBankItemSequenceNumber(cdSequenceNumber)
				cd.CheckDetailAddendumC[x].RecordNumber = cdAddendumARecordNumber
				cdAddendumCRecordNumber++
				if cdAddendumCRecordNumber > 99 {
					cdAddendumCRecordNumber = 1
				}
			}
			cdSequenceNumber++

			cashLetterItemsCount = cashLetterItemsCount + 1
			cashLetterItemsCount = cashLetterItemsCount + len(cd.CheckDetailAddendumA) + len(cd.CheckDetailAddendumB) + len(cd.CheckDetailAddendumC)
			cashLetterItemsCount = cashLetterItemsCount + len(cd.ImageViewDetail) + len(cd.ImageViewData) + len(cd.ImageViewAnalysis)
			cashLetterTotalAmount = cashLetterTotalAmount + cd.ItemAmount
			cashLetterImagesCount = cashLetterImagesCount + len(cd.ImageViewDetail)
		}

		// Returns Items
		for _, rd := range b.Returns {

			// Sequence  Number
			rdSequenceNumber := 1
			// Record Numbers
			rdAddendumARecordNumber := 1
			rdAddendumDRecordNumber := 1

			// Set ReturnDetail Sequence Numbers
			rd.SetEceInstitutionItemSequenceNumber(rdSequenceNumber)

			// Set Addenda SequenceNumber and RecordNumber
			for i := range rd.ReturnDetailAddendumA {
				rd.ReturnDetailAddendumA[i].SetBOFDItemSequenceNumber(rdSequenceNumber)
				rd.ReturnDetailAddendumA[i].RecordNumber = rdAddendumARecordNumber
				rdAddendumARecordNumber++
				if rdAddendumARecordNumber > 9 {
					rdAddendumARecordNumber = 1
				}
			}

			for x := range rd.ReturnDetailAddendumD {
				rd.ReturnDetailAddendumD[x].SetEndorsingBankItemSequenceNumber(rdSequenceNumber)
				rd.ReturnDetailAddendumA[x].RecordNumber = rdAddendumDRecordNumber
				rdAddendumDRecordNumber++
				if rdAddendumDRecordNumber > 99 {
					rdAddendumDRecordNumber = 1
				}
			}
			rdSequenceNumber++

			cashLetterItemsCount = cashLetterItemsCount + 1
			cashLetterItemsCount = cashLetterItemsCount + len(rd.ReturnDetailAddendumA) + len(rd.ReturnDetailAddendumB) + len(rd.ReturnDetailAddendumC) + len(rd.ReturnDetailAddendumD)
			cashLetterItemsCount = cashLetterItemsCount + len(rd.ImageViewDetail) + len(rd.ImageViewData) + len(rd.ImageViewAnalysis)
			cashLetterTotalAmount = cashLetterTotalAmount + rd.ItemAmount
			cashLetterImagesCount = cashLetterImagesCount + len(rd.ImageViewDetail)
		}
		// Validate Bundle
		if err := b.Validate(); err != nil {
			return err
		}
		// Build Bundle
		if err := b.build(); err != nil {
			return err
		}

		bundleSequenceNumber++
	}

	// build a CashLetterControl record
	clc := NewCashLetterControl()
	clc.CashLetterBundleCount = cashLetterBundleCount
	clc.CashLetterItemsCount = cashLetterItemsCount
	clc.CashLetterTotalAmount = cashLetterTotalAmount
	clc.CashLetterImagesCount = cashLetterImagesCount
	clc.ECEInstitutionName = cl.GetHeader().ECEInstitutionRoutingNumber
	clc.CreditTotalIndicator = creditIndicator
	cl.CashLetterControl = clc
	return nil
}

// Create creates a CashLetter of Bundles containing CheckDetail or ReturnDetail
func (cl *CashLetter) Create() error {
	if err := cl.build(); err != nil {
		return err
	}
	return cl.Validate()
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

// AddBundle appends a Bundle to the CashLetter
func (cl *CashLetter) AddBundle(bundle *Bundle) []*Bundle {
	cl.Bundles = append(cl.Bundles, bundle)
	return cl.Bundles
}

// GetBundles returns a slice of Bundle for the CashLetter
func (cl *CashLetter) GetBundles() []*Bundle {
	return cl.Bundles
}

// AddRoutingNumberSummary appends a RoutingNumberSummary to the CashLetter
func (cl *CashLetter) AddRoutingNumberSummary(rns *RoutingNumberSummary) []*RoutingNumberSummary {
	cl.RoutingNumberSummary = append(cl.RoutingNumberSummary, rns)
	return cl.RoutingNumberSummary
}

// GetRoutingNumberSummary returns a slice of RoutingNumberSummary for the CashLetter
func (cl *CashLetter) GetRoutingNumberSummary() []*RoutingNumberSummary {
	return cl.RoutingNumberSummary
}

// AddCreditItem appends a CreditItem to the CashLetter
func (cl *CashLetter) AddCreditItem(ci *CreditItem) []*CreditItem {
	cl.CreditItems = append(cl.CreditItems, ci)
	return cl.CreditItems
}

// GetCreditItems returns a slice of CreditItem for the CashLetter
func (cl *CashLetter) GetCreditItems() []*CreditItem {
	return cl.CreditItems
}
