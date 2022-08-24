package main

import (
	"flag"
	"log"
)

func main() {
	TelegramBot(mustToken())
}

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
