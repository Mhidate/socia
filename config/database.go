package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// func LoadEnv() {
// 	// Cari path ke root proyek (assume eksekusi dari `cmd/main.go`)
// 	rootPath, err := filepath.Abs("../") // Kembali ke root proyek
// 	if err != nil {
// 		log.Fatalf("Gagal mendapatkan path root proyek: %v", err)
// 	}

// 	envPath := filepath.Join(rootPath, ".env") // Path lengkap ke .env
// 	// Load file .env dari root proyek
// 	err = godotenv.Load(envPath)
// 	if err != nil {
// 		log.Println("Peringatan: Tidak dapat memuat file .env, menggunakan environment variable default")
// 	}
// }

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
	fmt.Println("Berhasil terhubung ke database")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
