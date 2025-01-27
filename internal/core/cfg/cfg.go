package cfg

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type Config struct {
	AppPort                      string
	GrpcPort                     string
	LotofHubMSvcUsersGrpcAddress string
}

var (
	once     sync.Once
	instance *Config
)

func Inst() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env file found, loading from OS environment variables.")
		}

		instance = &Config{
			AppPort:                      getEnv("APP_PORT", "8081"),
			GrpcPort:                     getEnv("GRPC_PORT", "50060"),
			LotofHubMSvcUsersGrpcAddress: getEnv("LOTOF_HUB_USERS_SVC_GRPC_ADDRESS", "localhost:50061"),
		}
	})
	return instance
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
