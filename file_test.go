// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

// mockFile creates an imagecashletter file
func mockFile() *File {
	mockFile := NewFile()
	mockFile.SetHeader(mockFileHeader())
	clh := mockCashLetterHeader()
	mockCashLetter := NewCashLetter(clh)
	mockFile.AddCashLetter(mockCashLetter)
	if err := mockFile.Create(); err != nil {
		panic(err)
	}
	return mockFile
}

func TestFileCreate(t *testing.T) {
	file := mockFile()
	if err := file.Validate(); err != nil {
		t.Error("File does not validate and will break other tests: ", err)
	}
}

func TestFile__FileFromJSON(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("test", "testdata", "icl-valid.json"))
	if err != nil {
		t.Fatal(err)
	}

	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	if err := file.Validate(); err != nil {
		t.Fatal(err)
	}

	// error conditions

	if _, err := FileFromJSON(nil); err == nil {
		t.Error("expected error")
	}

	if _, err := FileFromJSON([]byte(`{invalid-json`)); err == nil {
		t.Error("expected error")
	}
}
