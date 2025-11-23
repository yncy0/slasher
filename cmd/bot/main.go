package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/yncy0/slasher/internal/config"
)

func main() {
	log.Print("LOG: Starting Discord Bot...")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("ERROR: Cannot load Config - %v", err)
		return
	}

	dg, err := discordgo.New("Bot " + cfg.DiscordBotToken)
	if err != nil {
		log.Fatalf("ERROR: Cannot load Discord Bot - %v", err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Fatalf("ERROR: Cannot open Discord Bot - %v", err)
		return
	}

	log.Printf("RUN: Discord bot is Online! - %s", dg.State.User)

	defer dg.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
