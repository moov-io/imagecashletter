package main

import (
	"io/ioutil"
	"os"
	"testing"
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
	// tmp, err := ioutil.TempFile("", "x9-writeX9-test")
	tmp, err := ioutil.TempFile("", "icl-writeICL-test")
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
