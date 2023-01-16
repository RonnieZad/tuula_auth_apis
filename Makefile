proto:
	# protoc pkg/pb/*.proto --go-grpc_out=:.
	# protoc pkg/pb/*.proto --go_out=:.
	# protoc pkg/pb/*.proto --go_out=plugins=grpc:.
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/pb/*.proto

server:
	go run cmd/main.go