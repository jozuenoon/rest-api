
all: api.pb.go api.pb.gw.go api.swagger.json

api.pb.go: api.proto
	protoc -I=. -I=/usr/local/include/ \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
		--gogofaster_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:. \
		--lint_out=. $<

api.pb.gw.go: payment_service.yaml api.proto
	protoc -I=. -I=/usr/local/include/ \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true,grpc_api_configuration=payment_service.yaml,v=2:. api.proto

api.swagger.json: payment_service.yaml api.proto
	protoc -I/usr/local/include -I. \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
		--swagger_out=logtostderr=true,grpc_api_configuration=payment_service.yaml,v=2:. api.proto