package commands

import (
	"disgord/cmd/discord/base"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	base.Commands["ping"] = base.SlashCommand{
		Command: discordgo.ApplicationCommand{
			Type:        discordgo.ChatApplicationCommand,
			Name:        "ping",
			Description: "Replies with pong 🏓",
		},
		Run: teste,
	}
}

func teste(s *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Println("Comando ping")
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
			if err != nil {
				fmt.Println("Erro ao responder ao comando:", err)
			}
}
