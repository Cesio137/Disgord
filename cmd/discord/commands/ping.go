package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	SlashCommands["ping"] = SlashCommand{
		Command: discordgo.ApplicationCommand{
			Type:        discordgo.ChatApplicationCommand,
			Name:        "ping",
			Description: "Replies with pong 🏓",
		},
		Run: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong 🏓",
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				log.Println("Error trying to respond the ping command:", err)
			}
		},
	}
}
