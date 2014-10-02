package nettools

import (
	"testing"
)

var dottedPortTests = map[string]string{
	"97.98.99.100:25958":                              "abcdef",
	"[6162:6364:6566:6768:6970:7172:7374:7576]:25958": "abcdefghipqrstuvef",
	"[6162:64:6566:6768:6970:7172:7374:7576]:25958":   "ab\x00defghipqrstuvef",
}

func TestDottedPort(t *testing.T) {
	for k, v := range dottedPortTests {
		s := DottedPortToBinary(k)
		if s != v {
			t.Fatalf("DottedPortToBinary got %v wanted %v", s, v)
		}
	}
}

func TestBinaryToDottedPort(t *testing.T) {
	for k, v := range dottedPortTests {
		s := BinaryToDottedPort(v)
		if s != k {
			t.Fatalf("BinaryToDottedPort got %v wanted %v", s, k)
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
// BenchmarkDottedPort	 5000000	       683 ns/op
