// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestImageCashLetter_ReadCrashers will attempt to parse files which have previously been reported
// as crashing. These files are typically generated via fuzzing, but might also be reported by users.
func TestImageCashLetter_ReadCrashers(t *testing.T) {
	root := filepath.Join("test", "testdata", "crashers")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if (err != nil && !errors.Is(err, filepath.SkipDir)) || info == nil || info.IsDir() {
			return nil // Ignore SkipDir and directories
		}
		if strings.HasSuffix(path, ".output") {
			return nil // go-fuzz makes these which contain the panic's trace
		}

		fd, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("problem opening %s: %v", path, err)
		}

		// Read out test file with multiple option patterns and ensure we don't panic
		require.NotPanics(t, func() {
			_, _ = NewReader(fd).Read()
			_, _ = NewReader(fd, ReadVariableLineLengthOption()).Read()
		})

		if testing.Verbose() {
			t.Logf("read and parsed %s", fd.Name())
		}

		return nil
	})
	require.NoError(t, err)
}
