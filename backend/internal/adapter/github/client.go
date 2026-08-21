package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

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

func (c Client) Fetch(ctx context.Context) (GithubData, error) {
	var wg sync.WaitGroup
	var total, followers, repos int
	var weeks []Week
	var errCalendar, errProfile error

	wg.Add(2)

	go func() {
		defer wg.Done()
		total, weeks, errCalendar = c.fetchContributions(ctx)
	}()

	go func() {
		defer wg.Done()
		followers, repos, errProfile = c.fetchProfile(ctx)
	}()

	wg.Wait()

	if errCalendar != nil {
		return GithubData{}, fmt.Errorf("contributions: %w", errCalendar)
	}
	if errProfile != nil {
		return GithubData{}, fmt.Errorf("profile: %w", errProfile)
	}

	return GithubData{
		Followers:     followers,
		Repos:         repos,
		Contributions: total,
		Calendar:      weeks,
	}, nil
}

func (c Client) fetchContributions(ctx context.Context) (int, []Week, error) {
	const weeksLimit = 21

	now := time.Now().UTC()
	from := now.AddDate(-1, 0, 0)

	requestBody := graphQLRequest{
		Query: contributionsQuery,
		Variables: map[string]any{
			"username": c.cfg.GithubUsername,
			"from":     from.Format(time.RFC3339),
			"to":       now.Format(time.RFC3339),
		},
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return 0, nil, fmt.Errorf("marshal query: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.github.com/graphql", bytes.NewReader(payload))
	if err != nil {
		return 0, nil, fmt.Errorf("build request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.cfg.GuthibToken)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return 0, nil, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Printf("close response body: %v", err)
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("read body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf("unexpected status %d: %s", response.StatusCode, body)
	}

	var parsed contributionsResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return 0, nil, fmt.Errorf("unmarshal: %w", err)
	}

	if len(parsed.Errors) > 0 {
		return 0, nil, fmt.Errorf("graphql error: %s", parsed.Errors[0].Message)
	}

	calendar := parsed.Data.User.ContributionsCollection.ContributionCalendar
	weeks, err := toWeeks(calendar.Weeks)
	if err != nil {
		return 0, nil, err
	}

	if len(weeks) > weeksLimit {
		weeks = weeks[len(weeks)-weeksLimit:]
	}

	return calendar.TotalContributions, weeks, nil
}

func (c Client) fetchProfile(ctx context.Context) (int, int, error) {
	reqURL := "https://api.github.com/users/" + c.cfg.GithubUsername

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, 0, fmt.Errorf("build request: %w", err)
	}

	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("Authorization", "Bearer "+c.cfg.GuthibToken)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return 0, 0, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Printf("close response body: %v", err)
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("read body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("unexpected status %d: %s", response.StatusCode, body)
	}

	var user rawUser
	if err := json.Unmarshal(body, &user); err != nil {
		return 0, 0, fmt.Errorf("unmarshal: %w", err)
	}

	return user.Followers, user.PublicRepos, nil
}
