serve:
	go run main.go

curl-get:
	@echo "Testing get 'local-storage/1/test.jpg' by curl"
	@echo ...
	curl localhost:9090/images/1/test.jpg --output ./curl-got.jpg

curl-post:
	@echo "Testing post 'curl-got.jpg' by curl"
	@echo ...
	curl localhost:9090/images/2/test.jpg -X POST --data-binary @curl-got.jpg