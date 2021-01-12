.PHONY: all proto
all: proto build_for_docker
build_for_docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
proto:
    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/auth_service.proto
   