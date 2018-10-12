// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// mockUserGeneral creates a UserGeneral
func mockUserGeneral() *UserGeneral {
	ug := new(UserGeneral)
	ug.OwnerIdentifierIndicator = 3
	ug.OwnerIdentifier = "230918276"
	ug.OwnerIdentifierModifier = "ZZAB"
	ug.UserRecordFormatType = "000"
	ug.FormatTypeVersionLevel = "1"
	ug.LengthUserData = "0000038"
	ug.UserData = "This is a payment for your information"
	return ug
}
