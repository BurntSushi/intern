// Package intern provides a simple interface for interning strings. Strings
// are mapped to integers, which may be used to index a slice. A string interner
// is useful when any particular string key can be used many times (e.g.,
// pairwise data corresponding to all combinations of some string identifiers).
package intern

// Atom represents a unique identifier for a particular string.
type Atom int

// Interner represents the state of a string interner, mapping strings to
// monotonically increasing integers. A new Interner value *must* be created
// with the NewInterner function.
//
// An Interner may not be used from multiple goroutines simultaneously.
//
// Interner satisfies all interfaces defined in the `encoding` standard library
// package.
type Interner struct {
	interner
}

// interner is the unexported representation of an Interner. The indirection
// is used for encoding/decoding.
type interner struct {
	Atms map[string]Atom
	Next Atom
}

// NewInterner creates a new interner.
func NewInterner() *Interner {
	return &Interner{interner{make(map[string]Atom, 1000), 0}}
}

// Atom returns the atom of `s` if it exists. If `s` hasn't been interned
// yet, it is interned and its atom is returned. Therefore, Atom is useful
// when the set of strings to be interned isn't known at program initialization.
func (in *Interner) Atom(s string) Atom {
	if a, ok := in.Atms[s]; ok {
		return a
	}
	in.Atms[s] = in.Next
	in.Next++
	return in.Next - 1
}

// Atoms interns many strings at once. While it should be faster than a
// series of calls to Atom, it is mainly provided as a convenience function
// when the set of strings is known ahead of time.
func (in *Interner) Atoms(ss ...string) []Atom {
	atoms := make([]Atom, len(ss))
	for i, s := range ss {
		atom, ok := in.Atms[s]
		if !ok {
			atom = in.Next
			in.Next++
			in.Atms[s] = atom
		}
		atoms[i] = atom
	}
	return atoms
}

// Exists returns true if and only if the given string has been interned.
func (in *Interner) Exists(s string) bool {
	_, ok := in.Atms[s]
	return ok
}

// Len returns the number of strings that have been interned.
func (in *Interner) Len() int {
	return len(in.Atms)
}

// Strings returns a slice copy of all the strings that have been interned.
func (in *Interner) Strings() []string {
	strs := make([]string, 0, in.Len())
	for str, _ := range in.Atms {
		strs = append(strs, str)
	}
	return strs
}
