package base

import (
	"log"
	"disgord/cmd/discord/events"

	"github.com/bwmarrin/discordgo"
)

func RegisterEvents(s *discordgo.Session) {
	log.Println("Registering events!")
	for _, ev := range events.Events {
		s.AddHandler(ev)
	}
}
