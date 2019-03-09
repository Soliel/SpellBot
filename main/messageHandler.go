package main

import (
	"github.com/Soliel/SpellBot/command"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func filterMessages(s *discordgo.Session, m *discordgo.MessageCreate) command.CommandMessage {
	var commandMsg command.CommandMessage

	if m.Author.ID == s.State.User.ID {
		return commandMsg
	}

	if len(m.Content) < len(conf.BotPrefix) {
		return commandMsg
	}

	if m.Content[:len(conf.BotPrefix)] != conf.BotPrefix {
		return commandMsg
	}

	content := m.Content[len(conf.BotPrefix):]
	if len(content) < 1 {
		return commandMsg
	}

	commandName := content[:strings.Index(content, " ")]
	commandName = strings.ToLower(commandName)

	content = content[len(commandName)+1:]

	commandMsg = command.CommandMessage{Command: commandName, Content: content}

	return commandMsg
}
