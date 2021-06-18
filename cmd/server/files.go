// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/moov-io/base"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/imagecashletter"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
)

var (
	errNoFileId       = errors.New("no File ID found")
	errNoCashLetterId = errors.New("no CashLetter ID found")
)

func addFileRoutes(logger log.Logger, r *mux.Router, repo ICLFileRepository) {
	r.Methods("GET").Path("/files").HandlerFunc(getFiles(logger, repo))
	r.Methods("POST").Path("/files/create").HandlerFunc(createFile(logger, repo))
	r.Methods("GET").Path("/files/{fileId}").HandlerFunc(getFile(logger, repo))
	r.Methods("POST").Path("/files/{fileId}").HandlerFunc(updateFileHeader(logger, repo))
	r.Methods("DELETE").Path("/files/{fileId}").HandlerFunc(deleteFile(logger, repo))

	r.Methods("GET").Path("/files/{fileId}/contents").HandlerFunc(getFileContents(logger, repo))
	r.Methods("GET").Path("/files/{fileId}/validate").HandlerFunc(validateFile(logger, repo))

	r.Methods("POST").Path("/files/{fileId}/cashLetters").HandlerFunc(addCashLetterToFile(logger, repo))
	r.Methods("DELETE").Path("/files/{fileId}/cashLetters/{cashLetterId}").HandlerFunc(removeCashLetterFromFile(logger, repo))
}

func getFileId(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["fileId"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoFileId)
		return ""
	}
	return v
}

func getCashLetterId(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["cashLetterId"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoCashLetterId)
		return ""
	}
	return v
}

func getFiles(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		files, err := repo.getFiles() // TODO(adam): implement soft and hard limits
		if err != nil {
			err = logger.LogErrorf("error getting ICL files: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}
		logger.Logf("found %d files", len(files))

		w.Header().Set("X-Total-Count", fmt.Sprintf("%d", len(files)))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(files)
	}
}

func createFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		req := imagecashletter.NewFile()
		if req.ID == "" {
			req.ID = base.ID()
		}

		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err = logger.LogErrorf("error reading request body: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		h := r.Header.Get("Content-Type")
		if strings.Contains(h, "application/json") {
			file, err := imagecashletter.FileFromJSON(bs)
			if err != nil {
				err = logger.LogErrorf("error creating file from JSON: %v", err).Err()
				moovhttp.Problem(w, err)
				return
			} else {
				req = file
			}
		} else {
			f, err := imagecashletter.NewReader(bytes.NewReader(bs), imagecashletter.ReadVariableLineLengthOption()).Read()
			if err != nil {
				err = logger.LogErrorf("error reading image cache letter: %v", err).Err()
				moovhttp.Problem(w, err)
				return
			} else {
				req = &f
			}
		}
		if req.ID == "" {
			req.ID = base.ID()
		}

		// Save the ICL file
		if err := repo.saveFile(req); err != nil {
			err = logger.LogErrorf("problem saving file %s: %v", req.ID, err).Err()
			moovhttp.Problem(w, err)
			return
		}
		logger.Logf("created file=%s", req.ID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(req)
	}
}

func getFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("problem reading file=%s: %v", fileId, err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Logf("file %q was not found", fileId)
			http.NotFound(w, r)
			return
		}

		logger.Log("rendering file")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(file)
	}
}

func updateFileHeader(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		var req imagecashletter.FileHeader
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			err = logger.LogErrorf("error reading request body: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Logf("file %q was not found", fileId)
			http.NotFound(w, r)
			return
		}

		file.Header = req
		if err := repo.saveFile(file); err != nil {
			err = logger.LogErrorf("error saving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}
		logger.Log("updated FileHeader")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(file)
	}
}

func deleteFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Logf("file %q was not found", fileId)
			http.NotFound(w, r)
			return
		}

		if err := repo.deleteFile(fileId); err != nil {
			err = logger.LogErrorf("error deleting file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		logger.Log("deleted file")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}

func getFileContents(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Logf("file %q was not found", fileId)
			http.NotFound(w, r)
			return
		}

		logger.Log("rendering file contents")

		opts := []imagecashletter.WriterOption{
			imagecashletter.WriteVariableLineLengthOption(),
		}

		w.Header().Set("Content-Type", "text/plain")
		if err := imagecashletter.NewWriter(w, opts...).Write(file); err != nil {
			err = logger.LogErrorf("problem rendering file contents: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func validateFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Logf("file %q was not found", fileId)
			http.NotFound(w, r)
			return
		}

		if err := file.Create(); err != nil { // Create calls Validate
			err = logger.LogErrorf("file=%s was invalid: %v", fileId, err).Err()
			moovhttp.Problem(w, err)
			return
		}

		logger.Log("validated file")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}

func addCashLetterToFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		var req imagecashletter.CashLetter
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			err = logger.LogErrorf("error reading request body: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Logf("file %q was not found", fileId)
			http.NotFound(w, r)
			return
		}

		file.CashLetters = append(file.CashLetters, req)
		if err := repo.saveFile(file); err != nil {
			err = logger.LogErrorf("error saving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		logger.Logf("added CashLetter=%s to file", req.ID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(file)
	}
}

func removeCashLetterFromFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		cashLetterId := getCashLetterId(w, r)
		if cashLetterId == "" {
			logger.LogError(errNoCashLetterId)
			return
		}
		logger = logger.Set("cashLetterID", log.String(cashLetterId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Logf("file %q was not found", fileId)
			http.NotFound(w, r)
			return
		}

		for i := 0; i < len(file.CashLetters); i++ {
			if file.CashLetters[i].ID == cashLetterId {
				file.CashLetters = append(file.CashLetters[:i], file.CashLetters[i+1:]...)
				i--
			}
		}
		if err := repo.saveFile(file); err != nil {
			err = logger.LogErrorf("error saving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		logger.Log("removed CashLetter from file")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}
