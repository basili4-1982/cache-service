protogen:
	 rm -rf ./internal/proto
	 mkdir -p ./internal/proto
	 protoc --proto_path=./proto \
	 --go_out=./internal \
	 --go-grpc_out=./internal \
	 ./proto/**/*.proto

up:
	./docker/up-stage.sh