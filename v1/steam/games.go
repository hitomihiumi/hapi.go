package steam

import (
	"context"
	"os"

	"hapi.go/util"

	"github.com/Jleagle/steam-go/steamapi"
	_ "github.com/joho/godotenv/autoload"
)

//encore:api public method=GET path=/v1/steam/user/:id/games
func Games(ctx context.Context, id int64) (*GamesResponse, error) {
	client := steamapi.NewClient()
	client.SetKey(os.Getenv("STEAM_API_KEY"))

	games, err := client.GetOwnedGames(id)
	if err != nil {
		return nil, err
	}

	ownedGames := make([]util.OwnedGame, len(games.Games))
	for i, g := range games.Games {
		ownedGames[i] = util.OwnedGame{
			AppID:                    g.AppID,
			Name:                     g.Name,
			PlaytimeForever:          g.PlaytimeForever,
			PlaytimeWindows:          g.PlaytimeWindows,
			PlaytimeMac:              g.PlaytimeMac,
			PlaytimeLinux:            g.PlaytimeLinux,
			ImgIconURL:               g.ImgIconURL,
			ImgLogoURL:               g.ImgLogoURL,
			HasCommunityVisibleStats: g.HasCommunityVisibleStats,
		}
	}

	res := util.ProcessGamesAssetsOwned(ownedGames)

	return &GamesResponse{Data: res, StatusCode: 200}, nil
}

type GamesResponse struct {
	Data       []util.ProcessedGameOwned `json:"data"`
	StatusCode int                       `json:"status"`
}
