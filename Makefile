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
proto_course_front:
	protoc -I protos protos/coursesservice/*.proto \
	--js_out=import_style=commonjs:./service.front/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./service.front/proto
proto_auth_front:
	protoc -I protos protos/authservice/*.proto \
	--js_out=import_style=commonjs:./service.front/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./service.front/proto
	
proto_front: proto_auth_front proto_course_front
proto_backend: proto_authservice proto_coursesservice

proto: proto_backend proto_front

build: build_auth build_courses build_docker_compose

build_auth:
	cd service.auth && make $(build_cmd)
build_courses:
	cd service.courses && make $(build_cmd)
build_docker_compose:
	docker-compose build