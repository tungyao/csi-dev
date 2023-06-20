docker:
	go build .
	docker build -t sim.tungyao.nfs:v1 .
docker-no-cachge:
	docker build -t sim.tungyao.nfs:v1 --no-cache .
deploy:
	kubectl apply -f deploy.yaml
delete:
	kubectl delete po nfs
logs:
	kubectl logs nfs -f
pvc:
	kubectl apply -f pvc.yaml
pvc-delete:
	kubectl delete pvc pvc