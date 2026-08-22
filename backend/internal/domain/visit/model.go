package visit

import "time"

type Visit struct {
	ID          int64
	Path        string
	VisitorHash string
	VisitedAt   time.Time
	CreatedAt   time.Time
}

type DailyStat struct {
	Date   string `json:"date"`
	Total  int    `json:"total"`
	Unique int    `json:"unique"`
}

type Stats struct {
	TotalVisits  int         `json:"total_visits"`
	UniqueVisits int         `json:"unique_visits"`
	Daily        []DailyStat `json:"daily"`
}
