all: generate

protoc:
	protoc --go_out=. --go-grpc_out=. .\csi_file.proto

package:
	docker build --tag tungyao/csi-dev:v0.1 .