package nettools

import (
	"testing"
)

var dottedPortTests = map[string]string{"97.98.99.100:25958": "abcdef"}

func TestDottedPort(t *testing.T) {
	for k, v := range dottedPortTests {
		s := DottedPortToBinary(k)
		if s != v {
			t.Fatalf("DottedPortToBinary got %v wanted %v", s, v)
		}
	}

}

func BenchmarkDottedPort(t *testing.B) {
	for i := 0; i < t.N; i++ {
		for k := range dottedPortTests {
			DottedPortToBinary(k)
		}

	}
}

// # MacBookAir 1.7Ghz i5
// $ go test -bench=.*
// BenchmarkDottedPort	  500000	      6645 ns/op
