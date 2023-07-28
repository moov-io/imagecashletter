package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/moov-io/imagecashletter"
)

// TestFileCreate tests creating an ICL File
func TestFileWrite(t *testing.T) {
	testFileWrite(t)
}

/*//BenchmarkTestFileCreate benchmarks creating an ICL File
func BenchmarkTestFileWrite(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileWrite(b)
	}
}*/

// testFileWrite creates an ICL File
func testFileWrite(t testing.TB) {
	tmp, err := os.CreateTemp("", "icl-writeICL-test")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(tmp.Name())

	write(tmp.Name())

	s, err := os.Stat(tmp.Name())
	if err != nil {
		t.Fatal(err.Error())
	}
	if s.Size() <= 0 {
		t.Fatal("expected non-empty file")
	}
}

// TestReadValidationOpts
func TestReadValidationOpts(t *testing.T) {
	tmp, err := os.CreateTemp("", "config")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(tmp.Name())

	f, err := os.Create(tmp.Name())
	if err != nil {
		t.Fatal(err.Error())
	}

	buf, _ := json.Marshal(&imagecashletter.ValidateOpts{})
	f.Write(buf)
	f.Close()

	if opt := readValidationOpts(tmp.Name()); opt == nil {
		t.Fatal("unable to create config")
	}

	if opt := readValidationOpts(""); opt != nil {
		t.Fatal("does not to create any config")
	}

	*flagSkipValidation = true
	if opt := readValidationOpts(""); opt == nil {
		t.Fatal("unable to create config")
	}

}
