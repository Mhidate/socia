package config

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Cari path ke root proyek (assume eksekusi dari `cmd/main.go`)
	rootPath, err := filepath.Abs("../") // Kembali ke root proyek
	if err != nil {
		log.Fatalf("Gagal mendapatkan path root proyek: %v", err)
	}

	envPath := filepath.Join(rootPath, ".env") // Path lengkap ke .env
	// Load file .env dari root proyek
	err = godotenv.Load(envPath)
	if err != nil {
		log.Println("Peringatan: Tidak dapat memuat file .env, menggunakan environment variable default")
	}
}
