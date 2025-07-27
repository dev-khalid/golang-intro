package userservice

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type User struct { 
	ID    string    `json:"id,omitempty"` // Make this optional in request body
	Name  string `json:"name"`
	Email string `json:"email"`
}

var userCache = make(map[string]User)
var cacheMutex sync.RWMutex


func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user = User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" {
		errorMessage, err := json.MarshalIndent(map[string]string{"error": "name and email are required"}, "", "  ")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, string(errorMessage), http.StatusBadRequest)
		return
	}


	cacheMutex.Lock()
	user.ID = generateRandomID()
	userCache[user.ID] = user
	cacheMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	createdUser, err := json.MarshalIndent(user, "", ""); 
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// err = json.NewEncoder(w).Encode(user); if err != nil { 
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Write(createdUser)
}


func GetUser(w http.ResponseWriter, r *http.Request) { 
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, ok := userCache[id]

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) { 
	id := r.PathValue("id")

	if id == "" { 
		http.Error(w, "Invalid User Id", http.StatusBadRequest)
		return
	}

	_, ok := userCache[id]; if !ok { 
		http.Error(w, "No user found", http.StatusNotFound)
		return
	}

	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) { 
	role := r.URL.Query().Get("role")

	if role != "Admin" {
		errorMessage, err := json.MarshalIndent(map[string]string {
			"message": "You are not allowed to get all users.",
		}, "", " "); if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, string(errorMessage), http.StatusUnauthorized)
		return
	}

	// We should return this in an array 
	var users []User

	for _,v := range userCache {
		users = append(users, v)
	}
	// userData, err := json.MarshalIndent(map[string][]User{"users": users}, "", " "); if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(userData)
	err := json.NewEncoder(w).Encode(users); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generateRandomID() string {
	return strconv.Itoa(rand.Intn(999999-100000) + 100000)
}