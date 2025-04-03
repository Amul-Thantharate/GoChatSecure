package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

// Client represents a connected user
type Client struct {
	conn     *websocket.Conn
	username string
	room     string
}

var (
	addr       = flag.String("addr", "localhost:8080", "Server address")
	certFile   = flag.String("cert", "server.crt", "SSL Certificate")
	keyFile    = flag.String("key", "server.key", "SSL Key")
	secretKey  = "supersecretkey"
	clients    = make(map[*websocket.Conn]*Client)
	rooms      = make(map[string]map[*websocket.Conn]*Client)
	messageBus = make(chan []byte)
	history    = make(map[string][]string) // Stores last 20 messages per room
	mutex      sync.Mutex
)

const maxHistory = 20

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// JWT Claims
type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Authenticate user via JWT token
func authenticate(r *http.Request) (string, error) {
	tokenStr := r.URL.Query().Get("token")
	if tokenStr == "" {
		return "", http.ErrNoCookie
	}
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return "", jwt.ErrInvalidKey
	}
	return claims.Username, nil
}

// Handle WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	username, err := authenticate(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	room := r.URL.Query().Get("room")
	if room == "" {
		room = "default"
	}

	// Register client in room
	mutex.Lock()
	if rooms[room] == nil {
		rooms[room] = make(map[*websocket.Conn]*Client)
	}
	client := &Client{conn, username, room}
	rooms[room][conn] = client
	clients[conn] = client
	history[room] = append(history[room], username+" joined the room")
	if len(history[room]) > maxHistory {
		history[room] = history[room][1:]
	}
	mutex.Unlock()

	log.Printf("User %s joined room %s", username, room)

	// Handle incoming messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		messageText := username + ": " + string(msg)

		// Private Message (@username message)
		if strings.HasPrefix(string(msg), "@") {
			parts := strings.SplitN(string(msg[1:]), " ", 2)
			if len(parts) < 2 {
				continue
			}
			targetUser, message := parts[0], parts[1]
			mutex.Lock()
			for _, client := range rooms[room] {
				if client.username == targetUser {
					client.conn.WriteMessage(websocket.TextMessage, []byte("(Private) "+username+": "+message))
				}
			}
			mutex.Unlock()
			continue
		}

		// Broadcast to all in the room
		messageBus <- []byte(messageText)
	}

	// Client disconnects
	mutex.Lock()
	delete(rooms[room], conn)
	delete(clients, conn)
	history[room] = append(history[room], username+" left the room")
	if len(history[room]) > maxHistory {
		history[room] = history[room][1:]
	}
	mutex.Unlock()
	log.Printf("User %s left room %s", username, room)
}

// Handle broadcasting messages
func handleMessages() {
	for msg := range messageBus {
		mutex.Lock()
		for _, client := range clients {
			roomClients := rooms[client.room]
			for conn := range roomClients {
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					conn.Close()
					delete(roomClients, conn)
				}
			}
		}
		mutex.Unlock()
	}
}

// Generate JWT token
func generateJWT(username string) (string, error) {
	if secretKey == "" {
		return "", errors.New("secret key is missing")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// Auth handler for generating JWT tokens
func authHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}
	token, err := generateJWT(username)
	if err != nil {
		http.Error(w, "Failed to generate token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func main() {
	flag.Parse()
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/auth", authHandler)
	go handleMessages()

	// Load SSL Certificates
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("Failed to load SSL certificates: %v", err)
	}

	server := &http.Server{
		Addr:      *addr,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
	}

	log.Printf("Server started at https://%s", *addr)
	log.Fatal(server.ListenAndServeTLS("", ""))
}
