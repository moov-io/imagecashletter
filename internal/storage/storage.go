// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package storage

import (
	"errors"
	"sync"

	"github.com/moov-io/imagecashletter"
)

type ICLFileRepository interface {
	GetFiles() ([]*imagecashletter.File, error)
	GetFile(fileId string) (*imagecashletter.File, error)

	SaveFile(file *imagecashletter.File) error
	DeleteFile(fileId string) error
}

type memoryICLFileRepository struct {
	mu    sync.Mutex
	files map[string]*imagecashletter.File
}

func NewInMemoryRepo() ICLFileRepository {
	return &memoryICLFileRepository{
		files: make(map[string]*imagecashletter.File),
	}
}

func (r *memoryICLFileRepository) GetFiles() ([]*imagecashletter.File, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var out []*imagecashletter.File
	for _, v := range r.files {
		f := *v
		out = append(out, &f)
	}
	return out, nil
}

func (r *memoryICLFileRepository) GetFile(fileId string) (*imagecashletter.File, error) {
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

func (r *memoryICLFileRepository) SaveFile(file *imagecashletter.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if file.ID == "" {
		return errors.New("empty ICL File ID")
	}
	r.files[file.ID] = file
	return nil
}

func (r *memoryICLFileRepository) DeleteFile(fileId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if fileId == "" {
		return errors.New("empty ICL File Id")
	}

	delete(r.files, fileId)

	return nil
}
