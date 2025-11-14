package discord

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
	"hapi.go/bot"
)

type UserType struct {
	discordgo.User `json:",inline"`
	AvatarURL      string `json:"avatar_url"`
	BannerURL      string `json:"banner_url"`
}

//encore:api public method=GET path=/v1/discord/user/:id
func User(ctx context.Context, id string) (*UserResponse, error) {
	if bot.Get() == nil {
		if err := bot.Start(); err != nil {
			return nil, err
		}
	}

	session := bot.Get()
	if session == nil {
		return nil, ErrNoSession
	}

	user, err := session.User(id)
	if err != nil {
		return nil, err
	}

	res := UserType{
		User:      *user,
		AvatarURL: user.AvatarURL("1024"),
		BannerURL: user.BannerURL("1024"),
	}

	return &UserResponse{Data: res, StatusCode: 200}, nil
}

type UserResponse struct {
	Data       UserType `json:"data"`
	StatusCode int      `json:"status"`
}

var ErrNoSession = fmt.Errorf("discord session not started")
