package config

import (
	"encoding/json"
	"fmt"
	"log"
	"movieTicket/models"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config struct to hold configuration values
type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
		TimeZone string `json:"timezone"`
	} `json:"database"`
}

// Global variables
var (
	DB          *gorm.DB
	AppConfig   *Config
	DBAvailable = true // Flag to check DB status
	mu          sync.Mutex
)

// LoadConfig reads the config.json file
func LoadConfig() {
	file, err := os.Open("config/config.json") // Ensure the file is in the config folder
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	AppConfig = &Config{}
	err = json.NewDecoder(file).Decode(AppConfig)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
}

// InitDB initializes the PostgreSQL database connection
func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		AppConfig.Database.Host,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.DBName,
		AppConfig.Database.Port,
		AppConfig.Database.SSLMode,
		AppConfig.Database.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("⚠️  Database connection failed! Switching to in-memory storage...")
		DBAvailable = false // Set DB status as unavailable
		return
	}

	DB = db
	DBAvailable = true // Set DB status as available

	// Run auto-migrations for all models
	err = DB.AutoMigrate(&models.Ticket{}, &models.Seat{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL successfully!")
}

// InitConfig initializes configuration and database and returns the loaded config
func InitConfig() *Config {
	LoadConfig()
	InitDB()
	log.Println("✅ Configuration and Database initialized successfully!")
	return AppConfig
}
