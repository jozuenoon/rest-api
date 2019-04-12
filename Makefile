

.PHONY: bin
bin:
	go build  -o bin/server cmd/main.go

test:
	ginkgo .
	ginkgo paymentservice

.PHONY: run
run:
	go run cmd/*.go

swagger-gen:
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
		-i /local/paymentapi/api.swagger.json \
		-l go \
		-o /local/client