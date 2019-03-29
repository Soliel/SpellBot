package data

import (
	"database/sql"
)

//Player is a container for all player information.
type Player struct {
	PlayerID              string //Effectively the Discord User ID.
	MaximumMana           int
	ManaRegen             int //PerSecond Amount
	SpellSkill            int
	ChantingSkill         int
	MainElementalAffinity string
	MainElementTier       int
	SubElementalAffinity  string
	SubElementTier        int
	MageRank              int
	MageExperience        int
}

//GetPlayerByID is used to return a Player struct from the database.
func GetPlayerByID(playerID string) (Player, error) {
	rows := db.QueryRow("SELECT * FROM PlayerTable WHERE PlayerID = $1", playerID)
	player := Player{}

	err := rows.Scan(&player.PlayerID,
		&player.MaximumMana,
		&player.ManaRegen,
		&player.SpellSkill,
		&player.ChantingSkill,
		&player.MainElementalAffinity,
		&player.SubElementalAffinity,
		&player.MainElementTier,
		&player.SubElementTier,
		&player.MageRank,
		&player.MageExperience,
	)
	if err == sql.ErrNoRows {
		return Player{}, nil
	}
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

//InsertNewPlayer update the database with a new player.
func InsertNewPlayer(playerID string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//oh look a comment
	stmt, err := tx.Prepare("INSERT INTO PlayerTable(PlayerID) VALUES($1)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(playerID)
	if err != nil {
		return err
	}

	tx.Commit()
	stmt.Close()
	return nil
}

/*UpdatePlayer will overwrite a Player in the database with the values from the struct.
 *This is best used after you've gotten a player struct from GetPlayerByID
 */
func UpdatePlayer(player Player) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		UPDATE PlayerTable
			SET MaximumMana = $1,
			ManaRegen = $2,
			SpellSkill = $3,
			ChantingSkill = $4,
			MainAffinity = $5,
			SubAffinity = $6,
			MainAffinityTier = $7,
			SubAffinityTier = $8,
			MageRank = $9,
			Experience = $10,
		WHERE PlayerID = $11`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		player.MaximumMana,
		player.ManaRegen,
		player.SpellSkill,
		player.ChantingSkill,
		player.MainElementalAffinity,
		player.SubElementalAffinity,
		player.MainElementTier,
		player.SubElementTier,
		player.MageRank,
		player.MageExperience,
		player.PlayerID,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

//RemovePlayer will scrub a player's stats and existing spells from the database.
func RemovePlayer(playerID string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("DELETE FROM PlayerTable WHERE PlayerID = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(playerID)
	if err != nil {
		return err
	}

	stmt.Close()
	tx.Commit()

	return nil
}
