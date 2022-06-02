package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		ID:          0,
		Name:        "",
		Description: "test desc",
		Price:       1.22,
		SKU:         "aaa-aaa-aaa",
	}

	v := NewValidation()
	errs := v.Validate(p)
	assert.Len(t, errs, 1)
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		ID:          0,
		Name:        "test name",
		Description: "test desc",
		SKU:         "aaa-aaa-aaa",
	}

	v := NewValidation()
	errs := v.Validate(p)
	assert.Len(t, errs, 1)
}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		ID:          0,
		Name:        "test name",
		Description: "test desc",
		Price:       1.22,
		SKU:         "aaa-aaa",
	}

	v := NewValidation()
	errs := v.Validate(p)
	assert.Len(t, errs, 1)
}

func TestValidProductDoesNOTReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc-efg-hji",
	}

	v := NewValidation()
	errs := v.Validate(p)
	assert.Nil(t, errs)
}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		{
			Name: "abc",
		},
	}

	b := bytes.NewBufferString("")
	errs := ToJSON(ps, b)
	assert.NoError(t, errs)
}
