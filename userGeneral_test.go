// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "testing"

// mockUserGeneral creates a UserGeneral
func mockUserGeneral() *UserGeneral {
	ug := NewUserGeneral()
	ug.OwnerIdentifierIndicator = 3
	ug.OwnerIdentifier = "230918276"
	ug.OwnerIdentifierModifier = "ZZ1"
	ug.UserRecordFormatType = "000"
	ug.FormatTypeVersionLevel = "1"
	ug.LengthUserData = "0000038"
	ug.UserData = "This is a payment for your information"
	return ug
}

// TestMockUserGeneral creates a UserGeneral
func TestMockUserGeneral(t *testing.T) {
	ug := mockUserGeneral()
	if err := ug.Validate(); err != nil {
		t.Error("mockUserGeneral does not validate and will break other tests: ", err)
	}
	if ug.recordType != "68" {
		t.Error("recordType does not validate")
	}
	if ug.OwnerIdentifierIndicator != 3 {
		t.Error("OwnerIdentifierIndicator does not validate")
	}
	if ug.OwnerIdentifier != "230918276" {
		t.Error("OwnerIdentifier does not validate")
	}
	if ug.OwnerIdentifierModifier != "ZZ1" {
		t.Error("OwnerIdentifierModifier does not validate")
	}
	if ug.UserRecordFormatType != "000" {
		t.Error("UserRecordFormatType does not validate")
	}
	if ug.FormatTypeVersionLevel != "1" {
		t.Error("FormatTypeVersionLevel does not validate")
	}
	if ug.LengthUserData != "0000038" {
		t.Error("LengthUserData does not validate")
	}
	if ug.UserData != "This is a payment for your information" {
		t.Error("UserData does not validate")
	}
}
