## Demos:

### Serialization/Deserialization

make run-benchmark

make run-size

### TCP - Wireshark

sudo wireshark

make run-rest-server
make run-gateway


make run-grpc-server
make run-gateway --> No handshake

make run-gateway --> show handshake

### REST polling vs gRPC streaming

make run-rest-server
make run-grpc-server
make run-gateway


curl --request GET   --url http://localhost:9092/rest/hello -w "\n"
curl --request GET   --url http://localhost:9092/grpc/hello -w "\n"

### Minikube deploy

- kubectl create -f build/rest-server/deployment.yaml
- kubectl create -f build/rest-server/service.yaml

- kubectl create -f build/grpc-server/deployment.yaml
- kubectl create -f build/grpc-server/service.yaml

- kubectl create -f build/gateway/deployment.yaml
- kubectl create -f build/gateway/service.yaml
- kubectl create -f build/gateway/ingress.yaml