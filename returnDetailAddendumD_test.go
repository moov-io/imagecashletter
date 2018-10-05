// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"log"
	"strings"
	"testing"
	"time"
)

// mockReturnDetailAddendumD creates a ReturnDetailAddendumD
func mockReturnDetailAddendumD() ReturnDetailAddendumD {
	rdAddendumD := NewReturnDetailAddendumD()
	rdAddendumD.RecordNumber = 1
	rdAddendumD.EndorsingBankRoutingNumber = "121042882"
	rdAddendumD.BOFDEndorsementBusinessDate = time.Now()
	rdAddendumD.EndorsingBankItemSequenceNumber = "1              "
	rdAddendumD.TruncationIndicator = "Y"
	rdAddendumD.EndorsingBankConversionIndicator = "1"
	rdAddendumD.EndorsingBankCorrectionIndicator = 0
	rdAddendumD.ReturnReason = "A"
	rdAddendumD.UserField = ""
	rdAddendumD.EndorsingBankIdentifier = 0
	return rdAddendumD
}

// TestMockReturnDetailAddendumD creates a ReturnDetailAddendumD
func TestMockReturnDetailAddendumD(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	if err := rdAddendumD.Validate(); err != nil {
		t.Error("MockReturnDetailAddendumD does not validate and will break other tests: ", err)
	}
	if rdAddendumD.recordType != "35" {
		t.Error("recordType does not validate")
	}
	if rdAddendumD.RecordNumber != 1 {
		t.Error("RecordNumber does not validate")
	}
	if rdAddendumD.EndorsingBankRoutingNumber != "121042882" {
		t.Error("EndorsingBankRoutingNumber does not validate")
	}
	if rdAddendumD.EndorsingBankItemSequenceNumber != "1              " {
		t.Error("EndorsingBankItemSequenceNumber does not validate")
	}
	if rdAddendumD.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate")
	}
	if rdAddendumD.ReturnReason != "A" {
		t.Error("ReturnReason does not validate")
	}
	if rdAddendumD.EndorsingBankConversionIndicator != "1" {
		t.Error("EndorsingBankConversionIndicator does not validate")
	}
	if rdAddendumD.EndorsingBankCorrectionIndicator != 0 {
		t.Error("EndorsingBankCorrectionIndicator does not validate")
	}
	if rdAddendumD.UserField != "" {
		t.Error("UserField does not validate")
	}
	if rdAddendumD.EndorsingBankIdentifier != 0 {
		t.Error("EndorsingBankIdentifier does not validate")
	}
}

// TestParseReturnDetailAddendumD validates parsing a ReturnDetailAddendumD
func TestParseReturnDetailAddendumD(t *testing.T) {
	var line = "3501121042882201809051              Y10A                   0                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	cd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(cd)

	if err := r.parseReturnDetailAddendumD(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumD[0]

	if record.recordType != "35" {
		t.Errorf("RecordType Expected '35' got: %v", record.recordType)
	}
	if record.RecordNumberField() != "01" {
		t.Errorf("RecordNumber Expected '01' got: %v", record.RecordNumberField())
	}

	if record.EndorsingBankRoutingNumberField() != "121042882" {
		t.Errorf("EndorsingBankRoutingNumbeRoutingNumber Expected '121042882' got: %v",
			record.EndorsingBankRoutingNumberField())
	}
	if record.BOFDEndorsementBusinessDateField() != "20180905" {
		t.Errorf("BOFDEndorsementBusinessDate Expected '20180905' got: %v",
			record.BOFDEndorsementBusinessDateField())
	}
	if record.EndorsingBankItemSequenceNumberField() != "1              " {
		t.Errorf("EndorsingBankItemSequenceNumber Expected '1              ' got: %v",
			record.EndorsingBankItemSequenceNumberField())
	}
	if record.TruncationIndicatorField() != "Y" {
		t.Errorf("TruncationIndicator Expected 'Y' got: %v", record.TruncationIndicatorField())
	}
	if record.EndorsingBankConversionIndicatorField() != "1" {
		t.Errorf("EndorsingBankConversionIndicator  Expected '1' got: %v", record.EndorsingBankConversionIndicatorField())
	}
	if record.EndorsingBankCorrectionIndicatorField() != "0" {
		t.Errorf("EndorsingBankCorrectionIndicator Expected '0' got: %v", record.EndorsingBankCorrectionIndicatorField())
	}
	if record.ReturnReasonField() != "A" {
		t.Errorf("ReturnReason  Expected 'A' got: %v", record.ReturnReasonField())
	}
	if record.UserFieldField() != "                   " {
		t.Errorf("UserField Expected '                   ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != "                    " {
		t.Errorf("reserved Expected '                    ' got: %v", record.reservedField())
	}
}

// testRDAddendumDString validates that a known parsed ReturnDetailAddendumD can return to a string of the same value
func testRDAddendumDString(t testing.TB) {
	var line = "3501121042882201809051              Y10A                   0                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	rb := NewBundle(bh)
	r.currentCashLetter.AddBundle(rb)
	r.addCurrentBundle(rb)
	rd := mockReturnDetail()
	r.currentCashLetter.currentBundle.AddReturnDetail(rd)

	if err := r.parseReturnDetailAddendumD(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetReturns()[0].ReturnDetailAddendumD[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestRDAddendumDString tests validating that a known parsed ReturnDetailAddendumD can return to a string of the
// same value
func TestRDAddendumDString(t *testing.T) {
	testRDAddendumDString(t)
}

// BenchmarkRDAddendumDString benchmarks validating that a known parsed ReturnDetailAddendumD
// can return to a string of the same value
func BenchmarkRDAddendumDString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDAddendumDString(b)
	}
}

// TestRDAddendumDRecordType validation
func TestRDAddendumDRecordType(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.recordType = "00"
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDReturnLocationRoutingNumber validation
func TestRDAddendumDReturnLocationRoutingNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankRoutingNumber = "X"
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDTruncationIndicator validation
func TestRDAddendumDTruncationIndicator(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.TruncationIndicator = "A"
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TruncationIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDBOFDConversionIndicator validation
func TestRDAddendumDBOFDConversionIndicator(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankConversionIndicator = "99"
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankConversionIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDBOFDCorrectionIndicator validation
func TestRDAddendumDBOFDCorrectionIndicator(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankCorrectionIndicator = 10
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankCorrectionIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDReturnReason validation
func TestRDAddendumDReturnReason(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.ReturnReason = "--"
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnReason" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDUserField validation
func TestRDAddendumDUserField(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.UserField = "®©"
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDEndorsingBankIdentifier validation
func TestRDAddendumDEndorsingBankIdentifier(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankIdentifier = 9
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestRDAddendumDFIRecordType validation
func TestRDAddendumDFIRecordType(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.recordType = ""
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDFIRecordNumber validation
func TestRDAddendumDFIRecordNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.RecordNumber = 0
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "RecordNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDFIReturnLocationRoutingNumber validation
func TestRDAddendumDFIReturnLocationRoutingNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankRoutingNumber = ""
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDFIBOFDEndorsementDate validation
func TestRDAddendumDFIBOFDEndorsementDate(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.BOFDEndorsementBusinessDate = time.Time{}
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDEndorsementBusinessDate" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDFIBOFDItemSequenceNumber validation
func TestRDAddendumDFIBOFDItemSequenceNumber(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.EndorsingBankItemSequenceNumber = "               "
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankItemSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestRDAddendumDFITruncationIndicator validation
func TestRDAddendumDFITruncationIndicator(t *testing.T) {
	rdAddendumD := mockReturnDetailAddendumD()
	rdAddendumD.TruncationIndicator = ""
	if err := rdAddendumD.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TruncationIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
