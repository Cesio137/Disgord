package base

import (
	"disgord/cmd/discord/commands"
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(s *discordgo.Session) {
	log.Println("Refreshing commands!")
    appcmd, err := s.ApplicationCommands(s.State.User.ID, "")
    if err == nil {
        for _, cmd := range appcmd {
			s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		}
    }

	for _, slash_cmd := range commands.SlashCommands {
		cmd := slash_cmd.Command
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", &cmd)

		if err == nil {
			log.Printf("{/} %s command loaded!\n", cmd.Name)
		} else {
			log.Printf("❌ cannot load %s command!\n %s\n", cmd.Name, err.Error())
		}
	}
}
