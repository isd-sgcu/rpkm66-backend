proto:
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. user.proto
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. group.proto
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. file.proto
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. checkin.proto
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. baan.proto
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. event.proto

test:
	go vet ./...
	go test  -v -coverpkg ./src/app/... -coverprofile coverage.out -covermode count ./src/app/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

server:
	go run ./src/.

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

seed:
	go run ./src/. seed