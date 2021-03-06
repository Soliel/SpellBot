package data

import (
	"reflect"
	"testing"
)

func TestGetSpellByNameAndPlayer(t *testing.T) {
	err := OpenTestDB()
	if err != nil {
		t.Errorf(err.Error())
	}

	type args struct {
		name     string
		playerID string
	}
	tests := []struct {
		name    string
		args    args
		want    Spell
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSpellByNameAndPlayer(tt.args.name, tt.args.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpellByNameAndPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpellByNameAndPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertNewSpell(t *testing.T) {
	err := OpenTestDB()
	if err != nil {
		t.Errorf(err.Error())
	}

	InsertNewPlayer("123")

	InsertNewSpell(Spell{
		SpellID:     0,
		PlayerID:    "123",
		Name:        "duplicate spell",
		Description: "A spell for testing.",
		Element:     "Test",
		SpellType:   Attack,
		Complexity:  1,
		Damage:      1,
		ManaCost:    1,
		Efficency:   1,
		Casttime:    1,
		Cooldown:    1,
	})

	type args struct {
		spell Spell
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Spell Insert",
			args: args{
				spell: Spell{
					SpellID:     0,     //SpellID will be set by the database
					PlayerID:    "123", //123 is the test player. discord ids are time based so it cannot collide.
					Name:        "test spell",
					Description: "A spell for testing.",
					Element:     "Test",
					SpellType:   Attack,
					Complexity:  1,
					Damage:      1,
					ManaCost:    1,
					Efficency:   1,
					Casttime:    1,
					Cooldown:    1,
				},
			},
			wantErr: false,
		},
		{
			name: "Duplicate Spell insert",
			args: args{
				spell: Spell{
					SpellID:     0,
					PlayerID:    "123",
					Name:        "duplicate spell",
					Description: "A spell for testing.",
					Element:     "Test",
					SpellType:   Attack,
					Complexity:  1,
					Damage:      1,
					ManaCost:    1,
					Efficency:   1,
					Casttime:    1,
					Cooldown:    1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertNewSpell(tt.args.spell); (err != nil) != tt.wantErr {
				t.Errorf("InsertNewSpell() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	RemovePlayer("123")
}

func TestUpdateSpell(t *testing.T) {
	err := OpenTestDB()
	if err != nil {
		t.Errorf(err.Error())
	}

	type args struct {
		spell Spell
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateSpell(tt.args.spell); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSpell() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
