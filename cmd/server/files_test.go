// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/base"
	"github.com/moov-io/imagecashletter"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
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
	req := httptest.NewRequest("GET", "/files", nil)
	repo := &testICLFileRepository{
		file: &imagecashletter.File{
			ID: base.ID(),
		},
	}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("returns one file", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusOK, w.Code, w.Body)
		var files []*imagecashletter.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&files))
		require.Len(t, files, 1)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_createFile(t *testing.T) {
	w := httptest.NewRecorder()
	fd, _ := os.Open(filepath.Join("..", "..", "test", "testdata", "valid-ebcdic.x937"))
	req := httptest.NewRequest("POST", "/files/create", fd)
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusCreated, w.Code, w.Body.String())

	var resp imagecashletter.File
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	require.Equal(t, "Wave Money", resp.Header.ImmediateDestinationName)

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
}

func TestFiles_createFileJSON(t *testing.T) {
	w := httptest.NewRecorder()
	fd, _ := os.Open(filepath.Join("..", "..", "test", "testdata", "icl-valid.json"))
	req := httptest.NewRequest("POST", "/files/create", fd)
	req.Header.Set("Content-Type", "application/json")
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusCreated, w.Code, w.Body)
	var resp imagecashletter.File
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "US", resp.Header.CountryCode)

	// error case
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/files/create", strings.NewReader("{invalid-json"))
	req.Header.Set("content-type", "application/json")

	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
}

func TestFiles_getFile(t *testing.T) {
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	req := httptest.NewRequest("GET", "/files/foo", nil)

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.file = &imagecashletter.File{
			ID: base.ID(),
		}
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusOK, w.Code, w.Body)
		var file imagecashletter.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&file))
		assert.NotEmpty(t, file.ID)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_updateFileHeader(t *testing.T) {
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	f := readFile(t, "BNK20180905121042882-A.icl")
	f.ID = base.ID()

	t.Run("file not found", func(t *testing.T) {
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(f.Header))
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s", f.ID), &buf)
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(f.Header))
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s", f.ID), &buf)
		repo.file = &imagecashletter.File{
			ID: f.ID, // create a file without FileHeader so it's updated
		}
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusCreated, w.Code, w.Body)
		assert.Equal(t, repo.file.Header.CountryCode, f.Header.CountryCode)
	})
}

func TestFiles_deleteFile(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/files/foo", nil)
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		repo.file = &imagecashletter.File{}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusOK, w.Code, w.Body)
	})

	t.Run("repo error", func(t *testing.T) {
		repo.err = errors.New("bad error")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})

}

func TestFiles_getFileContents(t *testing.T) {
	req := httptest.NewRequest("GET", "/files/foo/contents", nil)
	router := mux.NewRouter()
	repo := &testICLFileRepository{}
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		w := httptest.NewRecorder()
		f := readFile(t, "BNK20180905121042882-A.icl")
		repo.file = f
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusOK, w.Code, w.Body)
		assert.Equal(t, "text/plain", w.Header().Get("Content-Type"), "unexpected content type")
	})

	t.Run("repo error", func(t *testing.T) {
		repo.err = errors.New("bad error")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})

}

func TestFiles_validateFile(t *testing.T) {
	req := httptest.NewRequest("GET", "/files/foo/validate", nil)
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	f := readFile(t, "BNK20180905121042882-A.icl")

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})

	t.Run("valid file", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.file = f
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusOK, w.Code, w.Body)
		assert.Contains(t, w.Body.String(), `"{\"error\": null}"`)
	})

	t.Run("invalid file", func(t *testing.T) {
		w := httptest.NewRecorder()
		// make the file invalid
		repo.file.Header = imagecashletter.NewFileHeader()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_addCashLetterToFile(t *testing.T) {
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	f := readFile(t, "BNK20180905121042882-A.icl")
	cashLetter := f.CashLetters[0]
	f.CashLetters = nil

	t.Run("file not found", func(t *testing.T) {
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(cashLetter))
		req := httptest.NewRequest("POST", "/files/foo/cashLetters", &buf)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(cashLetter))
		req := httptest.NewRequest("POST", "/files/foo/cashLetters", &buf)
		w := httptest.NewRecorder()
		repo.file = f
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusOK, w.Code, w.Body)
		var out imagecashletter.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&out))
		assert.Len(t, out.CashLetters, 1, "expected one cashLetter")
	})

	t.Run("repo error", func(t *testing.T) {
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(cashLetter))
		req := httptest.NewRequest("POST", "/files/foo/cashLetters", &buf)
		w := httptest.NewRecorder()
		repo.file = nil
		repo.err = errors.New("bad error")
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_removeCashLetterFromFile(t *testing.T) {
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	f := readFile(t, "BNK20180905121042882-A.icl")
	cashLetterId := base.ID()
	f.CashLetters[0].ID = cashLetterId
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/files/foo/cashLetters/%s", cashLetterId), nil)

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})

	t.Run("successful request", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.file = f
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusOK, w.Code, w.Body)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.file = nil
		repo.err = errors.New("bad error")
		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_createFile_Issue228(t *testing.T) {
	repo := &testICLFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	w := httptest.NewRecorder()
	fd, _ := os.Open(filepath.Join("..", "..", "test", "testdata", "issue228.json"))
	req := httptest.NewRequest("POST", "/files/create", fd)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	type apiError struct {
		Error string `json:"error"`
	}
	var wantErr apiError
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &wantErr))
	require.Contains(t, wantErr.Error, "CashLetterControl record is mandatory")
}

func readFile(t *testing.T, filename string) *imagecashletter.File {
	t.Helper()

	fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", filename))
	require.NoError(t, err)
	f, err := imagecashletter.NewReader(fd, imagecashletter.ReadVariableLineLengthOption()).Read()
	require.NoError(t, err)
	return &f
}
