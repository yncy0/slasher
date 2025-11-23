package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordBotToken string
	DiscordGuildID  string
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()
	log.Println("LOG: ENV variables loaded")

	config := &Config{
		DiscordBotToken: os.Getenv("DISCORD_BOT_TOKEN"),
		DiscordGuildID:  os.Getenv("DISCORD_GUILD_ID"),
	}

	if config.DiscordBotToken == "" {
		return nil, fmt.Errorf("ERROR: DISCORD_BOT_TOKEN env not found")
	}
	if config.DiscordGuildID == "" {
		return nil, fmt.Errorf("ERROR: DISCORD_GUILD_ID env not found")
	}

	return config, nil
}
