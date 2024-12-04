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
	"github.com/stretchr/testify/require"
)

func TestIssue125(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "icl-valid.json"))
	require.NoError(t, err)

	file, err := FileFromJSON(bs)
	require.NoError(t, err)
	require.NotNil(t, file)

	var buf bytes.Buffer
	require.NoError(t, NewWriter(&buf).Write(file))

	lines := strings.Split(buf.String(), "\n")
	counts := make(map[string]int)
	for i := range lines {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		counts[line[:2]] += 1
	}

	// check each record count
	assert.Equal(t, 1, counts["01"])
	assert.Equal(t, 2, counts["10"])
	assert.Equal(t, 4, counts["20"])
	assert.Equal(t, 4, counts["25"])
	assert.Equal(t, 4, counts["26"])
	assert.Equal(t, 4, counts["27"])
	assert.Equal(t, 4, counts["28"])
	assert.Equal(t, 4, counts["31"])
	assert.Equal(t, 4, counts["32"])
	assert.Equal(t, 4, counts["33"])
	assert.Equal(t, 4, counts["34"])
	assert.Equal(t, 4, counts["35"])
	assert.Equal(t, 8, counts["50"])
	assert.Equal(t, 8, counts["52"])
	assert.Equal(t, 8, counts["54"])
	assert.Equal(t, 4, counts["70"])
	assert.Equal(t, 2, counts["90"])
	assert.Equal(t, 1, counts["99"])
}
