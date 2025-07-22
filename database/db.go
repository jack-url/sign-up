package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := "root:Born2001$@tcp(127.0.0.1:3306)/signup?parseTime=true"

	var err error
	DB, err = sql.Open("mysql", dsn) // ✅ assign to global DB

	if err != nil {
		log.Fatalf("❌ Error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	fmt.Println("✅ Connected to MySQL database")
}
