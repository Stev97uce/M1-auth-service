package auth

import (
	"encoding/json"
	"net/http"
	"bytes"
	"time"
	"io"

	"github.com/google/uuid"
	"auth-service/internal/session"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	SessionStore *session.RedisClient
}

func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// agreagr el link de la ec2 del micro2
	payload, _ := json.Marshal(creds)
	resp, err := http.Post("http://user-profile-service:8000/login", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		http.Error(w, "User service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, string(body), resp.StatusCode)
		return
	}

	token := uuid.New().String()
	a.SessionStore.SetSession(token, creds.Username)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: time.Now().Add(a.SessionStore.TTL),
	})
	w.Write([]byte("Login successful"))
}
