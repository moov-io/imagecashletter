// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/admin"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/http/bind"
	"github.com/moov-io/base/log"
	"github.com/moov-io/imagecashletter"
	"github.com/moov-io/imagecashletter/internal/files"
	v2files "github.com/moov-io/imagecashletter/internal/files/v2"
	"github.com/moov-io/imagecashletter/internal/storage"
)

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ICL"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ICL"), "Admin HTTP listen address")

	flagLogFormat = flag.String("log.format", "", "Format for log lines (Options: json, plain")
)

func main() {
	flag.Parse()

	var logger log.Logger
	if strings.ToLower(*flagLogFormat) == "json" {
		logger = log.NewJSONLogger()
	} else {
		logger = log.NewDefaultLogger()
	}
	logger = logger.Set("package", log.String("main"))

	logger.Logf("Starting moov-io/imagecashletter server version %s", imagecashletter.Version)

	// Channel for errors
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Start Admin server (with Prometheus metrics)
	adminServer, err := admin.New(admin.Opts{
		Addr: *adminAddr,
	})
	if err != nil {
		errs <- fmt.Errorf("creating admin server: %w", err)
	}
	adminServer.AddVersionHandler(imagecashletter.Version) // Setup 'GET /version'
	go func() {
		logger.Logf("admin server listening on %s", adminServer.BindAddr())
		if err := adminServer.Listen(); err != nil {
			if err == http.ErrServerClosed {
				logger.Log("admin server closed")
				return
			}
			err = logger.LogErrorf("problem starting admin http: %v", err).Err()
			errs <- err
		}
	}()
	defer adminServer.Shutdown()

	// Persistence layer shared between API versions for interoperability
	repository := storage.NewInMemoryRepo()

	// Setup business HTTP routes. Base ValidateOpts (from env) are merged with any
	// per-request opts (e.g. query params on create) for file creation.
	serverValidateOpts := getValidateOptsFromEnv()

	router := mux.NewRouter()
	moovhttp.AddCORSHandler(router)
	addPingRoute(router)
	files.AppendRoutes(logger, router, repository, serverValidateOpts)
	v2files.NewController(logger, repository, serverValidateOpts).AddRoutes(router)

	// Start business HTTP server
	readTimeout, _ := time.ParseDuration("30s")
	writTimeout, _ := time.ParseDuration("30s")
	idleTimeout, _ := time.ParseDuration("60s")

	serve := &http.Server{
		Addr:    *httpAddr,
		Handler: router,
		TLSConfig: &tls.Config{
			InsecureSkipVerify:       false,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
		},
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readTimeout,
		WriteTimeout:      writTimeout,
		IdleTimeout:       idleTimeout,
	}
	shutdownServer := func() {
		if err := serve.Shutdown(context.TODO()); err != nil {
			logger.LogErrorf("shutdown error: %v", err)
		}
	}

	// Start business logic HTTP server
	go func() {
		if certFile, keyFile := os.Getenv("HTTPS_CERT_FILE"), os.Getenv("HTTPS_KEY_FILE"); certFile != "" && keyFile != "" {
			logger.Logf("binding to %s for secure HTTP server", *httpAddr)
			if err := serve.ListenAndServeTLS(certFile, keyFile); err != nil {
				if err == http.ErrServerClosed {
					logger.Log("secure http server closed")
					return
				}
				logger.LogErrorf("http server error: %v", err)
			}
		} else {
			logger.Logf("binding to %s for HTTP server", *httpAddr)
			if err := serve.ListenAndServe(); err != nil {
				if err == http.ErrServerClosed {
					logger.Log("http server closed")
					return
				}
				logger.LogErrorf("http server error: %v", err)
			}
		}
	}()

	// Block/Wait for an error
	if err := <-errs; err != nil {
		logger.LogError(err)
		shutdownServer()
	}
}

func addPingRoute(r *mux.Router) {
	r.Methods("GET").Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		moovhttp.SetAccessControlAllowHeaders(w, r.Header.Get("Origin"))
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PONG"))
	})
}

// getValidateOptsFromEnv returns base ValidateOpts derived from environment
// variables (or nil). These are intended to be passed to the file controllers
// as defaults, and will be merged (OR) with any per-request opts on creates.
func getValidateOptsFromEnv() *imagecashletter.ValidateOpts {
	var opts imagecashletter.ValidateOpts
	changed := false
	for _, key := range []struct {
		env string
		set *bool
	}{
		{"SKIP_ALL_ON_FILE_CREATE", &opts.SkipAll},
		{"SKIP_COUNT_VALIDATION_ON_FILE_CREATE", &opts.SkipCountValidation},
	} {
		if v := os.Getenv(key.env); v != "" {
			if b, err := strconv.ParseBool(v); err == nil && b {
				*key.set = true
				changed = true
			}
		}
	}
	if !changed {
		return nil
	}
	return &opts
}
