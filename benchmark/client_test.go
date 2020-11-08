package benchmark

import "testing"

// go test -bench=BenchmarkServer -benchmem -benchtime=10s
func BenchmarkServer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		postRequest("http://localhost:8080/topics", "test 1")
	}
}

func BenchmarkServer2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		postRequest("http://localhost:8080/topics", "test 1")
	}
}
