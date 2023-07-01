# RPKM66 Backend

## Stacks
- golang
- gRPC

## Getting Start
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- golang 1.20 or [later](https://go.dev)
- docker
- makefile

### Installing
1. Clone the project from [RPKM66 Backend](https://github.com/isd-sgcu/rpkm66-backend)
2. Import project
3. Copy `config.example.toml` in `config` and paste it in the same location then remove `.example` from its name.
4. Download dependencies by `go mod download`

### Testing
1. Run `go test  -v -coverpkg ./... -coverprofile coverage.out -covermode count ./...` or `make test`

### Running
1. Run `docker-compose up -d` or `make compose-up`
2. Run `go run ./.` or `make server`

### Compile proto file
1. Run `make proto`

### Create Seed File
1. Create seeder file in `src/database/seeds`
2. Name seed func in pattern `<Name>Seed<Timestamp>`
   - ex `UserSeed1652002196085` `RoleSeed1651703066048`

### Seeding
- Run `go run ./. seed` or `make seed` (Seed all files)
- Run `go run ./. seed <name>` (Seed specific file)

