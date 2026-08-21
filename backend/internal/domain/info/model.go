package info

import (
	"github.com/TheQuorix/Personal-Website/internal/adapter/github"
	"github.com/TheQuorix/Personal-Website/internal/adapter/lastfm"
	"github.com/TheQuorix/Personal-Website/internal/adapter/openweather"
	"github.com/TheQuorix/Personal-Website/internal/adapter/steam"
)

type Info struct {
	Weather openweather.WeatherData `json:"weather"`
	Music   lastfm.TrackData        `json:"music"`
	Steam   steam.SteamData         `json:"steam"`
	Github  github.GithubData       `json:"github"`
}
