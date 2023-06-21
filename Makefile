.PHONY: gen-proto 
gen-proto: 
	protoc --go_out=. --go_opt=paths=source_relative dist/proto/*.proto

.PHONY: run-size
run-size:
	go run cmd/size/main.go

.PHONY: run-benchmark
run-benchmark:
	go test -bench=. ./... > out/benchmark.txt