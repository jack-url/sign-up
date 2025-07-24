package routes

import (
	"log"
	"net/http"
	"signup/handlers"
)

func SetupRoutes() *http.ServeMux {
	log.Println("🔧 Configuring HTTP routes...")

	mux := http.NewServeMux()

	// Create user endpoint
	log.Println("📝 Registering POST /signup endpoint")
	mux.HandleFunc("/signup", handlers.SignupHandler)

	// View all users endpoint
	log.Println("👥 Registering GET /users endpoint")
	mux.HandleFunc("/users", handlers.ViewUsers)

	// Update or delete specific user endpoint
	log.Println("🔄 Registering PUT/DELETE /users/ endpoint")
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🌐 Received %s request for %s", r.Method, r.URL.Path)
		// Debug log: Confirming entry into the /users/ handler
		log.Printf("DEBUG: Entering /users/ handler. Method: %s, Path: %s", r.Method, r.URL.Path)

		switch r.Method {
		case http.MethodPut:
			log.Printf("➡️  Routing to UpdateUser handler")
			handlers.UpdateUser(w, r)
		case http.MethodDelete:
			log.Printf("➡️  Routing to DeleteUser handler")
			handlers.DeleteUser(w, r)
		default:
			log.Printf("❌ Invalid method %s for %s", r.Method, r.URL.Path)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("✅ All routes configured successfully")
	return mux
}
