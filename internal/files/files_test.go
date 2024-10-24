// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package files

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gdamore/encoding"
	"github.com/moov-io/base"
	"github.com/moov-io/imagecashletter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileId(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	fileID := getFileId(w, req)

	assert.Empty(t, fileID)
	require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
}

func TestCashLetterId(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	cashLetterID := getCashLetterId(w, req)

	assert.Empty(t, cashLetterID)
	require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
}

func TestFiles_getFiles(t *testing.T) {
	repo := &testICLFileRepository{
		file: &imagecashletter.File{
			ID: base.ID(),
		},
	}
	env := newTestEnvironment(t, withRepo(repo))

	resp, files := env.listFiles(t)
	require.Equal(t, http.StatusOK, resp.Code, resp.Body)
	require.Len(t, files, 1)

	repo.err = errors.New("bad error")
	resp, _ = env.listFiles(t)
	require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
}

func TestFiles_determineBufferSize(t *testing.T) {
	env := t.Name()

	size := determineBufferSize(env, 10001)
	require.Equal(t, 10001, size)

	t.Setenv(env, "452181")

	size = determineBufferSize(env, 10001)
	require.Equal(t, 452181, size)
}

func TestFiles_createFile(t *testing.T) {
	repo := &testICLFileRepository{}
	env := newTestEnvironment(t, withRepo(repo))

	resp, file := env.createFile(t, "", openTestFile(t, "valid-ebcdic.x937"))
	require.Equal(t, http.StatusCreated, resp.Code, resp.Body)
	require.Equal(t, "Wave Money", file.Header.ImmediateDestinationName)

	// error case
	repo.err = errors.New("bad error")
	resp, _ = env.createFile(t, "", openTestFile(t, "valid-ebcdic.x937"))
	require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
}

func TestFiles_create_missingBundleControl(t *testing.T) {
	env := newTestEnvironment(t)

	wantBundleControl := imagecashletter.NewBundleControl()
	wantBundleControl.BundleItemsCount = 4
	wantBundleControl.BundleTotalAmount = 400000
	wantBundleControl.MICRValidTotalAmount = 400000
	wantBundleControl.BundleImagesCount = 4
	wantBundleControl.CreditTotalIndicator = 0

	resp, file := env.createFile(t, "application/json", openTestFile(t, "missing-bundle-control.json"))
	require.Equal(t, http.StatusCreated, resp.Code, resp.Body)
	require.Len(t, file.CashLetters, 1)
	require.Equal(t, "118507", file.CashLetters[0].CashLetterHeader.CashLetterID)
	require.Len(t, file.CashLetters[0].Bundles, 1)
	require.Len(t, file.CashLetters[0].Bundles[0].Checks, 4)
	require.Equal(t, *wantBundleControl, *file.CashLetters[0].Bundles[0].BundleControl)

	// GET the file
	resp, rawFile := env.getFileContents(t, file.ID)
	require.Equal(t, http.StatusOK, resp.Code, resp.Body)

	// POST the fetched file
	resp, newFile := env.createFile(t, "application/octet-stream", bytes.NewReader(rawFile))
	require.Equal(t, http.StatusCreated, resp.Code, resp.Body)
	require.Len(t, newFile.CashLetters, 1)
	require.Equal(t, "118507", newFile.CashLetters[0].CashLetterHeader.CashLetterID)
	require.Len(t, newFile.CashLetters[0].Bundles, 1)
	require.Len(t, newFile.CashLetters[0].Bundles[0].Checks, 4)
	require.Equal(t, *wantBundleControl, *newFile.CashLetters[0].Bundles[0].BundleControl)
}

func TestFiles_createFileJSON(t *testing.T) {
	env := newTestEnvironment(t)

	resp, file := env.createFile(t, "application/json", openTestFile(t, "icl-valid.json"))
	require.Equal(t, http.StatusCreated, resp.Code, resp.Body)
	require.Equal(t, "US", file.Header.CountryCode)

	// error case
	resp, _ = env.createFile(t, "application/json", strings.NewReader("{invalid-json"))

	require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
}

func TestFiles_getFile(t *testing.T) {
	fileID := base.ID()
	repo := &testICLFileRepository{
		file: &imagecashletter.File{
			ID: fileID,
		},
	}
	env := newTestEnvironment(t, withRepo(repo))

	t.Run("file not found", func(t *testing.T) {
		resp, _ := env.getFile(t, "foo")
		require.Equal(t, http.StatusNotFound, resp.Code, resp.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		resp, file := env.getFile(t, fileID)
		require.Equal(t, http.StatusOK, resp.Code, resp.Body)
		require.Equal(t, fileID, file.ID)
	})

	t.Run("repo error", func(t *testing.T) {
		defer func() {
			repo.err = nil
		}()
		repo.err = errors.New("bad error")
		resp, _ := env.getFile(t, fileID)
		require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
	})
}

func TestFiles_updateFileHeader(t *testing.T) {
	env := newTestEnvironment(t)

	f := parseTestFile(t, "BNK20180905121042882-A.icl")
	f.ID = base.ID()
	f.Header.UserField = "before"
	require.NoError(t, env.repo.SaveFile(f))

	t.Run("file not found", func(t *testing.T) {
		resp, _ := env.updateFileHeader(t, "foo", f.Header)

		require.Equal(t, http.StatusNotFound, resp.Code, resp.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		newHeader := f.Header
		newHeader.UserField = "after"

		resp, file := env.updateFileHeader(t, f.ID, newHeader)
		require.Equal(t, http.StatusCreated, resp.Code, resp.Body)
		assert.Equal(t, "after", file.Header.UserField)

		// check the repo as well
		updated, err := env.repo.GetFile(f.ID)
		require.NoError(t, err)
		assert.Equal(t, "after", updated.Header.UserField)
	})
}

func TestFiles_deleteFile(t *testing.T) {
	env := newTestEnvironment(t)

	f1 := base.ID()
	f2 := base.ID()
	require.NoError(t, env.repo.SaveFile(&imagecashletter.File{ID: f1}))
	require.NoError(t, env.repo.SaveFile(&imagecashletter.File{ID: f2}))

	t.Run("file not found", func(t *testing.T) {
		resp := env.deleteFile(t, "foo")

		require.Equal(t, http.StatusNotFound, resp.Code, resp.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		resp := env.deleteFile(t, f1)

		require.Equal(t, http.StatusOK, resp.Code, resp.Body)

		listResp, files := env.listFiles(t)
		require.Equal(t, http.StatusOK, listResp.Code, listResp.Body)
		require.Len(t, files, 1)
		require.Equal(t, f2, files[0].ID)
	})

	t.Run("repo error", func(t *testing.T) {
		repo := &testICLFileRepository{
			err: errors.New("bad error"),
		}
		mockEnv := newTestEnvironment(t, withRepo(repo))
		resp := mockEnv.deleteFile(t, f1)

		require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
	})

}

func TestFiles_getFileContents(t *testing.T) {
	env := newTestEnvironment(t)
	f := parseTestFile(t, "BNK20180905121042882-A.icl")
	f.ID = base.ID()
	f.Header.UserField = "user"
	require.NoError(t, env.repo.SaveFile(f))

	t.Run("file not found", func(t *testing.T) {
		resp, _ := env.getFileContents(t, "foo")

		require.Equal(t, http.StatusNotFound, resp.Code, resp.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		resp, file := env.getFileContents(t, f.ID)

		require.Equal(t, http.StatusOK, resp.Code, resp.Body)
		require.Equal(t, "text/plain", resp.Header().Get("Content-Type"), "unexpected content type")
		wantUserField, _ := encoding.EBCDIC.NewEncoder().String("user")
		require.Contains(t, string(file), wantUserField)
	})

	t.Run("repo error", func(t *testing.T) {
		repo := &testICLFileRepository{
			err: errors.New("bad error"),
		}
		mockEnv := newTestEnvironment(t, withRepo(repo))
		resp, _ := mockEnv.getFileContents(t, "foo")

		require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
	})
}

func TestFiles_validateFile(t *testing.T) {
	env := newTestEnvironment(t)
	f := parseTestFile(t, "BNK20180905121042882-A.icl")
	f.ID = base.ID()
	require.NoError(t, env.repo.SaveFile(f))

	t.Run("file not found", func(t *testing.T) {
		resp := env.validateFile(t, "foo")

		require.Equal(t, http.StatusNotFound, resp.Code, resp.Body)
	})

	t.Run("valid file", func(t *testing.T) {
		resp := env.validateFile(t, f.ID)

		require.Equal(t, http.StatusOK, resp.Code, resp.Body)
		assert.Contains(t, resp.Body.String(), `"{\"error\": null}"`)
	})

	t.Run("invalid file", func(t *testing.T) {
		invalidFile := *f
		invalidFile.ID = base.ID()
		invalidFile.Header = imagecashletter.NewFileHeader()
		require.NoError(t, env.repo.SaveFile(&invalidFile))
		resp := env.validateFile(t, invalidFile.ID)

		require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
		require.Contains(t, resp.Body.String(), "TestFileIndicator  is a mandatory field")
	})

	t.Run("repo error", func(t *testing.T) {
		repo := &testICLFileRepository{
			err: errors.New("bad error"),
		}
		mockEnv := newTestEnvironment(t, withRepo(repo))
		resp := mockEnv.validateFile(t, "foo")

		require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
	})
}

func TestFiles_addCashLetterToFile(t *testing.T) {
	env := newTestEnvironment(t)

	f := parseTestFile(t, "BNK20180905121042882-A.icl")
	cashLetter := f.CashLetters[0]
	f.CashLetters = nil
	f.ID = base.ID()
	require.NoError(t, env.repo.SaveFile(f))

	t.Run("file not found", func(t *testing.T) {
		resp, _ := env.addCashLetter(t, "foo", cashLetter)

		require.Equal(t, http.StatusNotFound, resp.Code, resp.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		resp, file := env.addCashLetter(t, f.ID, cashLetter)

		require.Equal(t, http.StatusOK, resp.Code, resp.Body)
		assert.Len(t, file.CashLetters, 1, "expected one cashLetter")
	})

	t.Run("repo error", func(t *testing.T) {
		repo := &testICLFileRepository{
			err: errors.New("bad error"),
		}
		mockEnv := newTestEnvironment(t, withRepo(repo))
		resp, _ := mockEnv.addCashLetter(t, "foo", cashLetter)

		require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
	})
}

func TestFiles_removeCashLetterFromFile(t *testing.T) {
	env := newTestEnvironment(t)

	f := parseTestFile(t, "BNK20180905121042882-A.icl")
	f.ID = base.ID()
	cashLetterId := base.ID()
	f.CashLetters[0].ID = cashLetterId
	require.NoError(t, env.repo.SaveFile(f))

	t.Run("file not found", func(t *testing.T) {
		resp := env.removeCashLetter(t, "foo", cashLetterId)

		require.Equal(t, http.StatusNotFound, resp.Code, resp.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		resp := env.removeCashLetter(t, f.ID, cashLetterId)

		require.Equal(t, http.StatusOK, resp.Code, resp.Body)
		file, err := env.repo.GetFile(f.ID)
		require.NoError(t, err)
		require.Len(t, file.CashLetters, 1)
		require.NotEqual(t, cashLetterId, file.CashLetters[0].ID)
	})

	t.Run("repo error", func(t *testing.T) {
		repo := &testICLFileRepository{
			err: errors.New("bad error"),
		}
		mockEnv := newTestEnvironment(t, withRepo(repo))
		resp := mockEnv.removeCashLetter(t, f.ID, cashLetterId)

		require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
	})
}

func TestFiles_createFile_Issue228(t *testing.T) {
	env := newTestEnvironment(t)

	resp, _ := env.createFile(t, "application/json", openTestFile(t, "issue228.json"))

	require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
	require.Contains(t, resp.Body.String(), "CashLetterControl record is mandatory")
}
