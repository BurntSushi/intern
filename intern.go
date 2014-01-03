// Package intern provides a simple interface for interning strings. Strings
// are mapped to integers, which may be used to index a slice. A string interner
// is useful when any particular string key can be used many times (e.g.,
// pairwise data corresponding to all combinations of some string identifiers).
//
// This package's interface supports two different modes of using the provided
// string interner. The simpler and slower approach is to only use the Index
// method, which will automatically intern a string for you if it doesn't
// already exist. The Index method may be called from multiple goroutines
// simultaneously.
//
// The second approach---which ought to be faster---uses the Atomize method
// to intern a bulk of strings initially and the RIndex function to access
// the index of a string. RIndex may not be called concurrently with Atomize
// or Index, but it may be called concurrently with other calls to RIndex.
//
// Note that all three methods---Index, Atomize and RIndex---may be used
// interchangeably, so long as the contracts governing their concurrent use
// are maintained.
package intern

import (
	"fmt"
	"sync"
)

// Interner represents the state of a string interner, mapping strings to
// monotonically increasing integers. A new Interner value *must* be created
// with the NewInterner function.
//
// An Interner may be used from multiple goroutines simultaneously. See the
// documentation for RIndex for a special case.
//
// Interner satisfies the encoding.TextMarshaler and encoding.TextUnmarshaler
// interfaces.
type Interner struct {
	indices map[string]int
	next    int
	lock    *sync.Mutex
}

// NewInterner creates a new interner.
func NewInterner() *Interner {
	return &Interner{make(map[string]int, 1000), 0, new(sync.Mutex)}
}

// Index returns the index of `s` if it exists. If `s` hasn't been interned
// yet, it is interned and its index is returned. Therefore, Index is useful
// when the set of strings to be interned isn't known at program initialization.
func (in *Interner) Index(s string) int {
	in.lock.Lock()
	defer in.lock.Unlock()

	if idx, ok := in.indices[s]; ok {
		return idx
	}
	in.indices[s] = in.next
	in.next++
	return in.next - 1
}

// Atomize interns many strings at once. While it should be faster than a
// series of calls to Index, it is mainly provided as a convenience function
// when the set of strings is known ahead of time.
func (in *Interner) Atomize(ss ...string) {
	in.lock.Lock()
	defer in.lock.Unlock()

	for _, s := range ss {
		if _, ok := in.indices[s]; !ok {
			in.indices[s] = in.next
			in.next++
		}
	}
}

// RIndex provides read-only access to the state of the interner. Namely, if
// `s` hasn't been interned, this function panics. This function is provided
// for performance critical sections of code. Namely, RIndex cannot be called
// concurrently with Index or Atomize. (But multiple goroutines may call RIndex
// simulanteously.)
func (in *Interner) RIndex(s string) int {
	n, ok := in.indices[s]
	if !ok {
		panic(fmt.Sprintf("RIndex called with uninterned string '%s'.", s))
	}
	return n
}

// Len returns the number of strings that have been interned.
func (in *Interner) Len() int {
	in.lock.Lock()
	defer in.lock.Unlock()

	return len(in.indices)
}

// Strings returns a slice copy of all the strings that have been interned.
func (in *Interner) Strings() []string {
	in.lock.Lock()
	defer in.lock.Unlock()

	strs := make([]string, 0, in.Len())
	for str, _ := range in.indices {
		strs = append(strs, str)
	}
	return strs
}
