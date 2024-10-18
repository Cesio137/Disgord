package commands

import "github.com/bwmarrin/discordgo"

type SlashCommand struct {
	Command discordgo.ApplicationCommand
	Run     func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var SlashCommands map[string]SlashCommand = make(map[string]SlashCommand)