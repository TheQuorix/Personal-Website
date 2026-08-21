package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const iconFetchConcurrency = 5

var (
	gridIDCacheMu sync.RWMutex
	gridIDCache   = make(map[int]int)
)

func (c Client) toGames(ctx context.Context, raw []rawGame) []Game {
	filtered := make([]rawGame, 0, len(raw))
	for _, g := range raw {
		if !excludedAppIDs[g.AppID] {
			filtered = append(filtered, g)
		}
	}

	icons := make([]string, len(filtered))

	var wg sync.WaitGroup
	sem := make(chan struct{}, iconFetchConcurrency)

	for i, g := range filtered {
		wg.Add(1)
		go func(i int, g rawGame) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			icons[i] = c.resolveIcon(ctx, g.AppID, g.ImgIconURL)
		}(i, g)
	}
	wg.Wait()

	games := make([]Game, len(filtered))
	for i, g := range filtered {
		games[i] = Game{
			AppID:           g.AppID,
			Name:            g.Name,
			IconURL:         icons[i],
			Playtime2Weeks:  g.Playtime2Weeks,
			PlaytimeForever: g.PlaytimeForever,
		}
	}

	return games
}

func (c Client) resolveIcon(ctx context.Context, appID int, iconHash string) string {
	sourceURL := ""

	if c.cfg.SteamGridDbApiKey != "" {
		if gridID, err := c.steamGridGameID(ctx, appID); err == nil {
			if url, err := c.fetchSteamGridIcon(ctx, gridID); err == nil && url != "" {
				sourceURL = url
			}
		}
	}

	if sourceURL == "" {
		sourceURL = steamIconURL(appID, iconHash)
	}
	if sourceURL == "" {
		return ""
	}

	key := fmt.Sprintf("%d", appID)
	localURL, err := c.imageCache.Get(ctx, key, sourceURL)
	if err != nil {
		return sourceURL
	}

	return localURL
}

func steamIconURL(appID int, iconHash string) string {
	if iconHash == "" {
		return ""
	}
	return fmt.Sprintf("https://media.steampowered.com/steamcommunity/public/images/apps/%d/%s.jpg", appID, iconHash)
}

func (c Client) steamGridGameID(ctx context.Context, appID int) (int, error) {
	gridIDCacheMu.RLock()
	id, ok := gridIDCache[appID]
	gridIDCacheMu.RUnlock()
	if ok {
		return id, nil
	}

	reqURL := fmt.Sprintf("https://www.steamgriddb.com/api/v2/games/steam/%d", appID)
	body, err := c.sgdbGet(ctx, reqURL)
	if err != nil {
		return 0, err
	}

	var result struct {
		Success bool `json:"success"`
		Data    struct {
			ID int `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("unmarshal game lookup: %w", err)
	}
	if !result.Success {
		return 0, fmt.Errorf("steamgriddb: game not found for appid %d", appID)
	}

	gridIDCacheMu.Lock()
	gridIDCache[appID] = result.Data.ID
	gridIDCacheMu.Unlock()

	return result.Data.ID, nil
}

func (c Client) fetchSteamGridIcon(ctx context.Context, gridID int) (string, error) {
	reqURL := fmt.Sprintf("https://www.steamgriddb.com/api/v2/icons/game/%d", gridID)
	body, err := c.sgdbGet(ctx, reqURL)
	if err != nil {
		return "", err
	}

	var result struct {
		Success bool `json:"success"`
		Data    []struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("unmarshal icons: %w", err)
	}
	if !result.Success || len(result.Data) == 0 {
		return "", fmt.Errorf("steamgriddb: no icons for game %d", gridID)
	}

	return result.Data[0].URL, nil
}

func (c Client) sgdbGet(ctx context.Context, reqURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.cfg.SteamGridDbApiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("steamgriddb: unexpected status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
