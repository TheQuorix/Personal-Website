package github

import (
	"fmt"
	"time"
)

func toWeeks(raw []struct {
	ContributionDays []struct {
		Date              string `json:"date"`
		ContributionCount int    `json:"contributionCount"`
		ContributionLevel string `json:"contributionLevel"`
	} `json:"contributionDays"`
}) ([]Week, error) {
	weeks := make([]Week, 0, len(raw))

	for _, w := range raw {
		days := make([]Day, 0, len(w.ContributionDays))
		for _, d := range w.ContributionDays {
			date, err := time.Parse("2006-01-02", d.Date)
			if err != nil {
				return nil, fmt.Errorf("parse date %q: %w", d.Date, err)
			}
			days = append(days, Day{
				Date:  date,
				Count: d.ContributionCount,
				Level: levelToInt(d.ContributionLevel),
			})
		}
		weeks = append(weeks, Week{Days: days})
	}

	return weeks, nil
}

func levelToInt(level string) int {
	switch level {
	case "FIRST_QUARTILE":
		return 1
	case "SECOND_QUARTILE":
		return 2
	case "THIRD_QUARTILE":
		return 3
	case "FOURTH_QUARTILE":
		return 4
	default:
		return 0
	}
}
