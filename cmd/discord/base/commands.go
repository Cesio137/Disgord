package base

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type SlashCommand struct {
	Command discordgo.ApplicationCommand
	Run func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var Commands map[string]SlashCommand = make(map[string]SlashCommand)

func RegisterCommands(s *discordgo.Session) {
	fmt.Println("Refreshing commands!")
	appID := s.State.User.ID
    commands, err := s.ApplicationCommands(appID, "1253804989487386624")
    if err == nil {
        for _, cmd := range commands {
			s.ApplicationCommandDelete(appID, "", cmd.ID)
		}
    }

	for _, slash_cmd := range Commands {
		cmd := slash_cmd.Command
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "1253804989487386624", &cmd)

		if err == nil {
			fmt.Printf("{/} %s command loaded!\n", cmd.Name)
		} else {
			fmt.Printf("❌  cannot load %s command!\n", cmd.Name)
			fmt.Println(err.Error())
		}
	}
}

func HandleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return 
	}
	Commands[i.ApplicationCommandData().Name].Run(s, i)
}
