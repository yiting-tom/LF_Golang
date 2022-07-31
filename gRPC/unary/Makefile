GRPC_PATH := $(PWD)/calculator

gen-route-path:
	@echo "Generating route path..."
	mkdir -p $(GRPC_PATH)
	@echo "Done."

gen-pb: gen-route-path
	@echo "Generating protobuf files..."

	protoc route.proto \
		--proto_path $(GRPC_PATH) \
		--go_out=$(GRPC_PATH) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GRPC_PATH) \
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