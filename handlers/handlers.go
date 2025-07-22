package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"signup/database"
	"signup/model"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📨 Received signup request")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("📋 Raw JSON: %s", string(body))

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("✅ Parsed: username=%s, email=%s, password=%s", user.Username, user.Email, user.Password)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// 🔍 Log the query and values
	log.Printf("🔐 INSERT INTO users (username, email, password) VALUES (%q, %q, %q)", user.Username, user.Email, user.Password)

	stmt := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	_, err = database.DB.Exec(stmt, user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("❌ Exec error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	log.Println("✅ User inserted successfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "User created successfully",
		"username": user.Username,
	})
}
