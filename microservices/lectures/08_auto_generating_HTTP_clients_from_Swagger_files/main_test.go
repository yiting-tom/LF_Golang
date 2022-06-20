package main

import (
	"08_auto_generating_http_clients_from_swagger_files/sdk/client"
	"08_auto_generating_http_clients_from_swagger_files/sdk/client/products"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {
	cfg := client.DefaultTransportConfig().
		WithHost(":9090")

	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prods, err := c.Products.ListProducts(params)

	assert.Nil(t, err)
	assert.Len(t, prods.Payload, 2)
}
