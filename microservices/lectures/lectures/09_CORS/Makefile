.DEFAULT_GOAL := swagger

install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	@echo Ensure you have the swagger CLI or this command will fail.
	@echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@echo ....

	@echo remove old swagger files
	rm ./swagger.yaml

	@echo ...
	swagger generate spec -o ./swagger.yaml --scan-models

swagger-generate-client:
	make swagger
	@echo generate swagger client
	@echo ...
	swagger generate client -f ./swagger.yaml -A product-api -c sdk/client/
