package commands

import (
	"disgord/discord/base"
	"github.com/bwmarrin/discordgo"
)

func teste() {
	base.Commands = append(base.Commands, base.SlashCommand{
		Command: *discordgo.ApplicationCommand{
			Name: "ping",
			Description: "Replies with pong 🏓",
			Type: discordgo.ChatApplicationCommand,
		},
		Run: func(interaction discordgo.Interaction) {
			
		},
	})
}