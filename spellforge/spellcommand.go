package spellforge

import (
	"log"

	"github.com/soliel/SpellBot/command"
	"github.com/soliel/SpellBot/data"
)

var (
	//This map will get dumped and remade every hour.
	spellLimiter map[string]int
)

//CreateSpellCommand is the command users will use to forge a new spell
func CreateSpellCommand(ctx command.Context) {
	argMap, err := command.SeperateArgs(ctx.Content)
	if _, present := argMap[byte('n')]; !present {
		return
	}

	if _, present := argMap[byte('t')]; !present {
		return
	}

	player, err := data.GetPlayerByID(ctx.Author.ID)
	if err != nil {
		log.Printf("Create spell could not execute due to: %v", err)
	}

	if player == (data.Player{}) {
		data.InsertNewPlayer(ctx.Author.ID)
	}

	if value, present := spellLimiter[ctx.Author.ID]; present {
		if value >= 10 {
			ctx.Session.ChannelMessageSend(ctx.Channel.ID, "You have exceeded the spell creation limit for now. try again in: ")
			//TODO: Undo spell creation.
			return
		} else {
			spellLimiter[ctx.Author.ID]++
		}
	} else {
		spellLimiter[ctx.Author.ID] = 1
	}

}

func randomizeSpellComponents(argMap map[byte]string) *data.Spell {

}
