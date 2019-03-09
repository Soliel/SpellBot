package data

import (
	"database/sql"
	"errors"
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
	rows, err := db.Query("SELECT * FROM SpellTable WHERE PlayerID = ? AND SpellName = ?")
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

	err = tx.QueryRow("SELECT SpellID FROM SpellTable WHERE PlayerID = ? AND SpellName = ?").Scan(nil)
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
		?,?,?,?,?,?,?,?,?,?,?
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
		SET SpellName = ?,
		SpellDesc = ?,
		SpellElement = ?,
		SpellType = ?,
		SpellComplexity = ?,
		Damage = ?,
		ManaCost = ?,
		Efficency = ?,
		CastTime = ?,
		CoolDown = ?,
	WHERE SpellID = ?
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
