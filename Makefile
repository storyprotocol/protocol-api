LASTEST_COMMIT = $(shell git rev-parse --short HEAD)
TAG ?= ${USER}-local-${LASTEST_COMMIT}

help:
	@echo '  ecr-auth:            - Authenticate ECR'
	@echo '  buildserver:         - Build api server locally'
	@echo '  runserver:           - Run api server locally'
	@echo '  server:              - Build and then run api server locally'

	@echo '  build-{service}:     - Build specific service'
	@echo '  push-{service}:      - Push the current local image for the service to ECR'
	@echo '  lint:                - Run linter'

ecr-auth:
	aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ECR}

buildserver:
	cd api && CGO_ENABLED=0 go build --ldflags "-extldflags '-static -s'" -o build/server cmd/api/main.go

runserver:
	cd api && ./build/server --config=config/local.yaml

server:
	make buildserver && make runserver

build-%:
	docker-compose build $*

push-%: ecr-auth
	docker tag $* ${ECR}/$*:${TAG}
	docker tag $* ${ECR}/$*:latest
	docker push ${ECR}/$*:${TAG}
	docker push ${ECR}/$*:latest	

lint:
	golangci-lint run
