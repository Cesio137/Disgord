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
		Run: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			fmt.Println("Command ping")
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				fmt.Println("Error trying to respond the ping command:", err)
			}
		},
	}
}
