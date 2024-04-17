// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"strconv"
	"strings"
	"time"
)

// converters handles golang to imagecashletter type Converters
type converters struct{}

func (c *converters) parseNumField(r string) (s int) {
	s, _ = strconv.Atoi(strings.TrimSpace(r))
	return s
}

func (c *converters) parseStringField(r string) (s string) {
	s = strings.TrimSpace(r)
	return s
}

// stringToBytesField
func (c *converters) stringToBytesField(r string) (b []byte) {
	b = []byte(r)
	return b

}

// formatYYYYMMDDDate takes a time.Time and returns a string of YYYYMMDD
func (c *converters) formatYYYYMMDDDate(t time.Time) string {
	return t.Format("20060102")
}

// parseYYYMMDDDate returns a time.Time when passed time as YYYYMMDD
func (c *converters) parseYYYYMMDDDate(s string) time.Time {
	t, _ := time.Parse("20060102", s)
	return t
}

// formatSimpleTime returns a string of HHMM when  passed a time.Time
func (c *converters) formatSimpleTime(t time.Time) string {
	return t.Format("1504")
}

// parseSimpleTime returns a time.Time when passed a string of HHMM
func (c *converters) parseSimpleTime(s string) time.Time {
	t, _ := time.Parse("1504", s)
	return t
}

// alphaField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	if count := int(max - ln); validSize(count) {
		s += strings.Repeat(" ", count)
	}
	return s
}

// numericField right-justified, unsigned, and zero filled
func (c *converters) numericField(n int, max uint) string {
	s := strconv.Itoa(n)
	ln := uint(len(s))
	if ln > max {
		return s[ln-max:]
	}
	if count := int(max - ln); validSize(count) {
		s = strings.Repeat("0", count) + s
	}
	return s
}

// nbsmField is a numeric-blank/special MICR (NBSM) or numeric-blank/special MICR On-Us (NBSMOS)
// which are right-justified and blank filled
func (c *converters) nbsmField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[ln-max:]
	}
	if count := int(max - ln); validSize(count) {
		s = strings.Repeat(" ", count) + s
	}
	return s
}

// stringField slices to max length and zero filled
func (c *converters) stringField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	if count := int(max - ln); validSize(count) {
		s = strings.Repeat("0", count) + s
	}
	return s
}
