// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import "testing"

// mockFile creates an X9 file
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