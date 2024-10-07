# Template, will be updated when folder structure is finished

run:
	go run ../cmd/server/main.go

test:
	go test ../... -v

fmt:
	go fmt ../...

install-swagger:
	@echo "=== installing go-swagger ==="
	#brew tap go-swagger/go-swagger
	#brew install go-swagger
	GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

gen-swagger:
	@echo "=== cleaning previous swagger ==="
	rm -rf ../internal/swagger/docs/v1/swagger.json

	@echo "=== generating swagger ==="

	@echo "=== version 1 ==="
		swagger generate spec -w ../api/v1/http/swagger -o ../internal/swagger/docs/v1/swagger.json --scan-models

install-grpc:
	@echo "=== installing protoc-grpc ==="
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc/latest
	brew install protobuf

gen-grpc:
	@echo "=== generating grpc ==="

	@echo "=== generating 'gateway_service_gateway.proto' ==="
	#protoc -I=../api/v1/grpc/gateway_service_gateway/proto --go-grpc_out=paths=source_relative:../api/v1/grpc/gateway_service_gateway --go_out=paths=source_relative:../api/v1/grpc/gateway_service_gateway ../api/v1/grpc/gateway_service_gateway/proto/gateway_service_gateway.proto

install:
	pre-commit install
