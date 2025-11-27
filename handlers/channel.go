package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandlePackage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Channel has been created",
		},
	})

	if err != nil {
		log.Panicf("ERROR: Cannot send message - %v", err)
	}
}
