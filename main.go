package main

import (
	"disgord/cmd/discord/base"
	_ "disgord/cmd/discord/commands"
	"disgord/cmd/settings"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = base.AppBootstrap(settings.EnvSchema.BotToken, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
