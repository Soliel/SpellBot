package main

import (
	"github.com/Soliel/SpellBot/config"
	"reflect"
	"testing"

	"github.com/Soliel/SpellBot/command"
	"github.com/bwmarrin/discordgo"
)

func mockSessionWithStateUserID(ID string) *discordgo.Session {
	mockSession := &discordgo.Session{
		State: &discordgo.State{
			Ready: discordgo.Ready{
				User: &discordgo.User{
					ID: ID,
				},
			},
		},
	}
	return mockSession
}

func mockMessageCreate(content string, AuthorID string) *discordgo.MessageCreate {
	msg := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: &discordgo.User{
				ID: AuthorID,
			},
			Content: content,
		},
	}

	return msg
}

func Test_filterMessages(t *testing.T) {
	conf = &config.Config{
		BotPrefix: "test ",
	}

	type args struct {
		s *discordgo.Session
		m *discordgo.MessageCreate
	}
	tests := []struct {
		name string
		args args
		want command.CommandMessage
	}{
		{
			name: "Matching UserID",
			args: args{
				s: mockSessionWithStateUserID("12345"),
				m: mockMessageCreate("test valid command", "12345"),
			},
			want: command.CommandMessage{},
		},
		{
			name: "Message Length too Short",
			args: args{
				s: mockSessionWithStateUserID("12345"),
				m: mockMessageCreate("test", "123"),
			},
			want: command.CommandMessage{},
		},
		{
			name: "Message Invalid Prefix",
			args: args{
				s: mockSessionWithStateUserID("12345"),
				m: mockMessageCreate("tess valid command", "12346"),
			},
			want: command.CommandMessage{},
		},
		{
			name: "Valid prefix no content",
			args: args{
				s: mockSessionWithStateUserID("12345"),
				m: mockMessageCreate("test ", "12346"),
			},
			want: command.CommandMessage{},
		},
		{
			name: "Valid command",
			args: args{
				s: mockSessionWithStateUserID("12345"),
				m: mockMessageCreate("test Valid Command", "12346"),
			},
			want: command.CommandMessage{
				Command: "valid",
				Content: "Command",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterMessages(tt.args.s, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
