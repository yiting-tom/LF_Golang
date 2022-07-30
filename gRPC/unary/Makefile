gen-pb:
	@echo "Generating protobuf files..."
	protoc --go_out=.\
		--go_opt=paths=source_relative\
		--go-grpc_out=.\
		--go-grpc_opt=paths=source_relative\
		route.proto
	@echo "Done."

run-server:
	@echo "Running server..."
	go run ./route-server/server.go
	@echo "Done."

run-client:
	@echo "Running client..."
	go run ./main.go
	@echo "Done."