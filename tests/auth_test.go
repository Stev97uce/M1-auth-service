package tests

import (
	"auth-service/internal/auth"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoginHandler(t *testing.T) {
	// Crear un mock del SessionStore
	mockSessionStore := &MockRedisClient{
		TTL: 30 * time.Second,
	}

	handler := &auth.AuthHandler{
		SessionStore: mockSessionStore,
	}

	// Crear credenciales de prueba
	creds := auth.Credentials{
		Username: "testuser",
		Password: "testpass",
	}

	// Convertir credenciales a JSON
	jsonData, _ := json.Marshal(creds)

	// Crear request de prueba
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Crear response recorder
	w := httptest.NewRecorder()

	// Ejecutar handler
	handler.LoginHandler(w, req)

	// Verificar status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestLogoutHandler(t *testing.T) {
	// Crear un mock del SessionStore
	mockSessionStore := &MockRedisClient{
		TTL: 30 * time.Second,
	}

	handler := &auth.AuthHandler{
		SessionStore: mockSessionStore,
	}

	// Crear request con cookie
	req := httptest.NewRequest("POST", "/logout", nil)
	req.AddCookie(&http.Cookie{
		Name:  "session_token",
		Value: "test-token",
	})

	// Crear response recorder
	w := httptest.NewRecorder()

	// Ejecutar handler
	handler.LogoutHandler(w, req)

	// Verificar status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestLogoutHandlerNoToken(t *testing.T) {
	// Crear un mock del SessionStore
	mockSessionStore := &MockRedisClient{
		TTL: 30 * time.Second,
	}

	handler := &auth.AuthHandler{
		SessionStore: mockSessionStore,
	}

	// Crear request sin cookie
	req := httptest.NewRequest("POST", "/logout", nil)

	// Crear response recorder
	w := httptest.NewRecorder()

	// Ejecutar handler
	handler.LogoutHandler(w, req)

	// Verificar status code
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

// MockRedisClient es un mock simple para las pruebas
type MockRedisClient struct {
	TTL time.Duration
}

func (m *MockRedisClient) SetSession(token, username string) error {
	return nil
}

func (m *MockRedisClient) DeleteSession(token string) error {
	return nil
}

func (m *MockRedisClient) ValidateSession(token string) (string, error) {
	return "testuser", nil
}

func (m *MockRedisClient) GetTTL() time.Duration {
	return m.TTL
}
