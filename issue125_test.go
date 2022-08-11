// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIssue125(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "icl-valid.json"))
	if err != nil {
		t.Fatal(err)
	}

	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}
	if file == nil {
		t.Fatal("nil File")
	}

	var buf bytes.Buffer
	if err := NewWriter(&buf).Write(file); err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(buf.String(), "\n")
	counts := make(map[string]int)
	for i := range lines {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		counts[string(line[:2])] += 1
	}

	// check each record count
	assert.Equal(t, counts["01"], 1)
	assert.Equal(t, counts["10"], 2)
	assert.Equal(t, counts["20"], 4)
	assert.Equal(t, counts["25"], 4)
	assert.Equal(t, counts["26"], 4)
	assert.Equal(t, counts["27"], 4)
	assert.Equal(t, counts["28"], 4)
	assert.Equal(t, counts["31"], 4)
	assert.Equal(t, counts["32"], 4)
	assert.Equal(t, counts["33"], 4)
	assert.Equal(t, counts["34"], 4)
	assert.Equal(t, counts["35"], 4)
	assert.Equal(t, counts["50"], 8)
	assert.Equal(t, counts["52"], 8)
	assert.Equal(t, counts["54"], 8)
	assert.Equal(t, counts["70"], 4)
	assert.Equal(t, counts["90"], 2)
	assert.Equal(t, counts["99"], 1)
}
