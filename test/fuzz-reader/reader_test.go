// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzzreader

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCorpusSymlinks(t *testing.T) {
	// avoid symbolic link error on windows
	if runtime.GOOS == "windows" {
		t.Skip()
	}

	fds, err := os.ReadDir("corpus")
	require.NoError(t, err)
	require.NotEmpty(t, fds)

	for i := range fds {
		info, err := fds[i].Info()
		require.NoError(t, err)

		require.NotEqualValuesf(t, 0, info.Mode()&os.ModeSymlink, "%s isn't a symlink, move outside corpus/ and symlink into directory", fds[i].Name())
		path, err := os.Readlink(filepath.Join("corpus", fds[i].Name()))
		require.NoError(t, err)
		_, err = os.Stat(filepath.Join("corpus", path))
		require.NoError(t, err)
	}
}
