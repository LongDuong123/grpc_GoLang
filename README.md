# Book gRPC Service

This is a simple gRPC service for managing books written in Go.

## Requirements

- Go 1.22.1 
- Protocol Buffers v3 (protoc)
- Go gRPC libraries

1. Clone the repository: git clone https://github.com/LongDuong123/grpc_GoLang.git

2. Install dependencies : go mod tidy

The server will start listening for requests on port `8080`
The client will connect to the server running on localhost:8080 and send a request to create a book.