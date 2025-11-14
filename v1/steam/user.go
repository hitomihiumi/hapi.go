package steam

import (
	"context"
	"os"

	"github.com/Jleagle/steam-go/steamapi"
	_ "github.com/joho/godotenv/autoload"
	"hapi.go/util"
)

//encore:api public method=GET path=/v1/steam/user/:id
func User(ctx context.Context, id int64) (*UserResponse, error) {
	client := steamapi.NewClient()
	client.SetKey(os.Getenv("STEAM_API_KEY"))

	user, err := client.GetPlayer(id)
	if err != nil {
		return nil, err
	}

	res, err := util.ScrapeProfile(ctx, user)

	return &UserResponse{Data: res, StatusCode: 200}, nil
}

type UserResponse struct {
	Data       *util.ProcessedProfile `json:"data"`
	StatusCode int                    `json:"status"`
}
