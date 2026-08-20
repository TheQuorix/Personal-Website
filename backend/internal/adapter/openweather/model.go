package openweather

type rawWeatherResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
}

type Weather struct {
	Temp      float64
	FeelsLike float64
}
