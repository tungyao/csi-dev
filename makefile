all: generate

generate:
	protoc --go_out=. --go-grpc_out=. .\csi_file.proto