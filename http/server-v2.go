package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	userservice "github.com/dev-khalid/golang-intro/http/user-service"
)


func ServerV2() { 
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("POST /users", userservice.CreateUser)
	mux.HandleFunc("/users/{id}", userservice.GetUser)
	mux.HandleFunc("DELETE /users/{id}", userservice.DeleteUser)
	mux.HandleFunc("GET /users/all", userservice.GetAllUser)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}


func handleRoot(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprint(w, "Hello, World!")
	
}

func handleHealth(w http.ResponseWriter, r *http.Request) { 
	response := map[string]string{
		"status": "healthy",
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

