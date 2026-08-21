package steam

type SteamData struct {
	Recent []Game `json:"recent"`
	Top    []Game `json:"top"`
}

type Game struct {
	AppID           int    `json:"app_id"`
	Name            string `json:"name"`
	IconURL         string `json:"icon_url"`
	Playtime2Weeks  int    `json:"playtime_2weeks"`
	PlaytimeForever int    `json:"playtime_forever"`
}

type recentGamesResponse struct {
	Response struct {
		TotalCount int       `json:"total_count"`
		Games      []rawGame `json:"games"`
	} `json:"response"`
}

type ownedGamesResponse struct {
	Response struct {
		GameCount int       `json:"game_count"`
		Games     []rawGame `json:"games"`
	} `json:"response"`
}

type rawGame struct {
	AppID           int    `json:"appid"`
	Name            string `json:"name"`
	ImgIconURL      string `json:"img_icon_url"`
	Playtime2Weeks  int    `json:"playtime_2weeks"`
	PlaytimeForever int    `json:"playtime_forever"`
}
