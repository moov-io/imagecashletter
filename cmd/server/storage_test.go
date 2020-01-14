// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/imagecashletter"
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
	if err != nil || len(files) != 0 {
		t.Errorf("files=%#v error=%v", files, err)
	}

	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	f.ID = base.ID()

	if err := repo.saveFile(f); err != nil {
		t.Fatal(err)
	}

	files, err = repo.getFiles()
	if err != nil || len(files) != 1 {
		t.Errorf("files=%#v error=%v", files, err)
	}

	file, err := repo.getFile(f.ID)
	if err != nil {
		t.Error(err)
	}
	if file.ID != f.ID {
		t.Errorf("file mis-match")
	}

	if err := repo.deleteFile(f.ID); err != nil {
		t.Error(err)
	}
	files, err = repo.getFiles()
	if err != nil || len(files) != 0 {
		t.Errorf("files=%#v error=%v", files, err)
	}
}
