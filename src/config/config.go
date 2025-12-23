package config

import (
	"log"

	"github.com/spf13/viper"
)

// App configuration
var (
	IsDev   bool
	AppEnv  string
	AppHost string
	AppPort int
	AppURL  string

	// Database configuration
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int
)

func init() {
	loadConfig()

	// app config
	AppEnv = viper.GetString("APP_ENV")
	IsDev = AppEnv == "dev"
	AppHost = viper.GetString("APP_HOST")
	AppPort = viper.GetInt("APP_PORT")
	AppURL = viper.GetString("APP_URL")

	// database config
	DBHost = viper.GetString("DB_HOST")
	DBUser = viper.GetString("DB_USER")
	DBPassword = viper.GetString("DB_PASSWORD")
	DBName = viper.GetString("DB_NAME")
	DBPort = viper.GetInt("DB_PORT")

	// basic validation
	if AppPort == 0 {
		log.Fatal("❌ APP_PORT missing")
	}
	if DBHost == "" || DBUser == "" || DBName == "" {
		log.Fatal("❌ Database configuration missing")
	}

	log.Println("✅ Configuration initialized")
}

func loadConfig() {
	configPaths := []string{
		"./",     // normal run
		"../../", // test run
	}

	for _, path := range configPaths {
		viper.SetConfigFile(path + ".env")

		if err := viper.ReadInConfig(); err == nil {
			log.Printf("✅ Config loaded from %s.env\n", path)
			return
		}
	}

	log.Fatal("❌ Failed to load any .env config file")
}
