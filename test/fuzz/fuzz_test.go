// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
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
		if len(contents) > 2<<20 {
			t.Skip()
		}

		file, err := imagecashletter.NewReader(strings.NewReader(contents)).Read()
		if err != nil {
			// Still run validate/write on whatever was partially parsed
			_ = file.Validate()
			_ = imagecashletter.NewWriter(io.Discard).Write(&file)
			return
		}

		checkFileHeader(file.Header)
		_ = file.Validate()

		var buf bytes.Buffer
		w := imagecashletter.NewWriter(&buf, imagecashletter.WriteVariableLineLengthOption())
		if err := w.Write(&file); err == nil && buf.Len() > 0 {
			// Round-trip
			_, _ = imagecashletter.NewReader(bytes.NewReader(buf.Bytes())).Read()
		}

		w = imagecashletter.NewWriter(io.Discard, imagecashletter.WriteEbcdicEncodingOption())
		_ = w.Write(&file)
	})
}

func FuzzReaderWriterJSON(f *testing.F) {
	populateCorpus(f, false)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		file, err := imagecashletter.FileFromJSON([]byte(contents))
		if err != nil || file == nil {
			return
		}

		checkFileHeader(file.Header)
		_ = file.Validate()

		w := imagecashletter.NewWriter(io.Discard, imagecashletter.WriteVariableLineLengthOption())
		_ = w.Write(file)

		w = imagecashletter.NewWriter(io.Discard, imagecashletter.WriteEbcdicEncodingOption())
		_ = w.Write(file)
	})
}

func populateCorpus(f *testing.F, icl bool) {
	f.Helper()

	f.Add("")
	f.Add("{}")

	err := filepath.Walk(filepath.Join("..", "..", "test", "testdata"), func(path string, info fs.FileInfo, err error) error {
		if err != nil || info == nil || info.IsDir() {
			return nil
		}
		if strings.Contains(path, string(filepath.Separator)+"fuzz"+string(filepath.Separator)) {
			return nil
		}

		ext := filepath.Ext(strings.ToLower(path))
		if ((ext == ".x937" || ext == "") && icl) || (ext == ".json" && !icl) {
			bs, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			if len(bs) > 512*1024 {
				return nil
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
