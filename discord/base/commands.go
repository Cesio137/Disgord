package base

import "github.com/bwmarrin/discordgo"

type SlashCommand struct {
	Command *discordgo.ApplicationCommand
	Run func(interaction discordgo.Interaction)
}

type ISlashCommand interface {
	Register()
}

func (c SlashCommand) Register() {
	appCommand := c.Command

	_, err := Client.Sess.ApplicationCommandCreate(Client.Sess.State.User.ID, "", appCommand)
    if err != nil {
        log.Fatal("Erro ao criar comando:", err)
    }
}

var Commands []SlashCommand
