// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"sync"

	"github.com/moov-io/imagecashletter"
)

type ICLFileRepository interface {
	getFiles() ([]*imagecashletter.File, error)
	getFile(fileId string) (*imagecashletter.File, error)

	saveFile(file *imagecashletter.File) error
	deleteFile(fileId string) error
}

type memoryICLFileRepository struct {
	mu    sync.Mutex
	files map[string]*imagecashletter.File
}

func (r *memoryICLFileRepository) getFiles() ([]*imagecashletter.File, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var out []*imagecashletter.File
	for _, v := range r.files {
		f := *v
		out = append(out, &f)
	}
	return out, nil
}

func (r *memoryICLFileRepository) getFile(fileId string) (*imagecashletter.File, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.files {
		if r.files[i].ID == fileId {
			f := *r.files[i]
			return &f, nil
		}
	}
	return nil, nil
}

func (r *memoryICLFileRepository) saveFile(file *imagecashletter.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if file.ID == "" {
		return errors.New("empty ICL File ID")
	}
	r.files[file.ID] = file
	return nil
}

func (r *memoryICLFileRepository) deleteFile(fileId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if fileId == "" {
		return errors.New("empty ICL File Id")
	}

	delete(r.files, fileId)

	return nil
}
