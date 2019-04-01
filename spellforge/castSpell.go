package spellforge

import (
	"fmt"
	"github.com/soliel/SpellBot/command"
	"github.com/soliel/SpellBot/data"
)

//Cast spell is the command for showing off your spells
func CastSpell(ctx command.Context) {
	spell, err := data.GetSpellByNameAndPlayer(ctx.Content, ctx.Author.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Channel.ID, spell.CreateSpellEmbed(ctx))
	if err != nil {
		fmt.Println(err)
	}
}
