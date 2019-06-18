// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestHTTP__cleanMetricsPath(t *testing.T) {
	if v := cleanMetricsPath("/v1/customers/companies/1234"); v != "v1-customers-companies" {
		t.Errorf("got %q", v)
	}
	if v := cleanMetricsPath("/v1/customers/ping"); v != "v1-customers-ping" {
		t.Errorf("got %q", v)
	}
	if v := cleanMetricsPath("/v1/customers/customers/19636f90bc95779e2488b0f7a45c4b68958a2ddd"); v != "v1-customers-customers" {
		t.Errorf("got %q", v)
	}
	// A value which looks like moov/base.ID, but is off by one character (last letter)
	if v := cleanMetricsPath("/v1/customers/customers/19636f90bc95779e2488b0f7a45c4b68958a2ddz"); v != "v1-customers-customers-19636f90bc95779e2488b0f7a45c4b68958a2ddz" {
		t.Errorf("got %q", v)
	}
}
