package routes

import (
	"net/http"
	"signup/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// POST: signup user
	mux.HandleFunc("/signup", handlers.SignupHandler)

	// GET: view users
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.ViewUsers(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// PUT & DELETE by ID: /users/{id}
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			handlers.UpdateUser(w, r)
		case http.MethodDelete:
			handlers.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
