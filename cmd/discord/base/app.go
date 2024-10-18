package base

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type App struct {
	Sess *discordgo.Session
	Err  error
}

func AppBootstrap(token string, intents discordgo.Intent) *App {
	sess, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal("❌ Error trying to create a client.\n", err.Error())
		return &App{}
	}

	if intents == 0 {
		sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	} else {
		sess.Identify.Intents = intents
	}

	RegisterEvents(sess)

	err = sess.Open()
	if err != nil {
		log.Fatal("❌ Error trying to open session.\n", err.Error())
		return &App{}
	}

	RegisterCommands(sess)

	log.Println("➝ Online as", sess.State.User.Username)

	return &App{
		Sess: sess,
		Err:  err,
	}
}
