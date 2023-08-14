// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/imagecashletter"
	"github.com/stretchr/testify/require"
)

type testICLFileRepository struct {
	err error

	file *imagecashletter.File
}

func (r *testICLFileRepository) getFiles() ([]*imagecashletter.File, error) {
	if r.err != nil {
		return nil, r.err
	}
	return []*imagecashletter.File{r.file}, nil
}

func (r *testICLFileRepository) getFile(fileId string) (*imagecashletter.File, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.file, nil
}

func (r *testICLFileRepository) saveFile(file *imagecashletter.File) error {
	if r.err == nil { // only persist if we're not error'ing
		r.file = file
	}
	return r.err
}

func (r *testICLFileRepository) deleteFile(fileId string) error {
	return r.err
}

func TestMemoryStorage(t *testing.T) {
	repo := &memoryICLFileRepository{
		files: make(map[string]*imagecashletter.File),
	}

	files, err := repo.getFiles()
	require.NoError(t, err)
	require.Equal(t, 0, len(files))

	f := readFile(t, "BNK20180905121042882-A.icl")
	f.ID = base.ID()
	require.NoError(t, repo.saveFile(f))

	files, err = repo.getFiles()
	require.NoError(t, err)
	require.Equal(t, 1, len(files))

	file, err := repo.getFile(f.ID)
	require.NoError(t, err)
	require.Equal(t, f.ID, file.ID)

	require.NoError(t, repo.deleteFile(f.ID))
	files, err = repo.getFiles()
	require.NoError(t, err)
	require.Equal(t, 0, len(files))
}
