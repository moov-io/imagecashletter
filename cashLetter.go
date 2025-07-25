// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"errors"
	"fmt"
)

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
	msgMandatoryRecord         = "record is mandatory"
)

// CashLetter contains CashLetterHeader, CashLetterControl and Bundle records.
type CashLetter struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// CashLetterHeader is a Cash Letter Header Record
	CashLetterHeader *CashLetterHeader `json:"cashLetterHeader,omitempty"`
	// Bundles is an array of Bundle
	Bundles []*Bundle `json:"bundles,omitempty"`
	// Credits is an array of Credit
	Credits []*Credit `json:"credit,omitempty"`
	// CreditItems is an array of CreditItem
	CreditItems []*CreditItem `json:"creditItem,omitempty"`
	// RoutingNumberSummary is an array of RoutingNumberSummary
	RoutingNumberSummary []*RoutingNumberSummary `json:"routingNumberSummary,omitempty"`
	// currentBundle is the currentBundle being parsed
	currentBundle *Bundle
	// RoutingNumberSummary is an imagecashletter RoutingNumberSummary
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

func (cl *CashLetter) setRecordType() {
	if cl == nil {
		return
	}

	cl.CashLetterHeader.setRecordType()
	for i := range cl.Bundles {
		cl.Bundles[i].setRecordType()
	}
	for i := range cl.CreditItems {
		cl.CreditItems[i].setRecordType()
	}
	for i := range cl.RoutingNumberSummary {
		cl.RoutingNumberSummary[i].setRecordType()
	}
	cl.CashLetterControl.setRecordType()
}

// Validate performs ImageCashLetter validations and format rule checks and returns an error if not Validated
func (cl *CashLetter) Validate() error {
	if cl.CashLetterHeader == nil {
		return errors.New("nil CashLetterHeader")
	}

	if cl.CashLetterHeader.RecordTypeIndicator == "N" {
		if cl.GetBundles() != nil {
			return &CashLetterError{
				CashLetterID: cl.CashLetterHeader.CashLetterID,
				FieldName:    "RecordTypeIndicator",
				Msg:          fmt.Sprintf(msgCashLetterBundleEntries, cl.CashLetterHeader.RecordTypeIndicator),
			}
		}
	}
	switch cl.CashLetterHeader.CollectionTypeIndicator {
	case
		"00", "01", "02":
	default:
		if cl.GetRoutingNumberSummary() != nil {
			return &CashLetterError{
				CashLetterID: cl.CashLetterHeader.CashLetterID,
				FieldName:    "CollectionTypeIndicator",
				Msg:          fmt.Sprintf(msgCashLetterRoutingNumber, cl.CashLetterHeader.CollectionTypeIndicator),
			}
		}
	}

	if cl.CashLetterControl == nil {
		return &CashLetterError{
			CashLetterID: cl.CashLetterHeader.CashLetterID,
			FieldName:    "CashLetterControl",
			Msg:          msgMandatoryRecord,
		}
	}

	if err := cl.CashLetterControl.Validate(); err != nil {
		return err
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
	// creditIndicator
	creditIndicator := 0

	if len(cl.GetCreditItems()) > 0 {
		cashLetterItemsCount = cashLetterItemsCount + len(cl.GetCreditItems())
		creditIndicator = 1
	}
	// Bundles
	for _, b := range cl.Bundles {

		// Set Bundle Sequence Numbers
		if b.BundleHeader.BundleSequenceNumber != "" {
			bundleSequenceNumber = b.parseNumField(b.BundleHeader.BundleSequenceNumber)
		}
		b.BundleHeader.SetBundleSequenceNumber(bundleSequenceNumber)

		// Sequence  Number
		cdSequenceNumber := 1

		// Check Items
		for _, cd := range b.Checks {

			if cd.EceInstitutionItemSequenceNumber != "" {
				cdSequenceNumber = cd.parseNumField(cd.EceInstitutionItemSequenceNumber)
			}

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
				if cd.CheckDetailAddendumC[x].EndorsingBankItemSequenceNumber == "" {
					cd.CheckDetailAddendumC[x].SetEndorsingBankItemSequenceNumber(cdSequenceNumber)
				}
				cd.CheckDetailAddendumC[x].RecordNumber = cdAddendumCRecordNumber
				cdAddendumCRecordNumber++
				if cdAddendumCRecordNumber > 99 {
					cdAddendumCRecordNumber = 1
				}
			}
			cdSequenceNumber++

			cashLetterItemsCount = cashLetterItemsCount + 1
			cashLetterTotalAmount = cashLetterTotalAmount + cd.ItemAmount
			cashLetterImagesCount = cashLetterImagesCount + len(cd.ImageViewDetail)
		}

		rdSequenceNumber := 1

		// Returns Items
		for _, rd := range b.Returns {

			// Override the default sequence number if set
			if rd.EceInstitutionItemSequenceNumber != "" {
				rdSequenceNumber = rd.parseNumField(rd.EceInstitutionItemSequenceNumber)
			}

			// Set ReturnDetail Sequence Numbers
			rd.SetEceInstitutionItemSequenceNumber(rdSequenceNumber)

			// Record Numbers
			rdAddendumARecordNumber := 1
			rdAddendumDRecordNumber := 1

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
				rd.ReturnDetailAddendumD[x].RecordNumber = rdAddendumDRecordNumber
				rdAddendumDRecordNumber++
				if rdAddendumDRecordNumber > 99 {
					rdAddendumDRecordNumber = 1
				}
			}

			cashLetterItemsCount = cashLetterItemsCount + 1
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
	if cl.CashLetterControl.ECEInstitutionName != "" {
		clc.ECEInstitutionName = cl.CashLetterControl.ECEInstitutionName
	} else {
		clc.ECEInstitutionName = cl.GetHeader().ECEInstitutionRoutingNumber
	}
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
	if cl == nil {
		return nil
	}
	return cl.Bundles
}

// AddRoutingNumberSummary appends a RoutingNumberSummary to the CashLetter
func (cl *CashLetter) AddRoutingNumberSummary(rns *RoutingNumberSummary) []*RoutingNumberSummary {
	cl.RoutingNumberSummary = append(cl.RoutingNumberSummary, rns)
	return cl.RoutingNumberSummary
}

// GetRoutingNumberSummary returns a slice of RoutingNumberSummary for the CashLetter
func (cl *CashLetter) GetRoutingNumberSummary() []*RoutingNumberSummary {
	if cl == nil {
		return nil
	}
	return cl.RoutingNumberSummary
}

// AddCredit appends a CreditItem to the CashLetter
func (cl *CashLetter) AddCredit(cr *Credit) []*Credit {
	cl.Credits = append(cl.Credits, cr)
	return cl.Credits
}

// AddCreditItem appends a CreditItem to the CashLetter
func (cl *CashLetter) AddCreditItem(ci *CreditItem) []*CreditItem {
	cl.CreditItems = append(cl.CreditItems, ci)
	return cl.CreditItems
}

// GetCredits returns a slice of Credit for the CashLetter
func (cl *CashLetter) GetCredits() []*Credit {
	if cl == nil {
		return nil
	}
	return cl.Credits
}

// GetCreditItems returns a slice of CreditItem for the CashLetter
func (cl *CashLetter) GetCreditItems() []*CreditItem {
	if cl == nil {
		return nil
	}
	return cl.CreditItems
}
