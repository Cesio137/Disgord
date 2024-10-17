package settings

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvConfig struct {
	BotToken       string `validate:"required"`
	WebhookLogsURL string `validate:"omitempty,url"`
}

var EnvSchema EnvConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error trying to load .env file")
		return
	}
	
	config := EnvConfig{
		BotToken:       os.Getenv("BOT_TOKEN"),
		WebhookLogsURL: os.Getenv("WEBHOOK_LOGS_URL"),
	}
	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("❌ Error on field '%s': %s\n", err.Field(), err.ActualTag())
		}
		os.Exit(1)
		return
	}
	EnvSchema = config
	fmt.Println("✔ Env vars loaded successfully!")
}
