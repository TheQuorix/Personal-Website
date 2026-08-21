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

type TrackData struct {
	Artist     string    `json:"artist"`
	Name       string    `json:"name"`
	ImageURL   string    `json:"image_url"`
	SongURL    string    `json:"song_url"`
	NowPlaying bool      `json:"now_playing"`
	Date       time.Time `json:"date"`
}
