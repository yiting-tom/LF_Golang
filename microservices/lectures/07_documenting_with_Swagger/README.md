# 07_documenting with Swagger

## Project structure

```sh
.
├── data
│   ├── json.go
│   ├── products.go
│   ├── products_test.go
│   └── validation.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── handlers
│   ├── delete.go
│   ├── get.go
│   ├── middleware.go
│   ├── post.go
│   ├── product.go
│   └── put.go
├── main.go
├── Makefile
├── README.md
└── swagger.yaml
```

## Generate the docs
- By Makefile
```sh
make swagger
```

- By CLI
```sh
swagger generate spec -o ./swagger.yaml --scan-models
```

## ReDoc UI
In browser at http://localhost:9090/docs