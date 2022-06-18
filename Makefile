proto:
	protoc --proto_path=src/proto --go_out=plugins=grpc:. user.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. dto.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. common.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. contact.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. location.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. team.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. organization.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. role.proto
	protoc --proto_path=src/proto --go_out=plugins=grpc:. permission.proto

test:
	go vet ./...
	go test  -v -coverpkg ./... -coverprofile coverage.out -covermode count ./...
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