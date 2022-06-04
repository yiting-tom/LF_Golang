serve:
	go run main.go

curl-get:
	@echo "Testing get 'local-storage/1/test.jpeg' by curl"
	@echo ...
	curl localhost:9090/images/1/test.jpeg --output ./curl-got.jpeg

curl-post:
	@echo "Testing post 'curl-got.jpeg' by curl"
	@echo ...
	curl localhost:9090/images/2/test.jpeg -X POST --data-binary @curl-got.jpeg