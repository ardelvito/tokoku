package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"backendmono/config" // Sesuaikan dengan path ke file config.go
)

func InitDB() (*gorm.DB, error) {
	// Load konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("could not load config: %v", err)
	}

	// URL koneksi PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Open koneksi ke database menggunakan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	log.Println("Successfully connected to database with GORM")
	return db, nil
}
