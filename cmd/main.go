package main

import (
	"log"
	"net/http"

	"auth-service/config"
	"auth-service/internal/auth"
	"auth-service/internal/session"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	redisClient := session.NewRedisClient(cfg)
	authHandler := &auth.AuthHandler{SessionStore: redisClient}
	router := mux.NewRouter()
	router.HandleFunc("/login", authHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", authHandler.LogoutHandler).Methods("POST")
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(auth.AuthMiddleware(redisClient))
	protected.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Access granted to protected route!"))
	}).Methods("GET")

	log.Println("Auth Service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
