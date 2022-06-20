package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1.22,
		Name:  "test-name",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: -1,
		SKU:   "abc-efg-hji",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestValidProductDoesNOTReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc-efg-hji",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Nil(t, err)
}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		{
			Name:  "abc",
			Price: 1.22,
			SKU:   "abc-efg-hji",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(ps, b)
	assert.NoError(t, err)
}
