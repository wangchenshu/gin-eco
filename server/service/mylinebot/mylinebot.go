package mylinebot

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var MyBot *linebot.Client

func Init() *linebot.Client {
	MyBot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	return MyBot
}
