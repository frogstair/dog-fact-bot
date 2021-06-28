package main

import (
	"dogfact/bot"
	"dogfact/fact"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	godotenv.Load()
	go fact.Update()
	bot.Start()
}
