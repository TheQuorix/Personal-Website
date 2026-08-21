package github

import "time"

type graphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

type rawUser struct {
	Followers   int `json:"followers"`
	PublicRepos int `json:"public_repos"`
}

type contributionsResponse struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int `json:"totalContributions"`
					Weeks              []struct {
						ContributionDays []struct {
							Date              string `json:"date"`
							ContributionCount int    `json:"contributionCount"`
							ContributionLevel string `json:"contributionLevel"`
						} `json:"contributionDays"`
					} `json:"weeks"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type Github struct {
	Followers     int    `json:"followers"`
	Repos         int    `json:"repos"`
	Contributions int    `json:"contributions"`
	Calendar      []Week `json:"calendar"`
}

type Week struct {
	Days []Day `json:"days"`
}

type Day struct {
	Date  time.Time `json:"date"`
	Count int       `json:"count"`
	Level int       `json:"level"`
}
