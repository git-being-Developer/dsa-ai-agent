package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig OpenAIConfig

type Config interface {
	loadConfig()
}

type OpenAIConfig struct {
	OPEN_AI_API string
}

func NewAppConfig() *OpenAIConfig {
	return &OpenAIConfig{
		OPEN_AI_API: "",
	}
}

func (c *OpenAIConfig) LoadConfig() {

	err := godotenv.Load("conf.env")
	if err != nil {
		log.Printf("No .env file found, continuing with env vars")
	}
	c.OPEN_AI_API = os.Getenv("OPEN_AI_KEY")
}
