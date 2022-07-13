package main

import (
	"flag"
	"log"
)

//"fmt"

func main() {
	t:= mustToken()

	// token = flags.Get(token)

	// tgClient = telegram.New(token)
	
	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token:=flag.String(
		"token-bot-token",
		"",
		"token for access to telegramm bot",
	)

	flag.Parse()

	if *token == ""{
		log.Fatal("token is not specified")
	}

	return *token
}