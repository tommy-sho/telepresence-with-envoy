
SERVICES := gateway backend

.PHONY: proto image-build push apply delete

proto:
	protoc  \
	--go_out=plugins=grpc:. \
    ./proto/*.proto
	cp ./proto/*.pb.go backend/genproto
	cp ./proto/*.pb.go gateway/genproto


image-build:
	for service in ${SERVICES}; do \
		docker image build -t gcr.io/my-first-project-236315/grpc-test/$$service:latest $$service/; \
	done

push:
	for service in ${SERVICES}; do \
		docker push gcr.io/my-first-project-236315/grpc-test/$$service:latest; \
	done

apply:
	kubectl apply -f kubernetes-manifest/deployment-and-service

delete:
	kubectl delete -f kubernetes-manifest/deployment-and-service
