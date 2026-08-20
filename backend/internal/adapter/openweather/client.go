package openweather

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

func (c Client) Fetch(ctx context.Context) (Weather, error) {
	reqURL := c.getUrl(c.cfg.WeatherCity, c.cfg.WeatherApiKey)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return Weather{}, fmt.Errorf("build request: %w", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Weather{}, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Printf("close response body: %v", err)
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Weather{}, fmt.Errorf("read body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("unexpected status %d: %s", response.StatusCode, body)
	}

	var raw rawWeatherResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return Weather{}, fmt.Errorf("unmarshal: %w", err)
	}

	return convert(raw), nil
}

func (c Client) getUrl(city string, apiKey string) string {
	params := url.Values{}
	params.Set("q", city)
	params.Set("appid", apiKey)
	params.Set("units", "metric")

	return "https://api.openweathermap.org/data/2.5/weather?" + params.Encode()
}
