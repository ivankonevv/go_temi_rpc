create-proto:
	protoc --proto_path=api/proto api/proto/*.proto --go_out=pkg/api/
	protoc --proto_path=api/proto api/proto/*.proto --go-grpc_out=pkg/api/

	protoc --proto_path=api/proto api/proto/*.proto --js_out=import_style=commonjs:js/proto/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:js/proto/

clean-proto:
	rm pkg/api/v1/*.go


create-cert_rpc:
	docker-compose up -d nginx && docker-compose run --rm  certbot certonly --webroot --webroot-path /var/www/certbot/ -v --dry-run -d rpc.nbmfscafdcasc.xyz

create-cert_api:
	docker-compose up -d nginx && docker-compose run --rm  certbot certonly --webroot --webroot-path /var/www/certbot/ -v --dry-run -d api.nbmfscafdcasc.xyz

create-cert_vue:
	docker-compose up -d nginx && docker-compose run --rm  certbot certonly --webroot --webroot-path /var/www/certbot/ -v --dry-run -d nbmfscafdcasc.xyz


run-with-renew:
	docker compose run --rm certbot renew

start:
	go run cmd/temi_rpc/server.go

up:
	docker-compose up --build