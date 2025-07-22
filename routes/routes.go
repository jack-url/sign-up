package routes

import (
	"net/http"
	"signup/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", handlers.SignupHandler)
	return mux
}
