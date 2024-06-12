// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"os"
)

const FRBCompatibilityMode = "FRB_COMPATIBILITY_MODE"

// Determine if FRB (Federal Reserve Bank) compatibility mode is enabled
func IsFRBCompatibilityModeEnabled() bool {
	_, ok := os.LookupEnv(FRBCompatibilityMode)
	return ok
}
