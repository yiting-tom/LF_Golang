# 10 File handling


## Run the code
- run the server
```sh
make serve
# or by yourself
go run main.go
```

- download file
```sh
make curl-get
# or by yourself
curl localhost:9090/images/1/test.jpeg --output <filename>
```

- upload file
```sh
make curl-post
# or by yourself
curl localhost:9090/images/2/test.jpeg -X POST --data-binary @<filename>
```