package main

import (
	"fmt"
	"log"
	"socia/config"
)

func main() {
	err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	defer config.CloseDB()
	fmt.Println("Aplikasi sedang berjalan.....")
}
