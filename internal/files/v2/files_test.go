package v2_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"path/filepath"
	"testing"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
	"github.com/moov-io/imagecashletter"
	openapi "github.com/moov-io/imagecashletter/client"
	v2 "github.com/moov-io/imagecashletter/internal/files/v2"
	"github.com/moov-io/imagecashletter/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestController_createJSONFile(t *testing.T) {
	router := newRouter(t)

	t.Run("returns JSON", func(t *testing.T) {
		rdr := getTestData(t, "icl-valid.json")

		resp, apiErr := createFile(t, router, rdr, "application/json")
		require.Empty(t, apiErr)

		var created imagecashletter.File
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&created))

		require.Contains(t, resp.Header().Get("Location"), created.ID)
		require.NotEmpty(t, created.ID)
		require.NotEmpty(t, created)
		require.Equal(t, "231380104", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 2)
		require.Equal(t, 400000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})

	t.Run("returns EBCDIC file", func(t *testing.T) {
		rdr := getTestData(t, "icl-valid.json")

		resp, apiErr := createFile(t, router, rdr, "application/octet-stream")
		require.Empty(t, apiErr)

		opts := []imagecashletter.ReaderOption{
			imagecashletter.ReadVariableLineLengthOption(),
			imagecashletter.ReadEbcdicEncodingOption(),
		}
		created, err := imagecashletter.NewReader(resp.Body, opts...).Read()
		require.NoError(t, err)

		require.Contains(t, resp.Header().Get("Content-Type"), "application/octet-stream")
		require.Contains(t, resp.Header().Get("Content-Disposition"), ".x9")
		require.NotEmpty(t, created)
		require.Equal(t, "231380104", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 2)
		require.Equal(t, 400000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})
}

func TestController_uploadEBCDICFile(t *testing.T) {
	router := newRouter(t)

	t.Run("uploads EBCDIC; returns ASCII", func(t *testing.T) {
		rdr := getTestData(t, "valid-ebcdic.x937")

		resp, apiErr := uploadFile(t, router, rdr, "application/octet-stream", "text/plain")
		require.Empty(t, apiErr)

		// now read back in without EBCDIC option
		created, err := imagecashletter.NewReader(resp.Body, imagecashletter.ReadVariableLineLengthOption()).Read()
		require.NoError(t, err)
		require.Contains(t, resp.Header().Get("Content-Type"), "text/plain")
		require.Contains(t, resp.Header().Get("Content-Disposition"), ".x9")
		require.NotEmpty(t, created)
		require.Equal(t, "061000146", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 1)
		require.Equal(t, 10000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})

	t.Run("uploads EBCDIC; returns EBCDIC", func(t *testing.T) {
		rdr := getTestData(t, "valid-ebcdic.x937")

		resp, apiErr := uploadFile(t, router, rdr, "application/octet-stream", "application/octet-stream")
		require.Empty(t, apiErr)

		// now read back in with EBCDIC option
		created, err := imagecashletter.NewReader(resp.Body,
			imagecashletter.ReadVariableLineLengthOption(),
			imagecashletter.ReadEbcdicEncodingOption(),
		).Read()
		require.NoError(t, err)
		require.Contains(t, resp.Header().Get("Content-Type"), "application/octet-stream")
		require.Contains(t, resp.Header().Get("Content-Disposition"), ".x9")
		require.NotEmpty(t, created)
		require.Equal(t, "061000146", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 1)
		require.Equal(t, 10000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})

	t.Run("uploads EBCDIC; returns JSON", func(t *testing.T) {
		rdr := getTestData(t, "valid-ebcdic.x937")

		resp, apiErr := uploadFile(t, router, rdr, "application/octet-stream", "application/json")
		require.Empty(t, apiErr)

		// now read back in without EBCDIC option
		var created imagecashletter.File
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&created))
		require.Contains(t, resp.Header().Get("Content-Type"), "application/json")
		require.NotEmpty(t, created.ID)
		require.Equal(t, "061000146", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 1)
		require.Equal(t, 10000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})
}

func TestController_uploadASCIIFile(t *testing.T) {
	router := newRouter(t)

	t.Run("uploads ASCII; returns ASCII", func(t *testing.T) {
		rdr := getTestData(t, "valid-ascii.x937")

		resp, apiErr := uploadFile(t, router, rdr, "text/plain", "text/plain")
		require.Empty(t, apiErr)

		// now read back in without EBCDIC option
		created, err := imagecashletter.NewReader(resp.Body, imagecashletter.ReadVariableLineLengthOption()).Read()
		require.NoError(t, err)
		require.Contains(t, resp.Header().Get("Content-Type"), "text/plain")
		require.Contains(t, resp.Header().Get("Content-Disposition"), ".x9")
		require.NotEmpty(t, created)
		require.Equal(t, "061000146", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 1)
		require.Equal(t, 10000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})

	t.Run("uploads EBCDIC; returns EBCDIC", func(t *testing.T) {
		rdr := getTestData(t, "valid-ascii.x937")

		resp, apiErr := uploadFile(t, router, rdr, "text/plain", "application/octet-stream")
		require.Empty(t, apiErr)

		// now read back in with EBCDIC option
		created, err := imagecashletter.NewReader(resp.Body,
			imagecashletter.ReadVariableLineLengthOption(),
			imagecashletter.ReadEbcdicEncodingOption(),
		).Read()
		require.NoError(t, err)
		require.Contains(t, resp.Header().Get("Content-Type"), "application/octet-stream")
		require.Contains(t, resp.Header().Get("Content-Disposition"), ".x9")
		require.NotEmpty(t, created)
		require.Equal(t, "061000146", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 1)
		require.Equal(t, 10000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})

	t.Run("uploads EBCDIC; returns JSON", func(t *testing.T) {
		rdr := getTestData(t, "valid-ascii.x937")

		resp, apiErr := uploadFile(t, router, rdr, "text/plain", "application/json")
		require.Empty(t, apiErr)

		// now read back in without EBCDIC option
		var created imagecashletter.File
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&created))
		require.Contains(t, resp.Header().Get("Content-Type"), "application/json")
		require.NotEmpty(t, created.ID)
		require.Equal(t, "061000146", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 1)
		require.Equal(t, 10000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})
}

func createFile(t *testing.T, router *mux.Router, body io.Reader, accept string) (*httptest.ResponseRecorder, openapi.Error) {
	req, err := http.NewRequest(http.MethodPost, "/v2/files", body)
	require.NoError(t, err)

	req.Header.Set("Accept", accept)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	var apiErr openapi.Error
	if w.Code >= 400 && w.Code < 500 {
		require.NoError(t, json.NewDecoder(w.Body).Decode(&apiErr))
	}

	return w, apiErr
}

func uploadFile(t *testing.T, router *mux.Router, file io.Reader, contentType, accept string) (*httptest.ResponseRecorder, openapi.Error) {
	// create the multipart-form writer
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)

	// set up headers for the "file" portion
	partHeaders := make(textproto.MIMEHeader)
	partHeaders.Set("Content-Disposition", `form-data; name="file"; filename="cashletter.x9"`)
	partHeaders.Set("Content-Type", contentType)

	// create a new multipart-form section with the headers
	part, err := mw.CreatePart(partHeaders)
	require.NoError(t, err)

	// copy the file contents to the form section
	_, err = io.Copy(part, file)
	require.NoError(t, err)
	require.NoError(t, mw.Close())

	req, err := http.NewRequest(http.MethodPost, "/v2/files", body)
	require.NoError(t, err)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("Accept", accept)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	var apiErr openapi.Error
	if w.Code >= 400 && w.Code < 500 {
		require.NoError(t, json.NewDecoder(w.Body).Decode(&apiErr))
	}

	return w, apiErr
}

func newRouter(t *testing.T) *mux.Router {
	c := v2.NewController(log.NewTestLogger(), storage.NewInMemoryRepo())
	router := mux.NewRouter()
	c.AddRoutes(router)

	return router
}

func getTestData(t *testing.T, filename string) io.Reader {
	fd, err := os.Open(filepath.Join("..", "..", "..", "test", "testdata", filename))
	require.NoError(t, err)
	t.Cleanup(func() {
		fd.Close()
	})

	return fd
}
