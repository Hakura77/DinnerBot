package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

const infPath string = "botinf.yaml"

func main() {
	token, err := loadToken()
	if err != nil {
		log.Fatalf("Unable to load token: %s", err.Error())
	}

	if err := LoadMeals(); err != nil {
		log.Fatalf("Unable to load meals: %s", err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch strings.ToLower(update.Message.Command()) {
			case commandHelp:
				msg.Text = Help(update.Message.CommandArguments())
			case commandOutOfFood:
				err := OutOfFood(update.Message.CommandArguments())
				if err != nil {
					msg.Text = err.Error()
					break
				}
				msg.Text = fmt.Sprintf("Successfully set food %s as out", update.Message.CommandArguments())
			case commandGotFood:
				msg.Text = GotFood(update.Message.CommandArguments())
			case commandPickMeal:
				msg.Text = PickMeal()
			case commandWhatNeeded:
				msg.Text = FoodNeeded()
			case commandReloadMeals:
				if err := LoadMeals(); err != nil {
					msg.Text = fmt.Sprintf("error loading meals: %s", err.Error())
				}
				msg.Text = "reloaded meals database"
			default:
				msg.Text = "I'm sorry, I don't know that command. Try /help"
			}
			_, _ = bot.Send(msg)
		}
	}
}


func loadToken() (token string, err error) {
	tokenFile, err := ioutil.ReadFile(infPath)
	if err != nil {
		return "", fmt.Errorf("unable to read %s: %w", infPath, err)
	}
	v := make(map[string]interface{})
	err = yaml.Unmarshal(tokenFile, &v)
	if err != nil {
		return "", fmt.Errorf("unable to unmarshal token file: %w", err)
	}
	token, ok := v["token"].(string)
	if !ok {
		return "", fmt.Errorf("unable to extract token from token file %s", infPath)
	}
	if len(token) == 0 {
		return "", fmt.Errorf("extracted empty token from token file %s", infPath)
	}
	return token, nil


}