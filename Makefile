create-proto:
	protoc --proto_path=api/proto api/proto/*.proto --go_out=pkg/api/
	protoc --proto_path=api/proto api/proto/*.proto --go-grpc_out=pkg/api/

	protoc --proto_path=api/proto api/proto/*.proto --js_out=import_style=commonjs:js/proto/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:js/proto/

clean-proto:
	rm pkg/api/v1/*.go

nginx-cert:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./cert/nginx.key -out ./cert/nginx.cert  -subj '/CN=nginx'

start:
	go run cmd/temi_rpc/server.go

up:
	docker-compose up --build