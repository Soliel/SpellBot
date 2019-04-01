package spellforge

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/soliel/SpellBot/command"
	"github.com/soliel/SpellBot/config"
	"github.com/soliel/SpellBot/data"
)

var (
	//This map will get dumped and remade every hour.
	spellLimiter map[string]int

	generator *rand.Rand
)

//CreateSpellCommand is the command users will use to forge a new spell
func CreateSpellCommand(ctx command.Context) {
	argMap, err := command.SeperateArgs(ctx.Content, "n", "d", "e")
	if _, present := argMap[byte('n')]; !present {
		return
	}

	/*if _, present := argMap[byte('t')]; !present {
		return
	}*/

	player, err := data.GetPlayerByID(ctx.Author.ID)
	if err != nil {
		log.Printf("Create spell could not execute due to: %v", err)
	}

	if player == (data.Player{}) {
		data.InsertNewPlayer(ctx.Author.ID)
	}

	/*if value, present := spellLimiter[ctx.Author.ID]; present {
		if value >= 10 {
			ctx.Session.ChannelMessageSend(ctx.Channel.ID, "You have exceeded the spell creation limit for now. try again in: ")
			//TODO: Undo spell creation.
			return
		} else {
			spellLimiter[ctx.Author.ID]++
		}
	} else {
		spellLimiter[ctx.Author.ID] = 1
	}*/

	spell := randomizeSpellComponents(argMap)
	spell.SpellType = 0
	spell.Name = argMap[byte('n')]
	if argMap[byte('d')] != "" {
		spell.Description = argMap[byte('d')]
	}
	if argMap[byte('e')] != "" {
		spell.Element = argMap[byte('e')]
	}

	spell.PlayerID = ctx.Author.ID

	err = data.InsertNewSpell(*spell)
	if err != nil {
		ctx.Session.ChannelMessageSend(ctx.Channel.ID, "Unable to create spell: Spell already exists in your spell book.")
		return
	}

	ctx.Session.ChannelMessageSend(ctx.Channel.ID, "You craft the spell: ")
	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Channel.ID, spell.CreateSpellEmbed(ctx))
	if err != nil {
		fmt.Println(err)
	}
}

func randomizeSpellComponents(argMap map[byte]string) *data.Spell {
	if generator == nil {
		generator = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	tierConf := config.TierMap[1].TierBaseSettings

	randDamage := getRandomNormal(tierConf.DamageVariance)
	randEfficency := getRandomNormal(tierConf.EfficencyVariance)
	randCastTime := getRandomNormal(tierConf.CastTimeVariance)
	randCooldown := getRandomNormal(tierConf.CooldownVarience)

	totalDamage := int(math.Floor(float64(tierConf.Damage)*(float64(randDamage)/100))) + tierConf.Damage
	if totalDamage < 1 {
		totalDamage = 1
	}

	totalEfficency := int(math.Floor(float64(tierConf.Efficency)*(float64(randEfficency)/100))) + tierConf.Efficency
	if totalEfficency < 10 {
		totalEfficency = 10
	}

	spell := data.Spell{
		Damage:     totalDamage,
		Efficency:  totalEfficency,
		Casttime:   tierConf.CastTime + randCastTime,
		Cooldown:   tierConf.Cooldown + randCooldown,
		Complexity: 1,
	}

	spell.ManaCost = calculateManaCost(spell)

	return &spell
}

func getRandomNormal(varience int) int {
	meanVarience := float64(varience / 2)
	i := -1
	for i <= 0 || i > varience {
		i = int(math.Floor(generator.NormFloat64()*(meanVarience*0.75) + meanVarience))
	}

	return i
}

func calculateManaCost(spell data.Spell) int {
	return int(math.Ceil((0.08/float64(spell.Efficency))*float64(spell.Damage) + 2))
}

func startCreationTicker() *time.Ticker {
	ticker := time.NewTicker(time.Hour)

	go func() {
		for range ticker.C {
			spellLimiter = make(map[string]int)
		}
	}()

	return ticker
}
