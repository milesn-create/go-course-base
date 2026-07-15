package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {

	expected := "London"
	format := 3
	geo := geo.GeoData{
		City: expected,
	}

	result, err := weather.GetWeather(geo, format)
	if err != nil {
		t.Errorf("Пришла ошибка  %v", err)

	}
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v , получилось %v", expected, result)
	}

}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "0 formt", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			expected := "London"
			geo := geo.GeoData{
				City: expected,
			}

			_, err := weather.GetWeather(geo, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидалось %v , получилось %v", weather.ErrWrongFormat, err)
			}

		})

	}

}
