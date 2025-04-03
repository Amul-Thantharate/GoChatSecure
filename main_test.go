package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWT(t *testing.T) {
	username := "testuser"
	token, err := generateJWT(username)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	// Parse and verify the token
	parsedToken, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		t.Fatalf("Failed to parse JWT: %v", err)
	}

	if !parsedToken.Valid {
		t.Fatal("Token is not valid")
	}

	claims, ok := parsedToken.Claims.(*JWTClaims)
	if !ok {
		t.Fatal("Failed to parse claims")
	}

	if claims.Username != username {
		t.Fatalf("Expected username %s, got %s", username, claims.Username)
	}
}

func TestAuthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth?username=testuser", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "token") {
		t.Errorf("handler response doesn't contain token: %v", rr.Body.String())
	}
}

func TestAuthHandlerMissingUsername(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestAuthenticate(t *testing.T) {
	username := "testuser"
	token, err := generateJWT(username)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	req, err := http.NewRequest("GET", "/ws?token="+url.QueryEscape(token), nil)
	if err != nil {
		t.Fatal(err)
	}

	authenticatedUser, err := authenticate(req)
	if err != nil {
		t.Fatalf("Authentication failed: %v", err)
	}

	if authenticatedUser != username {
		t.Fatalf("Expected username %s, got %s", username, authenticatedUser)
	}
}

func TestAuthenticateInvalidToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/ws?token=invalidtoken", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = authenticate(req)
	if err == nil {
		t.Fatal("Expected authentication to fail with invalid token")
	}
}
