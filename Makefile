DOCKER_IMAGE_NAME := rest-api
DOCKER_REGISTRY ?= jozuenoon

GIT_BRANCH := $(shell git branch | sed -n '/\* /s///p' 2>/dev/null)
GIT_COMMIT := $(shell git rev-parse HEAD 2>/dev/null)

BUILD_FLAGS := GOOS=linux GOARCH=amd64 CGO_ENABLED=0

.PHONY: bin
bin:
	$(BUILD_FLAGS) go build  -o bin/server cmd/main.go

test:
	ginkgo .
	ginkgo paymentservice

.PHONY: run
run:
	docker run -p8080:8080 -vdb:/db $(DOCKER_IMAGE_NAME):$(GIT_BRANCH)_$(GIT_COMMIT)

swagger-gen:
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
		-i /local/paymentapi/api.swagger.json \
		-l go \
		-o /local/client

build_docker:
	docker build -f cmd/Dockerfile -t $(DOCKER_IMAGE_NAME):$(GIT_BRANCH)_$(GIT_COMMIT) .

api.swagger.md: paymentapi/api.swagger.json
	swagger-markdown -i $< -o $@

api.swagger.pdf: api.swagger.md
	pandoc $<  -o $@