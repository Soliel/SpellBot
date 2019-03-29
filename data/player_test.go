/*
* 	This is a terrible way to test your db I'm too lazy to do it right.
 */
package data

import (
	"reflect"
	"testing"
)

func TestGetPlayerByID(t *testing.T) {
	err := OpenTestDB()
	if err != nil {
		t.Errorf(err.Error())
	}

	InsertNewPlayer("12345")

	type args struct {
		playerID string
	}
	tests := []struct {
		name    string
		args    args
		want    Player
		wantErr bool
	}{
		{
			name: "Valid get",
			args: args{
				playerID: "12345",
			},
			want: Player{
				PlayerID:              "12345",
				MaximumMana:           10,
				ManaRegen:             1,
				SpellSkill:            1,
				ChantingSkill:         1,
				MainElementalAffinity: "",
				SubElementalAffinity:  "",
				MainElementTier:       0,
				SubElementTier:        0,
				MageRank:              0,
				MageExperience:        0,
			},
			wantErr: false,
		},
		{
			name: "Player doesn't exist",
			args: args{
				playerID: "123456",
			},
			want:    Player{},
			wantErr: false,
		},
		{
			name: "Invalid player",
			args: args{
				playerID: "Hello!", //Why can a string be the wrong playerID? Because I'm lazy. again.
			},
			want:    Player{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPlayerByID(tt.args.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayerByID() = %v, want %v", got, tt.want)
			}
		})
	}

	RemovePlayer("12345")
}

func TestInsertNewPlayer(t *testing.T) {
	err := OpenTestDB()
	if err != nil {
		t.Errorf(err.Error())
	}

	InsertNewPlayer("11111")

	type args struct {
		playerID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test correct insert",
			args: args{
				playerID: "12345",
			},
			wantErr: false,
		},
		{
			name: "Test duplicate insert",
			args: args{
				playerID: "11111",
			},
			wantErr: true,
		},
		{
			name: "Test invalid insert",
			args: args{
				playerID: "invalidID",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertNewPlayer(tt.args.playerID); (err != nil) != tt.wantErr {
				t.Errorf("InsertNewPlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	RemovePlayer("12345")
	RemovePlayer("11111")
}

func TestUpdatePlayer(t *testing.T) {
	err := OpenTestDB()
	if err != nil {
		t.Errorf(err.Error())
	}

	InsertNewPlayer("12345")

	type args struct {
		player Player
	}
	tests := []struct {
		name    string
		args    args
		want    Player
		wantErr bool
	}{
		{
			name: "update exp",
			args: args{
				player: Player{
					PlayerID:              "12345",
					MaximumMana:           10,
					ManaRegen:             1,
					SpellSkill:            1,
					ChantingSkill:         1,
					MainElementalAffinity: "",
					SubElementalAffinity:  "",
					MainElementTier:       0,
					SubElementTier:        0,
					MageRank:              0,
					MageExperience:        200,
				},
			},
			want: Player{
				PlayerID:              "12345",
				MaximumMana:           10,
				ManaRegen:             1,
				SpellSkill:            1,
				ChantingSkill:         1,
				MainElementalAffinity: "",
				SubElementalAffinity:  "",
				MainElementTier:       0,
				SubElementTier:        0,
				MageRank:              0,
				MageExperience:        200,
			},
			wantErr: false,
		},
		{
			name: "Invalid Update",
			args: args{
				player: Player{
					PlayerID:              "Hello",
					MaximumMana:           10,
					ManaRegen:             1,
					SpellSkill:            1,
					ChantingSkill:         1,
					MainElementalAffinity: "",
					SubElementalAffinity:  "",
					MainElementTier:       0,
					SubElementTier:        0,
					MageRank:              0,
					MageExperience:        200,
				},
			},
			want: Player{
				PlayerID:              "",
				MaximumMana:           0,
				ManaRegen:             0,
				SpellSkill:            0,
				ChantingSkill:         0,
				MainElementalAffinity: "",
				SubElementalAffinity:  "",
				MainElementTier:       0,
				SubElementTier:        0,
				MageRank:              0,
				MageExperience:        0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdatePlayer(tt.args.player); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
			updateResult, _ := GetPlayerByID(tt.args.player.PlayerID)
			if !reflect.DeepEqual(updateResult, tt.want) {
				t.Errorf("UpdatePlayer() = %v, want %v", updateResult, tt.want)
			}
		})
	}

	RemovePlayer("12345")
}
