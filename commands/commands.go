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
		{
			Name:        "channel",
			Description: "Command for Interact with Discord Server Channles",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "add",
					Description: "Adding Channels",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "name",
							Description: "Naming Channels",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionChannel,
							Name:        "type",
							Description: "Type of Discord Channel",
						},
					},
				},
			},
		},
	}

	return commands, nil
}
