package steam

import (
	"context"
	"os"

	"github.com/Jleagle/steam-go/steamapi"
	_ "github.com/joho/godotenv/autoload"
)

//encore:api public method=GET path=/v1/steam/user/:id/games/achievements/:appid
func Achievements(ctx context.Context, id int64, appid int32) (*AchievementsResponse, error) {
	client := steamapi.NewClient()
	client.SetKey(os.Getenv("STEAM_API_KEY"))

	achievements, err := client.GetPlayerAchievements(uint64(id), uint32(appid))
	if err != nil {
		return nil, err
	}

	return &AchievementsResponse{Data: achievements, StatusCode: 200}, nil
}

type AchievementsResponse struct {
	Data       steamapi.PlayerAchievementsResponse `json:"data"`
	StatusCode int                                 `json:"status"`
}
