# 08 Auto generate HTTP clients from Swagger files

## Run the code
- [API Client](https://goswagger.io/generate/client.html)
```sh
swagger generate client -f ./swagger.yaml -A product-api -c sdk/client/
```
or by Makefile
```sh
make swagger-generate-client
```