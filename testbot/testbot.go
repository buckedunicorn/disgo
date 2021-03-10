package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chebyrash/promise"
	log "github.com/sirupsen/logrus"

	"github.com/DiscoOrg/disgo"
	"github.com/DiscoOrg/disgo/api"
	"github.com/DiscoOrg/disgo/internal/events"
)

func main() {
	token := os.Getenv("token")

	dgo, err := disgo.NewBuilder(token).
		SetLogLevel(log.InfoLevel).
		SetIntents(api.IntentsGuildMessages | api.IntentsGuildMembers).
		AddEventListeners(&events.ListenerAdapter{
			OnGuildMessageReceived: messageHandler,
		}).
		Build()
	if err != nil {
		return
	}

	err = dgo.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer dgo.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-s
}

func messageHandler(event events.GuildMessageReceivedEvent) {
	log.Printf("Message received: %v", event.Message.Content)
	if event.Message.Author.IsBot {
		return
	}

	switch event.Message.Content {
	case "ping":
		log.Print("hm")
		event.TextChannel.SendMessage("pong")
	case "pong":
		log.Print("hm2")
		event.TextChannel.SendMessage("ping")
	case "dm":
		event.Message.Author.OpenDMChannel().Then(func(channel promise.Any) promise.Any {
			return channel.(*api.DMChannel).SendMessage("helo")
		}).Then(func(_ promise.Any) promise.Any {
			return event.Message.AddReaction("✅")
		}).Catch(func(_ error) error {
			event.Message.AddReaction("❌")
			return nil
		})
	}
}