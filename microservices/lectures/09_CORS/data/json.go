package data

import (
	"encoding/json"
	"io"
)

// ToJSON encode the given interface to a string based JSON format.
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

// FromJSON decode the given string based JSON format
// in an io.Reader to the given interface.
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)

	return d.Decode(i)
}
