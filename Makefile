.PHONY: gen-proto 
gen-proto: 
	protoc --go_out=. --go_opt=paths=source_relative dist/proto/*.proto
	protoc --go_out=internal/pkg --go-grpc_out=internal/pkg internal/pkg/proto/*.proto

.PHONY: run-size
run-size:
	go run cmd/size/main.go

.PHONY: run-benchmark
run-benchmark:
	go test -bench=. ./... > out/benchmark.txt

.PHONY: run-rest-server
run-rest-server:
	go run cmd/rest-app/main.go

.PHONY: run-grpc-server
run-grpc-server:
	go run cmd/grpc-app/main.go

.PHONY: run-gateway
run-gateway:
	go run cmd/gateway-app/main.go