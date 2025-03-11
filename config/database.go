package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	// LoadEnv()
	connStr := os.Getenv("DATABASE_URL") // Ambil dari environment variable
	if connStr == "" {
		return fmt.Errorf("DATABASE_URL tidak ditemukan dalam environment variable")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi database: %v", err)
	}

	// Coba koneksi ke database
	if err := db.Ping(); err != nil {
		return fmt.Errorf("gagal koneksi ke database: %v", err)
	}

	DB = db
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
