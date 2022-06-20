package data

import (
	"bytes"
	"testing"

	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/utils"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1.22,
		SKU:   "abc-123-abc",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		Name: "abc",
		SKU:  "abc-123-abc",
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
	assert.Len(t, err, 1)
}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		{
			Name: "abc",
		},
	}

	b := bytes.NewBufferString("")
	err := utils.ToJSON(ps, b)
	assert.NoError(t, err)
}
