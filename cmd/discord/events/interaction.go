package events

import (
	"disgord/cmd/discord/commands"

	"github.com/bwmarrin/discordgo"
)

func init() {
	Events = append(Events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}
		commands.SlashCommands[i.ApplicationCommandData().Name].Run(s, i)
	})
}