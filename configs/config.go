package configs

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Port     int    `yaml:"port"`
	Driver   string `yaml:"driver"`
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	DB_Port  int    `yaml:"db_port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var defaultAppConfig AppConfig
	defaultAppConfig.Port = 8000
	defaultAppConfig.Driver = getEnv("DRIVER", "mysql")
	defaultAppConfig.Name = getEnv("NAME", "test")
	defaultAppConfig.Address = getEnv("ADDRESS", "localhost")
	defaultAppConfig.DB_Port = 3306
	defaultAppConfig.Username = getEnv("USERNAME", "admin")
	defaultAppConfig.Password = getEnv("PASSWORD", "admin")

	fmt.Println("connected to:", defaultAppConfig)

	return &defaultAppConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "cakcup" {
		return value
	}
	return fallback
}
