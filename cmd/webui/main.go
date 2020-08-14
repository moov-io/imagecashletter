// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/admin"
	"github.com/moov-io/base/http/bind"

	"github.com/moov-io/imagecashletter"
)

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ICL"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ICL"), "Admin HTTP listen address")

	flagLogFormat = flag.String("log.format", "", "Format for log lines (Options: json, plain")
)

func main() {
	flag.Parse()

	log.Printf("Starting moov-io/imagecashletter server version %s", imagecashletter.Version)

	// Channel for errors
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Start Admin server (with Prometheus metrics)
	adminServer := admin.NewServer(*adminAddr)
	adminServer.AddVersionHandler(imagecashletter.Version) // Setup 'GET /version'
	go func() {
		log.Printf("listening on %s", adminServer.BindAddr())
		if err := adminServer.Listen(); err != nil {
			err = fmt.Errorf("problem starting admin http: %v", err)
			log.Print(err)
			errs <- err
		}
	}()
	defer adminServer.Shutdown()

	// Setup business HTTP routes
	router := mux.NewRouter()
	addPingRoute(router)
	addAssetsPath(router, http.FileServer(http.Dir(filepath.Join("cmd", "webui", "assets"))))

	serve := &http.Server{
		Addr:    *httpAddr,
		Handler: router,
		TLSConfig: &tls.Config{
			InsecureSkipVerify:       false,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	shutdownServer := func() {
		if err := serve.Shutdown(context.TODO()); err != nil {
			log.Print(err)
		}
	}

	// Start business logic HTTP server
	go func() {
		if certFile, keyFile := os.Getenv("HTTPS_CERT_FILE"), os.Getenv("HTTPS_KEY_FILE"); certFile != "" && keyFile != "" {
			log.Printf("binding to %s for secure HTTP server", *httpAddr)
			if err := serve.ListenAndServeTLS(certFile, keyFile); err != nil {
				log.Print(err)
			}
		} else {
			log.Printf("binding to %s for HTTP server", *httpAddr)
			if err := serve.ListenAndServe(); err != nil {
				log.Print(err)
			}
		}
	}()

	// Block/Wait for an error
	if err := <-errs; err != nil {
		shutdownServer()
		log.Print(err)
	}
}

func addPingRoute(r *mux.Router) {
	r.Methods("GET").Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PONG"))
	})
}

func addAssetsPath(r *mux.Router, handler http.Handler) {
	r.Methods("GET").PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}
