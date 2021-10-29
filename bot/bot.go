package bot

import (
	"dogfact/fact"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var TWORD = "dog"

func Start() {
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	dg.AddHandler(onMsg)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	log.Println("Running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
	fmt.Println("")
	log.Println("Stopped")
}

func onMsg(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.Bot {
		return
	}

	message := strings.ToLower(m.Content)

	if strings.Contains(message, TWORD) {
		s.ChannelTyping(m.ChannelID)

		i := rand.Intn(len(fact.List))
		msg := fmt.Sprintf("Fact #%d: %s", i+1, fact.List[i])

		s.ChannelMessageSend(m.ChannelID, msg)

		log.Printf("Dispenced fact #%d", i+1)
		return
	}

	if message == "how many facts?" {
		fcount := fmt.Sprintf("I have %d facts for you", len(fact.List))
		s.ChannelMessageSend(m.ChannelID, fcount)
		log.Println("Asked how many facts")
	}
}
