package data

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
	rows := db.QueryRow("SELECT * FROM PlayerTable WHERE PlayerID = ?", playerID)
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
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

//InsertNewPlayer update the database with a new player.
func InsertNewPlayer(playerID string) error {
	stmt, err := db.Prepare("INSERT INTO PlayerTable(PlayerID) VALUES(?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(playerID)
	if err != nil {
		return err
	}

	return nil
}

//UpdatePlayer will overwrite a Player in the database with the values from the struct.
func UpdatePlayer(player Player) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		UPDATE PlayerTable
			SET MaximumMana = ?,
			ManaRegen = ?,
			SpellSkill = ?,
			ChantingSkill = ?,
			MainAffinity = ?,
			SubAffinity = ?,
			MainAffinityTier = ?,
			SubAffinityTier = ?,
			MageRank = ?,
			Experience = ?
		WHERE PlayerID = ?`)
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
