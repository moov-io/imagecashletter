package imagecashletter

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestIssue138(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("test", "testdata", "issue138.json"))
	require.NoError(t, err)
	f, err := FileFromJSON(b)
	require.NoError(t, err)

	// prior to this code change, Write() panicked when writing collated images
	var buf bytes.Buffer
	require.NotPanics(t, func() {
		assert.NoError(t, NewWriter(&buf, WriteVariableLineLengthOption()).Write(f))
	})

	// Makes sure records are written with collated images
	expectedRecordOrder := []string{"01", "10", "20", "25", "50", "52", "50", "52",
		"25", "26", "28", "50", "52", "50", "52", "70", "90", "99"}
	actualRecordOrder := getRecordsList(buf.Bytes())
	assert.Equal(t, expectedRecordOrder, actualRecordOrder)
}

func getRecordsList(f []byte) []string {
	var recordsList []string
	for {
		if len(f) < 6 { // inserted length field (4 bytes) + record type (2 bytes)
			break
		}
		length := int(binary.BigEndian.Uint32(f[0:4]))
		recordType := string(f[4:6])
		recordsList = append(recordsList, recordType)

		f = f[length+4:]
	}

	return recordsList
}
