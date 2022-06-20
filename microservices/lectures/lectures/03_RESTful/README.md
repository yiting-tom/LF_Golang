# 03 RESTful

## Project Structure
```go
.
├── data +
│   └── products.go +
├── go.mod
├── go.sum
├── handlers *
│   └── product.go +
├── main.go *
└── README.md
```

## Run the code
- Server
```bash
go run main.go
```

- Client
```bash
curl localhost:9090
curl -d '{"name": "Coca Cola", "description": "A can of Coca Cola", "price": 1}' localhost:9090
curl localhost:9090
```

## data
Provide a repository-liked mock data and functions.

### Define the Product **struct** and **tags**.
```go
type Product struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product
```
- A struct tag with the tag offset with backtick(**`**) characters.
- We can define a tag with encoding method, key and suffix(optional) like: `<method>:"<key>,<suffix>"`, 
- If we need to **ignore** some fields like private fields, we could use `"-"` as the key.
- When displaying it can omit the empty value by using the **suffix** `omitempty`.
The package [validator](https://github.com/go-playground/validator) provides more powerful functions. You must like it 😍.

### Define the ToJSON function
```go
func (p *Products) ToJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(p)
}
```
ToJSON function could serialize the `Products` variable into **JSON** format.

#### [json.NewEncoder](https://pkg.go.dev/encoding/json#NewEncoder)
NewEncoder returns a new encoder that write to the [io.Writer](https://pkg.go.dev/io#Writer)

#### [json.Encoder](https://pkg.go.dev/json#Encoder)
```go
type Encoder struct {
	// contains filtered or unexported fields
}
```
An Encoder writes JSON values to an output stream.

#### [Encoder.Encode](https://pkg.go.dev/encoding/json#Encoder.Encode)
Encode writes the JSON encoding of v to the stream, followed by a newline character.

## handlers/product.go
### Implement the ServeHTTP method
```go
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle GET method.
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// If no method is satisfied return an error
	w.WriteHeader(http.StatusMethodNotAllowed)
}
```
#### [ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter)
```go
type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}
```
A interface which contain:
- `Header() Header`: Header returns the header map that will be sent by WriteHeader.
- `Write([]byte) (int, error)`: Write writes the data to the connection as part of an HTTP reply.
- `WriteHeader(statusCode int)`: WriteHeader sends an HTTP response header with the provided status code.
#### [Header](https://pkg.go.dev/net/http#Header)
```go
type Header map[string][]string
```
- Represents the **key-value pairs** in an HTTP header.
- The keys should be in **canonical form**, as returned by [CanonicalHeaderKey](https://pkg.go.dev/net/http#CanonicalHeaderKey).

### Define hte getProducts function
```go
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fFetch the products from the db.
	productList := data.GetProducts()

	// Serialize the list to JSON.
	if err := productList.ToJSON(w); err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
```