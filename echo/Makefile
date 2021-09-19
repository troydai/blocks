build:
	@protoc \
		--go_out=. \
		--go_opt=module=github.com/troydai/blocks/echo \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/troydai/blocks/echo \
		idl/echo.proto
		