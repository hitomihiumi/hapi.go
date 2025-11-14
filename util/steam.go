package util

import (
	"context"
	"strings"

	"github.com/Jleagle/steam-go/steamapi"
	"github.com/PuerkitoBio/goquery"

	"net/http"
	"strconv"
)

type Assets struct {
	Capsule        string `json:"capsule"`
	Capsule2X      string `json:"capsule_2x"`
	SmallCapsule   string `json:"small_capsule"`
	SmallCapsule2X string `json:"small_capsule_2x"`
	Header         string `json:"header"`
	Header2X       string `json:"header_2x"`
	Logo           string `json:"logo"`
	Logo2X         string `json:"logo_2x"`
	Hero           string `json:"hero"`
	Hero2X         string `json:"hero_2x"`
}

type ProcessedGameRecently struct {
	steamapi.RecentlyPlayedGame `json:",inline"`
	Assets                      `json:"assets"`
}

func ProcessGamesAssetsRecently(games []steamapi.RecentlyPlayedGame) []ProcessedGameRecently {
	var processedGames []ProcessedGameRecently

	for _, game := range games {
		appIDStr := strconv.Itoa(game.AppID)
		processedGames = append(processedGames, ProcessedGameRecently{
			RecentlyPlayedGame: game,
			Assets: struct {
				Capsule        string `json:"capsule"`
				Capsule2X      string `json:"capsule_2x"`
				SmallCapsule   string `json:"small_capsule"`
				SmallCapsule2X string `json:"small_capsule_2x"`
				Header         string `json:"header"`
				Header2X       string `json:"header_2x"`
				Logo           string `json:"logo"`
				Logo2X         string `json:"logo_2x"`
				Hero           string `json:"hero"`
				Hero2X         string `json:"hero_2x"`
			}{
				Capsule:        `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/library_600x900.jpg`,
				Capsule2X:      `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/library_600x900_2x.jpg`,
				SmallCapsule:   `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/capsule_231x87.jpg`,
				SmallCapsule2X: `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/capsule_231x87_2x.jpg`,
				Header:         `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/header.jpg`,
				Header2X:       `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/header_2x.jpg`,
				Logo:           `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/logo.png`,
				Logo2X:         `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/logo_2x.png`,
				Hero:           `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/library_hero.jpg`,
				Hero2X:         `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/library_hero_2x.jpg`,
			},
		})
	}

	return processedGames
}

type OwnedGame struct {
	AppID                    int    `json:"appid"`
	Name                     string `json:"name"`
	PlaytimeForever          int    `json:"playtime_forever"`
	PlaytimeWindows          int    `json:"playtime_windows_forever"`
	PlaytimeMac              int    `json:"playtime_mac_forever"`
	PlaytimeLinux            int    `json:"playtime_linux_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	ImgLogoURL               string `json:"img_logo_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
}

type ProcessedGameOwned struct {
	OwnedGame `json:",inline"`
	Assets    `json:"assets"`
}

func ProcessGamesAssetsOwned(games []OwnedGame) []ProcessedGameOwned {
	var processedGames []ProcessedGameOwned
	for _, game := range games {
		appIDStr := strconv.Itoa(game.AppID)
		processedGames = append(processedGames, ProcessedGameOwned{
			OwnedGame: game,
			Assets: struct {
				Capsule        string `json:"capsule"`
				Capsule2X      string `json:"capsule_2x"`
				SmallCapsule   string `json:"small_capsule"`
				SmallCapsule2X string `json:"small_capsule_2x"`
				Header         string `json:"header"`
				Header2X       string `json:"header_2x"`
				Logo           string `json:"logo"`
				Logo2X         string `json:"logo_2x"`
				Hero           string `json:"hero"`
				Hero2X         string `json:"hero_2x"`
			}{
				Capsule:        `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/library_600x900.jpg`,
				Capsule2X:      `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/library_600x900_2x.jpg`,
				SmallCapsule:   `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/capsule_231x87.jpg`,
				SmallCapsule2X: `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/capsule_231x87_2x.jpg`,
				Header:         `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/header.jpg`,
				Header2X:       `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/header_2x.jpg`,
				Logo:           `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/logo.png`,
				Logo2X:         `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/logo_2x.png`,
				Hero:           `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/library_hero.jpg`,
				Hero2X:         `https://cdn.cloudflare.steamstatic.com/steam/apps/` + appIDStr + `/library_hero_2x.jpg`,
			},
		})
	}

	return processedGames
}

type ProcessedProfile struct {
	steamapi.PlayerSummary `json:",inline"`
	Background             *string `json:"background,omitempty"`
	Frame                  *string `json:"frame,omitempty"`
	Level                  int     `json:"level"`
	AvatarFull             string  `json:"avatarfull"`
}

func ScrapeProfile(ctx context.Context, player steamapi.PlayerSummary) (*ProcessedProfile, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", player.ProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	profile := &ProcessedProfile{
		PlayerSummary: player,
		AvatarFull:    player.AvatarFull,
	}

	if bg, exists := doc.Find(".profile_background_image video").Attr("poster"); exists {
		profile.Background = &bg
	} else if bg, exists := doc.Find(".profile_animated_background video").Attr("poster"); exists {
		profile.Background = &bg
	}

	if frame, exists := doc.Find(".profile_avatar_frame img").Attr("src"); exists {
		profile.Frame = &frame
	}

	levelText := strings.TrimSpace(doc.Find(".friendPlayerLevelNum").Text())
	if level, err := strconv.Atoi(levelText); err == nil {
		profile.Level = level
	}

	if avatar, exists := doc.Find(".playerAvatarAutoSizeInner img").Attr("src"); exists {
		profile.AvatarFull = avatar
	}

	return profile, nil
}
