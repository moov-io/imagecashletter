// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"strings"
	"testing"
)

// mockCheckDetail creates a CheckDetail
func mockCheckDetail() *CheckDetail {
	cd := NewCheckDetail()
	cd.AuxiliaryOnUs = "123456789"
	cd.ExternalProcessingCode = ""
	cd.PayorBankRoutingNumber = "03130001"
	cd.PayorBankCheckDigit = "2"
	cd.OnUs = "5558881"
	cd.ItemAmount = 100000 // 1000.00
	cd.EceInstitutionItemSequenceNumber = "1              "
	cd.DocumentationTypeIndicator = "G"
	cd.ReturnAcceptanceIndicator = "D"
	cd.MICRValidIndicator = 1
	cd.BOFDIndicator = "Y"
	cd.AddendumCount = 3
	cd.CorrectionIndicator = 0
	cd.ArchiveTypeIndicator = "B"
	return cd
}

// TestMockCheckDetail creates a CheckDetail
func TestMockCheckDetail(t *testing.T) {
	cd := mockCheckDetail()
	if err := cd.Validate(); err != nil {
		t.Error("mockCheckDetail does not validate and will break other tests: ", err)
	}
	if cd.recordType != "25" {
		t.Error("recordType does not validate")
	}
	if cd.AuxiliaryOnUs != "123456789" {
		t.Error("AuxiliaryOnUs does not validate")
	}
	if cd.ExternalProcessingCode != "" {
		t.Error("ExternalProcessingCode does not validate")
	}
	if cd.PayorBankRoutingNumber != "03130001" {
		t.Error("PayorBankRoutingNumber does not validate")
	}
	if cd.PayorBankCheckDigit != "2" {
		t.Error("PayorBankCheckDigit does not validate")
	}
	if cd.OnUs != "5558881" {
		t.Error("OnUs does not validate")
	}
	if cd.ItemAmount != 100000 {
		t.Error("ItemAmount does not validate")
	}
	if cd.EceInstitutionItemSequenceNumber != "1              " {
		t.Error("EceInstitutionItemSequenceNumber does not validate")
	}
	if cd.DocumentationTypeIndicator != "G" {
		t.Error("DocumentationTypeIndicator does not validate")
	}
	if cd.ReturnAcceptanceIndicator != "D" {
		t.Error("ReturnAcceptanceIndicator does not validate")
	}
	if cd.MICRValidIndicator != 1 {
		t.Error("MICRValidIndicator does not validate")
	}
	if cd.BOFDIndicator != "Y" {
		t.Error("BOFDIndicator does not validate")
	}
	if cd.AddendumCount != 3 {
		t.Error("AddendumCount does not validate")
	}
	if cd.CorrectionIndicator != 0 {
		t.Error("CorrectionIndicator does not validate")
	}
	if cd.ArchiveTypeIndicator != "B" {
		t.Error("ArchiveTypeIndicator does not validate")
	}
}

// TestParseCheckDetail validates parsing a CheckDetail
func TestParseCheckDetail(t *testing.T) {
	var line = "25      123456789 031300012             555888100001000001              GD1Y030B"
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)

	if err := r.parseCheckDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0]

	if record.recordType != "25" {
		t.Errorf("RecordType Expected '25' got: %v", record.recordType)
	}
	if record.AuxiliaryOnUsField() != "      123456789" {
		t.Errorf("AuxiliaryOnUs Expected '      123456789' got: %v", record.AuxiliaryOnUsField())
	}
	if record.ExternalProcessingCodeField() != " " {
		t.Errorf("ExternalProcessingCodeField ' ' got: %v", record.ExternalProcessingCodeField())
	}
	if record.PayorBankRoutingNumberField() != "03130001" {
		t.Errorf("PayorBankRoutingNumber Expected '03130001' got: %v", record.PayorBankRoutingNumberField())
	}
	if record.PayorBankCheckDigitField() != "2" {
		t.Errorf("PayorBankCheckDigit Expected '2' got:'%v'", record.PayorBankCheckDigitField())
	}
	if record.OnUsField() != "             5558881" {
		t.Errorf("OnUs Expected '             5558881' got:'%v'", record.OnUsField())
	}
	if record.ItemAmountField() != "0000100000" {
		t.Errorf("ItemAmount Expected '0000100000' got:'%v'", record.ItemAmountField())
	}
	if record.EceInstitutionItemSequenceNumberField() != "1              " {
		t.Errorf("EceInstitutionItemSequenceNumber Expected '1              ' got:'%v'",
			record.EceInstitutionItemSequenceNumberField())
	}
	if record.DocumentationTypeIndicatorField() != "G" {
		t.Errorf("DocumentationTypeIndicator Expected 'G' got:'%v'", record.DocumentationTypeIndicatorField())
	}
	if record.ReturnAcceptanceIndicatorField() != "D" {
		t.Errorf("ReturnAcceptanceIndicator Expected 'D' got: '%v'", record.ReturnAcceptanceIndicatorField())
	}
	if record.MICRValidIndicatorField() != "1" {
		t.Errorf("MICRValidIndicator Expected '01' got:'%v'", record.MICRValidIndicatorField())
	}
	if record.BOFDIndicatorField() != "Y" {
		t.Errorf("BOFDIndicator Expected 'Y' got:'%v'", record.BOFDIndicatorField())
	}
	if record.AddendumCountField() != "03" {
		t.Errorf("AddendumCount Expected '03' got:'%v'", record.AddendumCountField())
	}
	if record.CorrectionIndicatorField() != "0" {
		t.Errorf("CorrectionIndicator Expected '0' got:'%v'", record.CorrectionIndicatorField())
	}
	if record.ArchiveTypeIndicatorField() != "B" {
		t.Errorf("ArchiveTypeIndicator Expected 'B' got:'%v'", record.ArchiveTypeIndicatorField())
	}
}

// testCDString validates that a known parsed CheckDetail can return to a string of the same value
func testCDString(t testing.TB) {
	var line = "25      123456789 031300012             555888100001000001              GD1Y030B"
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	if err := r.parseCheckDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCDString tests validating that a known parsed CheckDetail can return to a string of the same value
func TestCDString(t *testing.T) {
	testCDString(t)
}

// BenchmarkCDString benchmarks validating that a known parsed CheckDetail
// can return to a string of the same value
func BenchmarkCDString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDString(b)
	}
}

// TestCDRecordType validation
func TestCDRecordType(t *testing.T) {
	cd := mockCheckDetail()
	cd.recordType = "00"
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDDocumentationTypeIndicator validation
func TestCDDocumentationTypeIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.DocumentationTypeIndicator = "P"
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDDocumentationTypeIndicatorZ validation
func TestCDDocumentationTypeIndicatorZ(t *testing.T) {
	cd := mockCheckDetail()
	cd.DocumentationTypeIndicator = "Z"
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDReturnAcceptanceIndicator validation
func TestCDReturnAcceptanceIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.ReturnAcceptanceIndicator = "M"
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnAcceptanceIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDMICRValidIndicator validation
func TestCDMICRValidIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.MICRValidIndicator = 7
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "MICRValidIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDBOFDIndicator validation
func TestCDBOFDIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.BOFDIndicator = "B"
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDCorrectionIndicator validation
func TestCDCorrectionIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.CorrectionIndicator = 10
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CorrectionIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDArchiveTypeIndicator validation
func TestCDArchiveTypeIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.ArchiveTypeIndicator = "W"
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ArchiveTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestCDFIRecordType validation
func TestCDFIRecordType(t *testing.T) {
	cd := mockCheckDetail()
	cd.recordType = ""
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDFIPayorBankRoutingNumber validation
func TestCDFIPayorBankRoutingNumber(t *testing.T) {
	cd := mockCheckDetail()
	cd.PayorBankRoutingNumber = ""
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDFIPayorBankRoutingNumberZero validation
func TestCDFIPayorBankRoutingNumberZero(t *testing.T) {
	cd := mockCheckDetail()
	cd.PayorBankRoutingNumber = "00000000"
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDFIPayorBankCheckDigit validation
func TestCDFIPayorBankCheckDigit(t *testing.T) {
	cd := mockCheckDetail()
	cd.PayorBankCheckDigit = ""
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankCheckDigit" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDFIEceInstitutionItemSequenceNumber validation
func TestCDFIEceInstitutionItemSequenceNumber(t *testing.T) {
	cd := mockCheckDetail()
	cd.EceInstitutionItemSequenceNumber = "               "
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EceInstitutionItemSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDFIBOFDIndicator validation
func TestCDFIBOFDIndicator(t *testing.T) {
	cd := mockCheckDetail()
	cd.BOFDIndicator = ""
	if err := cd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
