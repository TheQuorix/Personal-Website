package openweather

func convert(raw rawWeatherResponse) Weather {
	return Weather{
		Temp:      raw.Main.Temp,
		FeelsLike: raw.Main.FeelsLike,
	}
}
