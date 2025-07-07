package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"auth-service/internal/event"
	"auth-service/internal/session"

	"github.com/google/uuid"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	SessionStore session.SessionStore
}

func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Agregar el link de la ec2 del micro
	payload, _ := json.Marshal(creds)
	resp, err := http.Post(os.Getenv("USER_PROFILE_SERVICE_URL")+"/login", "application/json", bytes.NewBuffer(payload))
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

	event.SendEventToKafka(creds.Username + " performed login")

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: time.Now().Add(a.SessionStore.GetTTL()),
	})
	w.Write([]byte("Login successful"))
}

func (a *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		http.Error(w, "Unauthorized - No token", http.StatusUnauthorized)
		return
	}

	err = a.SessionStore.DeleteSession(cookie.Value)
	if err != nil {
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}

	event.SendEventToKafka("User " + cookie.Value + " performed logout")

	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
	})

	w.Write([]byte("Logged out successfully"))
}

func (a *AuthHandler) RoleValidationHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		http.Error(w, "Unauthorized - No token", http.StatusUnauthorized)
		return
	}
	//Agregar el link de la ec2 del micro
	resp, err := http.Get(os.Getenv("USER_PROFILE_SERVICE_URL") + "/users/" + cookie.Value)
	if err != nil {
		http.Error(w, "User service unavailable", http.StatusServiceUnavailable)
		return
	}

	if resp.StatusCode == http.StatusOK {
		w.Write([]byte("Role validated"))
	} else {
		http.Error(w, "Invalid role", http.StatusUnauthorized)
	}
}
