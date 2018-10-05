// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

// TestX9FileRead validates reading an x9 file
func TestX9FileRead(t *testing.T) {
	f, err := os.Open("./test/testdata/20180905A.x9")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if err != nil {
		if p, ok := err.(*ParseError); ok {
			if e, ok := p.Err.(*BundleError); ok {
				if e.FieldName != "entries" {
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
	f, err := os.Open("./test/testdata/20180905A.x9")
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

// TestRecordTypeUnknown validates record type unknown
func TestRecordTypeUnknown(t *testing.T) {
	var line = "1735T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

//TestFileLineShort validates file line is short
func TestFileLineShort(t *testing.T) {
	line := "1 line is only 70 characters ........................................!"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()

	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.FieldName != "RecordLength" {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", e, e)
		}
	}
}

// TestFileFileHeaderErr validates error flows back from the parser
func TestFileFileHeaderErr(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = ""
	r := NewReader(strings.NewReader(fh.String()))
	// necessary to have a file control not nil
	r.File.Control = mockFileControl()
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTwoFileHeaders validates one file header
func TestTwoFileHeaders(t *testing.T) {
	var line = "0135T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	var twoHeaders = line + "\n" + line
	r := NewReader(strings.NewReader(twoHeaders))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.Msg != msgFileControl {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestCashLetterHeaderErr validates error flows back from the parser
func TestCashLetterHeaderErr(t *testing.T) {
	clh := mockCashLetterHeader()
	clh.DestinationRoutingNumber = ""
	r := NewReader(strings.NewReader(clh.String()))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestCashLetterHeaderDuplicate validates when two CashLetterHeader exists in a current CashLetter
func TestCashLetterHeaderDuplicate(t *testing.T) {
	// create a new CashLetter header string
	clh := mockCashLetterHeader()
	r := NewReader(strings.NewReader(clh.String()))
	// instantiate a CashLetter in the reader
	r.addCurrentCashLetter(NewCashLetter(clh))
	// read should fail because it is parsing a second CashLetter Header and there can only be one.
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.Msg != msgFileCashLetterInside {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBundleHeaderErr validates error flows back from the parser
func TestBundleHeaderErr(t *testing.T) {
	bh := mockBundleHeader()
	bh.DestinationRoutingNumber = ""
	r := NewReader(strings.NewReader(bh.String()))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBundleHeaderDuplicate validates when two BundleHeader exists in a current Bundle
func TestBundleHeaderDuplicate(t *testing.T) {
	// create a new CashLetter header string
	bh := mockBundleHeader()
	r := NewReader(strings.NewReader(bh.String()))
	// instantiate a CashLetter in the reader
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bhTwo := mockBundleHeader()
	r.addCurrentBundle(NewBundle(bhTwo))
	// read should fail because it is parsing a second CashLetter Header and there can only be one.
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.Msg != msgFileCashLetterInside {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}