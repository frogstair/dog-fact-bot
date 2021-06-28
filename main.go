package main

import (
	"dogfact/bot"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	godotenv.Load()
	bot.Start()
}
