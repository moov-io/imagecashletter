// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/moov-io/imagecashletter"

	"github.com/vincent-petithory/dataurl"
)

func decodeDataURI(input string) string {
	u, _ := dataurl.DecodeString(input)

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	base64.StdEncoding.Decode(dst, u.Data)

	return string(u.Data)
}

func parseContents(input string) (string, error) {
	r := strings.NewReader(input)
	file, err := imagecashletter.NewReader(r,
		imagecashletter.ReadVariableLineLengthOption(),
		imagecashletter.ReadEbcdicEncodingOption(),
	).Read()

	var buf bytes.Buffer
	if err2 := json.NewEncoder(&buf).Encode(file); err2 != nil {
		err = fmt.Errorf("original error: %v and json encoder error: %v", err, err2)
	}

	return buf.String(), err
}

func readAsJSON(input string) (string, error) {
	file, err := imagecashletter.FileFromJSON([]byte(input))

	pretty, err2 := json.MarshalIndent(file, "", " ")
	if err2 != nil {
		err = fmt.Errorf("original error: %v and json indent encoding error: %v", err, err2)
	}

	return string(pretty), err
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		input := args[0].String()

		// Decode the data URI if we're given one (file uploads)
		if strings.HasPrefix(input, "data:") {
			input = decodeDataURI(input)
		}

		// Handle either JSON or X9-formatted contents
		if json.Valid([]byte(input)) {
			pretty, err := readAsJSON(input)
			if err != nil {
				fmt.Printf("unable to parse JSON %s\n", err)
				return err.Error()
			}
			return pretty
		} else {
			parsed, err := parseContents(input)
			if err != nil {
				fmt.Printf("unable to parse contents %s\n", err)
				return err.Error()
			}

			pretty, err := prettyJson(parsed)
			if err != nil {
				fmt.Printf("unable to convert to json %s\n", err)
				return err.Error()
			}

			return pretty
		}
	})
	return jsonFunc
}

func main() {
	js.Global().Set("parseContents", jsonWrapper())
	<-make(chan bool)
}
