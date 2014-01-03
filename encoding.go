package intern

import (
	"encoding/json"
	"sync"
)

type transport struct {
	Indices map[string]int
	Next    int
}

// MarshalText satisfies the encoding.TextMarshaler interface.
func (in *Interner) MarshalText() ([]byte, error) {
	in.lock.Lock()
	defer in.lock.Unlock()

	return json.Marshal(transport{in.indices, in.next})
}

// UnmarshalText satisfies the encoding.TextUnmarshaler interface.
func (in *Interner) UnmarshalText(text []byte) error {
	var t transport
	if err := json.Unmarshal(text, &t); err != nil {
		return err
	}
	*in = Interner{t.Indices, t.Next, new(sync.Mutex)}
	return nil
}
