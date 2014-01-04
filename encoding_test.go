package intern

import (
	"reflect"
	"testing"
)

func TestInternerText(t *testing.T) {
	in := NewInterner()
	in.Atoms("a", "b", "c")
	out := new(Interner)

	bs, err := in.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if err := out.UnmarshalText(bs); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("In != out: %v != %v", in, out)
	}
}

func TestInternerBinary(t *testing.T) {
	in := NewInterner()
	in.Atoms("a", "b", "c")
	out := new(Interner)

	bs, err := in.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if err := out.UnmarshalBinary(bs); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("In != out: %v != %v", in, out)
	}
}

func TestTableText(t *testing.T) {
	in := NewTable(0)
	atoms := in.Atoms("a", "b", "c")
	in.Set(atoms[0], atoms[1], 1)
	in.Set(atoms[0], atoms[2], 2)
	in.Set(atoms[1], atoms[2], 3)
	out := new(Table)

	bs, err := in.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if err := out.UnmarshalText(bs); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(in.In, out.In) {
		t.Fatalf("In != out: %v != %v", in, out)
	}
	if !reflect.DeepEqual(in.Table, out.Table) || in.CapAtoms != out.CapAtoms {
		t.Fatalf("In != out: %v != %v", in, out)
	}
}

func TestTableBinary(t *testing.T) {
	in := NewTable(0)
	atoms := in.Atoms("a", "b", "c")
	in.Set(atoms[0], atoms[1], 1)
	in.Set(atoms[0], atoms[2], 2)
	in.Set(atoms[1], atoms[2], 3)
	out := new(Table)

	bs, err := in.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if err := out.UnmarshalBinary(bs); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(in.In, out.In) {
		t.Fatalf("In != out: %v != %v", in, out)
	}
	if !reflect.DeepEqual(in.Table, out.Table) || in.CapAtoms != out.CapAtoms {
		t.Fatalf("In != out: %v != %v", in, out)
	}
}
