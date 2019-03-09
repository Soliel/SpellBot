package command

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestHandler_Register(t *testing.T) {
	handler := CreateHandler()
	type args struct {
		name    string
		command commandFunc
	}
	tests := []struct {
		name    string
		handler Handler
		args    args
		wantErr bool
	}{
		{
			name:    "Valid Registration",
			handler: *handler,
			args: args{
				name:    "test",
				command: func(Context) {},
			},
			wantErr: false,
		},
		{
			name:    "Registration with Space",
			handler: *handler,
			args: args{
				name:    "test command",
				command: func(Context) {},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.Register(tt.args.name, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_RegisterDuplicate(t *testing.T) {
	handler := CreateHandler()
	handler.Register("test1", func(Context) {})

	type args struct {
		name    string
		command commandFunc
	}
	tests := []struct {
		name    string
		handler Handler
		args    args
		wantErr bool
	}{
		{
			name:    "Duplicate name",
			handler: *handler,
			args: args{
				name:    "test1",
				command: func(Context) {},
			},
			wantErr: true,
		},
		{
			name:    "Duplicate abbreviation",
			handler: *handler,
			args: args{
				name:    "test2",
				command: func(Context) {},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.Register(tt.args.name, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_HandleCommand(t *testing.T) {
	handler := CreateHandler()
	handler.Register("test", func(Context) {})

	type args struct {
		m       *discordgo.MessageCreate
		s       *discordgo.Session
		command CommandMessage
	}
	tests := []struct {
		name    string
		handler Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.HandleCommand(tt.args.m, tt.args.s, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("Handler.HandleCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
