package v2

import (
	"bufio"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
	"github.com/moov-io/imagecashletter"
	"github.com/moov-io/imagecashletter/internal/metrics"
	"github.com/moov-io/imagecashletter/internal/responder"
	"github.com/moov-io/imagecashletter/internal/storage"
)

var (
	maxReaderBufferSize = func() int {
		v, exists := os.LookupEnv("READER_BUFFER_SIZE")
		if exists {
			n, _ := strconv.ParseInt(v, 10, 32)
			return int(n)
		}
		return bufio.MaxScanTokenSize
	}()

	maxUploadSize = func() int64 {
		v, exists := os.LookupEnv("MAX_UPLOAD_SIZE")
		if exists {
			n, _ := strconv.ParseInt(v, 10, 64)
			return n
		}
		return 100 * 1024 * 1024 // 100MB default
	}()
)

type Controller struct {
	logger       log.Logger
	repo         storage.ICLFileRepository
	validateOpts *imagecashletter.ValidateOpts // base opts; merged with per-request for creates
}

// NewController constructs a v2 files controller. If validateOpts is non-nil it
// provides base options that are merged (per-request opts from the HTTP request
// take precedence via Merge) on file creation.
func NewController(logger log.Logger, fileRepo storage.ICLFileRepository, validateOpts *imagecashletter.ValidateOpts) Controller {
	return Controller{
		logger:       logger,
		repo:         fileRepo,
		validateOpts: validateOpts,
	}
}

func (c Controller) AddRoutes(router *mux.Router) {
	v2Routes := router.PathPrefix("/v2").Subrouter()

	v2Routes.
		Path("/files").
		Methods(http.MethodPost).
		HandlerFunc(c.createFile)

}

func (c Controller) createFile(w http.ResponseWriter, r *http.Request) {
	w = metrics.WrapResponseWriter(c.logger, w, r)
	respond := responder.NewResponder(c.logger, w, r)

	// Bound request body size to mitigate DoS via large uploads (G120)
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	var created *imagecashletter.File
	var err error

	contentType := r.Header.Get("Content-Type")

	// Per-request ValidateOpts (e.g. derived from the incoming HTTP request)
	// are merged with any Controller-level defaults.
	var requestOpts *imagecashletter.ValidateOpts
	// TODO: populate requestOpts from r when per-request configuration is needed
	// (for example from query parameters). Use c.validateOpts.Merge(requestOpts)
	// (already performed inside the fileFrom* helpers).

	switch {
	case strings.Contains(contentType, "application/json"):
		created, err = c.fileFromJSON(r, requestOpts)
	case strings.Contains(contentType, "multipart/form-data"):
		created, err = c.fileFromForm(r, requestOpts)
	default:
		err = fmt.Errorf("missing or unsupported Content-Type: %s", contentType)
	}

	if err != nil {
		c.logger.Error().LogErrorf("creating file: %v", err)
		respond.Error(http.StatusBadRequest, err)
		return
	}

	if err = c.repo.SaveFile(created); err != nil {
		c.logger.Error().LogErrorf("saving created file: %v", err)
		respond.Error(http.StatusInternalServerError, err)
		return
	}

	// TODO: Update to the v2 API endpoint once the GET file endpoint is implemented
	location := fmt.Sprintf("%s://%s/files/%s", r.URL.Scheme, r.URL.Host, created.ID)
	respond = respond.WithLocation(location)
	if expectingFile(r) {
		respond.File(http.StatusCreated, *created, fmt.Sprintf("%s.x9", created.ID))
		return
	}

	respond.JSON(http.StatusCreated, created)
}

func expectingFile(r *http.Request) bool {
	mimeType := r.Header.Get("Accept")
	return mimeType == "application/octet-stream" || mimeType == "text/plain"
}

func (c Controller) fileFromJSON(r *http.Request, requestOpts *imagecashletter.ValidateOpts) (*imagecashletter.File, error) {
	contents, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("reading request body: %w", err)
	}

	merged := c.validateOpts.Merge(requestOpts)
	file, err := imagecashletter.FileFromJSONWithOpts(contents, merged)
	if err != nil {
		return nil, fmt.Errorf("parsing request body: %w", err)
	}
	file.ID = uuid.NewString()

	return file, nil
}

func (c Controller) fileFromForm(r *http.Request, requestOpts *imagecashletter.ValidateOpts) (*imagecashletter.File, error) {
	mr, err := r.MultipartReader()
	if err != nil {
		return nil, fmt.Errorf("getting multipart reader: %w", err)
	}

	var part *multipart.Part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading multipart part: %w", err)
		}
		if p.FormName() == "file" {
			part = p
			break
		}
		p.Close()
	}
	if part == nil {
		return nil, fmt.Errorf("missing file part in multipart form")
	}
	defer part.Close()

	opts := []imagecashletter.ReaderOption{
		imagecashletter.ReadVariableLineLengthOption(),
		imagecashletter.BufferSizeOption(maxReaderBufferSize),
	}

	// Industry standard encoding is EBCDIC, so unless plain/text was
	// explicitly requested, default to EBCDIC.
	contentType := part.Header.Get("Content-Type")
	if contentType != "text/plain" {
		opts = append(opts, imagecashletter.ReadEbcdicEncodingOption())
	}

	merged := c.validateOpts.Merge(requestOpts)
	if merged != nil {
		opts = append(opts, imagecashletter.ReadValidateOpts(merged))
	}

	file, err := imagecashletter.NewReader(part, opts...).Read()
	if err != nil {
		return nil, fmt.Errorf("parsing file: %w", err)
	}
	file.ID = uuid.NewString()

	return &file, nil
}
