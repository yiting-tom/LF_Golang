package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "cool stuff",
		Price: 1.00,
		SKU:   "abc-def-ghi",
	}

	if err := p.Validate(); err != nil {
		t.Fatal(err)
	}
}
