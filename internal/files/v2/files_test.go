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
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
	"github.com/moov-io/imagecashletter"
	openapi "github.com/moov-io/imagecashletter/client"
	v2 "github.com/moov-io/imagecashletter/internal/files/v2"
	"github.com/moov-io/imagecashletter/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestController_createFile_errors(t *testing.T) {
	router := newRouter(t)

	t.Run("unsupported content type", func(t *testing.T) {
		rdr := strings.NewReader("real file")
		resp, apiErr := createFile(t, router, rdr, "application/msword", "application/json")
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Contains(t, apiErr.Error, "unsupported Content-Type")
	})

	t.Run("invalid json file", func(t *testing.T) {
		rdr := strings.NewReader("real file")
		resp, apiErr := createFile(t, router, rdr, "application/json", "application/json")
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Contains(t, apiErr.Error, "problem reading file")
	})

	t.Run("invalid ascii file", func(t *testing.T) {
		rdr := strings.NewReader("real file")
		resp, apiErr := uploadFile(t, router, rdr, "text/plain", "application/json")
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Contains(t, apiErr.Error, "parsing file")
	})

	t.Run("invalid ebcdic file", func(t *testing.T) {
		rdr := strings.NewReader("real file")
		resp, apiErr := uploadFile(t, router, rdr, "application/octet-stream", "application/json")
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Contains(t, apiErr.Error, "parsing file")
	})
}

func TestController_createJSONFile(t *testing.T) {
	router := newRouter(t)

	t.Run("returns JSON", func(t *testing.T) {
		rdr := getTestData(t, "icl-valid.json")

		resp, apiErr := createFile(t, router, rdr, "application/json", "application/json")
		require.Empty(t, apiErr)

		var created imagecashletter.File
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&created))

		require.NotEmpty(t, created.ID)
		require.Equal(t, "/v2/files/"+created.ID, resp.Header().Get("Location"))
		require.NotEmpty(t, created)
		require.Equal(t, "231380104", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 2)
		require.Equal(t, 400000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})

	t.Run("invalid accept header returns JSON", func(t *testing.T) {
		rdr := getTestData(t, "icl-valid.json")

		resp, apiErr := createFile(t, router, rdr, "application/json", "foo/bar")
		require.Empty(t, apiErr)

		var created imagecashletter.File
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&created))

		require.Contains(t, resp.Header().Get("Location"), "/v2/files/"+created.ID)
		require.NotEmpty(t, created.ID)
		require.NotEmpty(t, created)
		require.Equal(t, "231380104", created.Header.ImmediateDestination)
		require.Len(t, created.CashLetters, 2)
		require.Equal(t, 400000, created.CashLetters[0].CashLetterControl.CashLetterTotalAmount)
	})

	t.Run("returns EBCDIC file", func(t *testing.T) {
		rdr := getTestData(t, "icl-valid.json")

		resp, apiErr := createFile(t, router, rdr, "application/json", "application/octet-stream")
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

	t.Run("returns ASCII file", func(t *testing.T) {
		rdr := getTestData(t, "icl-valid.json")

		resp, apiErr := createFile(t, router, rdr, "application/json", "text/plain")
		require.Empty(t, apiErr)

		opts := []imagecashletter.ReaderOption{
			imagecashletter.ReadVariableLineLengthOption(),
		}
		created, err := imagecashletter.NewReader(resp.Body, opts...).Read()
		require.NoError(t, err)

		require.Contains(t, resp.Header().Get("Content-Type"), "text/plain")
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

		// inspect headers
		require.Contains(t, resp.Header().Get("Content-Type"), "text/plain")
		require.Contains(t, resp.Header().Get("Content-Disposition"), ".x9")
		location := resp.Header().Get("Location")
		require.True(t, strings.HasPrefix(location, "/v2/files/"))
		resourceID := strings.TrimPrefix(location, "/v2/files/")
		require.NotEmpty(t, resourceID)
		_, err := uuid.Parse(resourceID)
		require.NoError(t, err)

		// now read back in without EBCDIC option
		created, err := imagecashletter.NewReader(resp.Body, imagecashletter.ReadVariableLineLengthOption()).Read()
		require.NoError(t, err)
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

func createFile(t *testing.T, router *mux.Router, body io.Reader, contentType string, accept string, queryParams ...string) (*httptest.ResponseRecorder, openapi.Error) {
	urlStr := "https://some.domain.io/v2/files"
	if len(queryParams) > 0 && queryParams[0] != "" {
		urlStr += "?" + queryParams[0]
	}
	req, err := http.NewRequest(http.MethodPost, urlStr, body)
	require.NoError(t, err)

	req.Header.Set("Content-Type", contentType)
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

func uploadFile(t *testing.T, router *mux.Router, file io.Reader, contentType, accept string, queryParams ...string) (*httptest.ResponseRecorder, openapi.Error) {
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

	urlStr := "https://some.domain.io/v2/files"
	if len(queryParams) > 0 && queryParams[0] != "" {
		urlStr += "?" + queryParams[0]
	}
	req, err := http.NewRequest(http.MethodPost, urlStr, body)
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
	return newRouterWithOpts(t, nil)
}

func newRouterWithOpts(t *testing.T, validateOpts *imagecashletter.ValidateOpts) *mux.Router {
	c := v2.NewController(log.NewTestLogger(), storage.NewInMemoryRepo(), validateOpts)
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

// mismatchedAddendumJSONReader returns a JSON reader for a file containing a
// CheckDetail with AddendumCount that does not match the number of addenda
// records. This triggers count validation errors under default (strict) opts.
func mismatchedAddendumJSONReader(t *testing.T) io.Reader {
	t.Helper()

	bs, err := os.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "icl-valid.json"))
	require.NoError(t, err)

	var f imagecashletter.File
	require.NoError(t, json.NewDecoder(bytes.NewReader(bs)).Decode(&f))

	for _, cl := range f.CashLetters {
		for _, b := range cl.Bundles {
			for _, cd := range b.Checks {
				if len(cd.CheckDetailAddendumA)+len(cd.CheckDetailAddendumB)+len(cd.CheckDetailAddendumC) > 0 {
					cd.AddendumCount = 0
					out, err := json.Marshal(f)
					require.NoError(t, err)
					return bytes.NewReader(out)
				}
			}
		}
	}
	t.Fatal("valid fixture contains no CheckDetail with addenda to create mismatch from")
	return nil
}

// addendumCountMismatchX937 returns raw X9.37 bytes with a deliberate
// AddendumCount mismatch (for exercising the binary upload path with/without opts).
func addendumCountMismatchX937(t *testing.T) []byte {
	t.Helper()

	bs, err := os.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "icl-valid.json"))
	require.NoError(t, err)

	var f imagecashletter.File
	require.NoError(t, json.NewDecoder(bytes.NewReader(bs)).Decode(&f))

	found := false
	for _, cl := range f.CashLetters {
		for _, b := range cl.Bundles {
			for _, cd := range b.Checks {
				if len(cd.CheckDetailAddendumA)+len(cd.CheckDetailAddendumB)+len(cd.CheckDetailAddendumC) > 0 {
					cd.AddendumCount = 0
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
	require.True(t, found, "fixture must contain a check with addenda to mismatch")

	f.SetValidation(&imagecashletter.ValidateOpts{SkipAll: true})
	require.NoError(t, f.Create())

	var buf bytes.Buffer
	require.NoError(t, imagecashletter.NewWriter(&buf, imagecashletter.WriteVariableLineLengthOption()).Write(&f))

	return buf.Bytes()
}

func TestController_createFile_withValidateOpts(t *testing.T) {
	t.Run("json: mismatch rejected by default", func(t *testing.T) {
		router := newRouter(t)
		resp, apiErr := createFile(t, router, mismatchedAddendumJSONReader(t), "application/json", "application/json", "skipAll=false")
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Contains(t, apiErr.Error, "does not match Addenda Records")
	})

	t.Run("json: SkipAll allows the file", func(t *testing.T) {
		router := newRouterWithOpts(t, &imagecashletter.ValidateOpts{SkipAll: true})
		resp, apiErr := createFile(t, router, mismatchedAddendumJSONReader(t), "application/json", "application/json")
		require.Empty(t, apiErr)
		require.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("json: SkipCountValidation allows the file", func(t *testing.T) {
		router := newRouterWithOpts(t, &imagecashletter.ValidateOpts{SkipCountValidation: true})
		resp, apiErr := createFile(t, router, mismatchedAddendumJSONReader(t), "application/json", "application/json")
		require.Empty(t, apiErr)
		require.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("json: SkipAll via per-request query param allows the file", func(t *testing.T) {
		router := newRouter(t)
		resp, apiErr := createFile(t, router, mismatchedAddendumJSONReader(t), "application/json", "application/json", "skipAll=true")
		require.Empty(t, apiErr)
		require.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("json: SkipCountValidation via per-request query param allows the file", func(t *testing.T) {
		router := newRouter(t)
		resp, apiErr := createFile(t, router, mismatchedAddendumJSONReader(t), "application/json", "application/json", "skipCountValidation=true")
		require.Empty(t, apiErr)
		require.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("json: controller SkipCount + per-request SkipAll merge", func(t *testing.T) {
		router := newRouterWithOpts(t, &imagecashletter.ValidateOpts{SkipCountValidation: true})
		resp, apiErr := createFile(t, router, mismatchedAddendumJSONReader(t), "application/json", "application/json", "skipAll=true")
		require.Empty(t, apiErr)
		require.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("upload: mismatch rejected by default", func(t *testing.T) {
		router := newRouter(t)
		resp, apiErr := uploadFile(t, router, bytes.NewReader(addendumCountMismatchX937(t)), "text/plain", "application/json", "skipAll=false")
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Contains(t, apiErr.Error, "does not match Addenda Records")
	})

	t.Run("upload: SkipAll allows the file", func(t *testing.T) {
		router := newRouterWithOpts(t, &imagecashletter.ValidateOpts{SkipAll: true})
		resp, apiErr := uploadFile(t, router, bytes.NewReader(addendumCountMismatchX937(t)), "text/plain", "application/json")
		require.Empty(t, apiErr)
		require.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("upload: SkipAll via per-request query param allows the file", func(t *testing.T) {
		router := newRouter(t)
		resp, apiErr := uploadFile(t, router, bytes.NewReader(addendumCountMismatchX937(t)), "text/plain", "application/json", "skipAll=true")
		require.Empty(t, apiErr)
		require.Equal(t, http.StatusCreated, resp.Code)
	})
}
