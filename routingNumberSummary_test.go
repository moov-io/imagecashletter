// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "testing"

// mockRoutingNumberSummary creates a RoutingNumberSummary
func mockRoutingNumberSummary() *RoutingNumberSummary {
	rns := NewRoutingNumberSummary()
	rns.CashLetterRoutingNumber = "231380104"
	rns.RoutingNumberTotalAmount = 100000
	rns.RoutingNumberItemCount = 1
	rns.UserField = ""
	return rns
}

// TestRoutingNumberSummary creates a ReturnRoutingNumberSummary
func TestRoutingNumberSummary(t *testing.T) {
	rns := mockRoutingNumberSummary()
	if err := rns.Validate(); err != nil {
		t.Error("mockRoutingNumberSummary does not validate and will break other tests: ", err)
	}
}
