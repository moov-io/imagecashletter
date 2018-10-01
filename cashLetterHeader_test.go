// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
	"time"
)

// mockCashLetterHeader creates a CashLetterHeader
func mockCashLetterHeader() *CashLetterHeader {
	clh := NewCashLetterHeader()
	clh.CollectionTypeIndicator = "01"
	clh.DestinationRoutingNumber = "231380104"
	clh.ECEInstitutionRoutingNumber = "121042882"
	clh.CashLetterBusinessDate = time.Now()
	clh.CashLetterCreationDate = time.Now()
	clh.CashLetterCreationTime = time.Now()
	clh.CashLetterRecordTypeIndicator = "I"
	clh.CashLetterDocumentationTypeIndicator = "G"
	clh.CashLetterID = "A1"
	clh.OriginatorContactName = "Contact Name"
	clh.OriginatorContactPhoneNumber = "5558675552"
	clh.FedWorkType = ""
	clh.ReturnsIndicator = ""
	clh.UserField = ""
	return clh
}

// TestMockCashLetterHeader creates a CashLetterHeader
func TestMockCashLetterHeader(t *testing.T) {
	clh := mockCashLetterHeader()
	if err := clh.Validate(); err != nil {
		t.Error("mockCashLetterHeader does not validate and will break other tests: ", err)
	}
	if clh.recordType != "10" {
		t.Error("recordType does not validate")
	}
	if clh.CollectionTypeIndicator != "01" {
		t.Error("CollectionTypeIndicator does not validate")
	}
	if clh.DestinationRoutingNumber != "231380104" {
		t.Error("DestinationRoutingNumber does not validate")
	}
	if clh.ECEInstitutionRoutingNumber != "121042882" {
		t.Error("ECEInstitutionRoutingNumber does not validate")
	}
	if clh.CashLetterRecordTypeIndicator != "I" {
		t.Error("RecordTypeIndicator does not validate")
	}
	if clh.CashLetterDocumentationTypeIndicator != "G" {
		t.Error("DocumentationTypeIndicator does not validate")
	}
	if clh.CashLetterID != "A1" {
		t.Error("CashLetterID does not validate")
	}
	if clh.OriginatorContactName != "Contact Name" {
		t.Error("OriginatorContactName does not validate")
	}
	if clh.OriginatorContactPhoneNumber != "5558675552" {
		t.Error("OriginatorContactPhoneNumber does not validate")
	}
	if clh.FedWorkType != "" {
		t.Error("FedWorkType does not validate")
	}
	if clh.ReturnsIndicator != "" {
		t.Error("ReturnsIndicator does not validate")
	}
	if clh.UserField != "" {
		t.Error("UserField does not validate")
	}
	if clh.reserved != "" {
		t.Error("Reserved does not validate")
	}
}

// TestParseCashLetterHeader validates parsing a CashLetterHeader
func TestParseCashLetterHeader(t *testing.T) {
	var line = "100123138010412104288220180905201809051523IGA1      Contact Name  5558675552    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseCashLetterHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.CashLetterHeader

	if record.recordType != "10" {
		t.Errorf("RecordType Expected '10' got: %v", record.recordType)
	}
	if record.CollectionTypeIndicatorField() != "01" {
		t.Errorf("CollectionTypeIndicator Expected '01' got: %v", record.CollectionTypeIndicatorField())
	}
	if record.DestinationRoutingNumberField() != "231380104" {
		t.Errorf("DestinationRoutingNumber '231380104' got: %v", record.DestinationRoutingNumberField())
	}
	if record.ECEInstitutionRoutingNumberField() != "121042882" {
		t.Errorf("ECEInstitutionRoutingNumber Expected '121042882' got: %v", record.ECEInstitutionRoutingNumberField())
	}
	if record.CashLetterBusinessDateField() != "20180905" {
		t.Errorf("CashLetterBusinessDate Expected '20180905' got:'%v'", record.CashLetterBusinessDateField())
	}
	if record.CashLetterCreationDateField() != "20180905" {
		t.Errorf("CashLetterCreationDate Expected '20180905' got:'%v'", record.CashLetterCreationDateField())
	}
	if record.CashLetterCreationTimeField() != "1523" {
		t.Errorf("CashLetterCreationTime Expected '1523' got:'%v'", record.CashLetterCreationTimeField())
	}
	if record.CashLetterRecordTypeIndicatorField() != "I" {
		t.Errorf("CashLetterRecordTypeIndicator Expected 'I' got: %v", record.CashLetterRecordTypeIndicatorField())
	}
	if record.CashLetterDocumentationTypeIndicatorField() != "G" {
		t.Errorf("CashLetterDocumentationTypeIndicator Expected 'G' got:'%v'", record.CashLetterDocumentationTypeIndicatorField())
	}
	if record.CashLetterIDField() != "A1      " {
		t.Errorf("CashLetterID Expected 'A1      ' got:'%v'", record.CashLetterIDField())
	}
	if record.OriginatorContactNameField() != "Contact Name  " {
		t.Errorf("OriginatorContactName Expected 'Contact Name  ' got: '%v'", record.OriginatorContactNameField())
	}
	if record.OriginatorContactPhoneNumberField() != "5558675552" {
		t.Errorf("OriginatorContactPhoneNumber Expected '5558675552' got: '%v'", record.OriginatorContactPhoneNumberField())
	}
	if record.FedWorkTypeField() != " " {
		t.Errorf("FedWorkType Expected ' ' got:'%v'", record.FedWorkTypeField())
	}
	if record.ReturnsIndicatorField() != " " {
		t.Errorf("ReturnsIndicator ' ' got:'%v'", record.ReturnsIndicatorField())
	}
	if record.UserFieldField() != " " {
		t.Errorf("UserField Expected ' ' got:'%v'", record.UserFieldField())
	}
	if record.reservedField() != " " {
		t.Errorf("reserved Expected ' ' got:'%v'", record.reservedField())
	}
}

// testCLHString validates that a known parsed CashLetterHeader can return to a string of the same value
func testCLHString(t testing.TB) {
	var line = "100123138010412104288220180905201809051523IGA1      Contact Name  5558675552    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseCashLetterHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.CashLetterHeader
	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCLHString tests validating that a known parsed CashLetterHeader can return to a string of the same value
func TestCLHString(t *testing.T) {
	testCLHString(t)
}

// BenchmarkCLHString benchmarks validating that a known parsed CashLetterHeader
// can return to a string of the same value
func BenchmarkCLHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCLHString(b)
	}
}

// TestCLHRecordType validation
func TestCLHRecordType(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.recordType = "00"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCHCollectionTypeIndicator validation
func TestCHCollectionTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CollectionTypeIndicator = "87"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CollectionTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCashLetterRecordTypeIndicator validation
func TestCashLetterRecordTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterRecordTypeIndicator = "W"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CashLetterRecordTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCashLetterDocumentationTypeIndicator validation
func TestCashLetterDocumentationTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterDocumentationTypeIndicator = "WAZ"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CashLetterDocumentationTypeIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCashLetterID validation
func TestCashLetterID(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterID = "--"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CashLetterID" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestOriginatorContactName validation
func TestOriginatorContactName(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.OriginatorContactName = "®©"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OriginatorContactName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestOriginatorContactPhoneNumber validation
func TestOriginatorContactPhoneNumber(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.OriginatorContactPhoneNumber = "--"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OriginatorContactPhoneNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFedWorkType validation
func TestFedWorkType(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.FedWorkType = "--"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FedWorkType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestReturnsIndicator validation
func TestReturnsIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.ReturnsIndicator = "A"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnsIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCLHUserField validation
func TestCLHUserFieldI(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.UserField = "®©"
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCLHFieldInclusionRecordType validates FieldInclusion
func TestCLHFieldInclusionRecordType(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.recordType = ""
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCollectionTypeIndicator validates FieldInclusion
func TestFieldInclusionCollectionTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CollectionTypeIndicator = ""
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCashLetterRecordTypeIndicator validates FieldInclusion
func TestFieldInclusionCashLetterRecordTypeIndicator(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterRecordTypeIndicator = ""
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionDestinationRoutingNumber validates FieldInclusion
func TestFieldInclusionDestinationRoutingNumber(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.DestinationRoutingNumber = ""
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionECEInstitutionRoutingNumber validates FieldInclusion
func TestFieldInclusionECEInstitutionRoutingNumber(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.ECEInstitutionRoutingNumber = ""
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCashLetterBusinessDate validates FieldInclusion
func TestFieldInclusionCashLetterBusinessDate(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterBusinessDate = time.Time{}
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCashLetterCreationDate validates FieldInclusion
func TestFieldInclusionCashLetterCreationDate(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterCreationDate = time.Time{}
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCashLetterCreationTime validates FieldInclusion
func TestFieldInclusionCashLetterCreationTime(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterCreationTime = time.Time{}
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFieldInclusionCashLetterID validates FieldInclusion
func TestFieldInclusionCashLetterID(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.CashLetterID = ""
	if err := clh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
