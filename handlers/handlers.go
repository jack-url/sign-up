package handlers

import (
	"encoding/json"
	"net/http"

	"signup/database"
	"signup/model"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":  "User created successfully",
		"username": user.Username,
	})
}

func ViewUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := database.DB.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password); err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Row iteration error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
