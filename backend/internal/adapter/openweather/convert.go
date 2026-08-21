package openweather

func convert(raw rawWeatherResponse) WeatherData {
	return WeatherData{
		Temp:      raw.Main.Temp,
		FeelsLike: raw.Main.FeelsLike,
	}
}
