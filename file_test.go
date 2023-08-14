// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFile creates an imagecashletter file
func mockFile() *File {
	f := NewFile()
	f.SetHeader(mockFileHeader())
	clh := mockCashLetterHeader()
	mockCashLetter := NewCashLetter(clh)
	mockCashLetter.CashLetterControl = mockCashLetterControl()
	f.AddCashLetter(mockCashLetter)
	if err := f.Create(); err != nil {
		panic(err)
	}
	return f
}

func TestFileCreate(t *testing.T) {
	file := mockFile()
	require.NoError(t, file.Validate())
}

func TestFile_FileFromJSON(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "icl-valid.json"))
	require.NoError(t, err)

	file, err := FileFromJSON(bs)
	require.NoError(t, err)
	require.NoError(t, file.Validate())

	// error conditions

	f, err := FileFromJSON(nil)
	require.Nil(t, f)
	require.Error(t, err)

	f, err = FileFromJSON([]byte(`{invalid-json`))
	require.Nil(t, f)
	require.Error(t, err)
}
