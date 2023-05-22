package bot

import (
	"dogfact/fact"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var TWORD = []string{"dog", "canine", "puppy", "bitch", "bark", "woof", "shiba"}

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

	if message == "how many facts do you have?" {
		fcount := fmt.Sprintf("I have %d facts for you", len(fact.List))
		s.ChannelMessageSend(m.ChannelID, fcount)
		log.Println("Asked how many facts")
		return
	}

	if strings.HasPrefix(message, "fact #") {
		message = strings.Trim(message, "fact #")
		message = strings.TrimFunc(message, func(r rune) bool {
			return r < '0' || r > '9'
		})
		num, err := strconv.Atoi(message)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Cannot understand the number "+message)
			return
		}

		num -= 1

		if num > len(fact.List)-1 {
			s.ChannelMessageSend(m.ChannelID, "Don't have as many facts!")
			return
		}

		msg := fmt.Sprintf("Fact #%s: %s", message, fact.List[num])

		s.ChannelMessageSend(m.ChannelID, msg)
		log.Printf("Dispenced fact #%d", num+1)
	}

	for _, t := range TWORD {
		if strings.Contains(message, t) {
			s.ChannelTyping(m.ChannelID)

			i := rand.Intn(len(fact.List))
			msg := fmt.Sprintf("Fact #%d: %s", i+1, fact.List[i])

			s.ChannelMessageSend(m.ChannelID, msg)

			log.Printf("Dispenced fact #%d", i+1)
			return
		}
	}
}
