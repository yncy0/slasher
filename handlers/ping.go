package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey I'am slasher bot!",
		},
	})

	if err != nil {
		log.Panicf("ERROR: Cannot send message - %v", err)
	}
}
