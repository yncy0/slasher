package handlers

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandleChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	var msg string

	if opt, ok := optionMap["add"]; ok {
		msg = fmt.Sprintf("I successfully add %v channel", opt)
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})

	if err != nil {
		log.Panicf("ERROR: Cannot send message - %v", err)
	}
}
