package main

import (
	"flag"
	"log"

	"github.com/koha90/tg-pocket-bot/clients/telegram"
	event_consumer "github.com/koha90/tg-pocket-bot/consumer/event-consumer"
	tg "github.com/koha90/tg-pocket-bot/events/telegram"
	"github.com/koha90/tg-pocket-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := tg.New(
		telegram.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stoped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is empty")
	}

	return *token
}
