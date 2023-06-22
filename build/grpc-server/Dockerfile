FROM golang:1.19.9-alpine3.18 AS builder

# Install all required dependencies for building
RUN apk update
RUN apk add git build-base
WORKDIR /service

# Install make
RUN apk update && apk upgrade && apk add --update alpine-sdk && apk add --no-cache make

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make compile-grpc-server

FROM alpine:3.14.2
COPY --from=builder ./service/build/out/grpc-app grpc-app
ENTRYPOINT ["./grpc-app"]