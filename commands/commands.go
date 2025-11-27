package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/yncy0/slasher/handlers"
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
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "type",
							Description: "Type of Discord Channel",
							Required:    true,
							Choices: []*discordgo.ApplicationCommandOptionChoice{
								{Name: "Text", Value: "text"},
								{Name: "Voice", Value: "voice"},
							},
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "delete",
					Description: "Deleting Channels",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionChannel,
							Name:        "channel",
							Description: "Discord Server Channel",
							Required:    true,
						},
					},
				},
			},
		},
	}

	return commands, nil
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping":    handlers.HandlePing,
	"channel": handlers.HandleChannel,
}

func InitHandlers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handleInteraction(s, i)
	})
}

func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	}
}

func GetOptions(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, opt := range options {
		if opt.Name == name {
			return opt
		}
	}

	return nil
}

func GetSubCommandName(i *discordgo.InteractionCreate) string {
	options := i.ApplicationCommandData().Options
	if len(options) > 0 && options[0].Type == discordgo.ApplicationCommandOptionSubCommand {
		return options[0].Name
	}

	return ""
}

func GetSubCommandOptions(i *discordgo.InteractionCreate) []*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options
	if len(options) > 0 && (options[0].Type == discordgo.ApplicationCommandOptionSubCommand || options[0].Type == discordgo.ApplicationCommandOptionSubCommandGroup) {
		return options[0].Options
	}

	return nil
}
