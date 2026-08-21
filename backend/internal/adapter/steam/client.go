package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"sync"

	"github.com/TheQuorix/Personal-Website/internal/config"
)

var excludedAppIDs = map[int]bool{
	365670: true, // Blender
	431960: true, // Wallpaper Engige
	629520: true, // Soundpad
}

type Client struct {
	httpClient *http.Client
	cfg        config.Config
	imageCache *ImageCache
}

func NewClient(httpClient *http.Client, cfg config.Config, imageCache *ImageCache) *Client {
	return &Client{
		httpClient: httpClient,
		cfg:        cfg,
		imageCache: imageCache,
	}
}

func (c Client) Fetch(ctx context.Context) (SteamData, error) {
	var wg sync.WaitGroup
	var recent, top []Game
	var errRecent, errTop error

	wg.Add(2)

	go func() {
		defer wg.Done()
		recent, errRecent = c.fetchRecentGames(ctx)
	}()

	go func() {
		defer wg.Done()
		top, errTop = c.fetchTopGames(ctx, 12)
	}()

	wg.Wait()

	if errRecent != nil {
		return SteamData{}, fmt.Errorf("recent games: %w", errRecent)
	}
	if errTop != nil {
		return SteamData{}, fmt.Errorf("top games: %w", errTop)
	}

	return SteamData{Recent: recent, Top: top}, nil
}

func (c Client) fetchRecentGames(ctx context.Context) ([]Game, error) {
	params := url.Values{}
	params.Set("key", c.cfg.SteamApiKey)
	params.Set("steamid", c.cfg.SteamID)
	params.Set("format", "json")
	params.Set("count", "3")

	body, err := c.doGet(ctx, "https://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v1/?"+params.Encode())
	if err != nil {
		return nil, err
	}

	var resp recentGamesResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal recent games: %w", err)
	}

	return c.toGames(ctx, resp.Response.Games), nil
}

func (c Client) fetchTopGames(ctx context.Context, limit int) ([]Game, error) {
	params := url.Values{}
	params.Set("key", c.cfg.SteamApiKey)
	params.Set("steamid", c.cfg.SteamID)
	params.Set("format", "json")
	params.Set("include_appinfo", "true")
	params.Set("include_played_free_games", "true")

	body, err := c.doGet(ctx, "https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/?"+params.Encode())
	if err != nil {
		return nil, err
	}

	var resp ownedGamesResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal owned games: %w", err)
	}

	raw := resp.Response.Games

	filtered := make([]rawGame, 0, len(raw))
	for _, g := range raw {
		if !excludedAppIDs[g.AppID] {
			filtered = append(filtered, g)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].PlaytimeForever > filtered[j].PlaytimeForever
	})

	if len(filtered) > limit {
		filtered = filtered[:limit]
	}

	return c.toGames(ctx, filtered), nil
}

func (c Client) doGet(ctx context.Context, reqURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, body)
	}

	return body, nil
}
