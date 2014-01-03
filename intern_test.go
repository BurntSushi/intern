package intern

import (
	"math/rand"
	"testing"
	"time"
)

const (
	numStrings = 10000
	shortLo = 5
	shortHi = 7
	longLo = 40
	longHi = 50
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestInterner(t *testing.T) {
	type test struct {
		strs []string
		indices map[string]int
		length int
	}
	tests := []test{
		{ []string{}, map[string]int{}, 0 },
		{ nil, nil, 0 },
		{ []string{"a", "b"}, map[string]int{"a": 0, "b": 1}, 2 },
		{ []string{"a", "b", "a"}, map[string]int{"a": 0, "b": 1}, 2 },
	}
	for i, test := range tests {
		in := NewInterner()
		in.Atomize(test.strs...)
		if in.Len() != test.length {
			t.Fatalf("[test %d]: Length should be %d but is %d.",
				i, test.length, in.Len())
		}
		for str, idx := range test.indices {
			if idx != in.Index(str) {
				t.Fatalf("[test %d]: Index for '%s' should be %d but is %d.",
					i, str, idx, in.Index(str))
			}
		}
	}
}

func BenchmarkInternerIndexLong(b *testing.B) {
	in := NewInterner()
	strs := randomStrings(numStrings, longLo, longHi)
	in.Atomize(strs...)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, str := range strs {
			in.Index(str)
		}
	}
}

func BenchmarkInternerRIndexLong(b *testing.B) {
	in := NewInterner()
	strs := randomStrings(numStrings, longLo, longHi)
	in.Atomize(strs...)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, str := range strs {
			in.RIndex(str)
		}
	}
}

func BenchmarkInternerIndexShort(b *testing.B) {
	in := NewInterner()
	strs := randomStrings(numStrings, shortLo, shortHi)
	in.Atomize(strs...)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, str := range strs {
			in.Index(str)
		}
	}
}

func BenchmarkInternerRIndexShort(b *testing.B) {
	in := NewInterner()
	strs := randomStrings(numStrings, shortLo, shortHi)
	in.Atomize(strs...)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, str := range strs {
			in.RIndex(str)
		}
	}
}

func randomStrings(count, min, max int) []string {
	var strs []string
	for i := 0; i < count; i++ {
		strs = append(strs, randomString(min, max))
	}
	return strs
}

func randomString(min, max int) string {
	var str []byte
	length := randomRange(min, max)
	for b := 0; b < length; b++ {
		str = append(str, byte(randomRange(int('A'), int('z'))))
	}
	return string(str)
}

func randomRange(lo, hi int) int {
	return lo + rand.Intn(hi - lo)
}
