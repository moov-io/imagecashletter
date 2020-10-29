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

	"github.com/stretchr/testify/require"

	"github.com/moov-io/base"
	"github.com/moov-io/imagecashletter"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func TestFileId(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	if v := getFileId(w, req); v != "" {
		t.Errorf("unexpected fileId=%s", v)
	}
	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected HTTP status: %d", w.Code)
	}
}

func TestCashLetterId(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	if v := getCashLetterId(w, req); v != "" {
		t.Errorf("unexpected fileId=%s", v)
	}
	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected HTTP status: %d", w.Code)
	}
}

func TestFiles__getFiles(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files", nil)

	repo := &testICLFileRepository{
		file: &imagecashletter.File{
			ID: base.ID(),
		},
	}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
	var files []*imagecashletter.File
	if err := json.NewDecoder(w.Body).Decode(&files); err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Errorf("unexpected %d ICL files: %#v", len(files), files)
	}

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}

func readFile(filename string) (*imagecashletter.File, error) {
	fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", filename))
	if err != nil {
		return nil, err
	}
	f, err := imagecashletter.NewReader(fd).Read()
	return &f, err
}

func TestFiles__createFile(t *testing.T) {
	fd, _ := os.Open(filepath.Join("..", "..", "test", "testdata", "valid-ascii.x937"))

	w := httptest.NewRecorder()
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

func TestFiles__createFileJSON(t *testing.T) {
	fd, _ := os.Open(filepath.Join("..", "..", "test", "testdata", "icl-valid.json"))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", fd)
	req.Header.Set("Content-Type", "application/json")

	repo := &testICLFileRepository{}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusCreated {
		t.Errorf("bogus HTTP status: %d: %s", w.Code, w.Body.String())
	}
	var resp imagecashletter.File
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Header.CountryCode != "US" {
		t.Errorf("CountryCode=%s", resp.Header.CountryCode)
	}

	// error case

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/files/create", strings.NewReader("{invalid-json"))
	req.Header.Set("content-type", "application/json")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %s", w.Code, w.Body.String())
	}
}

func TestFiles__getFile(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/foo", nil)

	repo := &testICLFileRepository{
		file: &imagecashletter.File{
			ID: base.ID(),
		},
	}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	var file imagecashletter.File
	if err := json.NewDecoder(w.Body).Decode(&file); err != nil {
		t.Fatal(err)
	}
	if file.ID == "" {
		t.Errorf("unexpected ICL file: %#v", file)
	}

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}

func TestFiles__updateFileHeader(t *testing.T) {
	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	f.ID = base.ID()

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(f.Header); err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s", f.ID), &buf)

	repo := &testICLFileRepository{
		file: &imagecashletter.File{
			ID: f.ID, // create a file without FileHeader so it's updated
		},
	}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusCreated {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	if repo.file.Header.CountryCode != f.Header.CountryCode {
		t.Errorf("repo.file.Header.CountryCode=%s expected=%s", repo.file.Header.CountryCode, f.Header.CountryCode)
	}
}

func TestFiles__deleteFile(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/files/foo", nil)

	repo := &testICLFileRepository{}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}

func TestFiles__getFileContents(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/foo/contents", nil)

	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	repo := &testICLFileRepository{file: f}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	if v := w.Header().Get("Content-Type"); v != "text/plain" {
		t.Errorf("unexpected Content-Type: %s", v)
	}

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}

func TestFiles__validateFile(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/foo/validate", nil)

	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	repo := &testICLFileRepository{file: f}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), `"{\"error\": null}"`) {
		t.Errorf("unexpected body: %v", w.Body.String())
	}

	// make the file invalid
	repo.file.Header = imagecashletter.NewFileHeader()
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}

	// repository error case
	repo.err = errors.New("bad error")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}

func TestFiles__addCashLetterToFile(t *testing.T) {
	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	cashLetter := f.CashLetters[0]
	f.CashLetters = nil
	repo := &testICLFileRepository{file: f}

	// encode our CashLetter into JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(cashLetter); err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/foo/cashLetters", &buf)

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	var out imagecashletter.File
	if err := json.NewDecoder(w.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	if len(out.CashLetters) != 1 {
		t.Errorf("CashLetters: %#v", out.CashLetters)
	}

	// error case
	repo.err = errors.New("bad error")
	if err := json.NewEncoder(&buf).Encode(cashLetter); err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}

func TestFiles__removeCashLetterFromFile(t *testing.T) {
	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	repo := &testICLFileRepository{file: f}

	cashLetterId := base.ID()
	repo.file.CashLetters[0].ID = cashLetterId

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/files/foo/cashLetters/%s", cashLetterId), nil)

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}
