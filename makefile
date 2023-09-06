docker:
	go build .
	docker build -t tungyao/nfs-csi:v1 .
docker-no-cachge:
	docker build -t tungyao/nfs-csi:v1 --no-cache .
deploy:
	kubectl apply -f ./sc/deploy.yaml
delete:
	kubectl delete -f ./sc/deploy.yaml
logs:
	kubectl logs nfs -f
pvc:
	kubectl apply -f ./sc/pvc.yaml
pvc-delete:
	kubectl delete pvc pvc
pod:
	kubectl apply -f ./sc/pod.yaml
pod-delete:
	kubectl delete -f ./sc/pod.yaml