all: generate

protoc:
	protoc --go_out=. --go-grpc_out=. .\csi_file.proto
build:
	go build -o app .
docker:
	docker build --tag tungyao-csi:v1 .
docker-no-cache:
	docker build --tag tungyao-csi:v1 --no-cache .
docker-remove:
	docker rmi tungyao-csi:v1
deploy:
	kubectl apply -f csi-deploy.yml
remove:
	kubectl delete po csi-pod
remove-all: build
	-kubectl delete po csi-pod
	-docker rmi tungyao-csi:v1
	-docker build --tag tungyao-csi:v1 .
status:
	@kubectl describe pod csi-pod
log:
	@kubectl logs csi-pod $p
echo:
	echo $a
pvc:
	kubectl apply -f csi-pvc.yaml
pvc-remove:
	kubectl delete -f csi-pvc.yaml
test:
	kubectl apply -f csi-test-pod.yaml