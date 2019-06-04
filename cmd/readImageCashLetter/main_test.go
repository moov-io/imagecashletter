package main

import "testing"

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
func testFileRead(t testing.TB) {
	main()
}
