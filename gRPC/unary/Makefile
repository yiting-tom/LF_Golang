GRPC_DIR := $(PWD)/calculator/
PROTO_FILE_NAME := calculator.proto

gen-route-path:
	@echo "Generating route path..."
	mkdir -p $(GRPC_DIR)
	@echo "Done."

gen-pb: gen-route-path
	@echo "Generating protobuf files..."

	protoc $(PROTO_FILE_NAME) \
		--proto_path $(GRPC_DIR) \
		--go_out=$(GRPC_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GRPC_DIR) \
		--go-grpc_opt=paths=source_relative

	@echo "Done."

serve:
	@echo "Running server..."
	go run $(PWD)/server/main.go
	@echo "Done."

args=1 + 2
request:
	@echo "Running client..."
	go run $(PWD)/client/main.go $(args)
	@echo "Done."