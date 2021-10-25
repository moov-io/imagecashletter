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

// mockCheckDetailAddendumA creates a CheckDetailAddendumA
func mockCheckDetailAddendumA() CheckDetailAddendumA {
	cdAddendumA := NewCheckDetailAddendumA()
	cdAddendumA.RecordNumber = 1
	cdAddendumA.ReturnLocationRoutingNumber = "121042882"
	cdAddendumA.BOFDEndorsementDate = time.Now()
	cdAddendumA.BOFDItemSequenceNumber = "1              "
	cdAddendumA.BOFDAccountNumber = "938383"
	cdAddendumA.BOFDBranchCode = "01"
	cdAddendumA.PayeeName = "Test Payee"
	cdAddendumA.TruncationIndicator = "Y"
	cdAddendumA.BOFDConversionIndicator = "1"
	cdAddendumA.BOFDCorrectionIndicator = 0
	cdAddendumA.UserField = ""
	return cdAddendumA
}

func TestCheckDetailAddendumParseErr(t *testing.T) {
	var c CheckDetailAddendumA
	c.Parse("asdshfaksjs")
	if c.RecordNumber != 0 {
		t.Errorf("c.RecordNumber=%d", c.RecordNumber)
	}
}

// TestMockCheckDetailAddendumA creates a CheckDetailAddendumA
func TestMockCheckDetailAddendumA(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	if err := cdAddendumA.Validate(); err != nil {
		t.Error("MockCheckDetailAddendumA does not validate and will break other tests: ", err)
	}
	if cdAddendumA.recordType != "26" {
		t.Error("recordType does not validate")
	}
	if cdAddendumA.RecordNumber != 1 {
		t.Error("RecordNumber does not validate")
	}
	if cdAddendumA.ReturnLocationRoutingNumber != "121042882" {
		t.Error("ReturnLocationRoutingNumber does not validate")
	}
	if cdAddendumA.BOFDAccountNumber != "938383" {
		t.Error("BOFDAccountNumber does not validate")
	}
	if cdAddendumA.BOFDBranchCode != "01" {
		t.Error("BOFDBranchCode does not validate")
	}
	if cdAddendumA.PayeeName != "Test Payee" {
		t.Error("PayeeName does not validate")
	}
	if cdAddendumA.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate")
	}
	if cdAddendumA.BOFDConversionIndicator != "1" {
		t.Error("BOFDConversionIndicator does not validate")
	}
	if cdAddendumA.BOFDCorrectionIndicator != 0 {
		t.Error("BOFDCorrectionIndicator does not validate")
	}
	if cdAddendumA.UserField != "" {
		t.Error("UserField does not validate")
	}
}

// TestParseCheckDetailAddendumA validates parsing a CheckDetailAddendumA
func TestParseCheckDetailAddendumA(t *testing.T) {
	var line = "261121042882201809051              938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumA(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumA[0]

	if record.recordType != "26" {
		t.Errorf("RecordType Expected '26' got: %v", record.recordType)
	}
	if record.RecordNumberField() != "1" {
		t.Errorf("RecordNumber Expected '1' got: %v", record.RecordNumberField())
	}
	if record.ReturnLocationRoutingNumberField() != "121042882" {
		t.Errorf("ReturnLocationRoutingNumber Expected '121042882' got: %v", record.ReturnLocationRoutingNumberField())
	}
	if record.BOFDEndorsementDateField() != "20180905" {
		t.Errorf("BOFDEndorsementDate Expected '20180905' got: %v", record.BOFDEndorsementDateField())
	}
	if record.BOFDItemSequenceNumberField() != "1              " {
		t.Errorf("BOFDItemSequenceNumber Expected '1               ' got: %v", record.BOFDItemSequenceNumberField())
	}
	if record.BOFDAccountNumberField() != "938383            " {
		t.Errorf("BOFDAccountNumber Expected '938383            ' got: %v", record.BOFDAccountNumberField())
	}
	if record.BOFDBranchCodeField() != "01   " {
		t.Errorf("BOFDBranchCode Expected '01   ' got: %v", record.BOFDBranchCodeField())
	}
	if record.PayeeNameField() != "Test Payee     " {
		t.Errorf("PayeeName Expected 'Test Payee     ' got: %v", record.PayeeNameField())
	}
	if record.TruncationIndicatorField() != "Y" {
		t.Errorf("TruncationIndicator Expected 'Y' got: %v", record.TruncationIndicatorField())
	}
	if record.BOFDConversionIndicatorField() != "1" {
		t.Errorf("BOFDConversionIndicator Expected '1' got: %v", record.BOFDConversionIndicatorField())
	}
	if record.BOFDCorrectionIndicatorField() != "0" {
		t.Errorf("BOFDCorrectionIndicator Expected '0' got: %v", record.BOFDCorrectionIndicatorField())
	}
	if record.UserFieldField() != " " {
		t.Errorf("UserField Expected ' ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != "   " {
		t.Errorf("reserved Expected '   ' got: %v", record.reservedField())
	}
}

// testCDAddendumAString validates that a known parsed CheckDetailAddendumA can return to a string of the same value
func testCDAddendumAString(t testing.TB) {
	var line = "261121042882201809051              938383            01   Test Payee     Y10    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumA(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumA[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCDAddendumAString tests validating that a known parsed CheckDetailAddendumA can return to a string of the
// same value
func TestCDAddendumAString(t *testing.T) {
	testCDAddendumAString(t)
}

// BenchmarkCDAddendumAString benchmarks validating that a known parsed CheckDetailAddendumA
// can return to a string of the same value
func BenchmarkCDAddendumAString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumAString(b)
	}
}

// TestCDAddendumARecordType validation
func TestCDAddendumARecordType(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.recordType = "00"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAReturnLocationRoutingNumber validation
func TestCDAddendumAReturnLocationRoutingNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = "X"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnLocationRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumABOFDAccountNumber validation
func TestCDAddendumABOFDAccountNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDAccountNumber = "®©"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDAccountNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumABOFDBranchCode validation
func TestCDAddendumABOFDBranchCode(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDBranchCode = "®©"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDBranchCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAPayeeName validation
func TestCDAddendumAPayeeName(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.PayeeName = "®©"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayeeName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumATruncationIndicator validation
func TestCDAddendumATruncationIndicator(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.TruncationIndicator = "A"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TruncationIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumABOFDConversionIndicator validation
func TestCDAddendumABOFDConversionIndicator(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDConversionIndicator = "99"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDConversionIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumABOFDCorrectionIndicator validation
func TestCDAddendumABOFDCorrectionIndicator(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDCorrectionIndicator = 10
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDCorrectionIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAUserField validation
func TestCDAddendumAUserField(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.UserField = "®©"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestCDAddendumAFIRecordType validation
func TestCDAddendumAFIRecordType(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.recordType = ""
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAFIRecordNumber validation
func TestCDAddendumAFIRecordNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.RecordNumber = 0
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "RecordNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAFIReturnLocationRoutingNumber validation
func TestCDAddendumAFIReturnLocationRoutingNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = ""
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnLocationRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAFIReturnLocationRoutingNumberZero validation
func TestCDAddendumAFIReturnLocationRoutingNumberZero(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = "000000000"
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnLocationRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAFIBOFDEndorsementDate validation
func TestCDAddendumAFIBOFDEndorsementDate(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDEndorsementDate = time.Time{}
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDEndorsementDate" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAFIBOFDItemSequenceNumber validation
func TestCDAddendumAFIBOFDItemSequenceNumber(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.BOFDItemSequenceNumber = "               "
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDItemSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumAFITruncationIndicator validation
func TestCDAddendumAFITruncationIndicator(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.TruncationIndicator = ""
	if err := cdAddendumA.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TruncationIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// End FieldInclusion

// TestAlphaFieldTrim validation
func TestAlphaFieldTrim(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.PayeeName = "Payee Name James Steel"
	if len(cdAddendumA.PayeeNameField()) > 15 {
		t.Error("Payee field is greater than max")
	}

}

// TestStringFieldTrim validation
func TestStringFieldTrim(t *testing.T) {
	cdAddendumA := mockCheckDetailAddendumA()
	cdAddendumA.ReturnLocationRoutingNumber = "12345678912345"
	if len(cdAddendumA.ReturnLocationRoutingNumberField()) > 15 {
		t.Error("ReturnLocationRoutingNumber field is greater than max")
	}

}
