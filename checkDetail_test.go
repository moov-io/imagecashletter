// Copyright 2018 The x9 Authors
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
	cd.EceInstitutionItemSequenceNumber = 1
	cd.DocumentationTypeIndicator = "G"
	cd.ReturnAcceptanceIndicator = "D"
	cd.MICRValidIndicator = 1
	cd.BOFDIndicator = "Y"
	cd.AddendumCount = 0
	cd.CorrectionIndicator = 0
	cd.ArchiveTypeIndicator = "B"
	return cd
}

// mockCheckDetailAddendum creates a CheckDetail
func mockCheckDetailAddendum() *CheckDetail {
	cd := NewCheckDetail()
	cd.AuxiliaryOnUs = "123456789"
	cd.ExternalProcessingCode = ""
	cd.PayorBankRoutingNumber = "03130001"
	cd.PayorBankCheckDigit = "2"
	cd.OnUs = "5558881"
	cd.ItemAmount = 100000 // 1000.00
	cd.EceInstitutionItemSequenceNumber = 1
	cd.DocumentationTypeIndicator = "G"
	cd.ReturnAcceptanceIndicator = "D"
	cd.MICRValidIndicator = 1
	cd.BOFDIndicator = "Y"
	cd.AddendumCount = 0
	cd.CorrectionIndicator = 0
	cd.ArchiveTypeIndicator = "B"
	cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
	cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
	cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())
	return cd
}


// testMockCheckDetail creates an ICL CheckDetail
func testMockCheckDetail(t testing.TB) {
	cd := mockCheckDetail()
	if err := cd.Validate(); err != nil {
		t.Error("mockBundleHeader does not validate and will break other tests: ", err)
	}
	if cd.recordType != "25" {
		t.Error("recordType does not validate and will break other tests")
	}
	if cd.AuxiliaryOnUs != "123456789" {
		t.Error("AuxiliaryOnUs does not validate and will break other tests")
	}
	if cd.ExternalProcessingCode != "" {
		t.Error("ExternalProcessingCode does not validate and will break other tests")
	}
	if cd.PayorBankRoutingNumber != "03130001" {
		t.Error("PayorBankRoutingNumber does not validate and will break other tests")
	}
	if cd.PayorBankCheckDigit != "2" {
		t.Error("PayorBankCheckDigit does not validate and will break other tests")
	}
	if cd.OnUs != "5558881" {
		t.Error("OnUs does not validate and will break other tests")
	}
	if cd.ItemAmount != 100000 {
		t.Error("ItemAmount does not validate and will break other tests")
	}
	if cd.EceInstitutionItemSequenceNumber != 1 {
		t.Error("EceInstitutionItemSequenceNumber does not validate and will break other tests")
	}
	if cd.DocumentationTypeIndicator != "G" {
		t.Error("DocumentationTypeIndicator does not validate and will break other tests")
	}
	if cd.ReturnAcceptanceIndicator != "D" {
		t.Error("ReturnAcceptanceIndicator does not validate and will break other tests")
	}
	if cd.MICRValidIndicator != 1 {
		t.Error("MICRValidIndicator does not validate and will break other tests")
	}
	if cd.BOFDIndicator != "Y" {
		t.Error("BOFDIndicator does not validate and will break other tests")
	}
	if cd.AddendumCount != 0 {
		t.Error("AddendumCount does not validate and will break other tests")
	}
	if cd.CorrectionIndicator != 0 {
		t.Error("CorrectionIndicator does not validate and will break other tests")
	}
	if cd.ArchiveTypeIndicator != "B" {
		t.Error("ArchiveTypeIndicator does not validate and will break other tests")
	}
}

// TestMockCheckDetail tests creating an ICL CheckDetail
func TestMockCheckDetail(t *testing.T) {
	testMockCheckDetail(t)
}

// BenchmarkMockCheckDetail benchmarks creating an ICL CheckDetail
func BenchmarkMockCheckDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCheckDetail(b)
	}
}

// parseCheckDetail validates parsing a CheckDetail
func parseCheckDetail(t testing.TB) {
	var line = "25      123456789 031300012             55588810000100000000000000000001GD1Y000B"
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))

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
	if record.EceInstitutionItemSequenceNumberField() != "000000000000001" {
		t.Errorf("EceInstitutionItemSequenceNumber Expected '000000000000001' got:'%v'",
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
	if record.AddendumCountField() != "00" {
		t.Errorf("AddendumCount Expected 'Y' got:'%v'", record.AddendumCountField())
	}
	if record.CorrectionIndicatorField() != "0" {
		t.Errorf("CorrectionIndicator Expected '0' got:'%v'", record.CorrectionIndicatorField())
	}
	if record.ArchiveTypeIndicatorField() != "B" {
		t.Errorf("ArchiveTypeIndicator Expected 'B' got:'%v'", record.ArchiveTypeIndicatorField())
	}
}

// TestParseCheckDetail test validates parsing a CheckDetail
func TestParseCheckDetail(t *testing.T) {
	parseCheckDetail(t)
}

// BenchmarkParseCheckDetail benchmark validates parsing a CheckDetail
func BenchmarkParseCheckDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseCheckDetail(b)
	}
}

// testCDString validates that a known parsed CheckDetail can return to a string of the same value
func testCDString(t testing.TB) {
	var line = "25      123456789 031300012             55588810000100000000000000000001GD1Y000B"
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	r.currentCashLetter.AddBundle(NewBundle(bh))
	r.addCurrentBundle(NewBundle(bh))
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
