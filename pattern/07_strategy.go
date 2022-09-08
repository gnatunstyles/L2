package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type WeatherToday interface {
	show(*Weather)
}

type Weather struct {
	season      string
	sky         string
	wet         int
	temperature WeatherToday
}

func newWeather(season string, temp WeatherToday) *Weather {
	return &Weather{
		season:      season,
		sky:         "clear",
		wet:         432,
		temperature: temp,
	}
}

func (w *Weather) setTemperature(t WeatherToday) {
	w.temperature = t
}
func (w *Weather) showTemperature() {
	w.temperature.show(w)
}

type winter struct{}

func (w *winter) show(weather *Weather) {
	if weather.sky == "clear" {
		fmt.Println(-25)
		return
	} else {
		fmt.Println(-40)
		return
	}
}

type summer struct{}

func (s *summer) show(weather *Weather) {
	if weather.sky == "clear" {
		fmt.Println(45)
		return
	} else {
		fmt.Println(30)
		return
	}
}
