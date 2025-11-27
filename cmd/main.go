package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/yncy0/slasher/commands"
	"github.com/yncy0/slasher/config"
)

func main() {
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

	comms, err := commands.Commands()
	if err != nil {
		log.Fatalf("ERROR: Cannot load commands - %v", err)
		return
	}

	for _, cmd := range comms {
		_, err := dg.ApplicationCommandCreate(cfg.DiscordApplicationID, cfg.DiscordGuildID, cmd)
		if err != nil {
			log.Fatal(err)
		}
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commands.InitHandlers(s, i)
	})

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
