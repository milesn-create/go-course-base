package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода пароля")
	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Println(geoData)
	weatherData, _ := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
