package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Commands() ([]*discordgo.ApplicationCommand, error) {
	var commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Checking the bot if online",
		},
	}

	return commands, nil
}

// var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
// 	"ping": handlers.OnHandlePing,
// }
