// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"strings"
	"testing"
	"time"
)

// mockUserPayeeEndorsement creates a UserPayeeEndorsement
func mockUserPayeeEndorsement() *UserPayeeEndorsement {
	upe := NewUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 3
	upe.OwnerIdentifier = "230918276"
	upe.OwnerIdentifierModifier = "ZZ1"
	upe.UserRecordFormatType = "001"
	upe.FormatTypeVersionLevel = "1"
	upe.LengthUserData = "0000290"
	upe.PayeeName = "Payee Name"

	upe.EndorsementDate = time.Now()
	upe.BankRoutingNumber = "121042882"
	upe.BankAccountNumber = "123456888"
	upe.CustomerIdentifier = "A234A"
	upe.CustomerContactInformation = "Home"
	upe.StoreMerchantProcessingSiteNumber = "12345678"
	upe.InternalControlSequenceNumber = "ZB17262ZB"
	upe.Time = time.Now()
	upe.OperatorName = "ZJK"
	upe.OperatorNumber = "12345"
	upe.ManagerName = "ZBK"
	upe.ManagerNumber = "12345"
	upe.EquipmentNumber = "123456789012345"
	upe.EndorsementIndicator = 1
	upe.UserField = ""
	return upe
}

// TestMockUserPayeeEndorsement creates a UserPayeeEndorsement
func TestMockUserPayeeEndorsement(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	if err := upe.Validate(); err != nil {
		t.Error("mockUserPayeeEndorsement does not validate and will break other tests: ", err)
	}
	if upe.recordType != "68" {
		t.Error("recordType does not validate")
	}
	if upe.OwnerIdentifierIndicator != 3 {
		t.Error("OwnerIdentifierIndicator does not validate")
	}
	if upe.OwnerIdentifier != "230918276" {
		t.Error("OwnerIdentifier does not validate")
	}
	if upe.OwnerIdentifierModifier != "ZZ1" {
		t.Error("OwnerIdentifierModifier does not validate")
	}
	if upe.UserRecordFormatType != "001" {
		t.Error("UserRecordFormatType does not validate")
	}
	if upe.FormatTypeVersionLevel != "1" {
		t.Error("FormatTypeVersionLevel does not validate")
	}
	if upe.LengthUserData != "0000290" {
		t.Error("LengthUserData does not validate")
	}

	_ = additionalUPEFields(upe, t)
}

func additionalUPEFields(upe *UserPayeeEndorsement, t *testing.T) string {
	if upe.PayeeName != "Payee Name" {
		t.Error("PayeeName does not validate")
	}
	if upe.BankRoutingNumber != "121042882" {
		t.Error("BankRoutingNumber does not validate")
	}
	if upe.BankAccountNumber != "123456888" {
		t.Error("BankAccountNumber does not validate")
	}
	if upe.CustomerIdentifier != "A234A" {
		t.Error("CustomerIdentifier does not validate")
	}
	if upe.CustomerContactInformation != "Home" {
		t.Error("CustomerContactInformation does not validate")
	}
	if upe.StoreMerchantProcessingSiteNumber != "12345678" {
		t.Error("StoreMerchantProcessingSiteNumber does not validate")
	}
	if upe.InternalControlSequenceNumber != "ZB17262ZB" {
		t.Error("InternalControlSequenceNumber does not validate")
	}
	if upe.OperatorName != "ZJK" {
		t.Error("OperatorName does not validate")
	}
	if upe.OperatorNumber != "12345" {
		t.Error("OperatorNumber does not validate")
	}
	if upe.ManagerName != "ZBK" {
		t.Error("ManagerName does not validate")
	}
	if upe.ManagerNumber != "12345" {
		t.Error("ManagerNumber does not validate")
	}
	if upe.EquipmentNumber != "123456789012345" {
		t.Error("EquipmentNumber does not validate")
	}
	if upe.EndorsementIndicator != 1 {
		t.Error("EndorsementIndicator does not validate")
	}
	if upe.UserField != "" {
		t.Error("UserField does not validate")
	}
	return ""
}

// TestUPEString validation
func TestUPEString(t *testing.T) {
	line := "683230918276ZZ1                 0011  0000290Payee Name                                        20181015121042882123456888           A234A               Home                                              12345678ZB17262ZB                2222ZJK                           12345ZBK                           123451234567890123451          "
	upe := NewUserPayeeEndorsement()
	r := NewReader(strings.NewReader(line))
	r.line = line
	upe.Parse(r.line)

	if upe.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestUPEParse validation
func TestUPEParse(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	if err := upe.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	line := upe.String()
	r := NewReader(strings.NewReader(line))
	r.line = line

	upe.Parse(r.line)

	if err := upe.Validate(); err == nil {
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestUPERecordType validation
func TestUPERecordType(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.recordType = "00"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEUserRecordFormatType validation
func TestUPEUserRecordFormatType(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.UserRecordFormatType = "001"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserRecordFormatType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEOwnerIdentifierIndicator validation
func TestUPEOwnerIdentifierIndicator(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 9
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifierIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEOwnerIdentifierModifier validation
func TestUPEOwnerIdentifierModifier(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierModifier = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifierModifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEUserRecordFormatTypeChar validation
func TestUPEUserRecordFormatTypeChar(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.UserRecordFormatType = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserRecordFormatType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEFormatTypeVersionLevel validation
func TestUPEFormatTypeVersionLevel(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.FormatTypeVersionLevel = "W"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FormatTypeVersionLevel" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPELengthUserData validation
func TestUPELengthUserData(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.LengthUserData = "W"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "LengthUserData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEOwnerIdentifierIndicatorZero validation
func TestUPEOwnerIdentifierIndicatorZero(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 0
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEOwnerIdentifierIndicator123 validation
func TestUPEOwnerIdentifierIndicator123(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 1
	upe.OwnerIdentifier = "W"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEOwnerIdentifierIndicatorFour validation
func TestUPEOwnerIdentifierIndicatorFour(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OwnerIdentifierIndicator = 4
	upe.OwnerIdentifier = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEPayeeName validation
func TestUPEPayeeName(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.PayeeName = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PayeeName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEBankRoutingNumber validation
func TestUPEBankRoutingNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.BankRoutingNumber = "W"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEBankAccountNumber validation
func TestUPEBankAccountNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.BankAccountNumber = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BankAccountNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPECustomerIdentifier validation
func TestUPECustomerIdentifier(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.CustomerIdentifier = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CustomerIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPECustomerContactInformation validation
func TestUPECustomerContactInformation(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.CustomerContactInformation = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CustomerContactInformation" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEStoreMerchantProcessingSiteNumber validation
func TestUPEStoreMerchantProcessingSiteNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.StoreMerchantProcessingSiteNumber = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "StoreMerchantProcessingSiteNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEInternalControlSequenceNumber validation
func TestUPEInternalControlSequenceNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.InternalControlSequenceNumber = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "InternalControlSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEOperatorName validation
func TestUPEOperatorName(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OperatorName = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OperatorName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEOperatorNumber validation
func TestUPEOperatorNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.OperatorNumber = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OperatorNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEManagerName validation
func TestUPEManagerName(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.ManagerName = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ManagerName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEManagerNumber validation
func TestUPEManagerNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.ManagerNumber = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ManagerNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEEquipmentNumber validation
func TestUPEEquipmentNumber(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.EquipmentNumber = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EquipmentNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEEndorsementIndicator validation
func TestUPEEndorsementIndicator(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.EndorsementIndicator = 7
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsementIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEUserField validation
func TestUPEUserField(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.UserField = "®©"
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestUPEFIRecordType validation
func TestUPEFIRecordType(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.recordType = ""
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEFIUserRecordFormatType validation
func TestUPEFIUserRecordFormatType(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.UserRecordFormatType = ""
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserRecordFormatType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEFIFormatTypeVersionLevel validation
func TestUPEFIFormatTypeVersionLevel(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.FormatTypeVersionLevel = ""
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FormatTypeVersionLevel" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUPEFILengthUserData validation
func TestUPEFILengthUserData(t *testing.T) {
	upe := mockUserPayeeEndorsement()
	upe.LengthUserData = ""
	if err := upe.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "LengthUserData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
