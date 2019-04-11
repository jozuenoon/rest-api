

.PHONY: bin
bin:
	go build  -o bin/server cmd/main.go

test:
	ginkgo .
	ginkgo paymentservice

.PHONY: run
run:
	go run cmd/payment/*.go