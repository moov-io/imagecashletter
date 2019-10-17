// Copyright 2019 The Moov Authors
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

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
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
		w = wrapResponseWriter(logger, w, r)

		files, err := repo.getFiles() // TODO(adam): implement soft and hard limits
		if err != nil {
			logger.Log("files", fmt.Sprintf("error getting ICL files: %v", err), "requestID", moovhttp.GetRequestID(r))
			moovhttp.Problem(w, err)
			return
		}
		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("found %d files", len(files)), "requestID", requestID)

		w.Header().Set("X-Total-Count", fmt.Sprintf("%d", len(files)))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(files)
	}
}

func createFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		req := imagecashletter.NewFile()
		if req.ID == "" {
			req.ID = base.ID()
		}

		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		h := r.Header.Get("Content-Type")
		if strings.Contains(h, "application/json") {
			file, err := imagecashletter.FileFromJSON(bs)
			if err != nil {
				moovhttp.Problem(w, err)
				return
			} else {
				req = file
			}
		} else {
			f, err := imagecashletter.NewReader(bytes.NewReader(bs)).Read()
			if err != nil {
				moovhttp.Problem(w, err)
				return
			} else {
				req = &f
			}
		}
		if req.ID == "" {
			req.ID = base.ID()
		}
		if err := repo.saveFile(req); err != nil {
			logger.Log("files", fmt.Sprintf("problem saving file %s: %v", req.ID, err), "requestID", moovhttp.GetRequestID(r))
			moovhttp.Problem(w, err)
			return
		}
		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("creatd file=%s", req.ID), "requestID", requestID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(req)
	}
}

func getFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			logger.Log("files", fmt.Sprintf("problem reading file=%s: %v", fileId, err), "requestID", moovhttp.GetRequestID(r))
			moovhttp.Problem(w, err)
			return
		}

		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("rendering file=%s", fileId), "requestID", requestID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(file)
	}
}

func updateFileHeader(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		var req imagecashletter.FileHeader
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		file.Header = req
		if err := repo.saveFile(file); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("updating FileHeader for file=%s", fileId), "requestID", requestID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(file)
	}
}

func deleteFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		if err := repo.deleteFile(fileId); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("deleted file=%s", fileId), "requestID", requestID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}

func getFileContents(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("rendering file=%s contents", fileId), "requestID", requestID)

		w.Header().Set("Content-Type", "text/plain")
		if err := imagecashletter.NewWriter(w).Write(file); err != nil {
			logger.Log("files", fmt.Sprintf("problem rendering file=%s contents: %v", fileId, err))
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func validateFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if err := file.Create(); err != nil { // Create calls Validate
			requestID := moovhttp.GetRequestID(r)
			logger.Log("files", fmt.Sprintf("file=%s was invalid: %v", fileId, err), "requestID", requestID)
			moovhttp.Problem(w, err)
			return
		}

		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("validated file=%s", fileId), "requestID", requestID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}

func addCashLetterToFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		var req imagecashletter.CashLetter
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		file.CashLetters = append(file.CashLetters, req)
		if err := repo.saveFile(file); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("added CashLetter=%s to file=%s", req.ID, fileId), "requestID", requestID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(file)
	}
}

func removeCashLetterFromFile(logger log.Logger, repo ICLFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId, cashLetterId := getFileId(w, r), getCashLetterId(w, r)
		if fileId == "" || cashLetterId == "" {
			return
		}

		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		for i := 0; i < len(file.CashLetters); i++ {
			if file.CashLetters[i].ID == cashLetterId {
				file.CashLetters = append(file.CashLetters[:i], file.CashLetters[i+1:]...)
				i--
			}
		}
		if err := repo.saveFile(file); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		requestID := moovhttp.GetRequestID(r)
		logger.Log("files", fmt.Sprintf("removed CashLetter=%s to file=%s", cashLetterId, fileId), "requestID", requestID)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}
