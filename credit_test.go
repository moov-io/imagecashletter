// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"
)

// mockCredit creates a Credit
func mockCredit() *Credit {
	cr := NewCredit()

	cr.AuxiliaryOnUs = "010910999940910"
	cr.ExternalProcessingCode = ""
	cr.PayorBankRoutingNumber = "999920060"
	cr.CreditAccountNumberOnUs = "50920060509383521210"
	cr.ItemAmount = 102088
	cr.ECEInstitutionItemSequenceNumber = "               "
	cr.DocumentationTypeIndicator = "G"
	cr.AccountTypeCode = "1"
	cr.SourceWorkCode = "3"
	cr.WorkType = " "
	cr.DebitCreditIndicator = " "

	return cr
}

// TestMockCredit creates a CreditItem
func TestMockCredit(t *testing.T) {
	ci := mockCredit()
	if err := ci.Validate(); err != nil {
		t.Error("mockCredit does not validate and will break other tests: ", err)
	}
	if ci.recordType != "61" {
		t.Error("recordType does not validate")
	}
	if ci.AuxiliaryOnUs != "010910999940910" {
		t.Error("AuxiliaryOnUs does not validate")
	}
	if ci.ExternalProcessingCode != "" {
		t.Error("ExternalProcessingCode does not validate")
	}
	if ci.PayorBankRoutingNumber != "999920060" {
		t.Error("PayorBankRoutingNumber does not validate")
	}
	if ci.CreditAccountNumberOnUs != "50920060509383521210" {
		t.Error("CreditAccountNumberOnUs does not validate")
	}
	if ci.ItemAmount != 102088 {
		t.Error("ItemAmount does not validate")
	}
	if ci.ECEInstitutionItemSequenceNumber != "               " {
		t.Error("ECEInstitutionItemSequenceNumber does not validate")
	}
	if ci.DocumentationTypeIndicator != "G" {
		t.Error("DocumentationTypeIndicator does not validate")
	}
	if ci.AccountTypeCode != "1" {
		t.Error("AccountTypeCode does not validate")
	}
	if ci.SourceWorkCode != "3" {
		t.Error("SourceWorkCode does not validate")
	}
	if ci.WorkType != " " {
		t.Error("WorkType does not validate")
	}
	if ci.DebitCreditIndicator != " " {
		t.Error("DebitCreditIndicator does not validate")
	}
}

func TestCreditCrash(t *testing.T) {
	cr := &Credit{}
	cr.Parse(`61010910999940910 999920060509200605093835212100000102088               G13     `)
	if cr.DocumentationTypeIndicator != "G" {
		t.Errorf("expected ci.DocumentationTypeIndicator=G")
	}
}

func TestParseCredit(t *testing.T) {
	var line = "61010910999940910 999920060509200605093835212100000102088               G13     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	cr := mockCredit()
	r.currentCashLetter.AddCredit(cr)
	if err := r.parseCredit(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentCashLetter.GetCredits()[0]

	if record.recordType != "61" {
		t.Errorf("RecordType Expected '61' got: %v", record.recordType)
	}
	if record.AuxiliaryOnUs != "010910999940910" {
		t.Errorf("AuxiliaryOnUs Expected '010910999940910' got: %v", record.AuxiliaryOnUs)
	}
	if record.ExternalProcessingCode != "" {
		t.Errorf("ExternalProcessingCode Expected '' got: %v", record.ExternalProcessingCode)
	}
	if record.PayorBankRoutingNumber != "999920060" {
		t.Errorf("PostingBankRoutingNumber Expected '999920060' got: %v", record.PayorBankRoutingNumber)
	}
	if record.CreditAccountNumberOnUs != "50920060509383521210" {
		t.Errorf("OnUs Expected '50920060509383521210' got: %v", record.CreditAccountNumberOnUs)
	}
	if record.ItemAmount != 102088 {
		t.Errorf("ItemAmount Expected '102088' got: %v", record.ItemAmount)
	}
	if record.ECEInstitutionItemSequenceNumber != "               " {
		t.Errorf("ECEInstitutionItemSequenceNumber Expected '               ' got: %v", record.ECEInstitutionItemSequenceNumber)
	}
	if record.DocumentationTypeIndicator != "G" {
		t.Errorf("DocumentationTypeIndicator Expected 'G' got: %v", record.DocumentationTypeIndicator)
	}
	if record.AccountTypeCode != "1" {
		t.Errorf("AccountTypeCode Expected '1' got: %v", record.AccountTypeCode)
	}
	if record.SourceWorkCode != "3" {
		t.Errorf("SourceWorkCode Expected '3' got: %v", record.SourceWorkCode)
	}
	if record.WorkType != " " {
		t.Errorf("WorkType Expected ' ' got: %v", record.WorkType)
	}
	if record.DebitCreditIndicator != " " {
		t.Errorf("DebitCreditIndicator Expected ' ' got: %v", record.DebitCreditIndicator)
	}
}

// testCIString validates parsing a CreditItem
func testCRString(t testing.TB) {
	var line = "61010910999940910 999920060509200605093835212100000102088               G13     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	cr := mockCredit()
	r.currentCashLetter.AddCredit(cr)
	if err := r.parseCredit(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentCashLetter.GetCredits()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCRString tests validating that a known parsed CheckDetail can return to a string of the same value
func TestCRString(t *testing.T) {
	testCRString(t)
}

// BenchmarkCRString benchmarks validating that a known parsed Credit
// can return to a string of the same value
func BenchmarkCRString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCRString(b)
	}
}

// TestCRRecordType validation
func TestCRRecordType(t *testing.T) {
	ci := mockCredit()
	ci.recordType = "00"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCRDocumentationTypeIndicator validation
func TestCRDocumentationTypeIndicator(t *testing.T) {
	ci := mockCredit()
	ci.DocumentationTypeIndicator = "P"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCRDocumentationTypeIndicatorZ validation
func TestCRDocumentationTypeIndicatorZ(t *testing.T) {
	ci := mockCredit()
	ci.DocumentationTypeIndicator = "Z"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCRDocumentationTypeIndicatorM validation
func TestCRDocumentationTypeIndicatorM(t *testing.T) {
	ci := mockCredit()
	ci.DocumentationTypeIndicator = "M"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCRSourceWorkCode validation
func TestCRSourceWorkCode(t *testing.T) {
	ci := mockCredit()
	ci.SourceWorkCode = "99"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "SourceWorkCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestCRFIRecordType validation
func TestCRFIRecordType(t *testing.T) {
	ci := mockCredit()
	ci.recordType = ""
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCRPayorBankRoutingNumber validation
func TestCRPayorBankRoutingNumber(t *testing.T) {
	ci := mockCredit()
	ci.PayorBankRoutingNumber = "000000000"
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCRCreditAccountNumberOnUs validation
func TestCRCreditAccountNumberOnUs(t *testing.T) {
	ci := mockCredit()
	ci.CreditAccountNumberOnUs = ""
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CreditAccountNumberOnUs" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCRItemAmount validation
func TestCRItemAmount(t *testing.T) {
	ci := mockCredit()
	ci.ItemAmount = 0
	if err := ci.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ItemAmount" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
