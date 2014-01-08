package intern

import (
	"github.com/BurntSushi/ty"
	"math"

	"reflect"
)

// Table corresponds to a dense square table of data, where values are floats 
// that correspond to pairwise string keys. Note that the table must by
// symmetric so that for all k1, k2 in Table, then
// Table[k1, k2] == Table[k2, k1].
//
// A Table may not be used from multiple goroutines simultaneously.
//
// Table satisfies all interfaces defined in the `encoding` standard library
// package.
type Table struct {
	table
}

// table is the unexported representation of a Table. The indirection
// is used for encoding/decoding.
type table struct {
	In       *Interner
	Table    []float64 // laid out in row-major order
	CapAtoms int
}

// NewTable creates a new dense table keyed by string pairs with float values.
// numAtoms is a "hint" for the length of the table. It may be zero.
func NewTable(numAtoms int) *Table {
	if numAtoms < 0 {
		numAtoms = 0
	}
	return &Table{table{
		NewInterner(),
		make([]float64, numAtoms*numAtoms),
		numAtoms,
	}}
}

// NewTableInterner creates a new dense table from a pre-existing string
// interner.
func NewTableInterner(in *Interner) *Table {
	return &Table{table{
		in,
		make([]float64, in.Len()*in.Len()),
		in.Len(),
	}}
}

// Get retrieves the value corresponding to the string pair (as atoms) given.
// The order of the string pair does not matter.
func (t *Table) Get(a1, a2 Atom) float64 {
	return t.Table[t.index(a1, a2)]
}

// Set sets the value corresponding to the string pair (as atoms) given.
// The order of the string pair does not matter.
func (t *Table) Set(a1, a2 Atom, v float64) {
	t.Table[t.index(a1, a2)] = v
	t.Table[t.index(a2, a1)] = v
}

// Atom interns a string and returns an Atom that may be used in the Get and
// Set methods.
func (t *Table) Atom(s string) Atom {
	a := t.In.Atom(s)
	if int(a) >= t.CapAtoms {
		table, capAtoms := ExpandSquareTable(t.Table, 1+int(a)-t.CapAtoms)
		t.Table, t.CapAtoms = table.([]float64), capAtoms
	}
	return a
}

// Atoms is a convenience function to intern many strings at once. The slice
// returned is in correspondence to the strings given.
func (t *Table) Atoms(ss ...string) []Atom {
	atoms := make([]Atom, len(ss))
	for i, s := range ss {
		atoms[i] = t.Atom(s)
	}
	return atoms
}

func (t *Table) index(row, col Atom) int {
	return int(row)*t.CapAtoms + int(col)
}

// ExpandSquareTable has a parametric type:
//
//	func ExpandSquareTable([]A, int) ([]A, int)
//
// ExpandSquareTable takes any slice holding a square table of data (row-major)
// and expands the length of the table to at least the length provided. A new
// slice is returned (with data from `slice` copied to it) along with the
// length of the table.
//
// This function is exported so that you may use it to build your own dense
// tables (since the Table in this package can only store floats).
func ExpandSquareTable(slice interface{}, leastLen int) (interface{}, int) {
	chk := ty.Check(
		new(func([]ty.A, int) []ty.A),
		slice, leastLen)
	vslice, tslice := chk.Args[0], chk.Returns[0]

	oldLen := int(math.Sqrt(float64(vslice.Cap())))
	newLen := oldLen * 2
	if newLen < leastLen {
		newLen = leastLen
	}
	rslice := reflect.MakeSlice(tslice, newLen*newLen, newLen*newLen)
	for r := 0; r < oldLen; r++ {
		for c := 0; c < oldLen; c++ {
			rslice.Index(r*newLen + c).Set(vslice.Index(r*oldLen + c))
		}
	}
	return rslice.Interface(), newLen
}
