build:
	@protoc \
		--go_out=. \
		--go_opt=module=github.com/troydai/blocks/echo \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/troydai/blocks/echo \
		idl/echo.proto
	@mkdir -p bin
	@go build -o bin/echo-server .
	@go build -o bin/echo-client ./client

run-client:
	@go build -o bin/echo-client ./client
	./bin/echo-client "`date`"
	

