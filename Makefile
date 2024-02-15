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

buildserver:
	cd api && CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build --ldflags "-extldflags '-static -s'" -o build/server cmd/api/main.go

runserver:
	cd api && ./build/server --config=config/local.yaml

server:
	make buildserver && make runserver

build:
	docker compose build $*

push-%: ecr-auth
	docker tag $* ${ECR}/${REPO}:${TAG}
	docker tag $* ${ECR}/${REPO}:latest
	docker push ${ECR}/${REPO}:${TAG}
	docker push ${ECR}/${REPO}:latest


lint:
	golangci-lint run

push-stg:
	PROFILE=staging REPO=story-stg-example REGION=us-west-2 ECR=478656756051.dkr.ecr.us-west-2.amazonaws.com make push-api

push-prod:
	PROFILE=prod REPO=story-prd-edge REGION=us-west-2 ECR=243963068353.dkr.ecr.us-west-2.amazonaws.com make push-api