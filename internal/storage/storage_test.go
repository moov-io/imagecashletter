// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/imagecashletter"
	"github.com/stretchr/testify/require"
)

func TestMemoryStorage(t *testing.T) {
	repo := &memoryICLFileRepository{
		files: make(map[string]*imagecashletter.File),
	}

	files, err := repo.GetFiles()
	require.NoError(t, err)
	require.Equal(t, 0, len(files))

	f := readFile(t, "BNK20180905121042882-A.icl")
	f.ID = base.ID()
	require.NoError(t, repo.SaveFile(f))

	files, err = repo.GetFiles()
	require.NoError(t, err)
	require.Equal(t, 1, len(files))

	file, err := repo.GetFile(f.ID)
	require.NoError(t, err)
	require.Equal(t, f.ID, file.ID)

	require.NoError(t, repo.DeleteFile(f.ID))
	files, err = repo.GetFiles()
	require.NoError(t, err)
	require.Equal(t, 0, len(files))
}

func readFile(t *testing.T, filename string) *imagecashletter.File {
	t.Helper()

	fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", filename))
	require.NoError(t, err)
	f, err := imagecashletter.NewReader(fd, imagecashletter.ReadVariableLineLengthOption()).Read()
	require.NoError(t, err)
	return &f
}
