docker:
	docker build -t sim.tungyao.nfs:v1 .
docker-no-cachge:
	docker build -t sim.tungyao.nfs:v1 --no-cache .
deploy:
	kubectl apply -f deploy.yaml