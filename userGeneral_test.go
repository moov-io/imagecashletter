// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"strings"
	"testing"
)

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

// TestUGString validation
func TestUGString(t *testing.T) {
	line := "683230918276ZZ1                 0001  0000038This is a payment for your information"
	ug := mockUserGeneral()
	if err := ug.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if ug.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestUGParse validation
func TestUGParse(t *testing.T) {
	ug := mockUserGeneral()
	if err := ug.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	line := ug.String()
	r := NewReader(strings.NewReader(line))
	r.line = line
	ug.Parse(r.line)

	if err := ug.Validate(); err == nil {
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestUGRecordType validation
func TestUGRecordType(t *testing.T) {
	ug := mockUserGeneral()
	ug.recordType = "00"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGUserRecordFormatType validation
func TestUGUserRecordFormatType(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserRecordFormatType = "001"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserRecordFormatType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGOwnerIdentifierIndicator validation
func TestUGOwnerIdentifierIndicator(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierIndicator = 9
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifierIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGOwnerIdentifierModifier validation
func TestUGOwnerIdentifierModifier(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierModifier = "®©"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifierModifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGUserRecordFormatTypeChar validation
func TestUGUserRecordFormatTypeChar(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserRecordFormatType = "®©"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserRecordFormatType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGFormatTypeVersionLevel validation
func TestUGFormatTypeVersionLevel(t *testing.T) {
	ug := mockUserGeneral()
	ug.FormatTypeVersionLevel = "W"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FormatTypeVersionLevel" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGLengthUserData validation
func TestUGLengthUserData(t *testing.T) {
	ug := mockUserGeneral()
	ug.LengthUserData = "W"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "LengthUserData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGUserData validation
func TestUGUserData(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserData = "A®©X"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGOwnerIdentifierIndicatorZero validation
func TestUGOwnerIdentifierIndicatorZero(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierIndicator = 0
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGOwnerIdentifierIndicator123 validation
func TestUGOwnerIdentifierIndicator123(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierIndicator = 1
	ug.OwnerIdentifier = "W"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGOwnerIdentifierIndicatorFour validation
func TestUGOwnerIdentifierIndicatorFour(t *testing.T) {
	ug := mockUserGeneral()
	ug.OwnerIdentifierIndicator = 4
	ug.OwnerIdentifier = "®©"
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OwnerIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestUGFIRecordType validation
func TestUGFIRecordType(t *testing.T) {
	ug := mockUserGeneral()
	ug.recordType = ""
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGFIUserRecordFormatType validation
func TestUGFIUserRecordFormatType(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserRecordFormatType = ""
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserRecordFormatType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGFIFormatTypeVersionLevel validation
func TestUGFIFormatTypeVersionLevel(t *testing.T) {
	ug := mockUserGeneral()
	ug.FormatTypeVersionLevel = ""
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FormatTypeVersionLevel" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGFILengthUserData validation
func TestUGFILengthUserData(t *testing.T) {
	ug := mockUserGeneral()
	ug.LengthUserData = ""
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "LengthUserData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUGFIUserData validation
func TestUGFIUserData(t *testing.T) {
	ug := mockUserGeneral()
	ug.UserData = ""
	if err := ug.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
