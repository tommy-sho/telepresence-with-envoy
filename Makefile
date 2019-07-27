
PROJECT  := $(shell gcloud projects list| sed -n 2P |cut -d' ' -f1)
ZONE     := asia-northeast1-c
CLUSTER  := telepresence-cluster
SERVICES := gateway backend

.PHONY: proto image-build apply delete

proto:
	protoc  \
	--go_out=plugins=grpc:. \
    ./proto/*.proto

init:
	gcloud container clusters create ${CLUSTER} --zone=${ZONE} --num-nodes=3 --preemptible
	gcloud container clusters get-credentials ${CLUSTER} --zone=${ZONE}

build:
	for service in ${SERVICES}; do \
		docker image build -t micro-service/$$service:latest $$service/; \
	done

push: build
	for service in ${SERVICES}; do \
		docker tag micro-service/$$service:latest gcr.io/${PROJECT}/telepresence/$$service:latest;\
		docker push gcr.io/${PROJECT}/telepresence/$$service:latest;\
		sed -i "" -e 's/micro-service/gcr.io\/${PROJECT}\/telepresence/g' ./manifest/deployment-$$service.yaml; \
	done

apply:
	kubectl apply -f manifest/

delete:
	kubectl delete -f manifest/