# User_Management_system
# Documentation for Setting Up a gRPC Server in Go
# Use it on Cmd Or Bash.

This documentation explains each step and command required to set up a gRPC server in Go. Follow these steps to configure your environment and create a gRPC project with Go.

## 1. Installing Required Tools

### Install Protocol Buffer Compiler Plugins for Go

1. **Install `protoc-gen-go`:**
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
- This command installs the `protoc-gen-go` plugin, which generates Go code for protocol buffers.

2. **Install `protoc-gen-go-grpc`:**
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   
   - This command installs the `protoc-gen-go-grpc` plugin, which generates Go code for gRPC services.

### Verify `protoc` Version
   
   protoc --version
   
   - Ensure that the protocol buffer compiler (`protoc`) is installed and displays the version.

## 2. Initializing a Go Module

1. **Create a Go Module:**
   ```bash
   go mod init user_project/server
   ```
   - Initializes a new Go module named `user_project/server`.

2. **Download Dependencies:**
   
   go mod tidy
   
   - Cleans up and synchronizes the module's dependencies.

## 3. Generating Go Code from `.proto` Files

1. **Compile Protocol Buffer File:**
   protoc --go_out=. --go-grpc_out=. proto/user.proto
   
   - `--go_out=.`: Generates Go code for protocol buffer message types.
   - `--go-grpc_out=.`: Generates Go code for gRPC services.
   - `proto/user.proto`: The input `.proto` file containing the protocol buffer definitions.

## 4. Installing Additional Dependencies

1. **Install MySQL Driver for Go:**
   go get -u github.com/go-sql-driver/mysql
   
   - Downloads and installs the MySQL driver for database interaction.

2. **Install gRPC for Go:**
   go get google.golang.org/grpc
   
   - Downloads and installs the gRPC package for Go.
