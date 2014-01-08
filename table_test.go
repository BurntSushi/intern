package intern

import "testing"

const numAtoms = 100

func TestTable(t *testing.T) {
	type test map[[2]string]float64
	tests := []test{
		map[[2]string]float64{
			[2]string{"a", "b"}: 1,
			[2]string{"b", "a"}: 1,
			[2]string{"a", "c"}: 2,
			[2]string{"z", "d"}: 3,
			[2]string{"y", "x"}: 4,
		},
	}
	for ti, test := range tests {
		tab := NewTable(0)
		for key, val := range test {
			tab.Set(tab.Atom(key[0]), tab.Atom(key[1]), val)
		}
		for key, val := range test {
			tabVal := tab.Get(tab.Atom(key[0]), tab.Atom(key[1]))
			if val != tabVal {
				t.Fatalf("[test %d]: Value for '%s' should be %d but is %d.",
					ti, key, val, tabVal)
			}
		}
	}
}

func BenchmarkTable(b *testing.B) {
	strs := randomStrings(numAtoms, 5, 7)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tab := NewTable(0)
		for j := 0; j < len(strs); j++ {
			for k := j + 1; k < len(strs); k++ {
				tab.Set(tab.Atom(strs[j]), tab.Atom(strs[k]), float64(i*j))
			}
		}
		for z := 0; z < 10; z++ {
			for j := 0; j < len(strs); j++ {
				for k := j + 1; k < len(strs); k++ {
					_ = tab.Get(tab.Atom(strs[j]), tab.Atom(strs[k]))
				}
			}
		}
	}
}

func BenchmarkTableAtoms(b *testing.B) {
	strs := randomStrings(numAtoms, 5, 7)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tab := NewTable(0)
		astrs := tab.Atoms(strs...)
		for j := 0; j < len(astrs); j++ {
			for k := j + 1; k < len(astrs); k++ {
				tab.Set(astrs[j], astrs[k], float64(i*j))
			}
		}
		for z := 0; z < 10; z++ {
			for j := 0; j < len(astrs); j++ {
				for k := j + 1; k < len(astrs); k++ {
					_ = tab.Get(astrs[j], astrs[k])
				}
			}
		}
	}
}
