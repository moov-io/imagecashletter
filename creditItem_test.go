// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"strings"
	"testing"
)

// mockCreditItem creates a CreditItem
func mockCreditItem() *CreditItem {
	ci := NewCreditItem()
	ci.AuxiliaryOnUs = "123456789"
	ci.ExternalProcessingCode = ""
	ci.PostingBankRoutingNumber = "031300012"
	ci.OnUs = "5558881"
	ci.ItemAmount = 100000 // 1000.00
	ci.CreditItemSequenceNumber = "1              "
	ci.DocumentationTypeIndicator = "G"
	ci.AccountTypeCode = "1"
	ci.SourceWorkCode = "01"
	ci.UserField = "                "
	return ci
}

// TestMockCreditItem creates a CreditItem
func TestMockCreditItem(t *testing.T) {
	ci := mockCreditItem()
	if err := ci.Validate(); err != nil {
		t.Error("mockCreditItem does not validate and will break other tests: ", err)
	}
	if ci.recordType != "62" {
		t.Error("recordType does not validate")
	}
	if ci.AuxiliaryOnUs != "123456789" {
		t.Error("AuxiliaryOnUs does not validate")
	}
	if ci.ExternalProcessingCode != "" {
		t.Error("ExternalProcessingCode does not validate")
	}
	if ci.PostingBankRoutingNumber != "031300012" {
		t.Error("PayorBankRoutingNumber does not validate")
	}
	if ci.OnUs != "5558881" {
		t.Error("OnUs does not validate")
	}
	if ci.ItemAmount != 100000 {
		t.Error("ItemAmount does not validate")
	}
	if ci.CreditItemSequenceNumber != "1              " {
		t.Error("CreditItemSequence does not validate")
	}
	if ci.DocumentationTypeIndicator != "G" {
		t.Error("DocumentationTypeIndicator does not validate")
	}
	if ci.AccountTypeCode != "1" {
		t.Error("AccountTypeCode does not validate")
	}
	if ci.SourceWorkCode != "01" {
		t.Error("SourceWorkCode does not validate")
	}
	if ci.UserField != "                " {
		t.Error("UserField does not validate")
	}
}

// TestParseCreditItem validates parsing a CreditItem
func TestParseCreditItem(t *testing.T) {
	var line = "62      123456789 031300012             5558881000000001000001              G101                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	ci := mockCreditItem()
	r.currentCashLetter.AddCreditItem(ci)
	if err := r.parseCreditItem(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentCashLetter.GetCreditItems()[0]

	if record.recordType != "62" {
		t.Errorf("RecordType Expected '62' got: %v", record.recordType)
	}
	if record.AuxiliaryOnUs != "123456789" {
		t.Errorf("AuxiliaryOnUs Expected '123456789' got: %v", record.AuxiliaryOnUs)
	}
	if record.ExternalProcessingCode != "" {
		t.Errorf("ExternalProcessingCode Expected '' got: %v", record.ExternalProcessingCode)
	}
	if record.PostingBankRoutingNumber != "031300012" {
		t.Errorf("PostingBankRoutingNumber Expected '031300012' got: %v", record.PostingBankRoutingNumber)
	}
	if record.OnUs != "5558881" {
		t.Errorf("OnUs Expected '5558881' got: %v", record.OnUs)
	}
	if record.ItemAmount != 100000 {
		t.Errorf("ItemAmount Expected '100000' got: %v", record.ItemAmount)
	}
	if record.CreditItemSequenceNumber != "1              " {
		t.Errorf("CreditItemSequenceNumber Expected '1              ' got: %v", record.CreditItemSequenceNumber)
	}
	if record.DocumentationTypeIndicator != "G" {
		t.Errorf("DocumentationTypeIndicator Expected 'G' got: %v", record.DocumentationTypeIndicator)
	}
	if record.AccountTypeCode != "1" {
		t.Errorf("AccountTypeCode Expected '1' got: %v", record.AccountTypeCode)
	}
	if record.SourceWorkCode != "01" {
		t.Errorf("SourceWorkCode Expected '01' got: %v", record.SourceWorkCode)
	}
	if record.UserField != "                " {
		t.Errorf("UserField Expected '                ' got: %v", record.UserField)
	}
}

// testCIString validates parsing a CreditItem
func testCIString(t testing.TB) {
	var line = "62      123456789 031300012             5558881000000001000001              G101                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	ci := mockCreditItem()
	r.currentCashLetter.AddCreditItem(ci)
	if err := r.parseCreditItem(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentCashLetter.GetCreditItems()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCIString tests validating that a known parsed CheckDetail can return to a string of the same value
func TestCIString(t *testing.T) {
	testCIString(t)
}

// BenchmarkCIString benchmarks validating that a known parsed CreditItem
// can return to a string of the same value
func BenchmarkCIString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCIString(b)
	}
}

// TestCIRecordType validation
func TestCIRecordType(t *testing.T) {
	ci := mockCreditItem()
	ci.recordType = "00"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCIDocumentationTypeIndicator validation
func TestCIDocumentationTypeIndicator(t *testing.T) {
	ci := mockCreditItem()
	ci.DocumentationTypeIndicator = "P"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCIDocumentationTypeIndicatorZ validation
func TestCIDocumentationTypeIndicatorZ(t *testing.T) {
	ci := mockCreditItem()
	ci.DocumentationTypeIndicator = "Z"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCIDocumentationTypeIndicatorM validation
func TestCIDocumentationTypeIndicatorM(t *testing.T) {
	ci := mockCreditItem()
	ci.DocumentationTypeIndicator = "M"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCIAccountTypeCode validation
func TestCIAccountTypeCode(t *testing.T) {
	ci := mockCreditItem()
	ci.AccountTypeCode = "Z"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "AccountTypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCISourceWorkCode validation
func TestCISourceWorkCode(t *testing.T) {
	ci := mockCreditItem()
	ci.SourceWorkCode = "99"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "SourceWorkCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCIUserField validation
func TestCIUserField(t *testing.T) {
	ci := mockCreditItem()
	ci.UserField = "®©"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestCIFIRecordType validation
func TestCIFIRecordType(t *testing.T) {
	ci := mockCreditItem()
	ci.recordType = ""
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCIFIPostingBankRoutingNumber validation
func TestCIFIPostingBankRoutingNumber(t *testing.T) {
	ci := mockCreditItem()
	ci.PostingBankRoutingNumber = ""
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PostingBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCIFIPostingBankRoutingNumberZero validation
func TestCIFIPostingBankRoutingNumberZero(t *testing.T) {
	ci := mockCreditItem()
	ci.PostingBankRoutingNumber = "000000000"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PostingBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCIFICreditItemSequenceNumber validation
func TestCIFICreditItemSequenceNumber(t *testing.T) {
	ci := mockCreditItem()
	ci.CreditItemSequenceNumber = ""
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CreditItemSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
