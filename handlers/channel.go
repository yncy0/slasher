package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HandleChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	subOptions := options[0]

	switch subOptions.Name {
	case "add":
		addChannel(s, i)
	case "delete":
		deleteChannel(s, i)
	}

}

func addChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options[0].Options

	var channelName string
	var channelType string

	for _, opt := range options {
		switch opt.Name {
		case "name":
			channelName = opt.StringValue()
		case "type":
			channelType = opt.StringValue()
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("New Channel Added: %s", channelName),
		Description: fmt.Sprintf("The %s Channel %s has been created", channelType, channelName),
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel Created",
			Embeds:  []*discordgo.MessageEmbed{embed},
		},
	})
}

func deleteChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options[0].Options

	var channelName *discordgo.Channel
	for _, opt := range options {
		switch opt.Name {
		case "channel":
			channelName = opt.ChannelValue(s)
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Channel Deleted",
		Description: fmt.Sprintf("Channel %s has been deleted", channelName.Name),
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel Delete",
			Embeds:  []*discordgo.MessageEmbed{embed},
		},
	})
}
