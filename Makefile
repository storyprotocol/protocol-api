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
	aws ecr get-login-password --region ${REGION} --profile=${PROFILE} | docker login --username AWS --password-stdin ${ECR}

compileswag:
	swag init -g ./api/cmd/api/main.go -o api/cmd/docs

updatereadme:
	rdme openapi --version=v0.1-beta \
      --key=rdme_xn8s9h3e5cb26106448296f9d6b812c0bf6347db950d48fde9554d4819f7a25b0899d6

buildserver:
	cd api && CGO_ENABLED=0 go build --ldflags "-extldflags '-static -s'" -o build/server cmd/api/main.go

runserver:
	cd api && ./build/server --config=config/local.yaml

server:
	make buildserver && make runserver

build:
	docker-compose build $*

push-%: ecr-auth
	docker tag $* ${ECR}/${REPO}:${TAG}
	docker tag $* ${ECR}/${REPO}:latest
	docker push ${ECR}/${REPO}:${TAG}
	docker push ${ECR}/${REPO}:latest


lint:
	golangci-lint run
