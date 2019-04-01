package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/soliel/SpellBot/command"
	"github.com/soliel/SpellBot/config"
	"github.com/soliel/SpellBot/data"
	"github.com/soliel/SpellBot/spellforge"
)

var (
	conf    *config.Config
	handler *command.Handler
)

func main() {
	loadConfBytes, err := ioutil.ReadFile("../Configuration Files/SpellBotTest.json")
	loadTierBytes, err := ioutil.ReadFile("../Configuration Files/TierTest.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	configuration, err := config.LoadConfig(loadConfBytes)
	err = config.LoadTierSettings(loadTierBytes)
	if err != nil {
		fmt.Println("Error loading configuration, ", err)
		return
	}
	conf = configuration

	dg, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		fmt.Println("Error starting discord session, ", err)
		return
	}

	err = data.InitDB(data.CreateDatabaseString(*conf))
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
	registerCommands()

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

func registerCommands() {
	handler.Register("createspell", spellforge.CreateSpellCommand)
	handler.Register("cast", spellforge.CastSpell)
}
