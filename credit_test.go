// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	require.NoError(t, ci.Validate(), "mockCredit does not validate and will break other tests")
	assert.Equal(t, "61", ci.recordType)
	assert.Equal(t, "010910999940910", ci.AuxiliaryOnUs)
	assert.Equal(t, "", ci.ExternalProcessingCode)
	assert.Equal(t, "999920060", ci.PayorBankRoutingNumber)
	assert.Equal(t, "50920060509383521210", ci.CreditAccountNumberOnUs)
	assert.Equal(t, 102088, ci.ItemAmount)
	assert.Equal(t, "               ", ci.ECEInstitutionItemSequenceNumber)
	assert.Equal(t, "G", ci.DocumentationTypeIndicator)
	assert.Equal(t, "1", ci.AccountTypeCode)
	assert.Equal(t, "3", ci.SourceWorkCode)
	assert.Equal(t, " ", ci.WorkType)
	assert.Equal(t, " ", ci.DebitCreditIndicator)
}

func TestCreditCrash(t *testing.T) {
	cr := &Credit{}
	cr.Parse(`61010910999940910 999920060509200605093835212100000102088               G13     `)
	assert.Equal(t, "G", cr.DocumentationTypeIndicator)
}

func TestParseCredit(t *testing.T) {
	var line = "61010910999940910 99992006050920060509383521210000010208812345          G13     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	cr := mockCredit()
	r.currentCashLetter.AddCredit(cr)
	require.NoError(t, r.parseCredit())
	require.Len(t, r.currentCashLetter.Credits, 2)

	record := r.currentCashLetter.GetCredits()[0]
	assert.Equal(t, "61", record.recordType)
	assert.Equal(t, "010910999940910", record.AuxiliaryOnUs)
	assert.Equal(t, "", record.ExternalProcessingCode)
	assert.Equal(t, "999920060", record.PayorBankRoutingNumber)
	assert.Equal(t, "50920060509383521210", record.CreditAccountNumberOnUs)
	assert.Equal(t, 102088, record.ItemAmount)
	assert.Equal(t, "               ", record.ECEInstitutionItemSequenceNumber)
	assert.Equal(t, "G", record.DocumentationTypeIndicator)
	assert.Equal(t, "1", record.AccountTypeCode)
	assert.Equal(t, "3", record.SourceWorkCode)
	assert.Equal(t, " ", record.WorkType)
	assert.Equal(t, " ", record.DebitCreditIndicator)

	assert.Equal(t, "12345          ", r.currentCashLetter.GetCredits()[1].ECEInstitutionItemSequenceNumber)
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
	require.NoError(t, r.parseCredit())
	require.Len(t, r.currentCashLetter.Credits, 2)
	record := r.currentCashLetter.GetCredits()[0]
	assert.Equal(t, line, record.String())
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
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "recordType", err.FieldName)
}

// TestCRDocumentationTypeIndicator validation
func TestCRDocumentationTypeIndicator(t *testing.T) {
	ci := mockCredit()
	ci.DocumentationTypeIndicator = "P"
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "DocumentationTypeIndicator", err.FieldName)
}

// TestCRDocumentationTypeIndicatorZ validation
func TestCRDocumentationTypeIndicatorZ(t *testing.T) {
	ci := mockCredit()
	ci.DocumentationTypeIndicator = "Z"
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "DocumentationTypeIndicator", err.FieldName)
}

// TestCRDocumentationTypeIndicatorM validation
func TestCRDocumentationTypeIndicatorM(t *testing.T) {
	ci := mockCredit()
	ci.DocumentationTypeIndicator = "M"
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "DocumentationTypeIndicator", err.FieldName)
}

// TestCRSourceWorkCode validation
func TestCRSourceWorkCode(t *testing.T) {
	ci := mockCredit()
	ci.SourceWorkCode = "99"
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "SourceWorkCode", err.FieldName)
}

// Field Inclusion

// TestCRFIRecordType validation
func TestCRFIRecordType(t *testing.T) {
	ci := mockCredit()
	ci.recordType = ""
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "recordType", err.FieldName)
}

// TestCRPayorBankRoutingNumber validation
func TestCRPayorBankRoutingNumber(t *testing.T) {
	ci := mockCredit()
	ci.PayorBankRoutingNumber = "000000000"
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "PayorBankRoutingNumber", err.FieldName)
}

// TestCRCreditAccountNumberOnUs validation
func TestCRCreditAccountNumberOnUs(t *testing.T) {
	ci := mockCredit()
	ci.CreditAccountNumberOnUs = ""
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "CreditAccountNumberOnUs", err.FieldName)
}

// TestCRItemAmount validation
func TestCRItemAmount(t *testing.T) {
	ci := mockCredit()
	ci.ItemAmount = 0
	var err *FieldError
	require.ErrorAs(t, ci.Validate(), &err)
	assert.Equal(t, "ItemAmount", err.FieldName)
}

func TestCredit_ECEInstitutionItemSequenceNumber(t *testing.T) {
	// valid ECEInstitutionItemSequenceNumber
	cr := mockCredit()
	cr.ECEInstitutionItemSequenceNumber = "123456789012345"
	require.NoError(t, cr.Validate())

	// empty ECEInstitutionItemSequenceNumber
	cr = mockCredit()
	cr.ECEInstitutionItemSequenceNumber = ""
	require.NoError(t, cr.Validate())

	// invalid ECEInstitutionItemSequenceNumber
	cr = mockCredit()
	cr.ECEInstitutionItemSequenceNumber = "®©"
	validateErr := cr.Validate()
	require.Error(t, validateErr)
	var err *FieldError
	require.ErrorAs(t, validateErr, &err)
	require.Equal(t, "ECEInstitutionItemSequenceNumber", err.FieldName)
}
