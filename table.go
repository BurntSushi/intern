package intern

import (
	"github.com/BurntSushi/ty"
	"math"

	"reflect"
)

type Table struct {
	table
}

type table struct {
	In       *Interner
	Table    []float64 // laid out in row-major order
	CapAtoms int
}

func NewTable(numAtoms int) *Table {
	return &Table{table{
		NewInterner(),
		make([]float64, numAtoms*numAtoms),
		numAtoms,
	}}
}

func (t *Table) Get(a1, a2 Atom) float64 {
	return t.Table[t.index(a1, a2)]
}

func (t *Table) Set(a1, a2 Atom, v float64) {
	t.Table[t.index(a1, a2)] = v
	t.Table[t.index(a2, a1)] = v
}

func (t *Table) Atom(s string) Atom {
	a := t.In.Atom(s)
	if int(a) >= t.CapAtoms {
		table, capAtoms := expandSquareTable(t.Table, 1+int(a)-t.CapAtoms)
		t.Table, t.CapAtoms = table.([]float64), capAtoms
	}
	return a
}

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

func expandSquareTable(slice interface{}, leastLen int) (interface{}, int) {
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
