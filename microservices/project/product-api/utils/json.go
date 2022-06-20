package utils

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the object to JSON string
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in the io.Reader to the given interface
// Returns ErrInvalidJSON when the JSON string is invalid
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
