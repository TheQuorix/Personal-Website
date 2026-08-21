package lastfm

import (
	"fmt"
	"strconv"
	"time"
)

func convert(raw rawTrack) (TrackData, error) {
	track := TrackData{
		Artist:     raw.Artist.Text,
		Name:       raw.Name,
		NowPlaying: raw.Attr.NowPlaying == "true",
		SongURL:    raw.Url,
	}

	if len(raw.Image) > 0 {
		track.ImageURL = raw.Image[len(raw.Image)-1].Text
	}

	if raw.Date.UTS != "" {
		uts, err := strconv.ParseInt(raw.Date.UTS, 10, 64)
		if err != nil {
			return TrackData{}, fmt.Errorf("invalid uts: %w", err)
		}
		track.Date = time.Unix(uts, 0)
	}

	return track, nil
}
