package main

import (
	"disgord/discord/base"
	"disgord/settings"
)

func main() {
	base.Client = base.App{}.Bootstrap(settings.EnvSchema.BotToken)
}
