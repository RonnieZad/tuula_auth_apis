proto:
	# protoc pkg/pb/*.proto --go-grpc_out=:.
	protoc pkg/pb/*.proto --go_out=:.
	# protoc pkg/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go