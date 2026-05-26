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
func mockFile(t *testing.T) *File {
	t.Helper()

	f := NewFile()
	f.SetHeader(mockFileHeader())
	clh := mockCashLetterHeader()
	mockCashLetter := NewCashLetter(clh)
	mockCashLetter.CashLetterControl = mockCashLetterControl()
	f.AddCashLetter(mockCashLetter)

	require.NoError(t, f.Create())

	return f
}

func TestFileCreate(t *testing.T) {
	file := mockFile(t)
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

func TestFileFromJSONWithOpts(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "icl-valid.json"))
	require.NoError(t, err)

	// Strict should work on valid
	file, err := FileFromJSONWithOpts(bs, nil)
	require.NoError(t, err)
	require.NoError(t, file.Validate())

	// With SkipAll should also succeed (and set opts on result)
	opts := &ValidateOpts{SkipAll: true}
	file, err = FileFromJSONWithOpts(bs, opts)
	require.NoError(t, err)
	require.NoError(t, file.Validate()) // still no error
	require.True(t, file.validateOpts != nil && file.validateOpts.SkipAll)
}

func TestFileCreateSkipAllLenient(t *testing.T) {
	// Under SkipAll, Create should succeed for degenerate/empty structures
	// that would fail strict validation (useful for partial archived data handling)
	f := NewFile()
	f.SetHeader(mockFileHeader())
	f.SetValidation(&ValidateOpts{SkipAll: true})

	// No CashLetters, no Bundles etc. -- would fail without SkipAll
	err := f.Create()
	require.NoError(t, err)

	// Also Validate is lenient
	require.NoError(t, f.Validate())
}
