all: grpcgen

grpcgen: grpcclean
	protoc --go_out=. \
		--go_opt=module=github.com/jandobki/beqoracle/server \
		--go-grpc_out . \
		--go-grpc_opt=module=github.com/jandobki/beqoracle/server \
		api/beqoracle.proto

grpcclean: 
	rm -rf internal/model/*

.PHONY: \
	grpcgen