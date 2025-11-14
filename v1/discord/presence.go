package discord

import (
	"context"
	"os"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
	"hapi.go/bot"
)

//encore:api public method=GET path=/v1/discord/user/:id/presence
func Presence(ctx context.Context, id string) (*PresenceResponse, error) {
	if bot.Get() == nil {
		if err := bot.Start(); err != nil {
			return nil, err
		}
	}

	session := bot.Get()
	if session == nil {
		return nil, ErrNoSession
	}

	presence, err := session.State.Presence(os.Getenv("BASE_GUILD_ID"), id)

	if err != nil {
		return nil, err
	}

	return &PresenceResponse{Data: presence.Activities, StatusCode: 200}, nil
}

type PresenceResponse struct {
	Data       []*discordgo.Activity `json:"data"`
	StatusCode int                   `json:"status"`
}
