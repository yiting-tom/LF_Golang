ROUTE_PATH := $(PWD)/route

gen-pb:
	@echo "Generating protobuf files..."

	protoc route.proto \
		--proto_path $(ROUTE_PATH) \
		--go_out=$(ROUTE_PATH) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(ROUTE_PATH) \
		--go-grpc_opt=paths=source_relative

	@echo "Done."

run-server:
	@echo "Running server..."
	go run ./route-server/server.go
	@echo "Done."

run-client:
	@echo "Running client..."
	go run ./main.go
	@echo "Done."