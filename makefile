all: generate

protoc:
	protoc --go_out=. --go-grpc_out=. .\csi_file.proto
build:
	go build -o app .
docker:
	docker build --tag tungyao-csi:v1 .
deploy:
	kubectl apply -f csi-deploy.yml
remove:
	kubectl delete po -n csi-dev csi-pod
status:
	@kubectl describe pod -n csi-dev csi-pod
log:
	@kubectl logs -n csi-dev csi-pod $p
echo:
	echo $a
pvc:
	kubectl apply -f csi-pvc.yaml