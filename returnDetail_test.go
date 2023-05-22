// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"log"
	"strings"
	"testing"
	"time"
)

// mockReturnDetail creates a ReturnDetail
func mockReturnDetail() *ReturnDetail {
	rd := NewReturnDetail()
	rd.PayorBankRoutingNumber = "03130001"
	rd.PayorBankCheckDigit = "2"
	rd.OnUs = "5558881"
	rd.ItemAmount = 100000
	rd.ReturnReason = "A"
	rd.AddendumCount = 4
	rd.DocumentationTypeIndicator = "G"
	rd.ForwardBundleDate = time.Now()
	rd.EceInstitutionItemSequenceNumber = "1              "
	rd.ExternalProcessingCode = ""
	rd.ReturnNotificationIndicator = "2"
	rd.ArchiveTypeIndicator = "B"
	rd.TimesReturned = 0
	return rd
}

func TestReturnDetailParse(t *testing.T) {
	var r ReturnDetail
	r.Parse("asshafaksjfas")
	if r.PayorBankRoutingNumber != "" {
		t.Errorf("r.PayorBankRoutingNumber=%s", r.PayorBankRoutingNumber)
	}
}

// TestMockReturnDetail creates a ReturnDetail
func TestMockReturnDetail(t *testing.T) {
	rd := mockReturnDetail()
	if err := rd.Validate(); err != nil {
		t.Error("mockReturnDetail does not validate and will break other tests: ", err)
	}
	if rd.recordType != "31" {
		t.Error("recordType does not validate")
	}
	if rd.PayorBankRoutingNumber != "03130001" {
		t.Error("PayorBankRoutingNumber does not validate")
	}
	if rd.PayorBankCheckDigit != "2" {
		t.Error("PayorBankCheckDigit does not validate")
	}
	if rd.OnUs != "5558881" {
		t.Error("OnUs does not validate")
	}
	if rd.ItemAmount != 100000 {
		t.Error("ItemAmount does not validate")
	}
	if rd.ReturnReason != "A" {
		t.Error("ReturnReason does not validate")
	}
	if rd.AddendumCount != 4 {
		t.Error("AddendumCount does not validate")
	}
	if rd.DocumentationTypeIndicator != "G" {
		t.Error("DocumentationTypeIndicator does not validate")
	}
	if rd.EceInstitutionItemSequenceNumber != "1              " {
		t.Error("EceInstitutionItemSequenceNumber does not validate")
	}
	if rd.ExternalProcessingCode != "" {
		t.Error("ExternalProcessingCode does not validate")
	}
	if rd.ReturnNotificationIndicator != "2" {
		t.Error("ReturnNotificationIndicator does not validate")
	}
	if rd.ArchiveTypeIndicator != "B" {
		t.Error("ArchiveTypeIndicator does not validate")
	}
	if rd.TimesReturned != 0 {
		t.Error("TimesReturned does not validate")
	}
}

// TestParseReturnDetail validates parsing a ReturnDetail
func TestParseReturnDetail(t *testing.T) {
	var line = "31031300012             55588810000100000A04G201809051               2B0        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)

	if err := r.parseReturnDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0]

	if record.recordType != "31" {
		t.Errorf("RecordType Expected '31' got: %v", record.recordType)
	}
	if record.PayorBankRoutingNumberField() != "03130001" {
		t.Errorf("PayorBankRoutingNumber Expected '03130001' got: %v", record.PayorBankRoutingNumberField())
	}
	if record.PayorBankCheckDigitField() != "2" {
		t.Errorf("PayorBank Expected '2' got: %v", record.PayorBankCheckDigitField())
	}
	if record.OnUsField() != "             5558881" {
		t.Errorf("OnUs Expected '             5558881' got: %v", record.OnUsField())
	}
	if record.ItemAmountField() != "0000100000" {
		t.Errorf("ItemAmount Expected '0000100000' got: %v", record.ItemAmountField())
	}
	if record.ReturnReasonField() != "A" {
		t.Errorf("ReturnReason Expected 'A' got: %v", record.ReturnReasonField())
	}
	if record.AddendumCountField() != "04" {
		t.Errorf("AddendumCount Expected '04' got: %v", record.AddendumCountField())
	}
	if record.DocumentationTypeIndicatorField() != "G" {
		t.Errorf("DocumentationTypeIndicator Expected 'G' got: %v", record.DocumentationTypeIndicatorField())
	}
	if record.EceInstitutionItemSequenceNumberField() != "1              " {
		t.Errorf("EceInstitutionItemSequenceNumber Expected '1              ' got: %v", record.EceInstitutionItemSequenceNumberField())
	}
	if record.ExternalProcessingCodeField() != " " {
		t.Errorf("ExternalProcessingCode Expected ' ' got: %v", record.ExternalProcessingCodeField())
	}
	if record.ReturnNotificationIndicatorField() != "2" {
		t.Errorf("ReturnNotificationIndicator Expected '2' got: %v", record.ReturnNotificationIndicatorField())
	}
	if record.ArchiveTypeIndicatorField() != "B" {
		t.Errorf("ArchiveTypeIndicator Expected 'R' got: %v", record.ArchiveTypeIndicatorField())
	}
	if record.TimesReturnedField() != "0" {
		t.Errorf("TimesReturned Expected '0' got: %v", record.TimesReturnedField())
	}

}

// testReturnDetailString validates that a known parsed ReturnDetail can return to a string of the same value
func testReturnDetailString(t testing.TB) {
	var line = "31031300012             55588810000100000A04G201809051               2B0        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)

	if err := r.parseReturnDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestReturnDetailString tests validating that a known parsed ReturnDetail can return to a string of the
// same value
func TestReturnDetailString(t *testing.T) {
	testReturnDetailString(t)
}

// BenchmarkReturnDetailString benchmarks validating that a known parsed ReturnDetailAddendumB
// can return to a string of the same value
func BenchmarkReturnDetailString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReturnDetailString(b)
	}
}

// TestRDRecordType validation
func TestRDRecordType(t *testing.T) {
	rd := mockReturnDetail()
	rd.recordType = "00"
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDDocumentationTypeIndicator validation
func TestRDDocumentationTypeIndicator(t *testing.T) {
	rd := mockReturnDetail()
	rd.DocumentationTypeIndicator = "P"
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDDocumentationTypeIndicatorZ validation
func TestRDDocumentationTypeIndicatorZ(t *testing.T) {
	rd := mockReturnDetail()
	rd.DocumentationTypeIndicator = "Z"
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDReturnNotificationIndicator validation
func TestRDReturnNotificationIndicator(t *testing.T) {
	rd := mockReturnDetail()
	rd.ReturnNotificationIndicator = "0"
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnNotificationIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDArchiveTypeIndicator validation
func TestRDArchiveTypeIndicator(t *testing.T) {
	rd := mockReturnDetail()
	rd.ArchiveTypeIndicator = "W"
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ArchiveTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDArchiveTypeIndicatorWithValidationOption validation
func TestRDArchiveTypeIndicatorWithValidationOption(t *testing.T) {
	rd := mockReturnDetail()
	rd.ArchiveTypeIndicator = "W"
	if err := rd.Validate(&ValidateOpts{ArchiveTypeIndicator: false}); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestRDTimesReturned validation
func TestRDTimesReturned(t *testing.T) {
	rd := mockReturnDetail()
	rd.TimesReturned = 5
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TimesReturned" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestReturnReasonInvalid validation
func TestReturnReasonInvalid(t *testing.T) {
	rd := mockReturnDetail()
	rd.ReturnReason = "88"
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnReason" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestRDFIRecordType validation
func TestRDFIRecordType(t *testing.T) {
	rd := mockReturnDetail()
	rd.recordType = ""
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDFIPayorBankRoutingNumber validation
func TestRDFIPayorBankRoutingNumber(t *testing.T) {
	rd := mockReturnDetail()
	rd.PayorBankRoutingNumber = ""
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDFIPayorBankRoutingNumberZero validation
func TestRDFIPayorBankRoutingNumberZero(t *testing.T) {
	rd := mockReturnDetail()
	rd.PayorBankRoutingNumber = "00000000"
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDFIPayorBankCheckDigit validation
func TestRDFIPayorBankCheckDigit(t *testing.T) {
	rd := mockReturnDetail()
	rd.PayorBankCheckDigit = ""
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayorBankCheckDigit" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDFIReturnReason validation
func TestRDFIReturnReason(t *testing.T) {
	rd := mockReturnDetail()
	rd.ReturnReason = ""
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnReason" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDFIEceInstitutionItemSequenceNumber validation
func TestRDFIEceInstitutionItemSequenceNumber(t *testing.T) {
	rd := mockReturnDetail()
	rd.EceInstitutionItemSequenceNumber = "               "
	if err := rd.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EceInstitutionItemSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
