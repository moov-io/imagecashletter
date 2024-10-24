package files

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
	"github.com/moov-io/imagecashletter"
	"github.com/moov-io/imagecashletter/internal/storage"
	"github.com/stretchr/testify/require"
)

type testEnvironment struct {
	router *mux.Router
	repo   storage.ICLFileRepository
}

type envOptionFunc func(*testEnvironment)

func withRepo(repo storage.ICLFileRepository) envOptionFunc {
	return func(env *testEnvironment) {
		env.repo = repo
	}
}

func newTestEnvironment(t *testing.T, opts ...envOptionFunc) *testEnvironment {
	t.Helper()

	env := &testEnvironment{
		router: mux.NewRouter(),
		repo:   storage.NewInMemoryRepo(),
	}

	for i := range opts {
		opts[i](env)
	}

	AppendRoutes(log.NewNopLogger(), env.router, env.repo)

	return env
}

func (env *testEnvironment) listFiles(t *testing.T) (*httptest.ResponseRecorder, []*imagecashletter.File) {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files", nil)
	env.router.ServeHTTP(w, req)
	w.Flush()

	var files []*imagecashletter.File
	if w.Code == http.StatusOK {
		require.NoError(t, json.NewDecoder(w.Body).Decode(&files))
	}

	return w, files
}

func (env *testEnvironment) getFile(t *testing.T, fileID string) (*httptest.ResponseRecorder, *imagecashletter.File) {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/"+fileID, nil)
	env.router.ServeHTTP(w, req)
	w.Flush()

	var file *imagecashletter.File
	if w.Code == http.StatusOK {
		require.NoError(t, json.NewDecoder(w.Body).Decode(&file))
	}

	return w, file
}

func (env *testEnvironment) getFileContents(t *testing.T, fileID string) (*httptest.ResponseRecorder, []byte) {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/"+fileID+"/contents", nil)
	env.router.ServeHTTP(w, req)
	w.Flush()

	return w, w.Body.Bytes()
}

func (env *testEnvironment) createFile(t *testing.T, contentType string, file io.Reader) (*httptest.ResponseRecorder, *imagecashletter.File) {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", file)
	req.Header.Set("Accept", "application/json")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	env.router.ServeHTTP(w, req)
	w.Flush()

	var resp *imagecashletter.File
	if w.Code == http.StatusCreated {
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	}

	return w, resp
}

func (env *testEnvironment) validateFile(t *testing.T, fileID string) *httptest.ResponseRecorder {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/"+fileID+"/validate", nil)
	env.router.ServeHTTP(w, req)
	w.Flush()

	return w
}

func (env *testEnvironment) updateFileHeader(t *testing.T, fileID string, header imagecashletter.FileHeader) (*httptest.ResponseRecorder, *imagecashletter.File) {
	t.Helper()

	w := httptest.NewRecorder()
	bs, _ := json.Marshal(header)
	req := httptest.NewRequest("POST", "/files/"+fileID, io.NopCloser(bytes.NewReader(bs)))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	env.router.ServeHTTP(w, req)
	w.Flush()

	var file *imagecashletter.File
	if w.Code == http.StatusCreated {
		require.NoError(t, json.NewDecoder(w.Body).Decode(&file))
	}

	return w, file
}

func (env *testEnvironment) addCashLetter(t *testing.T, fileID string, cashLetter imagecashletter.CashLetter) (*httptest.ResponseRecorder, *imagecashletter.File) {
	t.Helper()

	w := httptest.NewRecorder()
	bs, _ := json.Marshal(cashLetter)
	req := httptest.NewRequest("POST", "/files/"+fileID+"/cashLetters", io.NopCloser(bytes.NewReader(bs)))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	env.router.ServeHTTP(w, req)
	w.Flush()

	var resp *imagecashletter.File
	if w.Code == http.StatusOK {
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	}

	return w, resp
}

func (env *testEnvironment) removeCashLetter(t *testing.T, fileID, cashLetterID string) *httptest.ResponseRecorder {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/files/"+fileID+"/cashLetters/"+cashLetterID, nil)
	env.router.ServeHTTP(w, req)
	w.Flush()

	return w
}

func (env *testEnvironment) deleteFile(t *testing.T, fileID string) *httptest.ResponseRecorder {
	t.Helper()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/files/"+fileID, nil)
	env.router.ServeHTTP(w, req)
	w.Flush()

	return w
}

func openTestFile(t *testing.T, filename string) io.Reader {
	t.Helper()

	fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", filename))
	require.NoError(t, err)
	t.Cleanup(func() {
		fd.Close()
	})
	return fd
}

func parseTestFile(t *testing.T, filename string) *imagecashletter.File {
	t.Helper()

	f, err := imagecashletter.NewReader(
		openTestFile(t, filename),
		imagecashletter.ReadVariableLineLengthOption(),
	).Read()
	require.NoError(t, err)

	return &f
}
