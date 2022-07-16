package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/sdk/client"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/sdk/client/products"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	p, err := c.Products.ListProducts(params)

	assert.NoError(t, err)
	assert.IsType(t, &products.ListProductsOK{}, p)
}
