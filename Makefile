protogen:
	 rm -rf ./internal/proto
	 mkdir -p ./internal/proto
	 protoc --proto_path=./proto \
	 --go_out=./internal \
	 --go-grpc_out=./internal \
	 ./proto/**/*.proto

up-cache:
	docker run  -p 11211:11211  -d memcached memcached -m 64


start: up-cache
	   go mod tidy
	   go run ./cmd/main.go