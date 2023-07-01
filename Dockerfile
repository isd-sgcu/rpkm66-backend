# Base Image
FROM golang:1.20.5-bullseye as base

# Working directory
WORKDIR /app

# Setup credential
ENV GOPRIVATE=github.com/isd-sgcu/*

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN --mount=type=secret,id=netrcConf,target=/root/.netrc,required=true go mod download

# Copy the source code
COPY . .

# Build the application
RUN --mount=type=secret,id=netrcConf,target=/root/.netrc,required=true CGO_ENABLED=0 go build -o server ./cmd/main.go
# Create master image
FROM alpine AS master

# Working directory
WORKDIR /app

# Copy execute file
COPY --from=base /app/server ./

# Set ENV to production
ENV GO_ENV production

# Expose port 3000
EXPOSE 3000

# Run the application
ENTRYPOINT ["./server"]
