package visit

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
)

type Service struct {
	repo Repository
	salt string
}

func NewService(repo Repository, salt string) *Service {
	return &Service{repo: repo, salt: salt}
}

func (s *Service) hash(ip, ua string) string {
	sum := sha256.Sum256([]byte(ip + ua + s.salt))
	return hex.EncodeToString(sum[:])
}

func (s *Service) Track(ctx context.Context, path, ip, ua string) error {
	if path == "" {
		path = "/"
	}
	return s.repo.Insert(ctx, path, s.hash(ip, ua))
}

func (s *Service) GetStats(ctx context.Context) (*Stats, error) {
	total, err := s.repo.CountTotal(ctx)
	if err != nil {
		return nil, err
	}
	unique, err := s.repo.CountUnique(ctx)
	if err != nil {
		return nil, err
	}
	daily, err := s.repo.DailyStats(ctx, 30)
	if err != nil {
		return nil, err
	}
	return &Stats{TotalVisits: total, UniqueVisits: unique, Daily: daily}, nil
}
