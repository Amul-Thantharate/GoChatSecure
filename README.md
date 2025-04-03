# GoChatSecure 🔒💬

A secure, real-time chat application built with Go, featuring WebSockets, JWT authentication, and TLS encryption.

## 🌟 Features

- **Secure Communication** 🛡️ - TLS encryption for all connections
- **User Authentication** 🔑 - JWT-based authentication system
- **Real-time Messaging** ⚡ - WebSocket-based instant messaging
- **Chat Rooms** 🚪 - Support for multiple chat rooms
- **Private Messaging** 📨 - Direct messaging between users with @username syntax
- **Message History** 📜 - Stores recent messages for each room

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- SSL certificate and key for TLS encryption

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/GoChatSecure.git
   cd GoChatSecure
   ```

2. Generate SSL certificates (for development):
   ```bash
   openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes
   ```

3. Build the application:
   ```bash
   go build -o gochatsecure
   ```

### Running the Server

Start the server with default settings:
```bash
./gochatsecure
```

With custom settings:
```bash
./gochatsecure -addr=0.0.0.0:8443 -cert=path/to/cert.crt -key=path/to/key.key
```

## 🔧 Configuration

The server accepts the following command-line flags:

| Flag | Default | Description |
|------|---------|-------------|
| `-addr` | `localhost:8080` | Server address and port |
| `-cert` | `server.crt` | Path to SSL certificate file |
| `-key` | `server.key` | Path to SSL key file |

## 📝 API Documentation

### Authentication

**Endpoint**: `/auth`  
**Method**: GET  
**Parameters**: `username` - The username to authenticate  
**Response**: JSON containing a JWT token

Example:
```
GET https://localhost:8080/auth?username=john
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### WebSocket Connection

**Endpoint**: `/ws`  
**Parameters**:
- `token` - JWT token obtained from the auth endpoint
- `room` (optional) - Chat room to join (defaults to "default")

Example:
```
wss://localhost:8080/ws?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...&room=general
```

## 💬 Chat Commands

- **Regular message**: Type your message and press enter to send to everyone in the room
- **Private message**: Start your message with `@username ` to send a private message

Example:
```
@john Hey, how are you doing?
```

## 🔐 Security Considerations

- The default implementation uses a hardcoded secret key (`supersecretkey`). In production, you should use a strong, environment-specific secret.
- TLS certificates should be properly managed and renewed in a production environment.
- Consider implementing rate limiting to prevent abuse.

## 🏗️ Architecture

GoChatSecure uses a simple but effective architecture:

- **WebSockets** for bidirectional communication
- **Goroutines** for concurrent handling of connections and messages
- **Channels** for thread-safe message broadcasting
- **JWT** for stateless authentication
- **Mutex** for safe access to shared resources

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Gorilla WebSocket](https://github.com/gorilla/websocket) for the WebSocket implementation
- [golang-jwt](https://github.com/golang-jwt/jwt) for JWT authentication
