package info

import (
	"context"
	"fmt"
	"time"

	"github.com/TheQuorix/Personal-Website/internal/adapter/github"
	"github.com/TheQuorix/Personal-Website/internal/adapter/lastfm"
	"github.com/TheQuorix/Personal-Website/internal/adapter/openweather"
	"github.com/TheQuorix/Personal-Website/internal/adapter/steam"
)

type Poller struct {
	weatherClient *openweather.Client
	musicClient   *lastfm.Client
	steamClient   *steam.Client
	githubClient  *github.Client
}

func NewPoller(weatherClient *openweather.Client, musicClient *lastfm.Client, steamClient *steam.Client, githubClient *github.Client) *Poller {
	return &Poller{
		weatherClient: weatherClient,
		musicClient:   musicClient,
		steamClient:   steamClient,
		githubClient:  githubClient,
	}
}

func (p Poller) StartPolling(ctx context.Context) {
	go p.pollWeather(ctx, time.Minute)
	go p.pollMusic(ctx, 15*time.Second)
	go p.pollGithub(ctx, 10*time.Minute)
	go p.pollSteam(ctx, 5*time.Minute)
}

func (p Poller) pollWeather(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	p.updateWeather(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.updateWeather(ctx)
		}
	}
}

func (p Poller) updateWeather(parentCtx context.Context) {
	ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
	defer cancel()

	data, err := p.weatherClient.Fetch(ctx)
	if err != nil {
		fmt.Println("failed to get weather:", err)
		return
	}

	store.set(func(i *Info) {
		i.Weather = data
	})
}

func (p Poller) pollMusic(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	p.updateMusic(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.updateMusic(ctx)
		}
	}
}

func (p Poller) updateMusic(parentCtx context.Context) {
	ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
	defer cancel()

	track, err := p.musicClient.Fetch(ctx)
	if err != nil {
		fmt.Println("failed to get last track:", err)
		return
	}

	store.set(func(i *Info) {
		i.Music = track
	})
}

func (p Poller) pollGithub(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	p.updateGithub(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.updateGithub(ctx)
		}
	}
}

func (p Poller) updateGithub(parentCtx context.Context) {
	ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
	defer cancel()

	data, err := p.githubClient.Fetch(ctx)
	if err != nil {
		fmt.Println("failed to get github data:", err)
		return
	}

	store.set(func(i *Info) {
		i.Github = data
	})
}

func (p Poller) pollSteam(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	p.updateSteam(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.updateSteam(ctx)
		}
	}
}

func (p Poller) updateSteam(parentCtx context.Context) {
	ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
	defer cancel()

	data, err := p.steamClient.Fetch(ctx)
	if err != nil {
		fmt.Println("failed to get steam data:", err)
		return
	}

	store.set(func(i *Info) {
		i.Steam = data
	})
}
