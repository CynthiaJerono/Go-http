# Hello Go - Minimal Working Example

A simple HTTP API server built with Go (Golang) that demonstrates the fundamentals of building web services.

## Prerequisites

- **Go 1.21 or later** - Download from https://go.dev/dl/
- **Terminal/Command Prompt** - To run the application
- **Web Browser or curl** - To test the API

## Installation

### 1. Install Go

#### macOS
```bash
# Using Homebrew
brew install go

# Or download from https://go.dev/dl/
```

#### Linux
```bash
# Download and extract
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

#### Windows
1. Download the MSI installer from https://go.dev/dl/
2. Run the installer
3. Open Command Prompt and verify: `go version`

### 2. Verify Installation

```bash
go version
# Expected output: go version go1.21.0 darwin/amd64
```

## How to Run

### Step 1: Navigate to the Project Directory

```bash
cd hello-go
```

### Step 2: Initialize Go Module (if not already done)

```bash
go mod init hello-go
```

### Step 3: Run the Application

```bash
go run main.go
```

You should see output like:
```
Server starting on http://localhost:8080
Visit /hello to see the Hello World endpoint
```

### Step 4: Test the API

Open your web browser and visit:
- **Root endpoint**: http://localhost:8080/
- **Hello endpoint**: http://localhost:8080/hello

Or use curl in a new terminal:

```bash
# Test root endpoint
curl http://localhost:8080/

# Test hello endpoint (JSON response)
curl http://localhost:8080/hello
```

## Expected Output

### Root Endpoint (http://localhost:8080/)
```
Go API Server is running!
Visit /hello to see the Hello World endpoint
```

### Hello Endpoint (http://localhost:8080/hello)
```json
{
  "message": "Hello, World! Welcome to Go!",
  "timestamp": "2024-01-15T10:30:45.123456Z",
  "go_version": "1.21+"
}
```

## Project Structure

```
hello-go/
├── main.go    # Main application code
├── go.mod     # Go module file
└── README.md  # This file
```

## Key Concepts Demonstrated

1. **Package Declaration** - `package main` identifies an executable
2. **Imports** - Using standard library packages (`net/http`, `encoding/json`, `fmt`, `time`, `log`)
3. **HTTP Handlers** - Creating web server endpoints
4. **JSON Encoding** - Converting Go structs to JSON responses
5. **Struct Types** - Defining data structures with JSON tags
6. **Running the Server** - Using `http.ListenAndServe`

## Stopping the Server

Press `Ctrl+C` in the terminal where the server is running.

## Troubleshooting

### "port already in use"
Change the port in main.go:
```go
port := ":8081"  // Use 8081 instead of 8080
```

### "command not found: go"
Ensure Go is installed and in your PATH. Add to your shell profile:
```bash
export PATH=$PATH:/usr/local/go/bin
```

### VS Code errors
Install the Go extension:
1. Open VS Code
2. Go to Extensions (Ctrl+Shift+X)
3. Search for "Go" and install

## Learn More

- **Official Go Tutorial**: https://go.dev/tour/
- **Go by Example**: https://gobyexample.com/
- **Go Documentation**: https://go.dev/doc/

## License

MIT License - Feel free to use and modify!
