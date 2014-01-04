package intern

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// MarshalText satisfies the encoding.TextMarshaler interface.
func (in *Interner) MarshalText() ([]byte, error) {
	return json.Marshal(in.interner)
}

// UnmarshalText satisfies the encoding.TextUnmarshaler interface.
func (in *Interner) UnmarshalText(text []byte) error {
	var t interner
	if err := json.Unmarshal(text, &t); err != nil {
		return err
	}
	*in = Interner{t}
	return nil
}

// MarshalBinary satisfies the encoding.BinaryMarshaler interface.
func (in *Interner) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(in.interner); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary satisfies the encoding.BinaryUnmarshaler interface.
func (in *Interner) UnmarshalBinary(text []byte) error {
	var t interner
	dec := gob.NewDecoder(bytes.NewReader(text))
	if err := dec.Decode(&t); err != nil {
		return err
	}
	*in = Interner{t}
	return nil
}

// MarshalText satisfies the encoding.TextMarshaler interface.
func (t *Table) MarshalText() ([]byte, error) {
	return json.Marshal(t.table)
}

// UnmarshalText satisfies the encoding.TextUnmarshaler interface.
func (t *Table) UnmarshalText(text []byte) error {
	var it table
	if err := json.Unmarshal(text, &it); err != nil {
		return err
	}
	*t = Table{it}
	return nil
}

// MarshalBinary satisfies the encoding.BinaryMarshaler interface.
func (t *Table) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(t.table); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary satisfies the encoding.BinaryUnmarshaler interface.
func (t *Table) UnmarshalBinary(text []byte) error {
	var it table
	dec := gob.NewDecoder(bytes.NewReader(text))
	if err := dec.Decode(&it); err != nil {
		return err
	}
	*t = Table{it}
	return nil
}
