// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
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
		t.Error("mockUserGeneral does not validate and will break other tests: ", err)
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

}
