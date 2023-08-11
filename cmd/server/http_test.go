// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTP_cleanMetricsPath(t *testing.T) {
	require.Equal(t, "v1-customers-companies", cleanMetricsPath("/v1/customers/companies/1234"))
	require.Equal(t, "v1-customers-ping", cleanMetricsPath("/v1/customers/ping"))
	require.Equal(t, "v1-customers-customers", cleanMetricsPath("/v1/customers/customers/19636f90bc95779e2488b0f7a45c4b68958a2ddd"))

	// A value which looks like moov/base.ID, but is off by one character (last letter)
	require.Equal(t, "v1-customers-customers-19636f90bc95779e2488b0f7a45c4b68958a2ddz", cleanMetricsPath("/v1/customers/customers/19636f90bc95779e2488b0f7a45c4b68958a2ddz"))
}
