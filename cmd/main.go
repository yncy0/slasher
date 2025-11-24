package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/yncy0/slasher/commands"
	"github.com/yncy0/slasher/config"
	"github.com/yncy0/slasher/handlers"
)

func main() {
	log.Print("LOG: Starting Discord Bot...")

	log.Print("LOG: Loading Config...")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("ERROR: Cannot load Config - %v", err)
		return
	}

	log.Print("LOG: Loading Discord Bot...")
	dg, err := discordgo.New("Bot " + cfg.DiscordBotToken)
	if err != nil {
		log.Fatalf("ERROR: Cannot load Discord Bot - %v", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handlers.HandlePing(s, i)
	})

	log.Printf("LOG: Loading Commands...")
	comms, err := commands.Commands()
	if err != nil {
		log.Fatalf("ERROR: Cannot load commands - %v", err)
		return
	}

	_, err = dg.ApplicationCommandBulkOverwrite(cfg.DiscordApplicationID, cfg.DiscordGuildID, comms)
	if err != nil {
		log.Fatalf("ERROR: Cannot read commands - %v", err)
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
