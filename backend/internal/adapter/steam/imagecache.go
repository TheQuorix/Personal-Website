package steam

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type ImageCache struct {
	dir        string
	publicPath string
	httpClient *http.Client

	mu       sync.Mutex
	inFlight map[string]*sync.WaitGroup
}

func NewImageCache(dir, publicPath string, httpClient *http.Client) (*ImageCache, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("create image cache dir: %w", err)
	}
	return &ImageCache{
		dir:        dir,
		publicPath: strings.TrimRight(publicPath, "/"),
		httpClient: httpClient,
		inFlight:   make(map[string]*sync.WaitGroup),
	}, nil
}

func (ic *ImageCache) Get(ctx context.Context, key, sourceURL string) (string, error) {
	if sourceURL == "" {
		return "", fmt.Errorf("empty source url")
	}

	if u, ok := ic.lookupExisting(key); ok {
		return u, nil
	}

	ic.mu.Lock()
	if wg, ok := ic.inFlight[key]; ok {
		ic.mu.Unlock()
		wg.Wait()
		if u, ok := ic.lookupExisting(key); ok {
			return u, nil
		}
		return "", fmt.Errorf("download failed for key %s", key)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ic.inFlight[key] = wg
	ic.mu.Unlock()

	defer func() {
		ic.mu.Lock()
		delete(ic.inFlight, key)
		ic.mu.Unlock()
		wg.Done()
	}()

	return ic.download(ctx, key, sourceURL)
}

func (ic *ImageCache) lookupExisting(key string) (string, bool) {
	matches, _ := filepath.Glob(filepath.Join(ic.dir, key+".*"))
	if len(matches) == 0 {
		return "", false
	}
	return ic.publicPath + "/" + filepath.Base(matches[0]), true
}

func (ic *ImageCache) download(ctx context.Context, key, sourceURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, sourceURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := ic.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download image: unexpected status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	filename := key + extFromContentType(resp.Header.Get("Content-Type"), sourceURL)
	fullPath := filepath.Join(ic.dir, filename)

	tmp := fullPath + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return "", fmt.Errorf("write image: %w", err)
	}
	if err := os.Rename(tmp, fullPath); err != nil {
		return "", fmt.Errorf("finalize image: %w", err)
	}

	return ic.publicPath + "/" + filename, nil
}

func extFromContentType(contentType, sourceURL string) string {
	if contentType != "" {
		if exts, _ := mime.ExtensionsByType(contentType); len(exts) > 0 {
			return exts[0]
		}
	}
	if idx := strings.LastIndex(sourceURL, "."); idx != -1 && len(sourceURL)-idx <= 5 {
		return sourceURL[idx:]
	}
	return ".jpg"
}
