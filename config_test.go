// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMockBundleChecks creates a Bundle of checks
func TestFRBCompatibilityMode(t *testing.T) {
	t.Setenv(FRBCompatibilityMode, "false")
	assert.False(t, IsFRBCompatibilityModeEnabled())
	t.Setenv(FRBCompatibilityMode, "true")
	assert.True(t, IsFRBCompatibilityModeEnabled())
}
