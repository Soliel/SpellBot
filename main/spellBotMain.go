package main

import (
	"fmt"
	"github.com/Soliel/SpellBot/command"
	"github.com/Soliel/SpellBot/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	conf    *config.Config
	handler *command.Handler
)

func main() {
	conf, err := config.LoadConfig("SpellBotConfig.json")
	if err != nil {
		fmt.Println("Error loading configuration, ", err)
		return
	}

	dg, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		fmt.Println("Error starting discord session, ", err)
		return
	}

	dg.AddHandler(onMessageRecieved)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening communication with discord, ", err)
		return
	}

	handler = command.CreateHandler()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func onMessageRecieved(s *discordgo.Session, m *discordgo.MessageCreate) {
	command := filterMessages(s, m)
	if command.Command == "" {
		return
	}

	handler.HandleCommand(m, s, command)
}
