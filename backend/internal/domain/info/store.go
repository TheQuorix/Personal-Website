package info

import "sync"

type infoStore struct {
	mu   sync.RWMutex
	data Info
}

var store = &infoStore{}

func (p Poller) GetCached() Info {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return store.data
}

func (s *infoStore) set(update func(*Info)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	update(&s.data)
}
