// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/imagecashletter"
)

func FuzzReaderWriterICL(f *testing.F) {
	populateCorpus(f, true)

	f.Fuzz(func(t *testing.T, contents string) {
		file, _ := imagecashletter.NewReader(strings.NewReader(contents)).Read()

		checkFileHeader(file.Header)
		file.Validate()

		w := imagecashletter.NewWriter(io.Discard, imagecashletter.WriteVariableLineLengthOption())
		w.Write(&file)

		w = imagecashletter.NewWriter(io.Discard, imagecashletter.WriteEbcdicEncodingOption())
		w.Write(&file)
	})
}

func FuzzReaderWriterJSON(f *testing.F) {
	populateCorpus(f, false)

	f.Fuzz(func(t *testing.T, contents string) {
		file, _ := imagecashletter.FileFromJSON([]byte(contents))
		if file == nil {
			return
		}

		checkFileHeader(file.Header)
		file.Validate()

		w := imagecashletter.NewWriter(io.Discard, imagecashletter.WriteVariableLineLengthOption())
		w.Write(file)

		w = imagecashletter.NewWriter(io.Discard, imagecashletter.WriteEbcdicEncodingOption())
		w.Write(file)
	})
}

func populateCorpus(f *testing.F, icl bool) {
	f.Helper()

	err := filepath.Walk(filepath.Join("..", "..", "test", "testdata"), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(strings.ToLower(path))
		if ((ext == ".x937" || ext == "") && icl) || (ext == ".json" && !icl) {
			bs, err := os.ReadFile(path)
			if err != nil {
				f.Fatal(err)
			}
			f.Add(string(bs))
		}
		return nil
	})
	if err != nil {
		f.Fatal(err)
	}
}

func checkFileHeader(h imagecashletter.FileHeader) int {
	if h.ImmediateDestination != "" || h.ImmediateOrigin != "" {
		return 1
	}
	if !h.FileCreationDate.IsZero() || !h.FileCreationTime.IsZero() {
		return 1
	}
	if h.ImmediateDestinationName != "" || h.ImmediateOriginName != "" {
		return 1
	}
	if h.CountryCode != "" {
		return 1
	}
	return 0
}
