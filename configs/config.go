package configs

import (
	"company-name/constants"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Name            string
		Version         string
		Description     string
		Port            string
		Environment     string
		VerificationUrl string
	}
	DB struct {
		ConnectionString string
		Name             string
	}
	JWT struct {
		Secret     string
		Expiration int64
	}
	Email struct {
		Username string
		Password string
		Host     string
		Port     string
		From     string
	}
	FileStorage struct {
		Directory string
	}
}

var (
	configInstance *Config
	once           sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		var err error
		configInstance, err = initializeConfig()
		if err != nil {
			log.Fatalf(constants.ConfigFaildIntializeErrorMessages, err)
		}
	})

	return configInstance
}

func initializeConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(constants.ConfigNotFoundEnvFileErrorMessage)
	}

	config := &Config{}

	// APP
	config.App.Name = getEnv("APP_NAME", "Hakaya-Alaqzam")
	config.App.Version = getEnv("APP_VERSION", "1.0.0")
	config.App.Description = getEnv("APP_DESCRIPTION", "")
	config.App.Port = getEnv("APP_PORT", "8080")
	config.App.Environment = getEnv("APP_ENVIRONMENT", "development")
	config.App.VerificationUrl = getEnv("EMAIL_VERIFICATION_URL", "https://example.com/verify")

	// DB
	config.DB.ConnectionString = getEnv("DB_CONNECTION_STRING", "mongodb://localhost:27017")

	config.DB.Name = getEnv("DB_NAME", "hakaya-alaqzam")

	// JWT
	config.JWT.Secret = getEnv("JWT_SECRET", "SuperSecretKey")

	config.JWT.Expiration = getEnvAsInt("JWT_EXPIRATION_IN_MILLISECONDS", 3600000)

	// Email
	config.Email.Host = getEnv("EMAIL_HOST", "smtp.example.com")
	config.Email.Port = getEnv("EMAIL_PORT", "587")
	config.Email.Username = getEnv("EMAIL_USERNAME", "no-reply@example.com")
	config.Email.Password = getEnv("EMAIL_PASSWORD", "password")
	config.Email.From = getEnv("EMAIL_FROM", "no-reply@example.com")

	// File Storage
	config.FileStorage.Directory = getEnv("FILE_STORAGE_DIRECTORY", "uploads")

	return config, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return i
		}
		log.Printf(constants.ConfiginvalidValueMessageErrorMessage, key, fallback)
	}
	return fallback
}
