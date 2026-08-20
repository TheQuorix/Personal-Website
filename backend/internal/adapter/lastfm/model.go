package lastfm

import "time"

type lastfmResponse struct {
	RecentTraks struct {
		Track []rawTrack `json:"track"`
	} `json:"recenttracks"`
}

type rawTrack struct {
	Artist struct {
		Text string `json:"#text"`
	} `json:"artist"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Image []struct {
		Size string `json:"size"`
		Text string `json:"#text"`
	} `json:"image"`
	Attr struct {
		NowPlaying string `json:"nowplaying"`
	} `json:"@attr"`
	Date struct {
		UTS string `json:"uts"`
	} `json:"date"`
}

type Track struct {
	Artist     string
	Name       string
	ImageURL   string
	SongURL    string
	NowPlaying bool
	Date       time.Time
}
