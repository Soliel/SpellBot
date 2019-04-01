package data

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/soliel/SpellBot/command"
)

//Spell will store spell data to use with the rest of the bot.
type Spell struct {
	SpellID     int
	PlayerID    string
	Name        string
	Description string
	Element     string
	SpellType   SpellType
	Complexity  int
	Damage      int
	ManaCost    int
	Efficency   int
	Casttime    int
	Cooldown    int
	//SpellEffects []EffectObj //Planned Feature
}

//GetSpellByNameAndPlayer gets the first result spell of the player ID that matches the name specified
func GetSpellByNameAndPlayer(name string, playerID string) (Spell, error) {
	rows, err := db.Query("SELECT * FROM SpellTable WHERE PlayerID = $1 AND SpellName = $2", playerID, name)
	if err != nil {
		return Spell{}, err
	}

	defer rows.Close()
	spell := Spell{}

	rows.Next()
	err = rows.Scan(&spell.SpellID,
		&spell.Name,
		&spell.Description,
		&spell.Element,
		&spell.SpellType,
		&spell.Complexity,
		&spell.Damage,
		&spell.ManaCost,
		&spell.Efficency,
		&spell.Casttime,
		&spell.Cooldown,
		&spell.PlayerID,
	)
	if err != nil {
		return spell, err
	}

	rows.Close()
	return spell, nil
}

//InsertNewSpell Adds a new Spell Object to the database.
func InsertNewSpell(spell Spell) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//Prepare the statement to hopefully avoid sql injection.
	stmt1, err := tx.Prepare("SELECT SpellID, SpellName, PlayerID FROM SpellTable WHERE PlayerID = $1 AND SpellName = $2")
	err = stmt1.QueryRow(spell.PlayerID, spell.Name).Scan(nil, nil, nil)
	if err == nil {
		return errors.New("spell already exists")
	} else if err != sql.ErrNoRows {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO SpellTable (
		SpellName,
		SpellDesc,
		SpellElement,
		SpellType,
		SpellComplexity,
		Damage,
		ManaCost,
		Efficency,
		CastTime,
		CoolDown,
		PlayerID
	)VALUES(
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
	)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		spell.Name,
		spell.Description,
		spell.Element,
		spell.SpellType,
		spell.Complexity,
		spell.Damage,
		spell.ManaCost,
		spell.Efficency,
		spell.Casttime,
		spell.Cooldown,
		spell.PlayerID,
	)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

//UpdateSpell writes the Spell Struct to the database.
func UpdateSpell(spell Spell) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
	UPDATE SpellTable
		SET SpellName = $1,
		SpellDesc = $2,
		SpellElement = $3,
		SpellType = $4,
		SpellComplexity = $5,
		Damage = $6,
		ManaCost = $7,
		Efficency = $8,
		CastTime = $9,
		CoolDown = $10,
	WHERE SpellID = $11
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		spell.Name,
		spell.Description,
		spell.Element,
		spell.SpellType,
		spell.Complexity,
		spell.Damage,
		spell.ManaCost,
		spell.Efficency,
		spell.Casttime,
		spell.Cooldown,
		spell.SpellID,
	)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

//RemoveSpellByNameAndPlayerID is used for testing and potentially for admin duties.
func RemoveSpellByNameAndPlayerID(name string, playerID string) error {
	stmt, err := db.Prepare("DELETE FROM SpellTable WHERE SpellName = $1 AND PlayerID = $2")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, playerID)
	if err != nil {
		return err
	}

	stmt.Close()

	return nil
}

func (spell *Spell) CreateSpellEmbed(ctx command.Context) *discordgo.MessageEmbed {
	if spell.Element == "" {
		spell.Element = "None"
	}

	embed := &discordgo.MessageEmbed{
		Title:       spell.Name,
		Description: spell.Description,
		Color:       14030101,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    ctx.Author.Username,
			IconURL: ctx.Author.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{
			{"Spell Type", "Attack", true},
			{"Damage", strconv.FormatInt(int64(spell.Damage), 10), true},
			{"Element", spell.Element, true},
			{"Mana Cost", strconv.FormatInt(int64(spell.ManaCost), 10), true},
			{"Cast Time", strconv.FormatFloat(float64(spell.Casttime)/1000, 'f', 2, 64), true},
			{"Cooldown", strconv.FormatFloat(float64(spell.Cooldown)/1000, 'f', 2, 64), true},
		},
	}

	return embed
}
