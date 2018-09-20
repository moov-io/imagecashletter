// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "time"

// mockReturnDetail creates a ReturnDetail
func mockReturnDetail() *ReturnDetail {
	rd := NewReturnDetail()
	rd.PayorBankRoutingNumber = "03130001"
	rd.PayorBankCheckDigit = "2"
	rd.OnUs = "5558881"
	rd.ItemAmount = 100000
	rd.ReturnReason = "A"
	rd.AddendumCount = 3
	rd.DocumentationTypeIndicator = "G"
	rd.ForwardBundleDate = time.Now()
	rd.EceInstitutionItemSequenceNumber = "1              "
	rd.ExternalProcessingCode = ""
	rd.ReturnNotificationIndicator = 2
	rd.ArchiveTypeIndicator = "B"
	rd.TimesReturned = 0
	return rd
}
