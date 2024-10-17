package base

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type App struct {
	Sess *discordgo.Session
	Err  error
}

func AppBootstrap(token string, intents discordgo.Intent) *App {
	sess, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("❌ Error trying to create a client.")
		log.Fatal(err.Error())
		return &App{}
	}

	sess.AddHandler(HandleCommands)

	if intents == 0 {
		sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	} else {
		sess.Identify.Intents = intents
	}

	err = sess.Open()
	if err != nil {
		fmt.Println("❌ Error trying to create a open session.")
		log.Fatal(err.Error())
		return &App{}
	}
	
	defer sess.Close()

	RegisterCommands(sess)

	fmt.Println("➝ Online as", sess.State.User.Username)

	return &App{
		Sess: sess,
		Err:  err,
	}
}
