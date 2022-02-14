create-proto:
	protoc --proto_path=api/proto api/proto/*.proto --go_out=pkg/api/
	protoc --proto_path=api/proto api/proto/*.proto --go-grpc_out=pkg/api/

	protoc --proto_path=api/proto api/proto/*.proto --js_out=import_style=commonjs:js/proto/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:js/proto/

clean-proto:
	rm pkg/api/v1/*.go

start:
	go run cmd/temi_rpc/server.go

up:
	docker-compose up --build