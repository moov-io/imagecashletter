// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMockBundleChecks creates a Bundle of checks
func TestFRBCompatibilityMode(t *testing.T) {
	os.Setenv(FRBCompatibilityMode, "")
	assert.Equal(t, IsFRBCompatibilityModeEnabled(), true)
	os.Unsetenv(FRBCompatibilityMode)
	assert.Equal(t, IsFRBCompatibilityModeEnabled(), false)
}
