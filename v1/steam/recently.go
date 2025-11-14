package steam

import (
	"context"
	"os"

	"github.com/Jleagle/steam-go/steamapi"
	_ "github.com/joho/godotenv/autoload"
	"hapi.go/util"
)

//encore:api public method=GET path=/v1/steam/user/:id/games/recently
func Recently(ctx context.Context, id int64) (*RecentlyResponse, error) {
	client := steamapi.NewClient()
	client.SetKey(os.Getenv("STEAM_API_KEY"))

	games, err := client.GetRecentlyPlayedGames(id)
	if err != nil {
		return nil, err
	}

	res := util.ProcessGamesAssetsRecently(games)

	return &RecentlyResponse{Data: res, StatusCode: 200}, nil
}

type RecentlyResponse struct {
	Data       []util.ProcessedGameRecently `json:"data"`
	StatusCode int                          `json:"status"`
}
