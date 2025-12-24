package config

import (
	"log"

	"github.com/spf13/viper"
)

// App configuration
var (
	// app
	IsDev   bool
	AppEnv  string
	AppHost string
	AppPort int
	AppURL  string

	// database
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int

	// jwt ✅ (THIS WAS MISSING)
	JWTSecret           string
	JWTAccessExp        int // minutes
	JWTRefreshExp       int // days
	JWTResetPasswordExp int // minutes
	JWTVerifyEmailExp   int // minutes
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

	// jwt config ✅
	JWTSecret = viper.GetString("JWT_SECRET")
	JWTAccessExp = viper.GetInt("JWT_ACCESS_EXP_MINUTES")
	JWTRefreshExp = viper.GetInt("JWT_REFRESH_EXP_DAYS")
	JWTResetPasswordExp = viper.GetInt("JWT_RESET_PASSWORD_EXP_MINUTES")
	JWTVerifyEmailExp = viper.GetInt("JWT_VERIFY_EMAIL_EXP_MINUTES")

	// basic validation
	if AppPort == 0 {
		log.Fatal("❌ APP_PORT missing")
	}
	if DBHost == "" || DBUser == "" || DBName == "" {
		log.Fatal("❌ Database configuration missing")
	}
	if JWTSecret == "" {
		log.Fatal("❌ JWT_SECRET missing")
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
