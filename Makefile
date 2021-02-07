.PHONY: proto build

build_cmd := build_for_docker

proto_authservice:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	protos/authservice/auth_service.proto

proto_coursesservice:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	protos/coursesservice/coursesservice.proto

proto: proto_authservice proto_coursesservice

build: build_auth build_courses

build_auth:
	cd service.auth && make $(build_cmd)
build_courses:
	cd service.courses && make $(build_cmd)