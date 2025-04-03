# GoChatSecure 🔒💬

[![Go Reference](https://pkg.go.dev/badge/github.com/gorilla/websocket.svg)](https://pkg.go.dev/github.com/gorilla/websocket)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![WebSocket](https://img.shields.io/badge/WebSocket-Enabled-brightgreen.svg)](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket)

## 🌐 Overview

GoChatSecure is a robust and scalable chat application that leverages the power of Go's concurrency model and the efficiency of WebSockets to provide a seamless real-time communication experience. It's designed with security in mind, utilizing TLS encryption and JWT authentication to ensure that all interactions are private and secure. This project is perfect for developers looking to understand and implement secure, real-time communication systems.

[Quick Links](#quick-links) | [Features](#features) | [Installation](#installation) | [Usage](#usage) | [API](#api-documentation)

## 📋 Quick Links
- [Demo](#) (Coming soon!)
- [Documentation](#api-documentation)
- [Contributing Guidelines](CONTRIBUTING.md)
- [Security Policy](SECURITY.md)

## 🌟 Features

- **Secure Communication** 🛡️
  - TLS encryption for all connections
  - JWT-based authentication
  - Secure WebSocket protocol (wss://)

- **Real-time Features** ⚡
  - Instant messaging using WebSockets
  - Multiple chat room support
  - Private messaging with @username syntax
  - Message history retention

- **Developer Friendly** 🔧
  - Comprehensive test suite
  - Makefile automation
  - Cross-platform support
  - Docker support (coming soon)

## 🚀 Installation

### Prerequisites

- Go 1.16+ 
- OpenSSL (for certificate generation)
- Make (optional, for build automation)

### Quick Start

```bash
# Clone the repository
git clone https://github.com/Amul-Thantharate/GoChatSecure.git

# Navigate to project directory
cd GoChatSecure

# Install dependencies
go mod tidy

# Generate certificates
make generate-certs

# Build and run
make run
```

## 💻 Usage

### Command Line Flags

```bash
./gochat \
  -addr=localhost:8080 \    # Server address (default)
  -cert=server.crt \        # TLS certificate path
  -key=server.key          # TLS key path
```

### Client Authentication

```bash
# Get JWT token
curl -k https://localhost:8080/auth?username=yourname

# Connect to WebSocket (using token)
wscat -c "wss://localhost:8080/ws?token=YOUR_JWT_TOKEN&room=general"
```

### Chat Commands

```
# Regular message
Hello everyone!

# Private message
@john Hey, how are you?

# Join different room
/join room_name (Coming soon)
```

## 🔧 Development

### Build Commands

```bash
# Run tests
make test

# Development build with race detection
make race

# Cross-platform builds
make build-all      # Build for all platforms
make build-linux    # Linux only
make build-windows  # Windows only
make build-mac      # macOS only
```

### Project Structure

```
GoChatSecure/
├── main.go           # Server entry point
├── main_test.go      # Test suite
├── go.mod           # Go modules file
├── go.sum           # Module checksums
├── Makefile         # Build automation
├── server.crt       # TLS certificate
├── server.key       # TLS private key
└── README.md        # This file
```

## 📚 API Documentation

### REST Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/auth`  | GET    | Authenticate and get JWT |
| `/ws`    | GET    | WebSocket connection |

### WebSocket Events

| Event | Description |
|-------|-------------|
| `message` | Regular chat message |
| `private` | Private message |
| `join`    | User joined notification |
| `leave`   | User left notification |

## 🔐 Security

- TLS 1.3+ required for all connections
- JWT tokens expire after 24 hours
- Message sanitization
- Rate limiting (coming soon)

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gorilla WebSocket](https://github.com/gorilla/websocket) - WebSocket implementation
- [golang-jwt](https://github.com/golang-jwt/jwt) - JWT authentication
- [OpenSSL](https://www.openssl.org/) - TLS certificate generation
