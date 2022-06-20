package data

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the given interface into JSON.
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

// FromJSON deserializes the given JSON into the given interface.
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)

	return d.Decode(i)
}
