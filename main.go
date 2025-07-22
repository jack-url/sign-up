package main

import (
	"log"
	"net/http"
	"signup/database"
	"signup/routes"
)

func main() {
	log.Println("📦 Initializing DB...")
	database.InitDB()

	router := routes.SetupRoutes()

	log.Println("🚀 Server running at http://localhost:9000")
	err := http.ListenAndServe(":9000", router)
	if err != nil {
		log.Fatalf("❌ Server error: %v", err)
	}
}
