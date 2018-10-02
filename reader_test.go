// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// TestX9FileRead validates reading an x9 file
func TestX9FileRead(t *testing.T) {
	f, err := os.Open("./test/data/20180905A.x9")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if err != nil {
		if p, ok := err.(*ParseError); ok {
			if e, ok := p.Err.(*BundleError); ok {
				if e.FieldName != "" {
					t.Errorf("%T: %s", e, e)
				}
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}

	err2 := r.File.Validate()

	if err2 != nil {
		if e, ok := err2.(*FileError); ok {
			if e.FieldName != "BundleCount" {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestX9File validates reading an x9 file
func TestX9File(t *testing.T) {
	f, err := os.Open("./test/data/20180905A.x9")
	if err != nil {
		log.Panicf("Can not open local file: %s: \n", err)
	}
	r := NewReader(f)
	x9File, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}
	// ensure we have a validated file structure
	if x9File.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}
}
