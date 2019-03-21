package spellforge

import (
	"github.com/soliel/SpellBot/command"
	"github.com/soliel/SpellBot/data"
	"log"
)

func CreateSpell(ctx command.Context) {
	player, err := data.GetPlayerByID(ctx.Author.ID)
	if err != nil {
		log.Printf("Create spell could not execute due to: %v", err)
	}

	if player == data.Player{} {
		data.InsertNewPlayer(ctx.Author.ID)
	}
}