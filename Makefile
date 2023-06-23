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

.PHONY: compile-rest-server
compile-rest-server:
	go build -a -o ./build/out/rest-app cmd/rest-app/main.go

.PHONY: run-rest-server
run-rest-server:
	go run cmd/rest-app/main.go

.PHONY: package-rest-server
package-rest-server:
	docker build --file build/rest-server/Dockerfile -t hiteshpattanayak/rest-app:1.0 .

.PHONY: publish-rest-server
publish-rest-server:
	docker push hiteshpattanayak/rest-app:1.0

.PHONY: run-grpc-server
run-grpc-server:
	go run cmd/grpc-app/main.go

.PHONY: compile-grpc-server
compile-grpc-server:
	go build -a -o ./build/out/grpc-app cmd/grpc-app/main.go

.PHONY: package-grpc-server
package-grpc-server:
	docker build --file build/grpc-server/Dockerfile -t hiteshpattanayak/grpc-app:2.0 .

.PHONY: publish-grpc-server
publish-grpc-server:
	docker push hiteshpattanayak/grpc-app:2.0

.PHONY: run-gateway
run-gateway:
	go run cmd/gateway-app/main.go

.PHONY: compile-gateway-app
compile-gateway-app:
	go build -a -o ./build/out/gateway-app cmd/gateway-app/main.go

.PHONY: package-gateway-app
package-gateway-app:
	docker build --file build/gateway-app/Dockerfile -t hiteshpattanayak/gateway-app:1.0 .

.PHONY: publish-gateway-app
publish-gateway-app:
	docker push hiteshpattanayak/gateway-app:1.0