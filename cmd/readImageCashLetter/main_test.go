package main

import (
	"path/filepath"
	"testing"
)

func TestFileRead(t *testing.T) {
	testFileRead(t)
}

/*//BenchmarkTestFileCreate benchmarks creating an ICL File
func BenchmarkTestFileRead(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileRead(b)
	}
}*/

// FileCreate creates an ICL File
func resetFlags() {
	*fPath = "BNK20181015-A.icl"
	*flagJson = false
	*flagSkipValidation = false
	*flagEbcdic = false
	*flagVariableLength = false
}

func testFileRead(t testing.TB) {
	resetFlags()
	if err := readFile(*fPath); err != nil {
		t.Fatal(err)
	}
}

func TestReadTestdataEncodings(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		ebcdic   bool
		varlen   bool
		wantFail bool
	}{
		{name: "ascii variable-length", path: filepath.Join("..", "..", "test", "testdata", "valid-ascii.x937"), varlen: true},
		{name: "ebcdic variable-length", path: filepath.Join("..", "..", "test", "testdata", "valid-ebcdic.x937"), ebcdic: true, varlen: true},
		{name: "ascii without flags", path: filepath.Join("..", "..", "test", "testdata", "valid-ascii.x937"), wantFail: true},
		{name: "ebcdic without flags", path: filepath.Join("..", "..", "test", "testdata", "valid-ebcdic.x937"), wantFail: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			*flagEbcdic = tt.ebcdic
			*flagVariableLength = tt.varlen
			err := readFile(tt.path)
			if tt.wantFail {
				if err == nil {
					t.Fatalf("expected error reading %s without the required flags", tt.path)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
