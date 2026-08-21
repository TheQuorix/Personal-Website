package lastfm

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/TheQuorix/Personal-Website/internal/config"
)

type Client struct {
	httpClient *http.Client
	cfg        config.Config
}

func NewClient(httpClient *http.Client, cfg config.Config) *Client {
	return &Client{
		httpClient: httpClient,
		cfg:        cfg,
	}
}

func (c Client) Fetch(ctx context.Context) (TrackData, error) {
	reqURL := c.getUrl(c.cfg.LastFmUser, c.cfg.LastFmApiKey)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return TrackData{}, fmt.Errorf("build request: %w", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return TrackData{}, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Printf("close response body: %v", err)
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return TrackData{}, fmt.Errorf("read body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return TrackData{}, fmt.Errorf("unexpected status %d: %s", response.StatusCode, body)
	}

	var lastfmResponse lastfmResponse
	if err := json.Unmarshal(body, &lastfmResponse); err != nil {
		return TrackData{}, fmt.Errorf("unmarshal: %w", err)
	}

	tracks := lastfmResponse.RecentTraks.Track
	if len(tracks) == 0 {
		return TrackData{}, fmt.Errorf("no tracks found")
	}

	return convert(tracks[0])
}

func (c Client) getUrl(user string, apiKey string) string {
	params := url.Values{}
	params.Set("method", "user.getrecenttracks")
	params.Set("user", user)
	params.Set("api_key", apiKey)
	params.Set("format", "json")
	params.Set("limit", "1")

	return "https://ws.audioscrobbler.com/2.0/?" + params.Encode()
}
