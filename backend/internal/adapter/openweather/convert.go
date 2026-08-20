package openweather

func convert(rawWeather rawWeatherResponse) Weather {
	return Weather{
		Temp:      rawWeather.Main.Temp,
		FeelsLike: rawWeather.Main.FeelsLike,
	}
}
