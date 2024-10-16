package base

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

type App struct {
	Sess *discordgo.Session
	Err  error
}

type IApp interface {
	Bootstrap(token string) App
}

func (a App) Bootstrap(token string) App {
	a.Sess, a.Err = discordgo.New(token)
	if a.Err != nil {
		fmt.Println("Error trying to create a client.")
		log.Fatal(a.Err.Error())
		return a
	}
	a.Err = a.Sess.Open()
	if a.Err != nil {
		fmt.Println("Error trying to create a open session.")
		log.Fatal(a.Err.Error())
		return a
	}
	defer a.Sess.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

    return a
}

var Client App