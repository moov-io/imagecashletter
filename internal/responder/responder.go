package responder

import (
	"encoding/json"
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/imagecashletter"
	client "github.com/moov-io/imagecashletter/client"
)

// Responder is a helper for writing responses to an http.ResponseWriter.
type Responder struct {
	logger log.Logger
	w      http.ResponseWriter
	r      *http.Request
	optionalHeaders
}

type optionalHeaders struct {
	location string
}

func (h optionalHeaders) apply(w http.ResponseWriter) {
	if h.location != "" {
		w.Header().Set("Location", h.location)
	}
}

func NewResponder(logger log.Logger, w http.ResponseWriter, r *http.Request) *Responder {
	return &Responder{
		logger: logger,
		w:      w,
		r:      r,
	}
}

// WithLocation sets the Location header on the response. The Location header should be used for
// HTTP 201 responses to indicate the location of the newly created resource. While both absolute
// and relative URIs are allowed, the absolute URI is preferred.
func (r *Responder) WithLocation(location string) *Responder {
	r.optionalHeaders.location = location
	return r
}

// File writes the file to the http.ResponseWriter as an attachment. It should
// only be called when the request's Accept header was "application/octet-stream" or "text/plain".
func (r *Responder) File(status int, file imagecashletter.File, name string) {
	opts := []imagecashletter.WriterOption{
		imagecashletter.WriteVariableLineLengthOption(),
	}

	// determine which encoding the caller expects to set up the writer and response headers
	mimeType := r.r.Header.Get("Accept")
	switch mimeType {
	case "application/octet-stream":
		r.w.Header().Set("Content-Type", "application/octet-stream")
		opts = append(opts, imagecashletter.WriteEbcdicEncodingOption())
	case "text/plain":
		r.w.Header().Set("Content-Type", "text/plain")
	default:
		r.logger.LogErrorf("renderer: file method called for unsupported mime-type: %s", mimeType)
		r.w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.optionalHeaders.apply(r.w)
	r.w.Header().Set("Content-Disposition", "attachment; filename="+name)
	r.w.WriteHeader(status)

	if err := imagecashletter.NewWriter(r.w, opts...).Write(&file); err != nil {
		r.logger.LogErrorf("rendering file: %v", err)
		r.w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// JSON writes the resource to the http.ResponseWriter as JSON.
func (r *Responder) JSON(status int, resource any) {
	r.optionalHeaders.apply(r.w)
	r.w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.w.WriteHeader(status)
	if err := json.NewEncoder(r.w).Encode(resource); err != nil {
		r.logger.LogErrorf("problem encoding response: %v", err)
		r.w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Error sets the status code and writes an error to the http.ResponseWriter. Response body is
// omitted for 5xx errors.
func (r *Responder) Error(status int, err error) {
	if status >= 500 { // don't return body for internal errors
		r.w.WriteHeader(status)
		return
	}

	r.w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.w.WriteHeader(status)
	if err := json.NewEncoder(r.w).Encode(client.Error{Error: err.Error()}); err != nil {
		r.logger.LogErrorf("problem encoding response: %v", err)
		r.w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
