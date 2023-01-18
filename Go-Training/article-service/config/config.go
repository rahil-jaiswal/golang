package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// database
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// context
type Context struct {
	Timeout int64 `json:"timeout"`
}

// server
type Server struct {
	Address string `json:"address"`
}

// config
type Configuration struct {
	Debug    bool     `json:"debug"`
	Server   Server   `json:"server"`
	Context  Context  `json:"context"`
	Database Database `json:"database"`
}

func TranslateConfig(configFile string) *Configuration {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Println("Configfile Not Found", configFile)
		return nil
	}
	var config Configuration
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Unable to unmarshal", err)
		return nil
	}

	return &config
}
