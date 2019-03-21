package main

import (
	"fmt"
	"github.com/soliel/SpellBot/command"
	"github.com/soliel/SpellBot/config"
	"github.com/soliel/SpellBot/data"
	"log"
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
	config, err := config.LoadConfig("SpellBotConfig.json")
	if err != nil {
		fmt.Println("Error loading configuration, ", err)
		return
	}
	conf = config

	dg, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		fmt.Println("Error starting discord session, ", err)
		return
	}

	err = data.InitDB(data.CreateDatabaseString(*conf) + "/SpellBot")
	if err != nil {
		log.Panic("Unable to connect to the database: ", err)
	}

	dg.AddHandler(onMessageRecieved)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening communication with discord, ", err)
		return
	}

	handler = command.CreateHandler()

	for i, guild := range dg.State.Guilds {
		fmt.Print(guild.Name, ", ", guild.ID, ", ", i, "\n")
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func onMessageRecieved(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("Msg: %v | Author: %v\n", m.Content, m.Author)
	command := filterMessages(s, m)
	if command.Command == "" {
		return
	}

	handler.HandleCommand(m, s, command)
}
